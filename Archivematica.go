package go_am

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Ingest Functions
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

func (a *AMClient) GetCompletedIngests() (UUIDList, error) {
	completedTransfers := UUIDList{}
	endpoint := "/api/ingest/completed"
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

// Transfer Functions
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
