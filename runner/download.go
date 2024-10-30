package runner

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	fingerprintPath = "https://api.github.com/repos/EdOverflow/can-i-take-over-xyz/contents/fingerprints.json"

	subzyDir = "subzy"
)

type GitHubFileContent struct {
	Content string `json:"content"`
}

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

	// Get the file from the GitHub API instead of the raw URL
	resp, err := http.Get(fingerprintPath) // Use API URL
	if err != nil {
		return fmt.Errorf("downloadFingerprints: %v", err)
	}
	defer resp.Body.Close()

	bytes, err := decodeResponseFromApi(*resp)
	if err != nil {
		return fmt.Errorf("failed to decode base64 content: %v", err)
	}

	// Write the decoded content to the file
	_, err = out.Write(bytes)
	if err != nil {
		return fmt.Errorf("failed to write to fingerprints file: %v", err)
	}

	return nil
}

func CheckIntegrity() (bool, error) {
	// Fetch the upstream content via GitHub API
	resp, err := http.Get(fingerprintPath) // Use API URL
	if err != nil {
		return false, fmt.Errorf("downloadFingerprints: %v", err)
	}
	defer resp.Body.Close()

	// Decode the api response
	upstreamBytes, err := decodeResponseFromApi(*resp)
	if err != nil {
		return false, fmt.Errorf("failed to decode base64 content: %v", err)
	}

	// Get local file content
	fingerprintsLocal, err := GetFingerprintPath()
	if err != nil {
		return false, err
	}

	localBytes, err := os.ReadFile(fingerprintsLocal)
	if err != nil {
		return false, fmt.Errorf("failed to read local fingerprints file: %v", err)
	}

	// Calculate MD5 checksums for both upstream and local files
	h := md5.New()
	upstreamSum := h.Sum(upstreamBytes)

	h = md5.New()
	localSum := h.Sum(localBytes)

	// Compare the checksums
	return bytes.Equal(upstreamSum, localSum), nil
}

func decodeResponseFromApi(resp http.Response) ([]byte, error) {
	// Decode the JSON response from GitHub API
	var fileContent GitHubFileContent
	err := json.NewDecoder(resp.Body).Decode(&fileContent)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json response: %v", err)
	}
	upstreamBytes, err := base64.StdEncoding.DecodeString(fileContent.Content)
	if err != nil {
		return nil, err
	}
	return upstreamBytes, nil
}
