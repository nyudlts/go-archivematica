package go_archivematica

import (
	"encoding/base64"
	"regexp"
)

const version string = "0.2.0"

var uuidPtn = regexp.MustCompile("[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}")

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
