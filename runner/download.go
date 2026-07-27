package runner

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	fingerprintPath = "https://raw.githubusercontent.com/EdOverflow/can-i-take-over-xyz/master/fingerprints.json"
	subzyDir        = "subzy"
)

func GetFingerprintPath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("GetFingerprintPath: %v", err)
	}
	dirPath := filepath.Join(home, subzyDir)
	if _, err := os.Stat(dirPath); errors.Is(err, fs.ErrNotExist) {
		if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
			return "", err
		}
	}
	return path.Join(dirPath, "fingerprints.json"), nil
}

func DownloadFingerprints() error {
	fingerprintsPath, err := GetFingerprintPath()
	if err != nil {
		return err
	}

	out, err := os.OpenFile(fingerprintsPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("downloadFingerprints: %v", err)
	}
	defer out.Close()

	resp, err := http.Get(fingerprintPath)
	if err != nil {
		return fmt.Errorf("downloadFingerprints: %v", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("downloadFingerprints: %v", err)
	}

	return nil
}

func CleanupFingerprints() error {
	fingerprintsPath, err := GetFingerprintPath()
	if err != nil {
		return err
	}

	file, err := os.ReadFile(fingerprintsPath)
	if err != nil {
		return fmt.Errorf("CleanupFingerprints: %v", err)
	}

	var fingerprints []Fingerprint
	err = json.Unmarshal(file, &fingerprints)
	if err != nil {
		return fmt.Errorf("CleanupFingerprints: %v", err)
	}

	// Filter out Cargo Collective and fix UptimeRobot entries
	var cleaned []Fingerprint
	for _, fp := range fingerprints {
		// Skip Cargo Collective entries entirely
		if strings.ToLower(fp.Service) == "cargo collective" {
			fmt.Printf("[ - ] Removing false positive fingerprint: %s\n", fp.Service)
			continue
		}

		// Fix UptimeRobot "page not found" fingerprint
		if strings.ToLower(fp.Service) == "uptimerobot" && fp.Fingerprint == "page not found" {
			fmt.Printf("[ ~ ] Fixing UptimeRobot fingerprint (removing 'page not found')\n")
			fp.Fingerprint = ""
		}

		cleaned = append(cleaned, fp)
	}

	// Write cleaned fingerprints back
	cleanedJSON, err := json.MarshalIndent(cleaned, "", "  ")
	if err != nil {
		return fmt.Errorf("CleanupFingerprints: %v", err)
	}

	err = os.WriteFile(fingerprintsPath, cleanedJSON, os.ModePerm)
	if err != nil {
		return fmt.Errorf("CleanupFingerprints: %v", err)
	}

	return nil
}

func CheckIntegrity() (bool, error) {
	resp, err := http.Get(fingerprintPath)
	if err != nil {
		return false, fmt.Errorf("downloadFingerprints: %v", err)
	}
	defer resp.Body.Close()

	outBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	h := md5.New()
	upstreamSum := h.Sum(outBytes)

	fingerprintsLocal, err := GetFingerprintPath()
	if err != nil {
		return false, err
	}

	f, err := os.Open(fingerprintsLocal)
	if err != nil {
		return false, err
	}
	defer f.Close()

	localBytes := make([]byte, len(outBytes))
	_, err = f.Read(localBytes)
	if err != nil {
		return false, err
	}

	h = md5.New()
	localSum := h.Sum(localBytes)

	return bytes.Equal(upstreamSum, localSum), nil
}
