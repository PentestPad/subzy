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

	fmt.Println("[ * ] Loaded", len(subdomains), "targets")
	fmt.Println("[ * ] Loaded", len(fingerprints), "fingerprints")

	if config.Output != "" {
		fmt.Printf("[ * ] Output filename: %s\n", config.Output)
		fmt.Println(isEnabled(config.OnlyVuln), "Save only vulnerable subdomains")
	}

	fmt.Println(isEnabled(config.HTTPS), "HTTPS by default (--https)")
	fmt.Println("[", config.Concurrency, "]", "Concurrent requests (--concurrency)")
	fmt.Println(isEnabled(config.VerifySSL), "Check target only if SSL is valid (--verify_ssl)")
	fmt.Println("[", config.Timeout, "]", "HTTP request timeout (in seconds) (--timeout)")
	fmt.Println(isEnabled(config.HideFails), "Show only potentially vulnerable subdomains (--hide_fails)")

	const ExtraChannelCapacity = 5
	subdomainCh := make(chan string, config.Concurrency+ExtraChannelCapacity)
	resCh := make(chan *subdomainResult, config.Concurrency)

	var wg sync.WaitGroup
	wg.Add(config.Concurrency)

	var results []*subdomainResult
	go collectResults(resCh, &results, config)

	for i := 0; i < config.Concurrency; i++ {
		go processor(subdomainCh, resCh, config, &wg)
	}

	distributeSubdomains(subdomains, subdomainCh)
	wg.Wait()
	close(resCh)

	if config.Output != "" {
		if err := saveResults(config.Output, results); err != nil {
			return err
		}
	}

	return nil
}

func processor(subdomainCh <-chan string, resCh chan<- *subdomainResult, c *Config, wg *sync.WaitGroup) {
	defer wg.Done()
	for subdomain := range subdomainCh {
		result := c.checkSubdomain(subdomain)

		res := &subdomainResult{
			Subdomain:     subdomain,
			Status:        string(result.ResStatus),
			Engine:        result.Entry.Service,
			Documentation: result.Entry.Documentation,
		}

		if result.Status == aurora.Green("VULNERABLE") {
			fmt.Print("-----------------\n")
			fmt.Println("[", result.Status, "]", " - ", subdomain, " [", result.Entry.Service, "]")
			fmt.Println("[", aurora.Blue("DISCUSSION"), "]", " - ", result.Entry.Discussion)
			fmt.Println("[", aurora.Blue("DOCUMENTATION"), "]", " - ", result.Entry.Documentation)
			fmt.Print("-----------------\n")
		} else {
			if !c.HideFails {
				fmt.Println("[", result.Status, "]", " - ", subdomain)
			}
		}

		resCh <- res
	}
}

func distributeSubdomains(subdomains []string, subdomainCh chan<- string) {
	for _, subdomain := range subdomains {
		subdomainCh <- subdomain
	}
	close(subdomainCh)
}

func collectResults(resCh <-chan *subdomainResult, results *[]*subdomainResult, config *Config) {
	for r := range resCh {
		if config.Output != "" && (!config.OnlyVuln || r.Status == "vulnerable") {
			*results = append(*results, r)
		}
	}
}

func saveResults(filename string, results []*subdomainResult) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(results); err != nil {
		return err
	}

	fmt.Printf("[ * ] Saved output to %q\n", filename)
	return nil
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
