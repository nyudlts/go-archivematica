package go_archivematica

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (a *AMClient) GetEndpoints() (*Endpoints, error) {
	endpoint := "/api/v2/"
	url := fmt.Sprintf("%s%s", a.StorageServiceConfig.URL, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	get.Header.Add("Authorization", a.GetStorageServiceAuth())

	response, err := a.Client.Do(get)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var endpoints Endpoints
	if err := json.Unmarshal(body, &endpoints); err != nil {
		fmt.Println("Error parsing response body:", err)
		return nil, err
	}
	return &endpoints, nil

}
