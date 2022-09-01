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

// TestScript is documented here http://hl7.org/fhir/StructureDefinition/TestScript
// A structured set of tests against a FHIR server or client implementation to determine compliance against the FHIR specification.
type TestScript struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                  `bson:"url" json:"url"`
	Identifier        *Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                 `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                  `bson:"name" json:"name"`
	Title             *string                 `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus       `bson:"status" json:"status"`
	Experimental      *bool                   `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                 `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                 `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail         `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                 `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext          `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept       `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                 `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                 `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Origin            []TestScriptOrigin      `bson:"origin,omitempty" json:"origin,omitempty"`
	Destination       []TestScriptDestination `bson:"destination,omitempty" json:"destination,omitempty"`
	Metadata          *TestScriptMetadata     `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Fixture           []TestScriptFixture     `bson:"fixture,omitempty" json:"fixture,omitempty"`
	Profile           []Reference             `bson:"profile,omitempty" json:"profile,omitempty"`
	Variable          []TestScriptVariable    `bson:"variable,omitempty" json:"variable,omitempty"`
	Setup             *TestScriptSetup        `bson:"setup,omitempty" json:"setup,omitempty"`
	Test              []TestScriptTest        `bson:"test,omitempty" json:"test,omitempty"`
	Teardown          *TestScriptTeardown     `bson:"teardown,omitempty" json:"teardown,omitempty"`
}

// An abstract server used in operations within this test script in the origin element.
// The purpose of this element is to define the profile of an origin element used elsewhere in the script.  Test engines could then use the origin-profile mapping to offer a filtered list of test systems that can serve as the sender for the interaction.
type TestScriptOrigin struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Index             int         `bson:"index" json:"index"`
	Profile           Coding      `bson:"profile" json:"profile"`
}

// An abstract server used in operations within this test script in the destination element.
// The purpose of this element is to define the profile of a destination element used elsewhere in the script.  Test engines could then use the destination-profile mapping to offer a filtered list of test systems that can serve as the receiver for the interaction.
type TestScriptDestination struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Index             int         `bson:"index" json:"index"`
	Profile           Coding      `bson:"profile" json:"profile"`
}

// The required capability must exist and are assumed to function correctly on the FHIR server being tested.
type TestScriptMetadata struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Link              []TestScriptMetadataLink       `bson:"link,omitempty" json:"link,omitempty"`
	Capability        []TestScriptMetadataCapability `bson:"capability" json:"capability"`
}

// A link to the FHIR specification that this test is covering.
type TestScriptMetadataLink struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string      `bson:"url" json:"url"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
}

// Capabilities that must exist and are assumed to function correctly on the FHIR server being tested.
// When the metadata capabilities section is defined at TestScript.metadata or at TestScript.setup.metadata, and the server's conformance statement does not contain the elements defined in the minimal conformance statement, then all the tests in the TestScript are skipped.  When the metadata capabilities section is defined at TestScript.test.metadata and the server's conformance statement does not contain the elements defined in the minimal conformance statement, then only that test is skipped.  The "metadata.capabilities.required" and "metadata.capabilities.validated" elements only indicate whether the capabilities are the primary focus of the test script or not.  They do not impact the skipping logic.  Capabilities whose "metadata.capabilities.validated" flag is true are the primary focus of the test script.
type TestScriptMetadataCapability struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Required          bool        `bson:"required" json:"required"`
	Validated         bool        `bson:"validated" json:"validated"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Origin            []int       `bson:"origin,omitempty" json:"origin,omitempty"`
	Destination       *int        `bson:"destination,omitempty" json:"destination,omitempty"`
	Link              []string    `bson:"link,omitempty" json:"link,omitempty"`
	Capabilities      string      `bson:"capabilities" json:"capabilities"`
}

// Fixture in the test script - by reference (uri). All fixtures are required for the test script to execute.
type TestScriptFixture struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Autocreate        bool        `bson:"autocreate" json:"autocreate"`
	Autodelete        bool        `bson:"autodelete" json:"autodelete"`
	Resource          *Reference  `bson:"resource,omitempty" json:"resource,omitempty"`
}

// Variable is set based either on element value in response body or on header field value in the response headers.
// Variables would be set based either on XPath/JSONPath expressions against fixtures (static and response), or headerField evaluations against response headers. If variable evaluates to nodelist or anything other than a primitive value, then test engine would report an error.  Variables would be used to perform clean replacements in "operation.params", "operation.requestHeader.value", and "operation.url" element values during operation calls and in "assert.value" during assertion evaluations. This limits the places that test engines would need to look for placeholders "${}".  Variables are scoped to the whole script. They are NOT evaluated at declaration. They are evaluated by test engine when used for substitutions in "operation.params", "operation.requestHeader.value", and "operation.url" element values during operation calls and in "assert.value" during assertion evaluations.  See example testscript-search.xml.
type TestScriptVariable struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	DefaultValue      *string     `bson:"defaultValue,omitempty" json:"defaultValue,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Expression        *string     `bson:"expression,omitempty" json:"expression,omitempty"`
	HeaderField       *string     `bson:"headerField,omitempty" json:"headerField,omitempty"`
	Hint              *string     `bson:"hint,omitempty" json:"hint,omitempty"`
	Path              *string     `bson:"path,omitempty" json:"path,omitempty"`
	SourceId          *string     `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
}

// A series of required setup operations before tests are executed.
type TestScriptSetup struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            []TestScriptSetupAction `bson:"action" json:"action"`
}

// Action would contain either an operation or an assertion.
// An action should contain either an operation or an assertion but not both.  It can contain any number of variables.
type TestScriptSetupAction struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `bson:"assert,omitempty" json:"assert,omitempty"`
}

// The operation to perform.
type TestScriptSetupActionOperation struct {
	Id                *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *Coding                                       `bson:"type,omitempty" json:"type,omitempty"`
	Resource          *string                                       `bson:"resource,omitempty" json:"resource,omitempty"`
	Label             *string                                       `bson:"label,omitempty" json:"label,omitempty"`
	Description       *string                                       `bson:"description,omitempty" json:"description,omitempty"`
	Accept            *string                                       `bson:"accept,omitempty" json:"accept,omitempty"`
	ContentType       *string                                       `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Destination       *int                                          `bson:"destination,omitempty" json:"destination,omitempty"`
	EncodeRequestUrl  bool                                          `bson:"encodeRequestUrl" json:"encodeRequestUrl"`
	Method            *TestScriptRequestMethodCode                  `bson:"method,omitempty" json:"method,omitempty"`
	Origin            *int                                          `bson:"origin,omitempty" json:"origin,omitempty"`
	Params            *string                                       `bson:"params,omitempty" json:"params,omitempty"`
	RequestHeader     []TestScriptSetupActionOperationRequestHeader `bson:"requestHeader,omitempty" json:"requestHeader,omitempty"`
	RequestId         *string                                       `bson:"requestId,omitempty" json:"requestId,omitempty"`
	ResponseId        *string                                       `bson:"responseId,omitempty" json:"responseId,omitempty"`
	SourceId          *string                                       `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	TargetId          *string                                       `bson:"targetId,omitempty" json:"targetId,omitempty"`
	Url               *string                                       `bson:"url,omitempty" json:"url,omitempty"`
}

// Header elements would be used to set HTTP headers.
// This gives control to test-script writers to set headers explicitly based on test requirements.  It will allow for testing using:  - "If-Modified-Since" and "If-None-Match" headers.  See http://build.fhir.org/http.html#2.1.0.5.1 - "If-Match" header.  See http://build.fhir.org/http.html#2.1.0.11 - Conditional Create using "If-None-Exist".  See http://build.fhir.org/http.html#2.1.0.13.1 - Invalid "Content-Type" header for negative testing. - etc.
type TestScriptSetupActionOperationRequestHeader struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Field             string      `bson:"field" json:"field"`
	Value             string      `bson:"value" json:"value"`
}

// Evaluates the results of previous operations to determine if the server under test behaves appropriately.
// In order to evaluate an assertion, the request, response, and results of the most recently executed operation must always be maintained by the test engine.
type TestScriptSetupActionAssert struct {
	Id                        *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Label                     *string                      `bson:"label,omitempty" json:"label,omitempty"`
	Description               *string                      `bson:"description,omitempty" json:"description,omitempty"`
	Direction                 *AssertionDirectionType      `bson:"direction,omitempty" json:"direction,omitempty"`
	CompareToSourceId         *string                      `bson:"compareToSourceId,omitempty" json:"compareToSourceId,omitempty"`
	CompareToSourceExpression *string                      `bson:"compareToSourceExpression,omitempty" json:"compareToSourceExpression,omitempty"`
	CompareToSourcePath       *string                      `bson:"compareToSourcePath,omitempty" json:"compareToSourcePath,omitempty"`
	ContentType               *string                      `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Expression                *string                      `bson:"expression,omitempty" json:"expression,omitempty"`
	HeaderField               *string                      `bson:"headerField,omitempty" json:"headerField,omitempty"`
	MinimumId                 *string                      `bson:"minimumId,omitempty" json:"minimumId,omitempty"`
	NavigationLinks           *bool                        `bson:"navigationLinks,omitempty" json:"navigationLinks,omitempty"`
	Operator                  *AssertionOperatorType       `bson:"operator,omitempty" json:"operator,omitempty"`
	Path                      *string                      `bson:"path,omitempty" json:"path,omitempty"`
	RequestMethod             *TestScriptRequestMethodCode `bson:"requestMethod,omitempty" json:"requestMethod,omitempty"`
	RequestURL                *string                      `bson:"requestURL,omitempty" json:"requestURL,omitempty"`
	Resource                  *string                      `bson:"resource,omitempty" json:"resource,omitempty"`
	Response                  *AssertionResponseTypes      `bson:"response,omitempty" json:"response,omitempty"`
	ResponseCode              *string                      `bson:"responseCode,omitempty" json:"responseCode,omitempty"`
	SourceId                  *string                      `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	ValidateProfileId         *string                      `bson:"validateProfileId,omitempty" json:"validateProfileId,omitempty"`
	Value                     *string                      `bson:"value,omitempty" json:"value,omitempty"`
	WarningOnly               bool                         `bson:"warningOnly" json:"warningOnly"`
}

// A test in this script.
type TestScriptTest struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string                `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	Action            []TestScriptTestAction `bson:"action" json:"action"`
}

// Action would contain either an operation or an assertion.
// An action should contain either an operation or an assertion but not both.  It can contain any number of variables.
type TestScriptTestAction struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `bson:"assert,omitempty" json:"assert,omitempty"`
}

// A series of operations required to clean up after all the tests are executed (successfully or otherwise).
type TestScriptTeardown struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            []TestScriptTeardownAction `bson:"action" json:"action"`
}

// The teardown action will only contain an operation.
// An action should contain either an operation or an assertion but not both.  It can contain any number of variables.
type TestScriptTeardownAction struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         TestScriptSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
}

// This function returns resource reference information
func (r TestScript) ResourceRef() (string, *string) {
	return "TestScript", r.Id
}

type OtherTestScript TestScript

// MarshalJSON marshals the given TestScript as JSON into a byte slice
func (r TestScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestScript
		ResourceType string `json:"resourceType"`
	}{
		OtherTestScript: OtherTestScript(r),
		ResourceType:    "TestScript",
	})
}

// UnmarshalTestScript unmarshals a TestScript.
func UnmarshalTestScript(b []byte) (TestScript, error) {
	var testScript TestScript
	if err := json.Unmarshal(b, &testScript); err != nil {
		return testScript, err
	}
	return testScript, nil
}
