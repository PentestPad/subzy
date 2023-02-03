package runner

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/logrusorgru/aurora"
)

func Process(config *Config) error {

	fingerprints, err := Fingerprints()
	if err != nil {
		return fmt.Errorf("Process: %v", err)
	}

	config.initHTTPClient()
	config.loadFingerprints()
	subdomains := getSubdomains(config)

	fmt.Println("[ * ]", "Loaded", len(subdomains), "targets")
	fmt.Println("[ * ]", "Loaded", len(fingerprints), "fingerprints")
	if config.Output != "" {
		fmt.Printf("[ * ] Output filename: %s\n", config.Output)
		fmt.Println(isEnabled(config.OnlyVuln), "Save only vulnerable subdomains")
	}

	fmt.Println(isEnabled(config.HTTPS), "HTTPS by default (--https)")
	fmt.Println("[", config.Concurrency, "]", "Concurrent requests (--concurrency)")
	fmt.Println(isEnabled(config.VerifySSL), "Check target only if SSL is valid (--verify_ssl)")
	fmt.Println("[", config.Timeout, "]", "HTTP request timeout (in seconds) (--timeout)")
	fmt.Println(isEnabled(config.HideFails), "Show only potentially vulnerable subdomains (--hide_fails)")

	subdomainCh := make(chan string, config.Concurrency+5)
	resCh := make(chan *subdomainResult, config.Concurrency)

	var wg sync.WaitGroup
	wg.Add(config.Concurrency)

	var results []*subdomainResult
	go func() {
		for r := range resCh {
			if config.Output != "" {
				if config.OnlyVuln && r.Status != ResultVulnerable {
					continue
				}
				results = append(results, r)
			}
		}
	}()

	for i := 0; i < config.Concurrency; i++ {
		go processor(subdomainCh, resCh, config, &wg)
	}

	go func() {
		for _, subdomain := range subdomains {
			subdomainCh <- subdomain
		}
		close(subdomainCh)
	}()

	wg.Wait()
	close(resCh)

	if config.Output != "" {
		f, err := os.OpenFile(config.Output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer f.Close()

		enc := json.NewEncoder(f)
		enc.SetIndent("", "  ")

		if err := enc.Encode(results); err != nil {
			return err
		}

		fmt.Printf("[ * ] Saved output to %q\n", config.Output)
	}

	return nil
}

func processor(subdomainCh chan string, resCh chan *subdomainResult, c *Config, wg *sync.WaitGroup) {
	defer wg.Done()
	for subdomain := range subdomainCh {
		result := c.checkSubdomain(subdomain)
		resCh <- &subdomainResult{
			Subdomain:     subdomain,
			Status:        string(result.resStatus),
			Engine:        result.entry.Engine,
			Documentation: result.entry.Documentation,
		}

		if result.status == aurora.Green("VULNERABLE") {
			fmt.Print("-----------------\r\n")
			fmt.Println("[ ", result.status, " ]", " - ", subdomain, " [ ", result.entry.Engine, " ] ")
			fmt.Println("[ ", aurora.Blue("DISCUSSION"), " ]", " - ", result.entry.Discussion)
			fmt.Println("[ ", aurora.Blue("DOCUMENTATION"), " ]", " - ", result.entry.Documentation)

			fmt.Print("-----------------\r\n")

		} else {
			if !c.HideFails {
				fmt.Println("[ ", result.status, " ]", " - ", subdomain)
			}
		}
	}
}

func generator(subdomain string, subdomainCh chan string) {
	subdomainCh <- subdomain
}

func getSubdomains(c *Config) []string {
	if c.Target == "" {
		subdomains, err := readSubdomains(c.Targets)
		if err != nil {
			log.Fatalf("Error reading subdomains: %s", err)
		}

		return subdomains
	}

	return strings.Split(c.Target, ",")
}
