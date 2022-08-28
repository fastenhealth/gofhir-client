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

// ClaimResponse is documented here http://hl7.org/fhir/StructureDefinition/ClaimResponse
type ClaimResponse struct {
	Id                   *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               FinancialResourceStatusCodes    `bson:"status" json:"status"`
	Type                 CodeableConcept                 `bson:"type" json:"type"`
	SubType              *CodeableConcept                `bson:"subType,omitempty" json:"subType,omitempty"`
	Use                  Use                             `bson:"use" json:"use"`
	Patient              Reference                       `bson:"patient" json:"patient"`
	Created              string                          `bson:"created" json:"created"`
	Insurer              Reference                       `bson:"insurer" json:"insurer"`
	Requestor            *Reference                      `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Request              *Reference                      `bson:"request,omitempty" json:"request,omitempty"`
	Outcome              ClaimProcessingCodes            `bson:"outcome" json:"outcome"`
	Disposition          *string                         `bson:"disposition,omitempty" json:"disposition,omitempty"`
	PreAuthRef           *string                         `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	PreAuthPeriod        *Period                         `bson:"preAuthPeriod,omitempty" json:"preAuthPeriod,omitempty"`
	PayeeType            *CodeableConcept                `bson:"payeeType,omitempty" json:"payeeType,omitempty"`
	Item                 []ClaimResponseItem             `bson:"item,omitempty" json:"item,omitempty"`
	AddItem              []ClaimResponseAddItem          `bson:"addItem,omitempty" json:"addItem,omitempty"`
	Adjudication         []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Total                []ClaimResponseTotal            `bson:"total,omitempty" json:"total,omitempty"`
	Payment              *ClaimResponsePayment           `bson:"payment,omitempty" json:"payment,omitempty"`
	FundsReserve         *CodeableConcept                `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	FormCode             *CodeableConcept                `bson:"formCode,omitempty" json:"formCode,omitempty"`
	Form                 *Attachment                     `bson:"form,omitempty" json:"form,omitempty"`
	ProcessNote          []ClaimResponseProcessNote      `bson:"processNote,omitempty" json:"processNote,omitempty"`
	CommunicationRequest []Reference                     `bson:"communicationRequest,omitempty" json:"communicationRequest,omitempty"`
	Insurance            []ClaimResponseInsurance        `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Error                []ClaimResponseError            `bson:"error,omitempty" json:"error,omitempty"`
}
type ClaimResponseItem struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ItemSequence      int                             `bson:"itemSequence" json:"itemSequence"`
	NoteNumber        []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication" json:"adjudication"`
	Detail            []ClaimResponseItemDetail       `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ClaimResponseItemAdjudication struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `bson:"category" json:"category"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount            *Money           `bson:"amount,omitempty" json:"amount,omitempty"`
	Value             *string          `bson:"value,omitempty" json:"value,omitempty"`
}
type ClaimResponseItemDetail struct {
	Id                *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DetailSequence    int                                `bson:"detailSequence" json:"detailSequence"`
	NoteNumber        []int                              `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication    `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail         []ClaimResponseItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}
type ClaimResponseItemDetailSubDetail struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SubDetailSequence int                             `bson:"subDetailSequence" json:"subDetailSequence"`
	NoteNumber        []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ClaimResponseAddItem struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ItemSequence      []int                           `bson:"itemSequence,omitempty" json:"itemSequence,omitempty"`
	DetailSequence    []int                           `bson:"detailSequence,omitempty" json:"detailSequence,omitempty"`
	SubdetailSequence []int                           `bson:"subdetailSequence,omitempty" json:"subdetailSequence,omitempty"`
	Provider          []Reference                     `bson:"provider,omitempty" json:"provider,omitempty"`
	ProductOrService  CodeableConcept                 `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept               `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept               `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                       `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                          `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                         `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                          `bson:"net,omitempty" json:"net,omitempty"`
	BodySite          *CodeableConcept                `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite           []CodeableConcept               `bson:"subSite,omitempty" json:"subSite,omitempty"`
	NoteNumber        []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail            []ClaimResponseAddItemDetail    `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ClaimResponseAddItemDetail struct {
	Id                *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                       `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept                     `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Quantity          *Quantity                             `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                               `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                                `bson:"net,omitempty" json:"net,omitempty"`
	NoteNumber        []int                                 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication       `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail         []ClaimResponseAddItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}
type ClaimResponseAddItemDetailSubDetail struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                 `bson:"productOrService" json:"productOrService"`
	Modifier          []CodeableConcept               `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Quantity          *Quantity                       `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                          `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                         `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                          `bson:"net,omitempty" json:"net,omitempty"`
	NoteNumber        []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ClaimResponseTotal struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept `bson:"category" json:"category"`
	Amount            Money           `bson:"amount" json:"amount"`
}
type ClaimResponsePayment struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept  `bson:"type" json:"type"`
	Adjustment        *Money           `bson:"adjustment,omitempty" json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty"`
	Date              *string          `bson:"date,omitempty" json:"date,omitempty"`
	Amount            Money            `bson:"amount" json:"amount"`
	Identifier        *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
}
type ClaimResponseProcessNote struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Number            *int             `bson:"number,omitempty" json:"number,omitempty"`
	Type              *NoteType        `bson:"type,omitempty" json:"type,omitempty"`
	Text              string           `bson:"text" json:"text"`
	Language          *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}
type ClaimResponseInsurance struct {
	Id                  *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence            int         `bson:"sequence" json:"sequence"`
	Focal               bool        `bson:"focal" json:"focal"`
	Coverage            Reference   `bson:"coverage" json:"coverage"`
	BusinessArrangement *string     `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	ClaimResponse       *Reference  `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
}
type ClaimResponseError struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ItemSequence      *int            `bson:"itemSequence,omitempty" json:"itemSequence,omitempty"`
	DetailSequence    *int            `bson:"detailSequence,omitempty" json:"detailSequence,omitempty"`
	SubDetailSequence *int            `bson:"subDetailSequence,omitempty" json:"subDetailSequence,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
}
type OtherClaimResponse ClaimResponse

// MarshalJSON marshals the given ClaimResponse as JSON into a byte slice
func (r ClaimResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClaimResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherClaimResponse: OtherClaimResponse(r),
		ResourceType:       "ClaimResponse",
	})
}

// UnmarshalClaimResponse unmarshals a ClaimResponse.
func UnmarshalClaimResponse(b []byte) (ClaimResponse, error) {
	var claimResponse ClaimResponse
	if err := json.Unmarshal(b, &claimResponse); err != nil {
		return claimResponse, err
	}
	return claimResponse, nil
}
