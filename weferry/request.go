package weferry

import (
	"bytes"
	"errors"
	"net/http"
)

func sendHTTPRequest(method, url string, requestData []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, errors.New("request failed")
	}

	return resp, nil
}
