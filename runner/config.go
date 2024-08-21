package runner

import (
	"crypto/tls"
	"net/http"
	"time"
)

type Config struct {
	HTTPS        bool
	VerifySSL    bool
	Emoji        bool
	HideFails    bool
	OnlyVuln     bool
	Concurrency  int
	Timeout      int
	Targets      string
	Target       string
	Output       string
	client       *http.Client
	fingerprints []Fingerprint
	skipfingerprints []Fingerprint
}

func (s *Config) initHTTPClient() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !s.VerifySSL},
	}

	timeout := time.Duration(s.Timeout) * time.Second
	client := &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	s.client = client
}

func (c *Config) loadFingerprints() error {
	validFingerprints, skippedFingerprints, err := Fingerprints()
	if err != nil {
		return err
	}
	c.skipfingerprints = skippedFingerprints
	c.fingerprints = validFingerprints
	return nil
}
