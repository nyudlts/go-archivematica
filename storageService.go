package amatica

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// get a pipeline
func (a *AMClient) GetPipeline(pipelineId string) (*Pipeline, error) {
	var pipeline Pipeline

	endpoint := fmt.Sprintf("/api/v2/pipeline/%s", pipelineId)
	url := fmt.Sprintf("%s%s", a.SSHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &pipeline, err
	}

	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.SSAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return &pipeline, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &pipeline, err
	}

	if err := json.Unmarshal(body, &pipeline); err != nil {
		return &pipeline, err
	}

	return &pipeline, nil
}

// get all pipelines
func (a *AMClient) GetPipelines() (*Pipelines, error) {
	var pipelines Pipelines
	endpoint := "/api/v2/pipeline?description__startswith=Archivematica"
	url := fmt.Sprintf("%s%s", a.SSHost, endpoint)
	log.Printf("%s", url)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &pipelines, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.SSAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return &pipelines, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &pipelines, err
	}

	if err := json.Unmarshal(body, &pipelines); err != nil {
		return &pipelines, err
	}

	return &pipelines, nil
}

// get a space
func (a *AMClient) GetSpace(spaceUUID string) (*Space, error) {
	var space Space
	endpoint := fmt.Sprintf("/api/v2/space/%s/", spaceUUID)
	url := fmt.Sprintf("%s%s", a.SSHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &space, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.SSAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return &space, err
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

// get all spaces
func (a *AMClient) GetSpaces() (*Spaces, error) {
	var spaces Spaces
	endpoint := "/api/v2/space"
	url := fmt.Sprintf("%s%s", a.SSHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &spaces, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.SSAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return &spaces, err
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

// get all locations
func (a *AMClient) GetLocations() (*Locations, error) {
	var locations Locations
	endpoint := "/api/v2/location"
	url := fmt.Sprintf("%s%s", a.SSHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &locations, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.SSAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return &locations, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &locations, err
	}

	if err := json.Unmarshal(body, &locations); err != nil {
		return &locations, err
	}

	return &locations, nil
}

// get a location
func (a *AMClient) GetLocation(locationUUID string) (*Location, error) {
	var location Location
	endpoint := fmt.Sprintf("/api/v2/location/%s/", locationUUID)
	url := fmt.Sprintf("%s%s", a.SSHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &location, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.SSAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return &location, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &location, err
	}

	if err := json.Unmarshal(body, &location); err != nil {
		return &location, err
	}

	return &location, nil
}

// get a location by name
func (a *AMClient) GetLocationByName(locationName string) (Location, error) {
	locations, err := a.GetLocations()
	if err != nil {
		return Location{}, err
	}

	for _, location := range locations.Objects {
		if location.Description == locationName {
			return location, nil
		}
	}

	return Location{}, fmt.Errorf("no location found for `%s`", locationName)
}

// browse a location
func (a *AMClient) BrowseLocation(locationUUID string) (*LocationBrowser, error) {
	var locationBrowser LocationBrowser
	endpoint := fmt.Sprintf("/api/v2/location/%s/browse", locationUUID)
	url := fmt.Sprintf("%s%s", a.SSHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &locationBrowser, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.SSAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return &locationBrowser, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &locationBrowser, err
	}

	if err := json.Unmarshal(body, &locationBrowser); err != nil {
		return &locationBrowser, err
	}

	if err := decodeLocationBrowser(&locationBrowser); err != nil {
		return &locationBrowser, err
	}
	return &locationBrowser, nil
}
