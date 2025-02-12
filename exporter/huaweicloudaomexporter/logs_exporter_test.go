// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package huaweicloudaomexporter

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
)

func createSimpleLogData(numberOfLogs int) plog.Logs {
	logs := plog.NewLogs()
	logs.ResourceLogs().AppendEmpty() // Add an empty ResourceLogs
	rl := logs.ResourceLogs().AppendEmpty()
	rl.ScopeLogs().AppendEmpty() // Add an empty ScopeLogs
	sl := rl.ScopeLogs().AppendEmpty()

	for i := 0; i < numberOfLogs; i++ {
		ts := pcommon.Timestamp(int64(i) * time.Millisecond.Nanoseconds())
		logRecord := sl.LogRecords().AppendEmpty()
		logRecord.Body().SetStr("mylog")
		logRecord.Attributes().PutStr(conventions.AttributeServiceName, "myapp")
		logRecord.Attributes().PutStr("my-label", "myapp-type")
		logRecord.Attributes().PutStr(conventions.AttributeHostName, "myhost")
		logRecord.Attributes().PutStr("custom", "custom")
		logRecord.SetTimestamp(ts)
	}
	sl.LogRecords().AppendEmpty()

	return logs
}

func TestNewLogsExporter(t *testing.T) {
	got, err := newLogsExporter(exportertest.NewNopSettings(), &Config{
		Endpoint:        "https://lts-access.cn-north-4.myhuaweicloud.com",
		RegionId:        "cn-north-4",
		ProjectId:       "demo-project",
		LogGroupId:      "demo-group-id",
		LogStreamId:     "demo-stream-id",
		AccessKeyID:     "demo-ak",
		AccessKeySecret: "demo-sk",
	})
	assert.NoError(t, err)
	require.NotNil(t, got)

	err = got.ConsumeLogs(context.Background(), createSimpleLogData(3))
	assert.NoError(t, err)
	time.Sleep(time.Second * 4)
}

func TestTokenExporter(t *testing.T) {
	got, err := newLogsExporter(exportertest.NewNopSettings(), &Config{
		Endpoint:      "https://lts-access.cn-north-4.myhuaweicloud.com",
		RegionId:      "cn-north-4",
		ProjectId:     "demo-project-id",
		LogGroupId:    "demo-group-id",
		LogStreamId:   "demo-stream-id",
		TokenFilePath: filepath.Join("testdata", "config.yaml"),
	})
	assert.NoError(t, err)
	require.NotNil(t, got)
}

func TestNewFailsWithEmptyLogsExporterName(t *testing.T) {
	got, err := newLogsExporter(exportertest.NewNopSettings(), &Config{})
	assert.Error(t, err)
	require.Nil(t, got)
}
