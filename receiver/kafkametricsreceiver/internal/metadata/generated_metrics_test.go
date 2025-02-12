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

	enabledMetrics["kafka.brokers"] = true
	mb.RecordKafkaBrokersDataPoint(ts, 1)

	enabledMetrics["kafka.consumer_group.lag"] = true
	mb.RecordKafkaConsumerGroupLagDataPoint(ts, 1, "attr-val", "attr-val", 1)

	enabledMetrics["kafka.consumer_group.lag_sum"] = true
	mb.RecordKafkaConsumerGroupLagSumDataPoint(ts, 1, "attr-val", "attr-val")

	enabledMetrics["kafka.consumer_group.members"] = true
	mb.RecordKafkaConsumerGroupMembersDataPoint(ts, 1, "attr-val")

	enabledMetrics["kafka.consumer_group.offset"] = true
	mb.RecordKafkaConsumerGroupOffsetDataPoint(ts, 1, "attr-val", "attr-val", 1)

	enabledMetrics["kafka.consumer_group.offset_sum"] = true
	mb.RecordKafkaConsumerGroupOffsetSumDataPoint(ts, 1, "attr-val", "attr-val")

	enabledMetrics["kafka.partition.current_offset"] = true
	mb.RecordKafkaPartitionCurrentOffsetDataPoint(ts, 1, "attr-val", 1)

	enabledMetrics["kafka.partition.oldest_offset"] = true
	mb.RecordKafkaPartitionOldestOffsetDataPoint(ts, 1, "attr-val", 1)

	enabledMetrics["kafka.partition.replicas"] = true
	mb.RecordKafkaPartitionReplicasDataPoint(ts, 1, "attr-val", 1)

	enabledMetrics["kafka.partition.replicas_in_sync"] = true
	mb.RecordKafkaPartitionReplicasInSyncDataPoint(ts, 1, "attr-val", 1)

	enabledMetrics["kafka.topic.partitions"] = true
	mb.RecordKafkaTopicPartitionsDataPoint(ts, 1, "attr-val")

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

	mb.RecordKafkaBrokersDataPoint(ts, 1)
	mb.RecordKafkaConsumerGroupLagDataPoint(ts, 1, "attr-val", "attr-val", 1)
	mb.RecordKafkaConsumerGroupLagSumDataPoint(ts, 1, "attr-val", "attr-val")
	mb.RecordKafkaConsumerGroupMembersDataPoint(ts, 1, "attr-val")
	mb.RecordKafkaConsumerGroupOffsetDataPoint(ts, 1, "attr-val", "attr-val", 1)
	mb.RecordKafkaConsumerGroupOffsetSumDataPoint(ts, 1, "attr-val", "attr-val")
	mb.RecordKafkaPartitionCurrentOffsetDataPoint(ts, 1, "attr-val", 1)
	mb.RecordKafkaPartitionOldestOffsetDataPoint(ts, 1, "attr-val", 1)
	mb.RecordKafkaPartitionReplicasDataPoint(ts, 1, "attr-val", 1)
	mb.RecordKafkaPartitionReplicasInSyncDataPoint(ts, 1, "attr-val", 1)
	mb.RecordKafkaTopicPartitionsDataPoint(ts, 1, "attr-val")

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
		case "kafka.brokers":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of brokers in the cluster.", ms.At(i).Description())
			assert.Equal(t, "{brokers}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["kafka.brokers"] = struct{}{}
		case "kafka.consumer_group.lag":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Current approximate lag of consumer group at partition of topic", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("group")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("topic")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("partition")
			assert.True(t, ok)
			assert.EqualValues(t, 1, attrVal.Int())
			validatedMetrics["kafka.consumer_group.lag"] = struct{}{}
		case "kafka.consumer_group.lag_sum":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Current approximate sum of consumer group lag across all partitions of topic", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("group")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("topic")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["kafka.consumer_group.lag_sum"] = struct{}{}
		case "kafka.consumer_group.members":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Count of members in the consumer group", ms.At(i).Description())
			assert.Equal(t, "{members}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("group")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["kafka.consumer_group.members"] = struct{}{}
		case "kafka.consumer_group.offset":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Current offset of the consumer group at partition of topic", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("group")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("topic")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("partition")
			assert.True(t, ok)
			assert.EqualValues(t, 1, attrVal.Int())
			validatedMetrics["kafka.consumer_group.offset"] = struct{}{}
		case "kafka.consumer_group.offset_sum":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Sum of consumer group offset across partitions of topic", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("group")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("topic")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["kafka.consumer_group.offset_sum"] = struct{}{}
		case "kafka.partition.current_offset":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Current offset of partition of topic.", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("topic")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("partition")
			assert.True(t, ok)
			assert.EqualValues(t, 1, attrVal.Int())
			validatedMetrics["kafka.partition.current_offset"] = struct{}{}
		case "kafka.partition.oldest_offset":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Oldest offset of partition of topic", ms.At(i).Description())
			assert.Equal(t, "1", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("topic")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("partition")
			assert.True(t, ok)
			assert.EqualValues(t, 1, attrVal.Int())
			validatedMetrics["kafka.partition.oldest_offset"] = struct{}{}
		case "kafka.partition.replicas":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of replicas for partition of topic", ms.At(i).Description())
			assert.Equal(t, "{replicas}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("topic")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("partition")
			assert.True(t, ok)
			assert.EqualValues(t, 1, attrVal.Int())
			validatedMetrics["kafka.partition.replicas"] = struct{}{}
		case "kafka.partition.replicas_in_sync":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of synchronized replicas of partition", ms.At(i).Description())
			assert.Equal(t, "{replicas}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("topic")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			attrVal, ok = dp.Attributes().Get("partition")
			assert.True(t, ok)
			assert.EqualValues(t, 1, attrVal.Int())
			validatedMetrics["kafka.partition.replicas_in_sync"] = struct{}{}
		case "kafka.topic.partitions":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of partitions in topic.", ms.At(i).Description())
			assert.Equal(t, "{partitions}", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("topic")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["kafka.topic.partitions"] = struct{}{}
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

	mb.RecordKafkaBrokersDataPoint(ts, 1)
	mb.RecordKafkaConsumerGroupLagDataPoint(ts, 1, "attr-val", "attr-val", 1)
	mb.RecordKafkaConsumerGroupLagSumDataPoint(ts, 1, "attr-val", "attr-val")
	mb.RecordKafkaConsumerGroupMembersDataPoint(ts, 1, "attr-val")
	mb.RecordKafkaConsumerGroupOffsetDataPoint(ts, 1, "attr-val", "attr-val", 1)
	mb.RecordKafkaConsumerGroupOffsetSumDataPoint(ts, 1, "attr-val", "attr-val")
	mb.RecordKafkaPartitionCurrentOffsetDataPoint(ts, 1, "attr-val", 1)
	mb.RecordKafkaPartitionOldestOffsetDataPoint(ts, 1, "attr-val", 1)
	mb.RecordKafkaPartitionReplicasDataPoint(ts, 1, "attr-val", 1)
	mb.RecordKafkaPartitionReplicasInSyncDataPoint(ts, 1, "attr-val", 1)
	mb.RecordKafkaTopicPartitionsDataPoint(ts, 1, "attr-val")

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
