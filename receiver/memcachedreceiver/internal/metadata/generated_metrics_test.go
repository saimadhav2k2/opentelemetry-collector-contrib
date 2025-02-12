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

	assert.Equal(t, 0, observedLogs.Len())

	enabledMetrics := make(map[string]bool)

	enabledMetrics["memcached.bytes"] = true
	mb.RecordMemcachedBytesDataPoint(ts, 1)

	enabledMetrics["memcached.commands"] = true
	mb.RecordMemcachedCommandsDataPoint(ts, 1, AttributeCommand(1))

	enabledMetrics["memcached.connections.current"] = true
	mb.RecordMemcachedConnectionsCurrentDataPoint(ts, 1)

	enabledMetrics["memcached.connections.total"] = true
	mb.RecordMemcachedConnectionsTotalDataPoint(ts, 1)

	enabledMetrics["memcached.cpu.usage"] = true
	mb.RecordMemcachedCPUUsageDataPoint(ts, 1, AttributeState(1))

	enabledMetrics["memcached.current_items"] = true
	mb.RecordMemcachedCurrentItemsDataPoint(ts, 1)

	enabledMetrics["memcached.evictions"] = true
	mb.RecordMemcachedEvictionsDataPoint(ts, 1)

	enabledMetrics["memcached.network"] = true
	mb.RecordMemcachedNetworkDataPoint(ts, 1, AttributeDirection(1))

	enabledMetrics["memcached.operation_hit_ratio"] = true
	mb.RecordMemcachedOperationHitRatioDataPoint(ts, 1, AttributeOperation(1))

	enabledMetrics["memcached.operations"] = true
	mb.RecordMemcachedOperationsDataPoint(ts, 1, AttributeType(1), AttributeOperation(1))

	enabledMetrics["memcached.threads"] = true
	mb.RecordMemcachedThreadsDataPoint(ts, 1)

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

	assert.Equal(t, 0, observedLogs.Len())

	mb.RecordMemcachedBytesDataPoint(ts, 1)
	mb.RecordMemcachedCommandsDataPoint(ts, 1, AttributeCommand(1))
	mb.RecordMemcachedConnectionsCurrentDataPoint(ts, 1)
	mb.RecordMemcachedConnectionsTotalDataPoint(ts, 1)
	mb.RecordMemcachedCPUUsageDataPoint(ts, 1, AttributeState(1))
	mb.RecordMemcachedCurrentItemsDataPoint(ts, 1)
	mb.RecordMemcachedEvictionsDataPoint(ts, 1)
	mb.RecordMemcachedNetworkDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordMemcachedOperationHitRatioDataPoint(ts, 1, AttributeOperation(1))
	mb.RecordMemcachedOperationsDataPoint(ts, 1, AttributeType(1), AttributeOperation(1))
	mb.RecordMemcachedThreadsDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "memcached.bytes":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Current number of bytes used by this server to store items.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["memcached.bytes"] = struct{}{}
		case "memcached.commands":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Commands executed.", ms.At(i).Description())
			assert.Equal(t, "{commands}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("command")
			assert.True(t, ok)
			assert.Equal(t, "get", attrVal.Str())
			validatedMetrics["memcached.commands"] = struct{}{}
		case "memcached.connections.current":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The current number of open connections.", ms.At(i).Description())
			assert.Equal(t, "{connections}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["memcached.connections.current"] = struct{}{}
		case "memcached.connections.total":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of connections opened since the server started running.", ms.At(i).Description())
			assert.Equal(t, "{connections}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["memcached.connections.total"] = struct{}{}
		case "memcached.cpu.usage":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Accumulated user and system time.", ms.At(i).Description())
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
			validatedMetrics["memcached.cpu.usage"] = struct{}{}
		case "memcached.current_items":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of items currently stored in the cache.", ms.At(i).Description())
			assert.Equal(t, "{items}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["memcached.current_items"] = struct{}{}
		case "memcached.evictions":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Cache item evictions.", ms.At(i).Description())
			assert.Equal(t, "{evictions}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["memcached.evictions"] = struct{}{}
		case "memcached.network":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Bytes transferred over the network.", ms.At(i).Description())
			assert.Equal(t, "by", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "sent", attrVal.Str())
			validatedMetrics["memcached.network"] = struct{}{}
		case "memcached.operation_hit_ratio":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Hit ratio for operations, expressed as a percentage value between 0.0 and 100.0.", ms.At(i).Description())
			assert.Equal(t, "%", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("operation")
			assert.True(t, ok)
			assert.Equal(t, "increment", attrVal.Str())
			validatedMetrics["memcached.operation_hit_ratio"] = struct{}{}
		case "memcached.operations":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Operation counts.", ms.At(i).Description())
			assert.Equal(t, "{operations}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "hit", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("operation")
			assert.True(t, ok)
			assert.Equal(t, "increment", attrVal.Str())
			validatedMetrics["memcached.operations"] = struct{}{}
		case "memcached.threads":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of threads used by the memcached instance.", ms.At(i).Description())
			assert.Equal(t, "{threads}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["memcached.threads"] = struct{}{}
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

	mb.RecordMemcachedBytesDataPoint(ts, 1)
	mb.RecordMemcachedCommandsDataPoint(ts, 1, AttributeCommand(1))
	mb.RecordMemcachedConnectionsCurrentDataPoint(ts, 1)
	mb.RecordMemcachedConnectionsTotalDataPoint(ts, 1)
	mb.RecordMemcachedCPUUsageDataPoint(ts, 1, AttributeState(1))
	mb.RecordMemcachedCurrentItemsDataPoint(ts, 1)
	mb.RecordMemcachedEvictionsDataPoint(ts, 1)
	mb.RecordMemcachedNetworkDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordMemcachedOperationHitRatioDataPoint(ts, 1, AttributeOperation(1))
	mb.RecordMemcachedOperationsDataPoint(ts, 1, AttributeType(1), AttributeOperation(1))
	mb.RecordMemcachedThreadsDataPoint(ts, 1)

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
