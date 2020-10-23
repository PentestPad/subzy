## Subzy 

Subdomain takeover tool which works based on matching response fingerprings from [can-i-take-over-xyz](https://github.com/EdOverflow/can-i-take-over-xyz/blob/master/README.md) 

<a href="https://twitter.com/intent/follow?screen_name=return_0x">
        <img src="https://img.shields.io/twitter/follow/return_0x.svg?style=social&logo=twitter"
            alt="follow on Twitter"></a>


![Subzy subdomain takeover](https://i.imgur.com/QvZNFdF.png "Subzy subdomain takeover")

### Installation

```bash
go get -u -v github.com/lukasikic/subzy
go install -v github.com/lukasikic/subzy
```

If `$GOBIN` and `$GOPATH` are [properly set](https://github.com/golang/go/wiki/SettingGOPATH#bash), execute the program as:

```bash
subzy
``` 

If you get an error `exec format error: ./subzy`, you need to [install Golang](https://golang.org/doc/install) for your OS and compile the program by running `go build subzy.go` which will generate new `subzy` binary file

### Options

Only required flag is either `--target` or `--targets`  

`--target` (string) - Set single or multiple (comma separated) target subdomain/s  
`--targets` (string) - File name/path to list of subdomains    
`--concurrency` (integer) - Number of concurrent checks (default 10)    
`--hide_fails` (boolean) - Hide failed checks and invulnerable subdomains (default false)    
`--https` (boolean) - Use HTTPS by default if protocol not defined on targeted subdomain (default false)  
`--timeout` (integer) - HTTP request timeout in seconds (default 10)  
`--verify_ssl` (boolean) - If set to true, it won't check site with invalid SSL

### Usage

Target subdomain can have protocol defined, if not `http://` will be used by default if `--https` not specifically set to true.

-  List of subdomains
   - ````./subzy -targets list.txt````

- Single or multiple targets 
  - ```./subzy -target test.google.com```
  - ```./subzy -target test.google.com,https://test.yahoo.com```
