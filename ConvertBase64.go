package go_am

import (
	"encoding/base64"
)

func decodeLocationBrowser(encodedLocationBrowser *LocationBrowser) error {
	var err error

	encodedLocationBrowser.Directories, err = convertBase64Slice(encodedLocationBrowser.Directories)
	if err != nil {
		return err
	}

	encodedLocationBrowser.Entries, err = convertBase64Slice(encodedLocationBrowser.Entries)
	if err != nil {
		return err
	}

	encodedLocationBrowser.Properties, err = convertBase64Map(encodedLocationBrowser.Properties)
	if err != nil {
		return err
	}

	return nil
}

func convertBase64Slice(encodedStrings []string) ([]string, error) {
	decodedStrings := []string{}

	for _, encStr := range encodedStrings {
		decStr, err := base64.StdEncoding.DecodeString(encStr)
		if err != nil {
			return decodedStrings, err
		}
		decodedStrings = append(decodedStrings, string(decStr))
	}

	return decodedStrings, nil
}

func convertBase64Map(encodedMap map[string]map[string]int) (map[string]map[string]int, error) {
	var decodedMap = make(map[string]map[string]int)
	for k, v := range encodedMap {
		decodedKey, err := base64.StdEncoding.DecodeString(k)
		if err != nil {
			return decodedMap, err
		}
		decodedMap[string(decodedKey)] = v
	}
	return decodedMap, nil
}
