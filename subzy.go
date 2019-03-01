package main

import (
"./src"
"flag"
"fmt"
"os"
)

func main() {

	settings := src.Settings{}

	flag.StringVar(&settings.Targets, "targets", "list.txt", "File path to list of subdomains to be scanned")
	flag.BoolVar(&settings.Https, "https", false, "Force https protocol if not provided in the list (default false)")
	flag.IntVar(&settings.Concurrency, "concurrency", 10, "Number of concurrent checks")

	flag.Parse()

	if settings.Targets == "" {
		fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
		flag.PrintDefaults()
	}

	src.Process(settings)

}
