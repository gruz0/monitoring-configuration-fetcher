package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gruz0/monitoring-configuration-fetcher/internal/types"
)

type Client struct {
	configurationServiceURL string
	httpClient              *http.Client
}

func NewClient(configurationServiceURL string, httpClient *http.Client) *Client {
	return &Client{
		configurationServiceURL: configurationServiceURL,
		httpClient:              httpClient,
	}
}

func (c *Client) GetConfiguration() (types.Configuration, error) {
	url := c.configurationServiceURL + "/configurations"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return types.Configuration{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return types.Configuration{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return types.Configuration{}, fmt.Errorf("expected 200 status code, got %d", resp.StatusCode)
	}

	type responseStruct struct {
		Configuration types.Configuration
	}

	var result responseStruct

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return types.Configuration{}, err
	}

	return result.Configuration, nil
}
