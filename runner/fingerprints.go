package runner

import (
	"encoding/json"
	"fmt"
	"os"
)

type Fingerprint struct {
	Engine        string
	Status        string
	Fingerprint   string
	Discussion    string
	Documentation string
}

func Fingerprints() ([]Fingerprint, error) {

	var fingerprints []Fingerprint

	fingerPrintsPath, err := GetFingerprintPath()
	if err != nil {
		return nil, fmt.Errorf("Fingerprints: %v", err)
	}
	file, err := os.ReadFile(fingerPrintsPath)
	if err != nil {
		return nil, fmt.Errorf("Fingerprints: %v", err)
	}

	err = json.Unmarshal(file, &fingerprints)
	if err != nil {
		return nil, fmt.Errorf("Fingerprints: %v", err)
	}

	return fingerprints, nil
}
