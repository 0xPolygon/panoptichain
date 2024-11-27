package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/observer"
)

type parsedDesc struct {
	FQName         string
	Help           string
	ConstLabels    []string
	VariableLabels []string
	ObserverType   string
	MetricType     string
}

func parseDesc(line string) (*parsedDesc, error) {
	re := regexp.MustCompile(`Desc{fqName: "(.*)", help: "(.*)", constLabels: {(.*)}, variableLabels: {(.*)}}`)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 5 {
		return nil, fmt.Errorf("invalid input line")
	}

	return &parsedDesc{
		FQName:         matches[1],
		Help:           matches[2],
		ConstLabels:    parseLabels(matches[3]),
		VariableLabels: parseLabels(matches[4]),
	}, nil
}

func parseLabels(labels string) []string {
	labels = strings.TrimSpace(labels)
	if labels == "" {
		return []string{}
	}

	return strings.Split(labels, ",")
}

func printDesc(desc *parsedDesc) {
	fmt.Printf("\n### %s\n", desc.FQName)
	fmt.Printf("%s\n", desc.Help)
	fmt.Printf("\nMetric Type: %s\n", desc.MetricType)

	if len(desc.ConstLabels) > 0 {
		fmt.Printf("\nConstant Labels:\n")
		printLabels(desc.ConstLabels)
	}

	if len(desc.VariableLabels) > 0 {
		fmt.Printf("\nVariable Labels:\n")
		printLabels(desc.VariableLabels)
	}
}

func printLabels(labels []string) {
	for _, l := range labels {
		fmt.Printf("- %s\n", l)
	}
}

func main() {
	slog.Info("Starting export of observers")
	jsonMode := false
	markDownMode := true

	fauxEb := observer.NewEventBus()
	obs := observer.GetCompleteObserverSet()
	obs.Register(fauxEb)

	for _, o := range obs {
		observerType := reflect.ValueOf(o).Elem().Type()
		if markDownMode {
			fmt.Printf("\n## %s\n\n", observerType.Name())
		}

		for _, c := range o.GetCollectors() {
			descChan := make(chan *prometheus.Desc)
			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				desc := <-descChan
				d, _ := parseDesc(desc.String())
				d.ObserverType = observerType.Name()
				d.MetricType = reflect.ValueOf(c).Elem().Type().Name()
				dBytes, _ := json.Marshal(d)

				if jsonMode {
					fmt.Println(string(dBytes))
				}

				if markDownMode {
					printDesc(d)
				}

				wg.Done()
			}()

			c.Describe(descChan)
			wg.Wait()
		}
	}
}
