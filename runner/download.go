package runner

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"

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
