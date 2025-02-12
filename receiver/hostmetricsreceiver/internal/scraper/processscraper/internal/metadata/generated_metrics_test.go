// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestDefaultMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(loadConfig(t, "default"), settings, WithStartTime(start))

	assert.Equal(t, "[WARNING] `process.memory.physical_usage` should not be enabled: The metric is deprecated and will be removed in v0.70.0. Please use `process.memory.usage` instead. See https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver#transition-to-process-memory-metric-names-aligned-with-opentelemetry-specification for more details.", observedLogs.All()[0].Message)
	assert.Equal(t, "[WARNING] `process.memory.virtual_usage` should not be enabled: The metric is deprecated and will be removed in v0.70.0. Please use `process.memory.virtual` metric instead. See  https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver#transition-to-process-memory-metric-names-aligned-with-opentelemetry-specification for more details.", observedLogs.All()[1].Message)
	assert.Equal(t, 2, observedLogs.Len())

	enabledMetrics := make(map[string]bool)

	mb.RecordProcessContextSwitchesDataPoint(ts, 1, AttributeContextSwitchType(1))

	enabledMetrics["process.cpu.time"] = true
	mb.RecordProcessCPUTimeDataPoint(ts, 1, AttributeState(1))

	mb.RecordProcessCPUUtilizationDataPoint(ts, 1, AttributeState(1))

	enabledMetrics["process.disk.io"] = true
	mb.RecordProcessDiskIoDataPoint(ts, 1, AttributeDirection(1))

	mb.RecordProcessDiskOperationsDataPoint(ts, 1, AttributeDirection(1))

	enabledMetrics["process.memory.physical_usage"] = true
	mb.RecordProcessMemoryPhysicalUsageDataPoint(ts, 1)

	mb.RecordProcessMemoryUsageDataPoint(ts, 1)

	mb.RecordProcessMemoryUtilizationDataPoint(ts, 1)

	mb.RecordProcessMemoryVirtualDataPoint(ts, 1)

	enabledMetrics["process.memory.virtual_usage"] = true
	mb.RecordProcessMemoryVirtualUsageDataPoint(ts, 1)

	mb.RecordProcessOpenFileDescriptorsDataPoint(ts, 1)

	mb.RecordProcessPagingFaultsDataPoint(ts, 1, AttributePagingFaultType(1))

	mb.RecordProcessSignalsPendingDataPoint(ts, 1)

	mb.RecordProcessThreadsDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	sm := metrics.ResourceMetrics().At(0).ScopeMetrics()
	assert.Equal(t, 1, sm.Len())
	ms := sm.At(0).Metrics()
	assert.Equal(t, len(enabledMetrics), ms.Len())
	seenMetrics := make(map[string]bool)
	for i := 0; i < ms.Len(); i++ {
		assert.True(t, enabledMetrics[ms.At(i).Name()])
		seenMetrics[ms.At(i).Name()] = true
	}
	assert.Equal(t, len(enabledMetrics), len(seenMetrics))
}

func TestAllMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(loadConfig(t, "all_metrics"), settings, WithStartTime(start))

	assert.Equal(t, "[WARNING] `process.memory.physical_usage` should not be enabled: The metric is deprecated and will be removed in v0.70.0. Please use `process.memory.usage` instead. See https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver#transition-to-process-memory-metric-names-aligned-with-opentelemetry-specification for more details.", observedLogs.All()[0].Message)
	assert.Equal(t, "[WARNING] `process.memory.virtual_usage` should not be enabled: The metric is deprecated and will be removed in v0.70.0. Please use `process.memory.virtual` metric instead. See  https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver#transition-to-process-memory-metric-names-aligned-with-opentelemetry-specification for more details.", observedLogs.All()[1].Message)
	assert.Equal(t, 2, observedLogs.Len())

	mb.RecordProcessContextSwitchesDataPoint(ts, 1, AttributeContextSwitchType(1))
	mb.RecordProcessCPUTimeDataPoint(ts, 1, AttributeState(1))
	mb.RecordProcessCPUUtilizationDataPoint(ts, 1, AttributeState(1))
	mb.RecordProcessDiskIoDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordProcessDiskOperationsDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordProcessMemoryPhysicalUsageDataPoint(ts, 1)
	mb.RecordProcessMemoryUsageDataPoint(ts, 1)
	mb.RecordProcessMemoryUtilizationDataPoint(ts, 1)
	mb.RecordProcessMemoryVirtualDataPoint(ts, 1)
	mb.RecordProcessMemoryVirtualUsageDataPoint(ts, 1)
	mb.RecordProcessOpenFileDescriptorsDataPoint(ts, 1)
	mb.RecordProcessPagingFaultsDataPoint(ts, 1, AttributePagingFaultType(1))
	mb.RecordProcessSignalsPendingDataPoint(ts, 1)
	mb.RecordProcessThreadsDataPoint(ts, 1)

	metrics := mb.Emit(WithProcessCommand("attr-val"), WithProcessCommandLine("attr-val"), WithProcessExecutableName("attr-val"), WithProcessExecutablePath("attr-val"), WithProcessOwner("attr-val"), WithProcessParentPid(1), WithProcessPid(1))

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	attrCount++
	attrVal, ok := rm.Resource().Attributes().Get("process.command")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("process.command_line")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("process.executable.name")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("process.executable.path")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("process.owner")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("process.parent_pid")
	assert.True(t, ok)
	assert.EqualValues(t, 1, attrVal.Int())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("process.pid")
	assert.True(t, ok)
	assert.EqualValues(t, 1, attrVal.Int())
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "process.context_switches":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of times the process has been context switched.", ms.At(i).Description())
			assert.Equal(t, "{count}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "involuntary", attrVal.Str())
			validatedMetrics["process.context_switches"] = struct{}{}
		case "process.cpu.time":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total CPU seconds broken down by different states.", ms.At(i).Description())
			assert.Equal(t, "s", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("state")
			assert.True(t, ok)
			assert.Equal(t, "system", attrVal.Str())
			validatedMetrics["process.cpu.time"] = struct{}{}
		case "process.cpu.utilization":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Percentage of total CPU time used by the process since last scrape, expressed as a value between 0 and 1. On the first scrape, no data point is emitted for this metric.", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("state")
			assert.True(t, ok)
			assert.Equal(t, "system", attrVal.Str())
			validatedMetrics["process.cpu.utilization"] = struct{}{}
		case "process.disk.io":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Disk bytes transferred.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "read", attrVal.Str())
			validatedMetrics["process.disk.io"] = struct{}{}
		case "process.disk.operations":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of disk operations performed by the process.", ms.At(i).Description())
			assert.Equal(t, "{operations}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "read", attrVal.Str())
			validatedMetrics["process.disk.operations"] = struct{}{}
		case "process.memory.physical_usage":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Deprecated: use `process.memory.usage` metric instead. The amount of physical memory in use.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.memory.physical_usage"] = struct{}{}
		case "process.memory.usage":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The amount of physical memory in use.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.memory.usage"] = struct{}{}
		case "process.memory.utilization":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Percentage of total physical memory that is used by the process.", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["process.memory.utilization"] = struct{}{}
		case "process.memory.virtual":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Virtual memory size.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.memory.virtual"] = struct{}{}
		case "process.memory.virtual_usage":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Deprecated: Use `process.memory.virtual` metric instead. Virtual memory size.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.memory.virtual_usage"] = struct{}{}
		case "process.open_file_descriptors":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of file descriptors in use by the process.", ms.At(i).Description())
			assert.Equal(t, "{count}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.open_file_descriptors"] = struct{}{}
		case "process.paging.faults":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of page faults the process has made. This metric is only available on Linux.", ms.At(i).Description())
			assert.Equal(t, "{faults}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "major", attrVal.Str())
			validatedMetrics["process.paging.faults"] = struct{}{}
		case "process.signals_pending":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of pending signals for the process. This metric is only available on Linux.", ms.At(i).Description())
			assert.Equal(t, "{signals}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.signals_pending"] = struct{}{}
		case "process.threads":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Process threads count.", ms.At(i).Description())
			assert.Equal(t, "{threads}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["process.threads"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(loadConfig(t, "no_metrics"), settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())

	mb.RecordProcessContextSwitchesDataPoint(ts, 1, AttributeContextSwitchType(1))
	mb.RecordProcessCPUTimeDataPoint(ts, 1, AttributeState(1))
	mb.RecordProcessCPUUtilizationDataPoint(ts, 1, AttributeState(1))
	mb.RecordProcessDiskIoDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordProcessDiskOperationsDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordProcessMemoryPhysicalUsageDataPoint(ts, 1)
	mb.RecordProcessMemoryUsageDataPoint(ts, 1)
	mb.RecordProcessMemoryUtilizationDataPoint(ts, 1)
	mb.RecordProcessMemoryVirtualDataPoint(ts, 1)
	mb.RecordProcessMemoryVirtualUsageDataPoint(ts, 1)
	mb.RecordProcessOpenFileDescriptorsDataPoint(ts, 1)
	mb.RecordProcessPagingFaultsDataPoint(ts, 1, AttributePagingFaultType(1))
	mb.RecordProcessSignalsPendingDataPoint(ts, 1)
	mb.RecordProcessThreadsDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}

func loadConfig(t *testing.T, name string) MetricsSettings {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsSettings()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
