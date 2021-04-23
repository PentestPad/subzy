package src

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
)

var fingerprintPath = "https://raw.githubusercontent.com/LukaSikic/subzy/master/src/fingerprints.json"

func GetFingerprintPath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("GetFingerprintPath: %v", err)
	}
	return path.Join(home, "fingerprints.json"), nil
}

func downloadFingerprints() error {
	filePath, err := GetFingerprintPath()
	if err != nil {
		return fmt.Errorf("downloadFingerprints: %v", err)
	}
	out, err := os.Create(filePath)
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
		fmt.Errorf("downloadFingerprints: %v", err)
	}

	return nil
}

func CheckFingerprints() error {
	filePath, err := GetFingerprintPath()
	if err != nil {
		return fmt.Errorf("CheckFingerprints: %v", err)
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		downloadFingerprints()
	}

	return nil
}
