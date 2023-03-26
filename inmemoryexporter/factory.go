package inmemoryexporter

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

type Config struct {
}

// createDefaultConfig creates the default configuration for receiver.
func createDefaultConfig() component.Config {
	return &Config{}
}

// NewFactory creates a new OTLP receiver factory.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		"memory",
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, component.StabilityLevelStable),
		exporter.WithMetrics(createMetricsExporter, component.StabilityLevelStable),
		exporter.WithLogs(createLogsExporter, component.StabilityLevelBeta))
}
