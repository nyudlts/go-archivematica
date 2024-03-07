package amatica

import (
	"encoding/base64"
	"path/filepath"
	"regexp"

	"github.com/google/uuid"
)

var uuidPtn = regexp.MustCompile("[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}")

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

func ConvertUUIDToAMDirectory(uid string) (string, error) {
	uidDir := ""
	_, err := uuid.Parse(uid)
	if err != nil {
		return "", err
	}
	uidDir = filepath.Join(uidDir, uid[0:4], uid[4:8], uid[9:13], uid[14:18], uid[19:23], uid[24:28], uid[28:32], uid[32:36])
	return uidDir, nil
}

/*
func (a *AMClient) GetAIPLocation(sipUUID string) (string, error) {
	uuidPath, err := convertUUIDToAMDirectory(sipUUID)
	if err != nil {
		panic(err)
	}

	sipDir := filepath.Join(a.AIPStoreLocation, uuidPath)
	_, err = os.Stat(sipDir)
	if err != nil {
		return "", err
	}

	sipDirFiles, err := os.ReadDir(sipDir)
	if err != nil {
		return "", nil
	}

	aipDir := filepath.Join(sipDir, sipDirFiles[0].Name())
	fmt.Println("AIPDIR: ", aipDir)
	_, err = os.Stat(aipDir)
	if err != nil {
		return "", err
	}

	return aipDir, nil
}
*/
