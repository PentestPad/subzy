# Subzy
Subdomain takeover tool based on [can-i-take-over-xyz](https://github.com/EdOverflow/can-i-take-over-xyz/blob/master/README.md) fingerprints

![Subzy subdomain takeover](https://i.imgur.com/gw8RGo9.png "Logo Title Text 1")

## Installation
```git clone https://github.com/LukaSikic/subzy```

## Usage
``` 
$ ./subzy  --help

Usage of ./subzy:
  -concurrency int
        Number of concurrent checks (default 10)
  -https
        Force https protocol if not provided in the list (default false)
  -targets string
        File path to list of subdomains to be scanned (default "list.txt")

./subzy --targets=list.txt --https=false --concurrency=20
```