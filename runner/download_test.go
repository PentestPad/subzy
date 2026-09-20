package runner

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const validFingerprints = `[{"service":"Example","cname":["example.com"],"vulnerable":true}]`

func configureFingerprintTest(t *testing.T, handler http.HandlerFunc) {
	t.Helper()

	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	oldURL := fingerprintURL
	oldClient := fingerprintHTTPClient
	oldAttempts := fingerprintAttempts
	oldDelay := fingerprintRetryDelay
	fingerprintURL = server.URL
	fingerprintHTTPClient = server.Client()
	fingerprintAttempts = 1
	fingerprintRetryDelay = 0
	t.Cleanup(func() {
		fingerprintURL = oldURL
		fingerprintHTTPClient = oldClient
		fingerprintAttempts = oldAttempts
		fingerprintRetryDelay = oldDelay
	})
}

func TestCheckIntegrity(t *testing.T) {
	configureFingerprintTest(t, func(response http.ResponseWriter, _ *http.Request) {
		_, _ = response.Write([]byte(validFingerprints))
	})

	localPath := filepath.Join(t.TempDir(), "fingerprints.json")
	if err := os.WriteFile(localPath, []byte(validFingerprints), 0600); err != nil {
		t.Fatal(err)
	}

	matches, err := checkIntegrity(localPath)
	if err != nil {
		t.Fatalf("checkIntegrity returned an error: %v", err)
	}
	if !matches {
		t.Fatal("identical fingerprint files reported an integrity mismatch")
	}

	if err := os.WriteFile(localPath, []byte(`[{}]`), 0600); err != nil {
		t.Fatal(err)
	}
	matches, err = checkIntegrity(localPath)
	if err != nil {
		t.Fatalf("checkIntegrity returned an error: %v", err)
	}
	if matches {
		t.Fatal("different fingerprint files reported an integrity match")
	}
}

func TestDownloadDoesNotTruncateCacheOnNetworkFailure(t *testing.T) {
	configureFingerprintTest(t, func(response http.ResponseWriter, _ *http.Request) {
		response.WriteHeader(http.StatusBadGateway)
	})

	localPath := filepath.Join(t.TempDir(), "fingerprints.json")
	original := []byte(validFingerprints)
	if err := os.WriteFile(localPath, original, 0600); err != nil {
		t.Fatal(err)
	}

	if err := downloadFingerprints(localPath); err == nil {
		t.Fatal("downloadFingerprints unexpectedly succeeded")
	}
	contents, err := os.ReadFile(localPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != string(original) {
		t.Fatalf("cached fingerprints changed after failed download: %q", contents)
	}
}

func TestDownloadAtomicallyReplacesCache(t *testing.T) {
	configureFingerprintTest(t, func(response http.ResponseWriter, _ *http.Request) {
		_, _ = response.Write([]byte(validFingerprints))
	})

	localPath := filepath.Join(t.TempDir(), "fingerprints.json")
	if err := os.WriteFile(localPath, []byte(`[{}]`), 0600); err != nil {
		t.Fatal(err)
	}

	if err := downloadFingerprints(localPath); err != nil {
		t.Fatalf("downloadFingerprints returned an error: %v", err)
	}
	contents, err := os.ReadFile(localPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != validFingerprints {
		t.Fatalf("cached fingerprints = %q, want %q", contents, validFingerprints)
	}
}

func TestDownloadRejectsInvalidJSON(t *testing.T) {
	configureFingerprintTest(t, func(response http.ResponseWriter, _ *http.Request) {
		_, _ = response.Write([]byte(`<html>not fingerprints</html>`))
	})

	localPath := filepath.Join(t.TempDir(), "fingerprints.json")
	if err := downloadFingerprints(localPath); err == nil {
		t.Fatal("downloadFingerprints accepted invalid JSON")
	}
	if _, err := os.Stat(localPath); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("invalid download created a cache file: %v", err)
	}
}

func TestFetchFingerprintsRetries(t *testing.T) {
	attempts := 0
	configureFingerprintTest(t, func(response http.ResponseWriter, _ *http.Request) {
		attempts++
		if attempts < 3 {
			response.WriteHeader(http.StatusBadGateway)
			return
		}
		_, _ = response.Write([]byte(validFingerprints))
	})
	fingerprintAttempts = 3
	fingerprintRetryDelay = time.Millisecond

	if _, err := fetchFingerprints(); err != nil {
		t.Fatalf("fetchFingerprints did not recover: %v", err)
	}
	if attempts != 3 {
		t.Fatalf("got %d attempts, want 3", attempts)
	}
}
