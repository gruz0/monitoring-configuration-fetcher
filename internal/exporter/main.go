package exporter

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gruz0/monitoring-configuration-fetcher/internal/types"
)

const (
	prefix    = "SETTINGS"
	delimiter = "__"
	lineBreak = "\n"
)

type Exporter struct {
	outputDir string
}

func NewExporter(outputDir string) *Exporter {
	return &Exporter{outputDir: outputDir}
}

func (e *Exporter) Export(domains []types.Domain) error {
	var sb strings.Builder

	for _, domain := range domains {
		sb.Reset()

		sb.WriteString(buildKey("domain") + "=" + domain.Name + lineBreak)

		for _, plugin := range domain.Plugins {
			sb.WriteString(buildKey("plugins", plugin.Namespace, plugin.Name, "ENABLE") + "=1" + lineBreak)
		}

		if err := e.save(domain.Name, sb.String()); err != nil {
			return err
		}
	}

	return nil
}

func (e *Exporter) save(domainName string, content string) error {
	path := fmt.Sprintf("%s/%s.env", e.outputDir, domainName)

	if err := ioutil.WriteFile(path, []byte(content), 0600); err != nil {
		return err
	}

	return nil
}

func buildKey(s ...interface{}) string {
	result := make([]string, 0)

	for _, item := range s {
		result = append(result, fmt.Sprintf("%s", item))
	}

	return strings.ToUpper(prefix + delimiter + strings.Join(result, delimiter))
}
