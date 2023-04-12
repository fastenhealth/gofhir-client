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

package fhir401

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/fastenhealth/gofhir-models-gen
// PLEASE DO NOT EDIT BY HAND

// ServiceRequest is documented here http://hl7.org/fhir/StructureDefinition/ServiceRequest
// A record of a request for service such as diagnostic investigations, treatments, or operations to be performed.
type ServiceRequest struct {
	Id                      *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                    *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules           *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text                    *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension               []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier              []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical   []string          `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri         []string          `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn                 []Reference       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces                []Reference       `bson:"replaces,omitempty" json:"replaces,omitempty"`
	Requisition             *Identifier       `bson:"requisition,omitempty" json:"requisition,omitempty"`
	Status                  RequestStatus     `bson:"status" json:"status"`
	Intent                  RequestIntent     `bson:"intent" json:"intent"`
	Category                []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Priority                *RequestPriority  `bson:"priority,omitempty" json:"priority,omitempty"`
	DoNotPerform            *bool             `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	Code                    *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	OrderDetail             []CodeableConcept `bson:"orderDetail,omitempty" json:"orderDetail,omitempty"`
	QuantityQuantity        *Quantity         `bson:"quantityQuantity,omitempty" json:"quantityQuantity,omitempty"`
	QuantityRatio           *Ratio            `bson:"quantityRatio,omitempty" json:"quantityRatio,omitempty"`
	QuantityRange           *Range            `bson:"quantityRange,omitempty" json:"quantityRange,omitempty"`
	Subject                 Reference         `bson:"subject" json:"subject"`
	Encounter               *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	OccurrenceDateTime      *string           `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *Period           `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing           `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	AsNeededBoolean         *bool             `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept  `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	AuthoredOn              *string           `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester               *Reference        `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType           *CodeableConcept  `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Performer               []Reference       `bson:"performer,omitempty" json:"performer,omitempty"`
	LocationCode            []CodeableConcept `bson:"locationCode,omitempty" json:"locationCode,omitempty"`
	LocationReference       []Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	ReasonCode              []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference         []Reference       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Insurance               []Reference       `bson:"insurance,omitempty" json:"insurance,omitempty"`
	SupportingInfo          []Reference       `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Specimen                []Reference       `bson:"specimen,omitempty" json:"specimen,omitempty"`
	BodySite                []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Note                    []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	PatientInstruction      *string           `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	RelevantHistory         []Reference       `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}

// This function returns resource reference information
func (r ServiceRequest) ResourceRef() (string, *string) {
	return "ServiceRequest", r.Id
}

type OtherServiceRequest ServiceRequest

// MarshalJSON marshals the given ServiceRequest as JSON into a byte slice
func (r ServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherServiceRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherServiceRequest: OtherServiceRequest(r),
		ResourceType:        "ServiceRequest",
	})
}

// UnmarshalServiceRequest unmarshals a ServiceRequest.
func UnmarshalServiceRequest(b []byte) (ServiceRequest, error) {
	var serviceRequest ServiceRequest
	if err := json.Unmarshal(b, &serviceRequest); err != nil {
		return serviceRequest, err
	}
	return serviceRequest, nil
}
