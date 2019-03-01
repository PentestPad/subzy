package src

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
)

type Settings struct {
	Targets 	string
	Https		bool
	Concurrency	int
}

func Process(settings Settings) {

	subdomains, err := readSubdomains(settings.Targets)
	if err != nil {
		log.Fatalf("Error reading subdomains: %s", err)
	}


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

}

func processor(subdomainCh chan string, sizeCh chan string, settings Settings){
	for {
		subdomain := <-subdomainCh

		result := checkSubdomain(subdomain, settings)

		if result.status == aurora.Green("VULNERABLE") {
			fmt.Print("-----------------\r\n")
			fmt.Println("[ ", result.status, " ]", " - ", subdomain, " [ ", result.entry.engine, " ] ")
			fmt.Println("[ ", aurora.Blue("DISCUSSION"), " ]", " - ", result.entry.discussion)
			fmt.Println("[ ", aurora.Blue("DOCUMENTATION"), " ]", " - ", result.entry.documentation)

			fmt.Print("-----------------\r\n")

		} else {
			fmt.Println("[ ", result.status, " ]", " - ", subdomain)
		}

		sizeCh <- ""
	}
}

func generator(subdomain string, subdomainCh chan string) {
	subdomainCh <- subdomain
}