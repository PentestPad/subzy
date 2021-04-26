package src

import (
	"fmt"
	"log"
	"strings"

	"github.com/logrusorgru/aurora"
)

func Process(settings *Settings) error {

	fingerprints, err := Fingerprints()
	if err != nil {
		return fmt.Errorf("Process: %v", err)
	}

	subdomains := getSubdomains(settings)

	fmt.Println("[ * ]", "Loaded", len(subdomains), "targets")
	fmt.Println("[ * ]", "Loaded", len(fingerprints), "fingerprints")

	fmt.Println(isEnabled(settings.HTTPS), "HTTPS by default (--https)")
	fmt.Println("[", settings.Concurrency, "]", "Concurrent requests (--concurrency)")
	fmt.Println(isEnabled(settings.VerifySSL), "Check target only if SSL is valid (--verify_ssl)")
	fmt.Println("[", settings.Timeout, "]", "HTTP request timeout (in seconds) (--timeout)")
	fmt.Println(isEnabled(settings.HideFails), "Show only potentially vulnerable subdomains (--hide_fails)")

	subdomainCh := make(chan string)
	sizeCh := make(chan string)

	for i := 0; i < settings.Concurrency; i++ {
		go processor(subdomainCh, sizeCh, settings)
	}

	for _, subdomain := range subdomains {
		go generator(subdomain, subdomainCh)
	}

	for i := 0; i < len(subdomains); i++ {
		<-sizeCh
	}

	return nil
}

func isEnabled(setting bool) string {
	if setting == true {
		return "[ Yes ]"
	}
	return "[ No ]"
}

func processor(subdomainCh chan string, sizeCh chan string, settings *Settings) {
	for {
		subdomain := <-subdomainCh

		result := checkSubdomain(subdomain, *settings)

		if result.status == aurora.Green("VULNERABLE") {
			fmt.Print("-----------------\r\n")
			fmt.Println("[ ", result.status, " ]", " - ", subdomain, " [ ", result.entry.Engine, " ] ")
			fmt.Println("[ ", aurora.Blue("DISCUSSION"), " ]", " - ", result.entry.Discussion)
			fmt.Println("[ ", aurora.Blue("DOCUMENTATION"), " ]", " - ", result.entry.Documentation)

			fmt.Print("-----------------\r\n")

		} else {
			if !settings.HideFails {
				fmt.Println("[ ", result.status, " ]", " - ", subdomain)
			}
		}

		sizeCh <- ""
	}
}

func generator(subdomain string, subdomainCh chan string) {
	subdomainCh <- subdomain
}

func getSubdomains(settings *Settings) []string {
	if settings.Target == "" {
		subdomains, err := readSubdomains(settings.Targets)
		if err != nil {
			log.Fatalf("Error reading subdomains: %s", err)
		}

		return subdomains
	}

	return strings.Split(settings.Target, ",")
}
