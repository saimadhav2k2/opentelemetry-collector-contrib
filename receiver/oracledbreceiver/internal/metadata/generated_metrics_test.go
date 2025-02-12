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

	enabledMetrics["oracledb.cpu_time"] = true
	mb.RecordOracledbCPUTimeDataPoint(ts, 1)

	enabledMetrics["oracledb.dml_locks.limit"] = true
	mb.RecordOracledbDmlLocksLimitDataPoint(ts, "1")

	enabledMetrics["oracledb.dml_locks.usage"] = true
	mb.RecordOracledbDmlLocksUsageDataPoint(ts, "1")

	enabledMetrics["oracledb.enqueue_deadlocks"] = true
	mb.RecordOracledbEnqueueDeadlocksDataPoint(ts, "1")

	enabledMetrics["oracledb.enqueue_locks.limit"] = true
	mb.RecordOracledbEnqueueLocksLimitDataPoint(ts, "1")

	enabledMetrics["oracledb.enqueue_locks.usage"] = true
	mb.RecordOracledbEnqueueLocksUsageDataPoint(ts, "1")

	enabledMetrics["oracledb.enqueue_resources.limit"] = true
	mb.RecordOracledbEnqueueResourcesLimitDataPoint(ts, "1")

	enabledMetrics["oracledb.enqueue_resources.usage"] = true
	mb.RecordOracledbEnqueueResourcesUsageDataPoint(ts, "1")

	enabledMetrics["oracledb.exchange_deadlocks"] = true
	mb.RecordOracledbExchangeDeadlocksDataPoint(ts, "1")

	enabledMetrics["oracledb.executions"] = true
	mb.RecordOracledbExecutionsDataPoint(ts, "1")

	enabledMetrics["oracledb.hard_parses"] = true
	mb.RecordOracledbHardParsesDataPoint(ts, "1")

	enabledMetrics["oracledb.logical_reads"] = true
	mb.RecordOracledbLogicalReadsDataPoint(ts, "1")

	enabledMetrics["oracledb.parse_calls"] = true
	mb.RecordOracledbParseCallsDataPoint(ts, "1")

	enabledMetrics["oracledb.pga_memory"] = true
	mb.RecordOracledbPgaMemoryDataPoint(ts, "1")

	enabledMetrics["oracledb.physical_reads"] = true
	mb.RecordOracledbPhysicalReadsDataPoint(ts, "1")

	enabledMetrics["oracledb.processes.limit"] = true
	mb.RecordOracledbProcessesLimitDataPoint(ts, "1")

	enabledMetrics["oracledb.processes.usage"] = true
	mb.RecordOracledbProcessesUsageDataPoint(ts, "1")

	enabledMetrics["oracledb.sessions.limit"] = true
	mb.RecordOracledbSessionsLimitDataPoint(ts, "1")

	enabledMetrics["oracledb.sessions.usage"] = true
	mb.RecordOracledbSessionsUsageDataPoint(ts, "1", "attr-val", "attr-val")

	enabledMetrics["oracledb.tablespace_size.limit"] = true
	mb.RecordOracledbTablespaceSizeLimitDataPoint(ts, "1", "attr-val")

	enabledMetrics["oracledb.tablespace_size.usage"] = true
	mb.RecordOracledbTablespaceSizeUsageDataPoint(ts, "1", "attr-val")

	enabledMetrics["oracledb.transactions.limit"] = true
	mb.RecordOracledbTransactionsLimitDataPoint(ts, "1")

	enabledMetrics["oracledb.transactions.usage"] = true
	mb.RecordOracledbTransactionsUsageDataPoint(ts, "1")

	enabledMetrics["oracledb.user_commits"] = true
	mb.RecordOracledbUserCommitsDataPoint(ts, "1")

	enabledMetrics["oracledb.user_rollbacks"] = true
	mb.RecordOracledbUserRollbacksDataPoint(ts, "1")

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

	mb.RecordOracledbCPUTimeDataPoint(ts, 1)
	mb.RecordOracledbDmlLocksLimitDataPoint(ts, "1")
	mb.RecordOracledbDmlLocksUsageDataPoint(ts, "1")
	mb.RecordOracledbEnqueueDeadlocksDataPoint(ts, "1")
	mb.RecordOracledbEnqueueLocksLimitDataPoint(ts, "1")
	mb.RecordOracledbEnqueueLocksUsageDataPoint(ts, "1")
	mb.RecordOracledbEnqueueResourcesLimitDataPoint(ts, "1")
	mb.RecordOracledbEnqueueResourcesUsageDataPoint(ts, "1")
	mb.RecordOracledbExchangeDeadlocksDataPoint(ts, "1")
	mb.RecordOracledbExecutionsDataPoint(ts, "1")
	mb.RecordOracledbHardParsesDataPoint(ts, "1")
	mb.RecordOracledbLogicalReadsDataPoint(ts, "1")
	mb.RecordOracledbParseCallsDataPoint(ts, "1")
	mb.RecordOracledbPgaMemoryDataPoint(ts, "1")
	mb.RecordOracledbPhysicalReadsDataPoint(ts, "1")
	mb.RecordOracledbProcessesLimitDataPoint(ts, "1")
	mb.RecordOracledbProcessesUsageDataPoint(ts, "1")
	mb.RecordOracledbSessionsLimitDataPoint(ts, "1")
	mb.RecordOracledbSessionsUsageDataPoint(ts, "1", "attr-val", "attr-val")
	mb.RecordOracledbTablespaceSizeLimitDataPoint(ts, "1", "attr-val")
	mb.RecordOracledbTablespaceSizeUsageDataPoint(ts, "1", "attr-val")
	mb.RecordOracledbTransactionsLimitDataPoint(ts, "1")
	mb.RecordOracledbTransactionsUsageDataPoint(ts, "1")
	mb.RecordOracledbUserCommitsDataPoint(ts, "1")
	mb.RecordOracledbUserRollbacksDataPoint(ts, "1")

	metrics := mb.Emit(WithOracledbInstanceName("attr-val"))

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	attrCount++
	attrVal, ok := rm.Resource().Attributes().Get("oracledb.instance.name")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "oracledb.cpu_time":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Cumulative CPU time, in seconds", ms.At(i).Description())
			assert.Equal(t, "s", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["oracledb.cpu_time"] = struct{}{}
		case "oracledb.dml_locks.limit":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Maximum limit of active DML (Data Manipulation Language) locks.", ms.At(i).Description())
			assert.Equal(t, "{locks}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.dml_locks.limit"] = struct{}{}
		case "oracledb.dml_locks.usage":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Current count of active DML (Data Manipulation Language) locks.", ms.At(i).Description())
			assert.Equal(t, "{locks}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.dml_locks.usage"] = struct{}{}
		case "oracledb.enqueue_deadlocks":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of deadlocks between table or row locks in different sessions.", ms.At(i).Description())
			assert.Equal(t, "{deadlocks}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.enqueue_deadlocks"] = struct{}{}
		case "oracledb.enqueue_locks.limit":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Maximum limit of active enqueue locks.", ms.At(i).Description())
			assert.Equal(t, "{locks}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.enqueue_locks.limit"] = struct{}{}
		case "oracledb.enqueue_locks.usage":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Current count of active enqueue locks.", ms.At(i).Description())
			assert.Equal(t, "{locks}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.enqueue_locks.usage"] = struct{}{}
		case "oracledb.enqueue_resources.limit":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Maximum limit of active enqueue resources.", ms.At(i).Description())
			assert.Equal(t, "{resources}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.enqueue_resources.limit"] = struct{}{}
		case "oracledb.enqueue_resources.usage":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Current count of active enqueue resources.", ms.At(i).Description())
			assert.Equal(t, "{resources}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.enqueue_resources.usage"] = struct{}{}
		case "oracledb.exchange_deadlocks":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of times that a process detected a potential deadlock when exchanging two buffers and raised an internal, restartable error. Index scans are the only operations that perform exchanges.", ms.At(i).Description())
			assert.Equal(t, "{deadlocks}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.exchange_deadlocks"] = struct{}{}
		case "oracledb.executions":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of calls (user and recursive) that executed SQL statements", ms.At(i).Description())
			assert.Equal(t, "{executions}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.executions"] = struct{}{}
		case "oracledb.hard_parses":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of hard parses", ms.At(i).Description())
			assert.Equal(t, "{parses}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.hard_parses"] = struct{}{}
		case "oracledb.logical_reads":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of logical reads", ms.At(i).Description())
			assert.Equal(t, "{reads}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.logical_reads"] = struct{}{}
		case "oracledb.parse_calls":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of parse calls.", ms.At(i).Description())
			assert.Equal(t, "{parses}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.parse_calls"] = struct{}{}
		case "oracledb.pga_memory":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Session PGA (Program Global Area) memory", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.pga_memory"] = struct{}{}
		case "oracledb.physical_reads":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of physical reads", ms.At(i).Description())
			assert.Equal(t, "{reads}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.physical_reads"] = struct{}{}
		case "oracledb.processes.limit":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Maximum limit of active processes.", ms.At(i).Description())
			assert.Equal(t, "{processes}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.processes.limit"] = struct{}{}
		case "oracledb.processes.usage":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Current count of active processes.", ms.At(i).Description())
			assert.Equal(t, "{processes}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.processes.usage"] = struct{}{}
		case "oracledb.sessions.limit":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Maximum limit of active sessions.", ms.At(i).Description())
			assert.Equal(t, "{sessions}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.sessions.limit"] = struct{}{}
		case "oracledb.sessions.usage":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Count of active sessions.", ms.At(i).Description())
			assert.Equal(t, "{sessions}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("session_type")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("session_status")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["oracledb.sessions.usage"] = struct{}{}
		case "oracledb.tablespace_size.limit":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Maximum size of tablespace in bytes.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("tablespace_name")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["oracledb.tablespace_size.limit"] = struct{}{}
		case "oracledb.tablespace_size.usage":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Used tablespace in bytes.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("tablespace_name")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["oracledb.tablespace_size.usage"] = struct{}{}
		case "oracledb.transactions.limit":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Maximum limit of active transactions.", ms.At(i).Description())
			assert.Equal(t, "{transactions}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.transactions.limit"] = struct{}{}
		case "oracledb.transactions.usage":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Current count of active transactions.", ms.At(i).Description())
			assert.Equal(t, "{transactions}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.transactions.usage"] = struct{}{}
		case "oracledb.user_commits":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of user commits. When a user commits a transaction, the redo generated that reflects the changes made to database blocks must be written to disk. Commits often represent the closest thing to a user transaction rate.", ms.At(i).Description())
			assert.Equal(t, "{commits}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.user_commits"] = struct{}{}
		case "oracledb.user_rollbacks":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of times users manually issue the ROLLBACK statement or an error occurs during a user's transactions", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["oracledb.user_rollbacks"] = struct{}{}
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

	mb.RecordOracledbCPUTimeDataPoint(ts, 1)
	mb.RecordOracledbDmlLocksLimitDataPoint(ts, "1")
	mb.RecordOracledbDmlLocksUsageDataPoint(ts, "1")
	mb.RecordOracledbEnqueueDeadlocksDataPoint(ts, "1")
	mb.RecordOracledbEnqueueLocksLimitDataPoint(ts, "1")
	mb.RecordOracledbEnqueueLocksUsageDataPoint(ts, "1")
	mb.RecordOracledbEnqueueResourcesLimitDataPoint(ts, "1")
	mb.RecordOracledbEnqueueResourcesUsageDataPoint(ts, "1")
	mb.RecordOracledbExchangeDeadlocksDataPoint(ts, "1")
	mb.RecordOracledbExecutionsDataPoint(ts, "1")
	mb.RecordOracledbHardParsesDataPoint(ts, "1")
	mb.RecordOracledbLogicalReadsDataPoint(ts, "1")
	mb.RecordOracledbParseCallsDataPoint(ts, "1")
	mb.RecordOracledbPgaMemoryDataPoint(ts, "1")
	mb.RecordOracledbPhysicalReadsDataPoint(ts, "1")
	mb.RecordOracledbProcessesLimitDataPoint(ts, "1")
	mb.RecordOracledbProcessesUsageDataPoint(ts, "1")
	mb.RecordOracledbSessionsLimitDataPoint(ts, "1")
	mb.RecordOracledbSessionsUsageDataPoint(ts, "1", "attr-val", "attr-val")
	mb.RecordOracledbTablespaceSizeLimitDataPoint(ts, "1", "attr-val")
	mb.RecordOracledbTablespaceSizeUsageDataPoint(ts, "1", "attr-val")
	mb.RecordOracledbTransactionsLimitDataPoint(ts, "1")
	mb.RecordOracledbTransactionsUsageDataPoint(ts, "1")
	mb.RecordOracledbUserCommitsDataPoint(ts, "1")
	mb.RecordOracledbUserRollbacksDataPoint(ts, "1")

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
