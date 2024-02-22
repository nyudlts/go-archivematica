package amatica

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
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

	fmt.Println(response)

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

// Packages
func (a *AMClient) GetPackage(packageUUID uuid.UUID) (*Package, error) {
	var pack Package
	endpoint := fmt.Sprintf("/api/v2/file/%s", packageUUID)
	url := fmt.Sprintf("%s%s", a.SSHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.SSAPIKey)
	get.Header.Add("Authorization", auth)
	response, err := a.Client.Do(get)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &pack); err != nil {
		return &pack, err
	}

	return &pack, nil

}

// Get all Packages
func (a *AMClient) GetAllPackages() ([]Package, error) {
	packs := []Package{}
	packages, err := a.GetPackages(nil)
	if err != nil {
		return packs, err
	}

	for _, pack := range packages.Objects {
		packs = append(packs, pack)
	}

	complete := false
	for !complete {
		packages, err = a.GetPackages(&packages.Meta.Next)
		if err != nil {
			return packs, err
		}

		for _, pack := range packages.Objects {
			packs = append(packs, pack)
		}

		if packages.Meta.Next == "" {
			complete = true
		}
	}

	return packs, nil

}

// Get Packs by Type
func (a *AMClient) GetPackageType(packType string) ([]Package, error) {
	packs := []Package{}

	packages, err := a.GetPackages(nil)
	if err != nil {
		return packs, err
	}

	for _, pack := range packages.Objects {
		if pack.PackageType == packType {
			packs = append(packs, pack)
		}
	}

	complete := false

	for !complete {
		packages, err = a.GetPackages(&packages.Meta.Next)
		if err != nil {
			return packs, err
		}

		for _, pack := range packages.Objects {
			if pack.PackageType == packType {
				packs = append(packs, pack)
			}
		}

		if packages.Meta.Next == "" {
			complete = true
		}
	}

	return packs, nil
}

func (a *AMClient) GetPackageStatus(status string) ([]Package, error) {
	packs := []Package{}

	packages, err := a.GetPackages(nil)
	if err != nil {
		return packs, err
	}

	for _, pack := range packages.Objects {
		if pack.Status == status {
			packs = append(packs, pack)
		}
	}

	complete := false

	for !complete {
		packages, err = a.GetPackages(&packages.Meta.Next)
		if err != nil {
			return packs, err
		}

		for _, pack := range packages.Objects {
			if pack.Status == status {
				packs = append(packs, pack)
			}
		}

		if packages.Meta.Next == "" {
			complete = true
		}
	}

	return packs, nil
}

func (a *AMClient) GetPackages(params *string) (Packages, error) {
	var packages Packages
	endpoint := "/api/v2/file/"
	if params != nil {
		endpoint = *params
	}

	url := fmt.Sprintf("%s%s", a.SSHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return packages, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.SSAPIKey)
	get.Header.Add("Authorization", auth)
	response, err := a.Client.Do(get)
	if err != nil {
		return packages, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return packages, err
	}

	if err := json.Unmarshal(body, &packages); err != nil {
		return packages, err
	}

	return packages, nil
}

func (a *AMClient) RequestPackageDeletion(packageUUID uuid.UUID, deletionRequest string) (string, error) {
	//construct the endpoint
	endpoint := fmt.Sprintf("/api/v2/file/%s/delete_aip/", packageUUID)
	body := strings.NewReader(deletionRequest)

	//construct the url
	ss_url := fmt.Sprintf("%s%s", a.SSHost, endpoint)
	req, err := http.NewRequest("POST", ss_url, body)
	if err != nil {
		return "", err
	}

	//add the headers
	req.Header.Add("Authorization", fmt.Sprintf("Apikey %s:%s", a.Username, a.SSAPIKey))
	req.Header.Set("Content-Type", "application/json")

	fmt.Println(req)
	//execute the post
	resp, err := a.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil

}
