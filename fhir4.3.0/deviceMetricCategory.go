// Copyright 2022 - Fasten Health
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// DeviceMetricCategory is documented here http://hl7.org/fhir/ValueSet/metric-category
type DeviceMetricCategory int

const (
	DeviceMetricCategoryMeasurement DeviceMetricCategory = iota
	DeviceMetricCategorySetting
	DeviceMetricCategoryCalculation
	DeviceMetricCategoryUnspecified
)

func (code DeviceMetricCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceMetricCategory) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "measurement":
		*code = DeviceMetricCategoryMeasurement
	case "setting":
		*code = DeviceMetricCategorySetting
	case "calculation":
		*code = DeviceMetricCategoryCalculation
	case "unspecified":
		*code = DeviceMetricCategoryUnspecified
	default:
		return fmt.Errorf("unknown DeviceMetricCategory code `%s`", s)
	}
	return nil
}
func (code DeviceMetricCategory) String() string {
	return code.Code()
}
func (code DeviceMetricCategory) Code() string {
	switch code {
	case DeviceMetricCategoryMeasurement:
		return "measurement"
	case DeviceMetricCategorySetting:
		return "setting"
	case DeviceMetricCategoryCalculation:
		return "calculation"
	case DeviceMetricCategoryUnspecified:
		return "unspecified"
	}
	return "<unknown>"
}
func (code DeviceMetricCategory) Display() string {
	switch code {
	case DeviceMetricCategoryMeasurement:
		return "Measurement"
	case DeviceMetricCategorySetting:
		return "Setting"
	case DeviceMetricCategoryCalculation:
		return "Calculation"
	case DeviceMetricCategoryUnspecified:
		return "Unspecified"
	}
	return "<unknown>"
}
func (code DeviceMetricCategory) Definition() string {
	switch code {
	case DeviceMetricCategoryMeasurement:
		return "DeviceObservations generated for this DeviceMetric are measured."
	case DeviceMetricCategorySetting:
		return "DeviceObservations generated for this DeviceMetric is a setting that will influence the behavior of the Device."
	case DeviceMetricCategoryCalculation:
		return "DeviceObservations generated for this DeviceMetric are calculated."
	case DeviceMetricCategoryUnspecified:
		return "The category of this DeviceMetric is unspecified."
	}
	return "<unknown>"
}
