package importer

import (
	"net/http"

	"github.com/gruz0/monitoring-configuration-fetcher/internal/client"
	"github.com/gruz0/monitoring-configuration-fetcher/internal/types"
)

type Importer struct {
	configurationServiceURL string
	httpClient              *http.Client
}

func NewImporter(configurationServiceURL string, httpClient *http.Client) *Importer {
	return &Importer{
		configurationServiceURL: configurationServiceURL,
		httpClient:              httpClient,
	}
}

func (i *Importer) Import() (types.Configuration, error) {
	client := client.NewClient(i.configurationServiceURL, i.httpClient)

	result, err := client.GetConfiguration()

	if err != nil {
		return types.Configuration{}, err
	}

	return result, nil
}
