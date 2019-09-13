package main

import (
	"github.com/lukasikic/subzy/src"
	"flag"
	"fmt"
	"os"
)

func main() {

	settings := src.Settings{}

	flag.StringVar(&settings.Targets, "targets", "", "File path to list of subdomains to be scanned")
	flag.BoolVar(&settings.Https, "https", false, "Force https protocol if not no protocol defined for target (default false)")
	flag.IntVar(&settings.Concurrency, "concurrency", 10, "Number of concurrent checks")
	flag.StringVar(&settings.Target, "target", "", "Single or multiple subdomains separated by comma")
	flag.BoolVar(&settings.VerifySSL, "verify_ssl", false, "If set to true it won't check sites with insecure SSL and return HTTP Error")
	flag.IntVar(&settings.Timeout, "timeout", 10, "Request timeout in seconds")
	flag.BoolVar(&settings.HideFails, "hide_fails", false, "Don't display failed results")

	flag.Parse()

	if settings.Target == "" && settings.Targets == "" {
		fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	src.Process(settings)

}
