package opencast

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

func (op_req OpencastRequester) Validate() error {
	return nil
}

// baseRequest performs an HTTP request to the specified path with a query on the Opencast server.
// It creates an HTTP client with the configured timeout and includes basic authentication.
// Returns the HTTP response or an error if the request fails.
func (op_req OpencastRequester) baseRequestWithQuery(path string, query url.Values, request_type string) (*http.Response, error) {
	client := &http.Client{Timeout: time.Duration(op_req.Timeout * int(time.Millisecond))}

	url := op_req.URL.JoinPath(path)
	url.RawQuery = query.Encode()

	req, err := http.NewRequest(request_type, url.String(), nil)

	if err != nil {
		return nil, errors.New("")
	}

	req.SetBasicAuth(op_req.Username, op_req.Password)

	resp, err := client.Do(req)

	return resp, err
}

// baseRequest performs an HTTP request to the specified path on the Opencast server.
// It creates an HTTP client with the configured timeout and includes basic authentication.
// Returns the HTTP response or an error if the request fails.
func (op_req OpencastRequester) baseRequest(path string, request_type string) (*http.Response, error) {
	return op_req.baseRequestWithQuery(path, url.Values{}, request_type)
}

// GET performs an HTTP GET request to the specified path on the Opencast server.
// It reads and returns the response body.
// Returns the HTTP response or an error if the request or body reading fails.
func (op_req OpencastRequester) GET(path string) ([]byte, error) {
	resp, err := op_req.baseRequest(path, "GET")

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyText, nil
}

// Makes a request to the server with a query and directly parses the result in the specified JSON format
func (op_req OpencastRequester) GetJSONwithQuery(path string, query url.Values, data any) error {
	resp, err := op_req.baseRequestWithQuery(path, query, "GET")
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		return err
	}

	return nil
}

// Makes a request to the server and directly parses the result in the specified JSON format
func (op_req OpencastRequester) GetJSON(path string, data any) error {
	return op_req.GetJSONwithQuery(path, url.Values{}, data)
}

// checkConnection tests if the Opencast server is reachable.
// It sends a GET request to the base URL and returns true if successful.
// Returns true if the connection is successful, false otherwise.
func (op_req OpencastRequester) checkConnection() (bool, error) {
	client := &http.Client{Timeout: time.Duration(op_req.Timeout * int(time.Millisecond))}

	url := op_req.URL.String()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	_, err = client.Do(req)
	if err != nil {
		return false, nil
	} else {
		return true, nil
	}
}
