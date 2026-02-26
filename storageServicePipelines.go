package go_archivematica

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

/* Pipelines */

// get all pipelines
func (a *AMClient) GetPipelines() (*Pipelines, error) {
	var pipelines Pipelines
	endpoint := "/api/v2/pipeline"
	url := fmt.Sprintf("%s%s", a.StorageServiceConfig.URL, endpoint)

	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &pipelines, err
	}

	get.Header.Add("Authorization", a.GetStorageServiceAuth())
	get.Header.Set("Content-Type", "application/json")

	response, err := a.Client.Do(get)
	if err != nil {
		return &pipelines, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &pipelines, err
	}
	defer response.Body.Close()

	if err := json.Unmarshal(body, &pipelines); err != nil {
		return &pipelines, err
	}

	return &pipelines, nil
}

// get a pipeline
func (a *AMClient) GetPipeline(pipelineId string) (*Pipeline, error) {
	var pipeline Pipeline

	endpoint := fmt.Sprintf("/api/v2/pipeline/%s", pipelineId)
	url := fmt.Sprintf("%s%s", a.StorageServiceConfig.URL, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &pipeline, err
	}

	get.Header.Add("Authorization", a.GetStorageServiceAuth())
	get.Header.Set("Content-Type", "application/json")

	response, err := a.Client.Do(get)
	if err != nil {
		return &pipeline, err
	}

	if response.StatusCode != http.StatusOK {
		return &pipeline, fmt.Errorf("failed to get pipeline with id %s: status code %d", pipelineId, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &pipeline, err
	}
	defer response.Body.Close()

	if err := json.Unmarshal(body, &pipeline); err != nil {
		return &pipeline, err
	}

	return &pipeline, nil
}

func (a *AMClient) CreatePipeline(createPipelineRequest CreatePipelineRequest) error {

	endpoint := "/api/v2/pipeline/"
	url := fmt.Sprintf("%s%s", a.StorageServiceConfig.URL, endpoint)

	createPipelineRequest.UUID = uuid.New().String()
	createPipelineRequest.APIKey = a.StorageServiceConfig.Key
	createPipelineRequest.APIUsername = a.StorageServiceConfig.Username
	createPipelineRequest.CreateDefaultLocations = true

	bodyBytes, err := json.Marshal(createPipelineRequest)
	if err != nil {
		return err
	}

	fmt.Println("BODY:", string(bodyBytes))

	request, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		panic(err)
	}

	request.Header.Add("Authorization", a.GetStorageServiceAuth())
	request.Header.Set("Content-Type", "application/json")

	response, err := a.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fmt.Println("STATUS:", response.Status)
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println("RESPONSE:", string(responseBody))
	return nil
}

/*

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
func (a *AMClient) GetPackagesByType(packType string) ([]Package, error) {
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

func (a *AMClient) GetPackageByStatus(status string) ([]Package, error) {
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

func (a *AMClient) GetPackagesByPath(path string) ([]Package, error) {
	packs := []Package{}
	packages, err := a.GetPackages(nil)
	if err != nil {
		return packs, err
	}

	for _, pack := range packages.Objects {
		if strings.Contains(pack.CurrentPath, path) {
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
			if strings.Contains(pack.CurrentPath, path) {
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

func (a *AMClient) FilterPackages(packageType *string, packageStatus *string, packagePath *string) ([]Package, error) {
	filteredPacks := []Package{}
	allPacks, err := a.GetAllPackages()
	if err != nil {
		return filteredPacks, err
	}

	//filter by type
	if packageType == nil {
		filteredPacks = allPacks
	} else {
		for _, pack := range allPacks {
			if pack.PackageType == *packageType {
				filteredPacks = append(filteredPacks, pack)
			}
		}
	}

	//filter by status
	filteredPacks2 := []Package{}
	if packageStatus != nil {
		for _, pack := range filteredPacks {
			if pack.Status == *packageStatus {
				filteredPacks2 = append(filteredPacks2, pack)
			}
		}
	} else {
		filteredPacks2 = filteredPacks
	}

	filteredPacks3 := []Package{}
	//filter by path
	if packagePath != nil {
		for _, pack := range filteredPacks2 {
			if strings.Contains(pack.CurrentPath, *packagePath) {
				filteredPacks3 = append(filteredPacks3, pack)
			}
		}
	} else {
		filteredPacks3 = filteredPacks2
	}

	return filteredPacks3, nil
}
*/
