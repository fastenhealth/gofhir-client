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

// OperationKind is documented here http://hl7.org/fhir/ValueSet/operation-kind
type OperationKind int

const (
	OperationKindOperation OperationKind = iota
	OperationKindQuery
)

func (code OperationKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *OperationKind) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "operation":
		*code = OperationKindOperation
	case "query":
		*code = OperationKindQuery
	default:
		return fmt.Errorf("unknown OperationKind code `%s`", s)
	}
	return nil
}
func (code OperationKind) String() string {
	return code.Code()
}
func (code OperationKind) Code() string {
	switch code {
	case OperationKindOperation:
		return "operation"
	case OperationKindQuery:
		return "query"
	}
	return "<unknown>"
}
func (code OperationKind) Display() string {
	switch code {
	case OperationKindOperation:
		return "Operation"
	case OperationKindQuery:
		return "Query"
	}
	return "<unknown>"
}
func (code OperationKind) Definition() string {
	switch code {
	case OperationKindOperation:
		return "This operation is invoked as an operation."
	case OperationKindQuery:
		return "This operation is a named query, invoked using the search mechanism."
	}
	return "<unknown>"
}
