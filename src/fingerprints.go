package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Fingerprint struct {
	Engine         string
	Status         string
	Fingerprint    string
	Discussion     string
	Documentation  string
	False_Positive string
}

func Fingerprints() ([]Fingerprint, error) {

	var fingerprints []Fingerprint

	fingerPrintsPath, err := GetFingerprintPath()
	if err != nil {
		return nil, fmt.Errorf("Fingerprints: %v", err)
	}
	file, err := ioutil.ReadFile(fingerPrintsPath)
	if err != nil {
		return nil, fmt.Errorf("Fingerprints: %v", err)
	}

	err = json.Unmarshal(file, &fingerprints)
	if err != nil {
		return nil, fmt.Errorf("Fingerprints: %v", err)
	}

	return fingerprints, nil
}
