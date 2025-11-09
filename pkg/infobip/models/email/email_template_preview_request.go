package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// EmailTemplatePreviewRequest defines model for EmailTemplatePreviewRequest.
type EmailTemplatePreviewRequest struct {
	// Html HTML content for the preview.
	Html string `json:"html"`

	// Subject Subject line for the preview.
	Subject *string `json:"subject,omitempty"`

	// Preheader Preheader for the preview.
	Preheader *string `json:"preheader,omitempty"`
}

// checks if the EmailTemplatePreviewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplatePreviewRequest{}

// NewEmailTemplatePreviewRequest instantiates a new EmailTemplatePreviewRequest object
func NewEmailTemplatePreviewRequest(html string) *EmailTemplatePreviewRequest {
	this := EmailTemplatePreviewRequest{}
	this.Html = html
	return &this
}

// NewEmailTemplatePreviewRequestWithDefaults instantiates a new EmailTemplatePreviewRequest object with defaults
func NewEmailTemplatePreviewRequestWithDefaults() *EmailTemplatePreviewRequest {
	this := EmailTemplatePreviewRequest{}
	return &this
}

// GetHtml returns the Html field value
func (o *EmailTemplatePreviewRequest) GetHtml() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Html
}

// GetHtmlOk returns a tuple with the Html field value, and a boolean to check if the value has been set.
func (o *EmailTemplatePreviewRequest) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Html, true
}

// SetHtml sets field value
func (o *EmailTemplatePreviewRequest) SetHtml(v string) {
	o.Html = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *EmailTemplatePreviewRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
func (o *EmailTemplatePreviewRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailTemplatePreviewRequest) HasSubject() bool {
	return o != nil && !IsNil(o.Subject)
}

// SetSubject sets field value
func (o *EmailTemplatePreviewRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetPreheader returns the Preheader field value if set, zero value otherwise.
func (o *EmailTemplatePreviewRequest) GetPreheader() string {
	if o == nil || IsNil(o.Preheader) {
		var ret string
		return ret
	}
	return *o.Preheader
}

// GetPreheaderOk returns a tuple with the Preheader field value if set, nil otherwise
func (o *EmailTemplatePreviewRequest) GetPreheaderOk() (*string, bool) {
	if o == nil || IsNil(o.Preheader) {
		return nil, false
	}
	return o.Preheader, true
}

// HasPreheader returns a boolean if a field has been set.
func (o *EmailTemplatePreviewRequest) HasPreheader() bool {
	return o != nil && !IsNil(o.Preheader)
}

// SetPreheader sets field value
func (o *EmailTemplatePreviewRequest) SetPreheader(v string) {
	o.Preheader = &v
}

func (o EmailTemplatePreviewRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplatePreviewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["html"] = o.Html
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Preheader) {
		toSerialize["preheader"] = o.Preheader
	}
	return toSerialize, nil
}

type NullableEmailTemplatePreviewRequest struct {
	value *EmailTemplatePreviewRequest
	isSet bool
}

func (v NullableEmailTemplatePreviewRequest) Get() *EmailTemplatePreviewRequest {
	return v.value
}

func (v *NullableEmailTemplatePreviewRequest) Set(val *EmailTemplatePreviewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplatePreviewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplatePreviewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplatePreviewRequest(val *EmailTemplatePreviewRequest) *NullableEmailTemplatePreviewRequest {
	return &NullableEmailTemplatePreviewRequest{value: val, isSet: true}
}

func (v NullableEmailTemplatePreviewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplatePreviewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
