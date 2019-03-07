## Subzy 

Subdomain takeover tool which works based on matching response fingerprings from [can-i-take-over-xyz](https://github.com/EdOverflow/can-i-take-over-xyz/blob/master/README.md) 

![Twitter Follow](https://img.shields.io/twitter/follow/return_0x.svg?label=Follow&style=social) 

![Subzy subdomain takeover](https://i.imgur.com/QvZNFdF.png "Subzy subdomain takeover")

### Installation
___
Clone GitHub repo   
```git clone https://github.com/LukaSikic/subzy```  

Run program  
```./subzy/subzy```

If you get an error `exec format error: ./subzy`, you need to [install Golang](https://golang.org/doc/install) for your OS and compile the program by running `go build subzy.go` which will generate new `subzy` binary file

### Options
___
Only required flag is either `--target` or `--targets`  

`--target` (string) - Set single or multiple (comma separated) target subdomain/s  
`--targets` (string) - File name/path to list of subdomains    
`--concurrency` (integer) - Number of concurrent checks (default 10)    
`--hide_fails` (boolean) - Hide failed checks and invulnerable subdomains (default false)    
`--https` (boolean) - Use HTTPS by default if protocol not defined on targeted subdomain (default false)  
`--timeout` (integer) - HTTP request timeout in seconds (default 10)  
`--verify_ssl` (boolean) - If set to true, it won't check site with invalid SSL

### Usage
___
Target subdomain can have protocol defined, if not `http://` will be used by default if `--https` not specifically set to true.

-  List of subdomains
   - ````./subzy --targets=list.txt````

- Single or few subdomains 
  - ```./subzy --target=test.google.com```
  - ```./subzy --target=test.google.com,https://test.yahoo.com```
