package runner

import (
	"github.com/logrusorgru/aurora"
	"io"
	"strings"
)

type resultStatus string

const (
	ResultHTTPError     resultStatus = "http error"
	ResultResponseError              = "response error"
	ResultVulnerable                 = "vulnerable"
	ResultNotVulnerable              = "not vulnerable"
)

type Result struct {
	resStatus resultStatus
	status    aurora.Value
	entry     Fingerprint
}

func (c *Config) checkSubdomain(subdomain string) Result {
	if isValidUrl(subdomain) == false {
		if c.HTTPS {
			subdomain = "https://" + subdomain
		} else {
			subdomain = "http://" + subdomain
		}
	}

	resp, err := c.client.Get(subdomain)
	if err != nil {
		return Result{ResultHTTPError, aurora.Red("HTTP ERROR"), Fingerprint{}}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return Result{ResultResponseError, aurora.Red("RESPONSE ERROR"), Fingerprint{}}
	}
	resp.Body.Close()

	return c.matchResponse(string(body))
}

func (c *Config) matchResponse(body string) Result {
	for _, fingerprint := range c.fingerprints {
		if strings.Contains(body, fingerprint.Fingerprint) && fingerprint.Status != "Not vulnerable" {
			for _, false_positive_string := range fingerprint.False_Positive {
				if len(string(false_positive_string)) > 0 {

					if strings.Contains(body, string(false_positive_string)) {
						return Result{ResultNotVulnerable, aurora.Red("NOT VULNERABLE"), Fingerprint{}}
					}
				}
			}
			return Result{ResultVulnerable, aurora.Green("VULNERABLE"), fingerprint}
		}
	}

	return Result{ResultNotVulnerable, aurora.Red("NOT VULNERABLE"), Fingerprint{}}
}
