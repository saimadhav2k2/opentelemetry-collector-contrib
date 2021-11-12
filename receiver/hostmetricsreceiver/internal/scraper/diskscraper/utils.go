// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package diskscraper

import (
	"go.opentelemetry.io/collector/model/pdata"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/diskscraper/internal/metadata"
)

func initializeNumberDataPointAsInt(dataPoint pdata.NumberDataPoint, startTime, now pdata.Timestamp, deviceLabel string, directionLabel string, value int64) {
	attributes := dataPoint.Attributes()
	attributes.InsertString(metadata.Attributes.Device, deviceLabel)
	if directionLabel != "" {
		attributes.InsertString(metadata.Attributes.Direction, directionLabel)
	}
	dataPoint.SetStartTimestamp(startTime)
	dataPoint.SetTimestamp(now)
	dataPoint.SetIntVal(value)
}

func initializeNumberDataPointAsDouble(dataPoint pdata.NumberDataPoint, startTime, now pdata.Timestamp, deviceLabel string, directionLabel string, value float64) {
	attributes := dataPoint.Attributes()
	attributes.InsertString(metadata.Attributes.Device, deviceLabel)
	if directionLabel != "" {
		attributes.InsertString(metadata.Attributes.Direction, directionLabel)
	}
	dataPoint.SetStartTimestamp(startTime)
	dataPoint.SetTimestamp(now)
	dataPoint.SetDoubleVal(value)
}