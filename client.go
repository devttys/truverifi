package truverifi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const (
	APIHost = "https://app.truverifi.com/"
)

type Truverifi struct {
	apiKey string
}

func NewClient(apiKey string) *Truverifi {
	return &Truverifi{apiKey: apiKey}
}

func (t *Truverifi) GetAccount() (*AccountInfo, error) {
	data, err := t.sendRequest("GET", "/api/account", nil)
	if err != nil {
		return nil, err
	}

	var response AccountInfo
	err = json.Unmarshal(data, &response)

	return &response, err
}

func (t *Truverifi) GetLineStatus() (*LineStatus, error) {
	data, err := t.sendRequest("GET", "/api/line", nil)

	if err != nil {
		return nil, err
	}

	var response LineStatus
	err = json.Unmarshal(data, &response)

	response.Client = t

	return &response, err
}

func (t *Truverifi) CheckService(zip string, services []string) (*CheckServiceResponse, error) {
	data, err := t.sendRequest("POST", "/api/checkService", &CheckServiceRequest{Zip: zip, Services: services})
	if err != nil {
		return nil, err
	}

	var response CheckServiceResponse
	err = json.Unmarshal(data, &response)

	return &response, err
}

func (t *Truverifi) ChangeService(zip string, services []string) (*ChangeServiceResponse, error) {
	data, err := t.sendRequest("POST", "/api/line/changeService", &ChangeServiceRequest{Zip: zip, Services: services})
	if err != nil {
		return nil, err
	}

	var response ChangeServiceResponse
	err = json.Unmarshal(data, &response)

	return &response, err
}

func (t *Truverifi) Extend() (*ExtendResponse, error) {
	data, err := t.sendRequest("POST", "/api/line/extend", nil)
	if err != nil {
		return nil, err
	}

	var response ExtendResponse
	err = json.Unmarshal(data, &response)

	return &response, err
}

func (t *Truverifi) prepareReq(method string, requestPath string, payload interface{}) (*http.Request, error) {
	uri, err := url.Parse(APIHost)
	if err != nil {
		return nil, err
	}

	uri.Path = path.Join(uri.Path, requestPath)

	body := &bytes.Buffer{}
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		body = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, uri.String(), body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-Key", t.apiKey)

	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (t *Truverifi) sendRequest(method string, requestPath string, payload interface{}) ([]byte, error) {
	req, err := t.prepareReq(method, requestPath, payload)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var errorResponse ErrorResponse
	err = json.Unmarshal(data, &errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Error != "" {
		message := errorResponse.Message
		if message == "" {
			message = fmt.Sprintf("API code: %s", errorResponse.Error)
		}
		return nil, &ApiError{Code: errorResponse.Error, Message: message}
	}

	return data, err
}
