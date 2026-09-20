package runner

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	fingerprintURL = "https://raw.githubusercontent.com/EdOverflow/can-i-take-over-xyz/master/fingerprints.json"
	subzyDir       = "subzy"

	fingerprintHTTPClient = &http.Client{Timeout: 15 * time.Second}
	fingerprintAttempts   = 3
	fingerprintRetryDelay = 500 * time.Millisecond
)

func GetFingerprintPath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("GetFingerprintPath: %v", err)
	}
	dirPath := filepath.Join(home, subzyDir)
	if _, err := os.Stat(dirPath); errors.Is(err, fs.ErrNotExist) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return "", err
		}
	}
	return filepath.Join(dirPath, "fingerprints.json"), nil
}

// DownloadFingerprints fetches and validates the upstream fingerprints before
// atomically replacing the local copy. A failed request therefore cannot
// truncate a previously working cache.
func DownloadFingerprints() error {
	fingerprintsPath, err := GetFingerprintPath()
	if err != nil {
		return err
	}

	return downloadFingerprints(fingerprintsPath)
}

func downloadFingerprints(fingerprintsPath string) error {
	contents, err := fetchFingerprints()
	if err != nil {
		return fmt.Errorf("downloadFingerprints: %w", err)
	}

	temporary, err := os.CreateTemp(filepath.Dir(fingerprintsPath), ".fingerprints-*.tmp")
	if err != nil {
		return fmt.Errorf("downloadFingerprints: %w", err)
	}
	temporaryPath := temporary.Name()
	defer os.Remove(temporaryPath)

	if _, err := temporary.Write(contents); err != nil {
		temporary.Close()
		return fmt.Errorf("downloadFingerprints: %w", err)
	}
	if err := temporary.Sync(); err != nil {
		temporary.Close()
		return fmt.Errorf("downloadFingerprints: %w", err)
	}
	if err := temporary.Close(); err != nil {
		return fmt.Errorf("downloadFingerprints: %w", err)
	}
	if err := os.Rename(temporaryPath, fingerprintsPath); err != nil {
		return fmt.Errorf("downloadFingerprints: %w", err)
	}

	return nil
}

func CheckIntegrity() (bool, error) {
	fingerprintsPath, err := GetFingerprintPath()
	if err != nil {
		return false, err
	}

	return checkIntegrity(fingerprintsPath)
}

func checkIntegrity(fingerprintsPath string) (bool, error) {
	upstreamBytes, err := fetchFingerprints()
	if err != nil {
		return false, fmt.Errorf("checkIntegrity: %w", err)
	}

	localBytes, err := os.ReadFile(fingerprintsPath)
	if err != nil {
		return false, fmt.Errorf("checkIntegrity: %w", err)
	}

	upstreamSum := sha256.Sum256(upstreamBytes)
	localSum := sha256.Sum256(localBytes)

	return upstreamSum == localSum, nil
}

func fetchFingerprints() ([]byte, error) {
	var lastErr error

	for attempt := 1; attempt <= fingerprintAttempts; attempt++ {
		contents, err := fetchFingerprintsOnce()
		if err == nil {
			return contents, nil
		}
		lastErr = err

		if attempt < fingerprintAttempts {
			time.Sleep(time.Duration(attempt) * fingerprintRetryDelay)
		}
	}

	return nil, fmt.Errorf("fetch fingerprints after %d attempts: %w", fingerprintAttempts, lastErr)
}

func fetchFingerprintsOnce() ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, fingerprintURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "subzy")

	response, err := fingerprintHTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		_, _ = io.Copy(io.Discard, io.LimitReader(response.Body, 4096))
		return nil, fmt.Errorf("upstream returned %s", response.Status)
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var fingerprints []Fingerprint
	if err := json.Unmarshal(contents, &fingerprints); err != nil {
		return nil, fmt.Errorf("invalid upstream fingerprints: %w", err)
	}
	if len(fingerprints) == 0 {
		return nil, errors.New("upstream returned an empty fingerprint list")
	}

	return contents, nil
}
