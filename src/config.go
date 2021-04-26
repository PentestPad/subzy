package src

import (
	"flag"
	"fmt"
	"os"
)

type Settings struct {
	Targets     string
	HTTPS       bool
	Concurrency int
	Target      string
	Timeout     int
	VerifySSL   bool
	Emoji       bool
	HideFails   bool
}

func PrintHelp() {
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

// ParseConfiguration will return the pointer to the populated Settings struct
func ParseConfiguration() *Settings {
	opts := Settings{}

	flag.StringVar(&opts.Targets, "targets", "", "File path to list of subdomains to be scanned")
	flag.BoolVar(&opts.HTTPS, "https", false, "Force https protocol if not no protocol defined for target (default false)")
	flag.IntVar(&opts.Concurrency, "concurrency", 10, "Number of concurrent checks")
	flag.StringVar(&opts.Target, "target", "", "Single or multiple subdomains separated by comma")
	flag.BoolVar(&opts.VerifySSL, "verify_ssl", false, "If set to true it won't check sites with insecure SSL and return HTTP Error")
	flag.IntVar(&opts.Timeout, "timeout", 10, "Request timeout in seconds")
	flag.BoolVar(&opts.HideFails, "hide_fails", false, "Don't display failed results")

	flag.Parse()

	if opts.Target == "" && opts.Targets == "" {
		return nil
	}

	return &opts
}
