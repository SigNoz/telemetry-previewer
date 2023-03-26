package inmemoryreceiver

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/receiver"
)

type inMemoryTracesReceiver struct {
	nextConsumer consumer.Traces
}

func (r *inMemoryTracesReceiver) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (r *inMemoryTracesReceiver) Shutdown(_ context.Context) error {
	return nil
}

func (r *inMemoryTracesReceiver) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	return r.nextConsumer.ConsumeTraces(ctx, td)
}

func (r *inMemoryTracesReceiver) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func createTraces(
	_ context.Context,
	_ receiver.CreateSettings,
	_ component.Config,
	consumer consumer.Traces,
) (receiver.Traces, error) {
	return &inMemoryTracesReceiver{nextConsumer: consumer}, nil
}

type inMemoryMetricsReceiver struct {
	nextConsumer consumer.Metrics
}

func (r *inMemoryMetricsReceiver) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (r *inMemoryMetricsReceiver) Shutdown(_ context.Context) error {
	return nil
}

func (r *inMemoryMetricsReceiver) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	return r.nextConsumer.ConsumeMetrics(ctx, md)
}

func createMetrics(
	_ context.Context,
	_ receiver.CreateSettings,
	_ component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	return &inMemoryMetricsReceiver{nextConsumer: consumer}, nil
}

type inMemoryLogsReceiver struct {
	nextConsumer consumer.Logs
}

func (r *inMemoryLogsReceiver) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (r *inMemoryLogsReceiver) Shutdown(_ context.Context) error {
	return nil
}

func (r *inMemoryLogsReceiver) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	return r.nextConsumer.ConsumeLogs(ctx, ld)
}

func createLog(
	_ context.Context,
	_ receiver.CreateSettings,
	_ component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	return &inMemoryLogsReceiver{nextConsumer: consumer}, nil
}
