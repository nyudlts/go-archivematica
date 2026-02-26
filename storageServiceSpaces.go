package go_archivematica

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (a *AMClient) GetSpaces() (*Spaces, error) {
	var spaces Spaces
	endpoint := "/api/v2/space"
	url := fmt.Sprintf("%s%s", a.StorageServiceConfig.URL, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &spaces, err
	}

	get.Header.Add("Authorization", a.GetStorageServiceAuth())

	response, err := a.Client.Do(get)
	if err != nil {
		return &spaces, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return &spaces, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &spaces, err
	}

	if err := json.Unmarshal(body, &spaces); err != nil {
		return &spaces, err
	}

	return &spaces, nil
}

// get a space
func (a *AMClient) GetSpace(spaceUUID string) (*Space, error) {
	var space Space
	endpoint := fmt.Sprintf("/api/v2/space/%s/", spaceUUID)
	url := fmt.Sprintf("%s%s", a.StorageServiceConfig.URL, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &space, err
	}

	get.Header.Add("Authorization", a.GetStorageServiceAuth())

	response, err := a.Client.Do(get)
	if err != nil {
		return &space, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return &space, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &space, err
	}

	if err := json.Unmarshal(body, &space); err != nil {
		return &space, err
	}

	return &space, nil
}

func (a *AMClient) BrowseSpace(spaceUUID string) (*Browser, error) {
	endpoint := fmt.Sprintf("/api/v2/space/%s/browse/", spaceUUID)
	url := fmt.Sprintf("%s%s", a.StorageServiceConfig.URL, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	get.Header.Add("Authorization", a.GetStorageServiceAuth())

	response, err := a.Client.Do(get)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	spaceBrowser := Browser{}
	if err := json.Unmarshal(body, &spaceBrowser); err != nil {
		return nil, err
	}

	spaceBrowser.Directories, err = convertBase64Slice(spaceBrowser.Directories)
	if err != nil {
		return nil, err
	}

	spaceBrowser.Entries, err = convertBase64Slice(spaceBrowser.Entries)
	if err != nil {
		return nil, err
	}

	return &spaceBrowser, nil
}

func (a *AMClient) CreateSpace(createSpaceRequest CreateSpaceRequest) error {
	endpoint := "/api/v2/space/"
	url := fmt.Sprintf("%s%s", a.StorageServiceConfig.URL, endpoint)
	postBody, err := json.Marshal(createSpaceRequest)
	if err != nil {
		return err
	}

	post, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}

	post.Header.Add("Authorization", a.GetStorageServiceAuth())
	post.Header.Add("Content-Type", "application/json")

	response, err := a.Client.Do(post)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Println("Space created successfully:", string(body))

	return nil
}
