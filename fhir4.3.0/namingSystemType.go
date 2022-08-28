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

// NamingSystemType is documented here http://hl7.org/fhir/ValueSet/namingsystem-type
type NamingSystemType int

const (
	NamingSystemTypeCodesystem NamingSystemType = iota
	NamingSystemTypeIdentifier
	NamingSystemTypeRoot
)

func (code NamingSystemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *NamingSystemType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "codesystem":
		*code = NamingSystemTypeCodesystem
	case "identifier":
		*code = NamingSystemTypeIdentifier
	case "root":
		*code = NamingSystemTypeRoot
	default:
		return fmt.Errorf("unknown NamingSystemType code `%s`", s)
	}
	return nil
}
func (code NamingSystemType) String() string {
	return code.Code()
}
func (code NamingSystemType) Code() string {
	switch code {
	case NamingSystemTypeCodesystem:
		return "codesystem"
	case NamingSystemTypeIdentifier:
		return "identifier"
	case NamingSystemTypeRoot:
		return "root"
	}
	return "<unknown>"
}
func (code NamingSystemType) Display() string {
	switch code {
	case NamingSystemTypeCodesystem:
		return "Code System"
	case NamingSystemTypeIdentifier:
		return "Identifier"
	case NamingSystemTypeRoot:
		return "Root"
	}
	return "<unknown>"
}
func (code NamingSystemType) Definition() string {
	switch code {
	case NamingSystemTypeCodesystem:
		return "The naming system is used to define concepts and symbols to represent those concepts; e.g. UCUM, LOINC, NDC code, local lab codes, etc."
	case NamingSystemTypeIdentifier:
		return "The naming system is used to manage identifiers (e.g. license numbers, order numbers, etc.)."
	case NamingSystemTypeRoot:
		return "The naming system is used as the root for other identifiers and naming systems."
	}
	return "<unknown>"
}
