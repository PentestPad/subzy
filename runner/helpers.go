package runner

import "net/url"

func isEnabled(setting bool) string {
	if setting == true {
		return "[ Yes ]"
	}
	return "[ No ]"
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}

type subdomainResult struct {
	Subdomain     string `json:"subdomain"`
	Status        string `json:"status"`
	Engine        string `json:"engine"`
	Documentation string `json:"documentation"`
	Discussion    string `json:"discussion"`
}
