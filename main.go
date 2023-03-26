package main

import (
	"context"
	"fmt"

	"github.com/SigNoz/telemetry-previewer/inmemoryexporter"
	"github.com/mitchellh/mapstructure"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
)

func printTrace(trace ptrace.Traces) {
	for i := 0; i < trace.ResourceSpans().Len(); i++ {
		rs := trace.ResourceSpans().At(i)
		for j := 0; j < rs.ScopeSpans().Len(); j++ {
			ils := rs.ScopeSpans().At(j)
			for k := 0; k < ils.Spans().Len(); k++ {
				span := ils.Spans().At(k)
				span.Attributes().Range(func(k string, v pcommon.Value) bool {
					fmt.Println(k, v.AsString())
					return true
				})
			}
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	exporterf := inmemoryexporter.NewFactory()

	// Create a new exporter
	exporter, _ := exporterf.CreateTracesExporter(ctx, exporter.CreateSettings{}, nil)

	cfg := map[string]interface{}{
		"actions": []map[string]interface{}{{
			"key":    "attribute1",
			"value":  123,
			"action": "upsert"},
		},
	}
	processorCfg := &attributesprocessor.Config{}
	err := mapstructure.Decode(cfg, processorCfg)
	if err != nil {
		panic(err)
	}

	// Create a new attributes processor
	attrs, _ := attributesprocessor.NewFactory().CreateTracesProcessor(ctx, processor.CreateSettings{}, processorCfg, exporter)

	sampleTraces := GenerateTraces(2)

	attrs.ConsumeTraces(ctx, sampleTraces)

	t := exporter.(*inmemoryexporter.InMemoryExporter).GetTraces()
	for _, trace := range t {
		printTrace(trace)
	}
}
