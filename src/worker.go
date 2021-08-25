package src

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

type Result struct {
	status aurora.Value
	entry  Fingerprint
}

func checkSubdomain(subdomain string, settings *Settings) Result {

	if isValidUrl(subdomain) == false {
		if settings.HTTPS {
			subdomain = "https://" + subdomain
		} else {
			subdomain = "http://" + subdomain
		}
	}

	client := httpClient(settings)

	resp, err := client.Get(subdomain)
	if err != nil {
		return Result{aurora.Red("HTTP ERROR"), Fingerprint{}}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Result{aurora.Red("RESPONSE ERROR"), Fingerprint{}}
	}

	return matchResponse(string(body))
}

func matchResponse(body string) Result {
	fingerprints, err := Fingerprints()
	if err != nil {
		return Result{aurora.Red("Fingerprint error"), Fingerprint{}}
	}

	for _, fingerprint := range fingerprints {
		if strings.Contains(body, fingerprint.Fingerprint) {
			
			for _, false_positive_string := range fingerprint.False_Positive {
				if len(false_positive_string) > 0 {

					if strings.Contains(body, false_positive_string) {
						
						return Result{aurora.Red("NOT VULNERABLE"), Fingerprint{}}
					}
				}
			}
			return Result{aurora.Green("VULNERABLE"), fingerprint}

		}
	}

	return Result{aurora.Red("NOT VULNERABLE"), Fingerprint{}}

}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}

func httpClient(settings *Settings) *http.Client {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !settings.VerifySSL},
	}

	timeout := time.Duration(time.Duration(settings.Timeout) * time.Second)
	client := &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	return client

}
