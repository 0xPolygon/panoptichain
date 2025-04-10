package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/observer"
)

// desc stores the parsed prometheus.Desc.
type desc struct {
	FQName         string
	Help           string
	ConstLabels    []string
	VariableLabels []string
	ObserverType   string
	MetricType     string
}

// parseDesc parses a prometheus.Desc.
func parseDesc(line string) (*desc, error) {
	re := regexp.MustCompile(`Desc{fqName: "(.*)", help: "(.*)", constLabels: {(.*)}, variableLabels: {(.*)}}`)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 5 {
		return nil, fmt.Errorf("invalid input line")
	}

	return &desc{
		FQName:         matches[1],
		Help:           matches[2],
		ConstLabels:    parseLabels(matches[3]),
		VariableLabels: parseLabels(matches[4]),
	}, nil
}

// parseLabels parses a comma-separated string of labels.
func parseLabels(labels string) []string {
	labels = strings.TrimSpace(labels)
	if labels == "" {
		return []string{}
	}

	return strings.Split(labels, ",")
}

// printDesc prints parsed descriptions in Markdown format.
func printDesc(desc *desc) {
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

// printLabels prints labels as lists in Markdown format.
func printLabels(labels []string) {
	for _, l := range labels {
		fmt.Printf("- %s\n", l)
	}
}

func main() {
	slog.Info("Starting export of observers")

	jsonMode := flag.Bool("json", false, "Print output in JSON format")
	markdownMode := flag.Bool("md", false, "Print output in Markdown format")

	flag.Parse()

	if err := config.Init(flag.Args()); err != nil {
		slog.Error("Failed to initialize config", "error", err)
		return
	}

	eb := observer.NewEventBus()
	observers := observer.GetCompleteObserverSet()
	observers.Register(eb)

	for _, o := range observers {
		observerType := reflect.ValueOf(o).Elem().Type()

		if *markdownMode {
			fmt.Printf("\n## %s\n\n", observerType.Name())
		}

		for _, c := range o.GetCollectors() {
			descCh := make(chan *prometheus.Desc)
			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				desc := <-descCh
				d, _ := parseDesc(desc.String())
				d.ObserverType = observerType.Name()
				d.MetricType = reflect.ValueOf(c).Elem().Type().Name()
				bytes, _ := json.Marshal(d)

				if *jsonMode {
					fmt.Println(string(bytes))
				}

				if *markdownMode {
					printDesc(d)
				}

				wg.Done()
			}()

			c.Describe(descCh)
			wg.Wait()
		}
	}
}
