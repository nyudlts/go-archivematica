package go_archivematica

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (a *AMClient) GetLocations() (*Locations, error) {

	endpoint := "/api/v2/location/"
	url := fmt.Sprintf("%s%s", a.StorageServiceConfig.URL, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	get.Header.Add("Authorization", a.GetStorageServiceAuth())
	get.Header.Set("Content-Type", "application/json")

	response, err := a.Client.Do(get)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get locations: status code %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var locations = Locations{}
	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, err
	}

	return &locations, nil
}

func (a *AMClient) GetLocation(locationUUID string) (*Location, error) {
	endpoint := fmt.Sprintf("/api/v2/location/%s/", locationUUID)
	url := fmt.Sprintf("%s%s", a.StorageServiceConfig.URL, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	get.Header.Add("Authorization", a.GetStorageServiceAuth())
	get.Header.Set("Content-Type", "application/json")

	response, err := a.Client.Do(get)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get location with id %s: status code %d", locationUUID, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var location = Location{}
	if err := json.Unmarshal(body, &location); err != nil {
		return nil, err
	}

	return &location, nil
}

func (a *AMClient) BrowseLocation(locationUUID string) (*Browser, error) {
	endpoint := fmt.Sprintf("/api/v2/location/%s/browse/", locationUUID)
	url := fmt.Sprintf("%s%s", a.StorageServiceConfig.URL, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	get.Header.Add("Authorization", a.GetStorageServiceAuth())
	get.Header.Set("Content-Type", "application/json")

	response, err := a.Client.Do(get)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to browse location with id %s: status code %d", locationUUID, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	locationBrowser := Browser{}

	if err := json.Unmarshal(body, &locationBrowser); err != nil {
		return nil, err
	}

	locationBrowser.Directories, err = convertBase64Slice(locationBrowser.Directories)
	if err != nil {
		return nil, err
	}

	locationBrowser.Entries, err = convertBase64Slice(locationBrowser.Entries)
	if err != nil {
		return nil, err
	}

	locationBrowser.Properties, err = convertBase64Map(locationBrowser.Properties)
	if err != nil {
		return nil, err
	}

	return &locationBrowser, nil
}
