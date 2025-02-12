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

	enabledMetrics["postgresql.backends"] = true
	mb.RecordPostgresqlBackendsDataPoint(ts, 1, "attr-val")

	enabledMetrics["postgresql.bgwriter.buffers.allocated"] = true
	mb.RecordPostgresqlBgwriterBuffersAllocatedDataPoint(ts, 1)

	enabledMetrics["postgresql.bgwriter.buffers.writes"] = true
	mb.RecordPostgresqlBgwriterBuffersWritesDataPoint(ts, 1, AttributeBgBufferSource(1))

	enabledMetrics["postgresql.bgwriter.checkpoint.count"] = true
	mb.RecordPostgresqlBgwriterCheckpointCountDataPoint(ts, 1, AttributeBgCheckpointType(1))

	enabledMetrics["postgresql.bgwriter.duration"] = true
	mb.RecordPostgresqlBgwriterDurationDataPoint(ts, 1, AttributeBgDurationType(1))

	enabledMetrics["postgresql.bgwriter.maxwritten"] = true
	mb.RecordPostgresqlBgwriterMaxwrittenDataPoint(ts, 1)

	enabledMetrics["postgresql.blocks_read"] = true
	mb.RecordPostgresqlBlocksReadDataPoint(ts, 1, "attr-val", "attr-val", AttributeSource(1))

	enabledMetrics["postgresql.commits"] = true
	mb.RecordPostgresqlCommitsDataPoint(ts, 1, "attr-val")

	enabledMetrics["postgresql.connection.max"] = true
	mb.RecordPostgresqlConnectionMaxDataPoint(ts, 1)

	enabledMetrics["postgresql.database.count"] = true
	mb.RecordPostgresqlDatabaseCountDataPoint(ts, 1)

	enabledMetrics["postgresql.db_size"] = true
	mb.RecordPostgresqlDbSizeDataPoint(ts, 1, "attr-val")

	enabledMetrics["postgresql.index.scans"] = true
	mb.RecordPostgresqlIndexScansDataPoint(ts, 1)

	enabledMetrics["postgresql.index.size"] = true
	mb.RecordPostgresqlIndexSizeDataPoint(ts, 1)

	enabledMetrics["postgresql.operations"] = true
	mb.RecordPostgresqlOperationsDataPoint(ts, 1, "attr-val", "attr-val", AttributeOperation(1))

	enabledMetrics["postgresql.replication.data_delay"] = true
	mb.RecordPostgresqlReplicationDataDelayDataPoint(ts, 1, "attr-val")

	enabledMetrics["postgresql.rollbacks"] = true
	mb.RecordPostgresqlRollbacksDataPoint(ts, 1, "attr-val")

	enabledMetrics["postgresql.rows"] = true
	mb.RecordPostgresqlRowsDataPoint(ts, 1, "attr-val", "attr-val", AttributeState(1))

	enabledMetrics["postgresql.table.count"] = true
	mb.RecordPostgresqlTableCountDataPoint(ts, 1)

	enabledMetrics["postgresql.table.size"] = true
	mb.RecordPostgresqlTableSizeDataPoint(ts, 1)

	enabledMetrics["postgresql.table.vacuum.count"] = true
	mb.RecordPostgresqlTableVacuumCountDataPoint(ts, 1)

	enabledMetrics["postgresql.wal.age"] = true
	mb.RecordPostgresqlWalAgeDataPoint(ts, 1)

	enabledMetrics["postgresql.wal.lag"] = true
	mb.RecordPostgresqlWalLagDataPoint(ts, 1, AttributeWalOperationLag(1), "attr-val")

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

	mb.RecordPostgresqlBackendsDataPoint(ts, 1, "attr-val")
	mb.RecordPostgresqlBgwriterBuffersAllocatedDataPoint(ts, 1)
	mb.RecordPostgresqlBgwriterBuffersWritesDataPoint(ts, 1, AttributeBgBufferSource(1))
	mb.RecordPostgresqlBgwriterCheckpointCountDataPoint(ts, 1, AttributeBgCheckpointType(1))
	mb.RecordPostgresqlBgwriterDurationDataPoint(ts, 1, AttributeBgDurationType(1))
	mb.RecordPostgresqlBgwriterMaxwrittenDataPoint(ts, 1)
	mb.RecordPostgresqlBlocksReadDataPoint(ts, 1, "attr-val", "attr-val", AttributeSource(1))
	mb.RecordPostgresqlCommitsDataPoint(ts, 1, "attr-val")
	mb.RecordPostgresqlConnectionMaxDataPoint(ts, 1)
	mb.RecordPostgresqlDatabaseCountDataPoint(ts, 1)
	mb.RecordPostgresqlDbSizeDataPoint(ts, 1, "attr-val")
	mb.RecordPostgresqlIndexScansDataPoint(ts, 1)
	mb.RecordPostgresqlIndexSizeDataPoint(ts, 1)
	mb.RecordPostgresqlOperationsDataPoint(ts, 1, "attr-val", "attr-val", AttributeOperation(1))
	mb.RecordPostgresqlReplicationDataDelayDataPoint(ts, 1, "attr-val")
	mb.RecordPostgresqlRollbacksDataPoint(ts, 1, "attr-val")
	mb.RecordPostgresqlRowsDataPoint(ts, 1, "attr-val", "attr-val", AttributeState(1))
	mb.RecordPostgresqlTableCountDataPoint(ts, 1)
	mb.RecordPostgresqlTableSizeDataPoint(ts, 1)
	mb.RecordPostgresqlTableVacuumCountDataPoint(ts, 1)
	mb.RecordPostgresqlWalAgeDataPoint(ts, 1)
	mb.RecordPostgresqlWalLagDataPoint(ts, 1, AttributeWalOperationLag(1), "attr-val")

	metrics := mb.Emit(WithPostgresqlDatabaseName("attr-val"), WithPostgresqlIndexName("attr-val"), WithPostgresqlTableName("attr-val"))

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	attrCount++
	attrVal, ok := rm.Resource().Attributes().Get("postgresql.database.name")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("postgresql.index.name")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("postgresql.table.name")
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
		case "postgresql.backends":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of backends.", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["postgresql.backends"] = struct{}{}
		case "postgresql.bgwriter.buffers.allocated":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of buffers allocated.", ms.At(i).Description())
			assert.Equal(t, "{buffers}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["postgresql.bgwriter.buffers.allocated"] = struct{}{}
		case "postgresql.bgwriter.buffers.writes":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of buffers written.", ms.At(i).Description())
			assert.Equal(t, "{buffers}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("source")
			assert.True(t, ok)
			assert.Equal(t, "backend", attrVal.Str())
			validatedMetrics["postgresql.bgwriter.buffers.writes"] = struct{}{}
		case "postgresql.bgwriter.checkpoint.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of checkpoints performed.", ms.At(i).Description())
			assert.Equal(t, "{checkpoints}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "requested", attrVal.Str())
			validatedMetrics["postgresql.bgwriter.checkpoint.count"] = struct{}{}
		case "postgresql.bgwriter.duration":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total time spent writing and syncing files to disk by checkpoints.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("type")
			assert.True(t, ok)
			assert.Equal(t, "sync", attrVal.Str())
			validatedMetrics["postgresql.bgwriter.duration"] = struct{}{}
		case "postgresql.bgwriter.maxwritten":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of times the background writer stopped a cleaning scan because it had written too many buffers.", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["postgresql.bgwriter.maxwritten"] = struct{}{}
		case "postgresql.blocks_read":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of blocks read.", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("table")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("source")
			assert.True(t, ok)
			assert.Equal(t, "heap_read", attrVal.Str())
			validatedMetrics["postgresql.blocks_read"] = struct{}{}
		case "postgresql.commits":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of commits.", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["postgresql.commits"] = struct{}{}
		case "postgresql.connection.max":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Configured maximum number of client connections allowed", ms.At(i).Description())
			assert.Equal(t, "{connections}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["postgresql.connection.max"] = struct{}{}
		case "postgresql.database.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of user databases.", ms.At(i).Description())
			assert.Equal(t, "{databases}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["postgresql.database.count"] = struct{}{}
		case "postgresql.db_size":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The database disk usage.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["postgresql.db_size"] = struct{}{}
		case "postgresql.index.scans":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of index scans on a table.", ms.At(i).Description())
			assert.Equal(t, "{scans}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["postgresql.index.scans"] = struct{}{}
		case "postgresql.index.size":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The size of the index on disk.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["postgresql.index.size"] = struct{}{}
		case "postgresql.operations":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of db row operations.", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("table")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("operation")
			assert.True(t, ok)
			assert.Equal(t, "ins", attrVal.Str())
			validatedMetrics["postgresql.operations"] = struct{}{}
		case "postgresql.replication.data_delay":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The amount of data delayed in replication.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("replication_client")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["postgresql.replication.data_delay"] = struct{}{}
		case "postgresql.rollbacks":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of rollbacks.", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["postgresql.rollbacks"] = struct{}{}
		case "postgresql.rows":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The number of rows in the database.", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("database")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("table")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("state")
			assert.True(t, ok)
			assert.Equal(t, "dead", attrVal.Str())
			validatedMetrics["postgresql.rows"] = struct{}{}
		case "postgresql.table.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of user tables in a database.", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["postgresql.table.count"] = struct{}{}
		case "postgresql.table.size":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Disk space used by a table.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["postgresql.table.size"] = struct{}{}
		case "postgresql.table.vacuum.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of times a table has manually been vacuumed.", ms.At(i).Description())
			assert.Equal(t, "{vacuums}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["postgresql.table.vacuum.count"] = struct{}{}
		case "postgresql.wal.age":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Age of the oldest WAL file.", ms.At(i).Description())
			assert.Equal(t, "s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["postgresql.wal.age"] = struct{}{}
		case "postgresql.wal.lag":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Time between flushing recent WAL locally and receiving notification that the standby server has completed an operation with it.", ms.At(i).Description())
			assert.Equal(t, "s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("operation")
			assert.True(t, ok)
			assert.Equal(t, "flush", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("replication_client")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["postgresql.wal.lag"] = struct{}{}
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

	mb.RecordPostgresqlBackendsDataPoint(ts, 1, "attr-val")
	mb.RecordPostgresqlBgwriterBuffersAllocatedDataPoint(ts, 1)
	mb.RecordPostgresqlBgwriterBuffersWritesDataPoint(ts, 1, AttributeBgBufferSource(1))
	mb.RecordPostgresqlBgwriterCheckpointCountDataPoint(ts, 1, AttributeBgCheckpointType(1))
	mb.RecordPostgresqlBgwriterDurationDataPoint(ts, 1, AttributeBgDurationType(1))
	mb.RecordPostgresqlBgwriterMaxwrittenDataPoint(ts, 1)
	mb.RecordPostgresqlBlocksReadDataPoint(ts, 1, "attr-val", "attr-val", AttributeSource(1))
	mb.RecordPostgresqlCommitsDataPoint(ts, 1, "attr-val")
	mb.RecordPostgresqlConnectionMaxDataPoint(ts, 1)
	mb.RecordPostgresqlDatabaseCountDataPoint(ts, 1)
	mb.RecordPostgresqlDbSizeDataPoint(ts, 1, "attr-val")
	mb.RecordPostgresqlIndexScansDataPoint(ts, 1)
	mb.RecordPostgresqlIndexSizeDataPoint(ts, 1)
	mb.RecordPostgresqlOperationsDataPoint(ts, 1, "attr-val", "attr-val", AttributeOperation(1))
	mb.RecordPostgresqlReplicationDataDelayDataPoint(ts, 1, "attr-val")
	mb.RecordPostgresqlRollbacksDataPoint(ts, 1, "attr-val")
	mb.RecordPostgresqlRowsDataPoint(ts, 1, "attr-val", "attr-val", AttributeState(1))
	mb.RecordPostgresqlTableCountDataPoint(ts, 1)
	mb.RecordPostgresqlTableSizeDataPoint(ts, 1)
	mb.RecordPostgresqlTableVacuumCountDataPoint(ts, 1)
	mb.RecordPostgresqlWalAgeDataPoint(ts, 1)
	mb.RecordPostgresqlWalLagDataPoint(ts, 1, AttributeWalOperationLag(1), "attr-val")

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
