package exampleapiclient

import (
	"io/ioutil"
	"net/http"
	"time"
)

type APIClient struct {
	URL        string
	httpClient *http.Client
}

func NewAPIClient(url string, timeout time.Duration) APIClient {
	return APIClient{
		URL: url,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (apiClient APIClient) ToUpper(input string) (string, error) {

	req, err := http.NewRequest("GET", apiClient.URL, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Set("input", input)
	req.URL.RawQuery = q.Encode()

	resp, err := apiClient.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil

}
