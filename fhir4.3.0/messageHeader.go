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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// MessageHeader is documented here http://hl7.org/fhir/StructureDefinition/MessageHeader
type MessageHeader struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Destination       []MessageHeaderDestination `bson:"destination,omitempty" json:"destination,omitempty"`
	Sender            *Reference                 `bson:"sender,omitempty" json:"sender,omitempty"`
	Enterer           *Reference                 `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Author            *Reference                 `bson:"author,omitempty" json:"author,omitempty"`
	Source            MessageHeaderSource        `bson:"source" json:"source"`
	Responsible       *Reference                 `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Reason            *CodeableConcept           `bson:"reason,omitempty" json:"reason,omitempty"`
	Response          *MessageHeaderResponse     `bson:"response,omitempty" json:"response,omitempty"`
	Focus             []Reference                `bson:"focus,omitempty" json:"focus,omitempty"`
	Definition        *string                    `bson:"definition,omitempty" json:"definition,omitempty"`
}
type MessageHeaderDestination struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	Target            *Reference  `bson:"target,omitempty" json:"target,omitempty"`
	Endpoint          string      `bson:"endpoint" json:"endpoint"`
	Receiver          *Reference  `bson:"receiver,omitempty" json:"receiver,omitempty"`
}
type MessageHeaderSource struct {
	Id                *string       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string       `bson:"name,omitempty" json:"name,omitempty"`
	Software          *string       `bson:"software,omitempty" json:"software,omitempty"`
	Version           *string       `bson:"version,omitempty" json:"version,omitempty"`
	Contact           *ContactPoint `bson:"contact,omitempty" json:"contact,omitempty"`
	Endpoint          string        `bson:"endpoint" json:"endpoint"`
}
type MessageHeaderResponse struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        string       `bson:"identifier" json:"identifier"`
	Code              ResponseType `bson:"code" json:"code"`
	Details           *Reference   `bson:"details,omitempty" json:"details,omitempty"`
}
type OtherMessageHeader MessageHeader

// MarshalJSON marshals the given MessageHeader as JSON into a byte slice
func (r MessageHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMessageHeader
		ResourceType string `json:"resourceType"`
	}{
		OtherMessageHeader: OtherMessageHeader(r),
		ResourceType:       "MessageHeader",
	})
}

// UnmarshalMessageHeader unmarshals a MessageHeader.
func UnmarshalMessageHeader(b []byte) (MessageHeader, error) {
	var messageHeader MessageHeader
	if err := json.Unmarshal(b, &messageHeader); err != nil {
		return messageHeader, err
	}
	return messageHeader, nil
}
