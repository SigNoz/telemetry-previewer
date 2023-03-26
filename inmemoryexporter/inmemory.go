package inmemoryexporter

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// InMemoryExporter is an in-memory exporter that can be used for testing.
// It implements component.TracesExporter, component.MetricsExporter and component.LogsExporter interfaces.
type InMemoryExporter struct {
	// mu protects the below fields.
	mu sync.Mutex
	// traces is a slice of pdata.Traces that were received by this exporter.
	traces []ptrace.Traces
	// metrics is a slice of pdata.Metrics that were received by this exporter.
	metrics []pmetric.Metrics
	// logs is a slice of pdata.Logs that were received by this exporter.
	logs []plog.Logs
}

func createTracesExporter(_ context.Context, _ exporter.CreateSettings, config component.Config) (exporter.Traces, error) {
	return &InMemoryExporter{}, nil
}

func createMetricsExporter(_ context.Context, _ exporter.CreateSettings, config component.Config) (exporter.Metrics, error) {
	return &InMemoryExporter{}, nil
}

func createLogsExporter(_ context.Context, _ exporter.CreateSettings, config component.Config) (exporter.Logs, error) {
	return &InMemoryExporter{}, nil
}

func (e *InMemoryExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.traces = append(e.traces, td)
	return nil
}

// ConsumeMetrics implements component.MetricsExporter.
func (e *InMemoryExporter) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.metrics = append(e.metrics, md)
	return nil
}

// ConsumeLogs implements component.LogsExporter.
func (e *InMemoryExporter) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.logs = append(e.logs, ld)
	return nil
}

// GetTraces returns a slice of pdata.Traces that were received by this exporter.
func (e *InMemoryExporter) GetTraces() []ptrace.Traces {
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.traces
}

// GetMetrics returns a slice of pdata.Metrics that were received by this exporter.
func (e *InMemoryExporter) GetMetrics() []pmetric.Metrics {
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.metrics
}

// GetLogs returns a slice of pdata.Logs that were received by this exporter.
func (e *InMemoryExporter) GetLogs() []plog.Logs {
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.logs
}

// ResetTraces removes all traces that were received by this exporter.
func (e *InMemoryExporter) ResetTraces() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.traces = nil
}

// ResetMetrics removes all metrics that were received by this exporter.
func (e *InMemoryExporter) ResetMetrics() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.metrics = nil
}

// ResetLogs removes all logs that were received by this exporter.
func (e *InMemoryExporter) ResetLogs() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.logs = nil
}

// Shutdown is a no-op.
func (e *InMemoryExporter) Shutdown(ctx context.Context) error {
	return nil
}

// Start is a no-op.
func (e *InMemoryExporter) Start(ctx context.Context, host component.Host) error {
	return nil
}

func (e *InMemoryExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}
