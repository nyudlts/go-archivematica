package go_am

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

//* Ingest Functions *//

// Get The status of an ingest
func (a *AMClient) GetIngestStatus(id string) (IngestStatus, error) {
	ingestStatus := IngestStatus{}
	endpoint := fmt.Sprintf("/api/ingest/status/%s", id)
	url := fmt.Sprintf("%s%s", a.AMHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ingestStatus, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.AMAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return ingestStatus, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ingestStatus, err
	}

	if err := json.Unmarshal(body, &ingestStatus); err != nil {
		return ingestStatus, err
	}

	return ingestStatus, nil
}

// Get UUID List of completed ingests
func (a *AMClient) GetCompletedIngests() (UUIDList, error) {
	completedIngests := UUIDList{}
	endpoint := "/api/ingest/completed"
	url := fmt.Sprintf("%s%s", a.AMHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return completedIngests, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.AMAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return completedIngests, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return completedIngests, err
	}

	if err := json.Unmarshal(body, &completedIngests); err != nil {
		return completedIngests, err
	}

	return completedIngests, nil
}

// Get Waiting ingests
func (a *AMClient) GetWaitingIngests() (WaitingIngests, error) {
	waitingIngests := WaitingIngests{}
	endpoint := "/api/ingest/waiting"
	url := fmt.Sprintf("%s%s", a.AMHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return waitingIngests, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.AMAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return waitingIngests, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return waitingIngests, err
	}

	if err := json.Unmarshal(body, &waitingIngests); err != nil {
		return waitingIngests, err
	}

	return waitingIngests, nil

}

//* Transfer Functions *//

// Start a new transfer
func (a *AMClient) StartTransfer(location uuid.UUID, path string) (StartTransferResponse, error) {

	//encode the uuid and the path to base64
	pathValue := fmt.Sprintf("%s:%s", location.String(), path)
	encodedValue := base64.RawStdEncoding.EncodeToString([]byte(pathValue))

	//encode the params
	params := url.Values{}
	params.Add("name", path)
	params.Add("type", `standard`)
	params.Add("paths[]", fmt.Sprintf(`%s==`, encodedValue))
	body := strings.NewReader(params.Encode())

	//construct the uri and the post
	endpoint := "/api/transfer/start_transfer/"
	am_url := fmt.Sprintf("%s%s", a.AMHost, endpoint)
	req, err := http.NewRequest("POST", am_url, body)
	if err != nil {
		return StartTransferResponse{}, err
	}

	//set the header parameters
	req.Header.Set("Authorization", fmt.Sprintf("ApiKey archivematica:%s", a.AMAPIKey))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//execute the post
	resp, err := a.Client.Do(req)
	if err != nil {
		return StartTransferResponse{}, err
	}
	defer resp.Body.Close()

	//print the response body
	b, _ := io.ReadAll(resp.Body)
	startTransferResponse := StartTransferResponse{}
	if err := json.Unmarshal(b, &startTransferResponse); err != nil {
		return StartTransferResponse{}, err
	}

	return startTransferResponse, nil
}

// Approve a transfer
func (a *AMClient) ApproveTransfer(directory string, xtype string) error {
	params := url.Values{}
	params.Add("directory", directory)
	params.Add("type", xtype)
	params.Add("rows_id[]", ``)
	body := strings.NewReader(params.Encode())

	endpoint := "/api/transfer/approve/"
	am_url := fmt.Sprintf("%s%s", a.AMHost, endpoint)
	req, err := http.NewRequest("POST", am_url, body)
	if err != nil {
		return err
	}

	//set the header parameters
	req.Header.Set("Authorization", fmt.Sprintf("ApiKey archivematica:%s", a.AMAPIKey))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// print the response body
	//b, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(b))

	return nil
}

// Get the status of a transfer
func (a *AMClient) GetTransferStatus(id string) (TransferStatus, error) {
	transferStatus := TransferStatus{}
	endpoint := fmt.Sprintf("/api/transfer/status/%s", id)
	url := fmt.Sprintf("%s%s", a.AMHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return transferStatus, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.AMAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return transferStatus, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return transferStatus, err
	}

	if err := json.Unmarshal(body, &transferStatus); err != nil {
		return transferStatus, err
	}

	return transferStatus, nil
}

// Get completed Transfers
func (a *AMClient) GetCompletedTransfers() (UUIDList, error) {
	completedTransfers := UUIDList{}
	endpoint := "/api/transfer/completed"
	url := fmt.Sprintf("%s%s", a.AMHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return completedTransfers, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.AMAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return completedTransfers, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return completedTransfers, err
	}

	if err := json.Unmarshal(body, &completedTransfers); err != nil {
		return completedTransfers, err
	}

	return completedTransfers, nil
}

// Get Completed transfer map indexed by uuid
func (a *AMClient) GetCompletedTransfersMap(completedTransfers UUIDList) (map[string]TransferStatus, error) {
	transferMap := map[string]TransferStatus{}
	for _, completedTransfer := range completedTransfers.Results {
		xfrmd, err := a.GetTransferStatus(completedTransfer)
		if err != nil {
			return transferMap, err
		}
		transferMap[completedTransfer] = xfrmd
	}
	return transferMap, nil
}

// Get Transfers waiting approval
func (a *AMClient) GetUnapprovedTransfers() (UnapprovedTransfers, error) {
	unapprovedTransfers := UnapprovedTransfers{}
	endpoint := "/api/transfer/unapproved"
	url := fmt.Sprintf("%s%s", a.AMHost, endpoint)
	get, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return unapprovedTransfers, err
	}
	auth := fmt.Sprintf("Apikey %s:%s", a.Username, a.AMAPIKey)
	get.Header.Add("Authorization", auth)

	response, err := a.Client.Do(get)
	if err != nil {
		return unapprovedTransfers, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return unapprovedTransfers, err
	}

	if err := json.Unmarshal(body, &unapprovedTransfers); err != nil {
		return unapprovedTransfers, err
	}

	return unapprovedTransfers, nil
}

// get unapproved Transfers as map
func (a *AMClient) GetUnapprovedTransfersMap(unapprovedTransfers UnapprovedTransfers) (map[string]TransferStatus, error) {
	transferMap := map[string]TransferStatus{}
	for _, unapprovedTransfer := range unapprovedTransfers.Results {
		ts, err := a.GetTransferStatus(unapprovedTransfer.UUID.String())
		if err != nil {
			return transferMap, err
		}
		transferMap[unapprovedTransfer.UUID.String()] = ts
	}
	return transferMap, nil
}

// Delete a transfer
func (a *AMClient) DeleteTransfer(id uuid.UUID) error {

	endpoint := fmt.Sprintf("/api/transfer/unapproved/%s", id)
	reqUrl := fmt.Sprintf("%s%s", a.AMHost, endpoint)

	req, err := http.NewRequest("DELETE", reqUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ApiKey archivematica:%s", a.AMAPIKey))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

// Monitor
func (a *AMClient) Monitor() {
	fmt.Println("++ Archivematica Monitor ++")
	unapprovedTransfers, err := a.GetUnapprovedTransfers()
	if err != nil {
		panic(err)
	}
	fmt.Println("\n== Unapproved Transfers ==")
	for _, unapprovedTransfer := range unapprovedTransfers.Results {
		xfrmd, err := a.GetTransferStatus(unapprovedTransfer.UUID.String())
		if err != nil {
			panic(err)
		}
		fmt.Println(xfrmd)
	}

	completedTransfers, err := a.GetCompletedTransfers()
	if err != nil {
		panic(err)
	}

	fmt.Println("\n== Completed Transfers ==")
	for _, completedTransfer := range completedTransfers.Results {
		xfrmd, err := a.GetTransferStatus(completedTransfer)
		if err != nil {
			panic(err)
		}
		fmt.Println(xfrmd)
	}

	completedIngests, err := a.GetCompletedIngests()
	if err != nil {
		panic(err)
	}
	fmt.Println("\n== Completed Ingests ==")
	for _, completedIngest := range completedIngests.Results {
		igmd, err := a.GetIngestStatus(completedIngest)
		if err != nil {
			panic(err)
		}
		fmt.Println(igmd)
	}
	fmt.Println()
}
