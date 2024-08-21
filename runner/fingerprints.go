package runner

import (
	"encoding/json"
	"fmt"
	"os"
)

type Fingerprint struct {
	Engine        		string
	Status        		string
	Fingerprint  		string
	Discussion    		string
	Documentation 		string
	False_Positive 		[]string
}

func Fingerprints() ([]Fingerprint, []Fingerprint, error) {

	var allFingerprints []Fingerprint
	var validFingerprints []Fingerprint
	var skippedFingerprints []Fingerprint

	fingerPrintsPath, err := GetFingerprintPath()
	if err != nil {
		return nil, nil, fmt.Errorf("Fingerprints: %v", err)
	}
	file, err := os.ReadFile(fingerPrintsPath)
	if err != nil {
		return nil, nil, fmt.Errorf("Fingerprints: %v", err)
	}

	err = json.Unmarshal(file, &allFingerprints)
	if err != nil {
		return nil, nil, fmt.Errorf("Fingerprints: %v", err)
	}

	for _, fingerprint := range allFingerprints {
		if fingerprint.Fingerprint == "" {
			skippedFingerprints = append(skippedFingerprints, fingerprint)
		} else {
			validFingerprints = append(validFingerprints, fingerprint)
		}
	}


	return validFingerprints, skippedFingerprints, err
}
