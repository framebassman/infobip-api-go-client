package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// CreateEmailTemplateRequest defines model for CreateEmailTemplateRequest.
type CreateEmailTemplateRequest struct {
	// Attachments JSON string of attachments to be sent with the email template.
	Attachments *string `json:"attachments,omitempty"`

	// From Email address with optional sender name.
	From *string `json:"from,omitempty"`

	// Html HTML content of the email template.
	Html string `json:"html"`

	// LandingPage The identifier of an opt out landing late to be used and displayed when an end user clicks the unsubscribe link. Create a landing page in your Infobip account and use the ID number. For example, 1_23456. If not present, the default opt out landing page is used.
	LandingPage *string `json:"landingPage,omitempty"`

	// Name Name of the email template.
	Name *string `json:"name,omitempty"`

	// Preheader Preheader of the email template.
	Preheader *string `json:"preheader,omitempty"`

	// ReplyTo Email address to which recipients of the email can reply.
	ReplyTo *string `json:"replyTo,omitempty"`

	// Subject Subject of the email template.
	Subject *string `json:"subject,omitempty"`
}

// checks if the CreateEmailTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEmailTemplateRequest{}

// NewCreateEmailTemplateRequest instantiates a new CreateEmailTemplateRequest object
func NewCreateEmailTemplateRequest(html string) *CreateEmailTemplateRequest {
	this := CreateEmailTemplateRequest{}
	this.Html = html
	return &this
}

// NewCreateEmailTemplateRequestWithDefaults instantiates a new CreateEmailTemplateRequest object with defaults
func NewCreateEmailTemplateRequestWithDefaults() *CreateEmailTemplateRequest {
	this := CreateEmailTemplateRequest{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *CreateEmailTemplateRequest) GetAttachments() string {
	if o == nil || IsNil(o.Attachments) {
		var ret string
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
func (o *CreateEmailTemplateRequest) GetAttachmentsOk() (*string, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CreateEmailTemplateRequest) HasAttachments() bool {
	return o != nil && !IsNil(o.Attachments)
}

// SetAttachments sets field value
func (o *CreateEmailTemplateRequest) SetAttachments(v string) {
	o.Attachments = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *CreateEmailTemplateRequest) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
func (o *CreateEmailTemplateRequest) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *CreateEmailTemplateRequest) HasFrom() bool {
	return o != nil && !IsNil(o.From)
}

// SetFrom sets field value
func (o *CreateEmailTemplateRequest) SetFrom(v string) {
	o.From = &v
}

// GetHtml returns the Html field value
func (o *CreateEmailTemplateRequest) GetHtml() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Html
}

// GetHtmlOk returns a tuple with the Html field value, and a boolean to check if the value has been set.
func (o *CreateEmailTemplateRequest) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Html, true
}

// SetHtml sets field value
func (o *CreateEmailTemplateRequest) SetHtml(v string) {
	o.Html = v
}

// GetLandingPage returns the LandingPage field value if set, zero value otherwise.
func (o *CreateEmailTemplateRequest) GetLandingPage() string {
	if o == nil || IsNil(o.LandingPage) {
		var ret string
		return ret
	}
	return *o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value if set, nil otherwise
func (o *CreateEmailTemplateRequest) GetLandingPageOk() (*string, bool) {
	if o == nil || IsNil(o.LandingPage) {
		return nil, false
	}
	return o.LandingPage, true
}

// HasLandingPage returns a boolean if a field has been set.
func (o *CreateEmailTemplateRequest) HasLandingPage() bool {
	return o != nil && !IsNil(o.LandingPage)
}

// SetLandingPage sets field value
func (o *CreateEmailTemplateRequest) SetLandingPage(v string) {
	o.LandingPage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateEmailTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
func (o *CreateEmailTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateEmailTemplateRequest) HasName() bool {
	return o != nil && !IsNil(o.Name)
}

// SetName sets field value
func (o *CreateEmailTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetPreheader returns the Preheader field value if set, zero value otherwise.
func (o *CreateEmailTemplateRequest) GetPreheader() string {
	if o == nil || IsNil(o.Preheader) {
		var ret string
		return ret
	}
	return *o.Preheader
}

// GetPreheaderOk returns a tuple with the Preheader field value if set, nil otherwise
func (o *CreateEmailTemplateRequest) GetPreheaderOk() (*string, bool) {
	if o == nil || IsNil(o.Preheader) {
		return nil, false
	}
	return o.Preheader, true
}

// HasPreheader returns a boolean if a field has been set.
func (o *CreateEmailTemplateRequest) HasPreheader() bool {
	return o != nil && !IsNil(o.Preheader)
}

// SetPreheader sets field value
func (o *CreateEmailTemplateRequest) SetPreheader(v string) {
	o.Preheader = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *CreateEmailTemplateRequest) GetReplyTo() string {
	if o == nil || IsNil(o.ReplyTo) {
		var ret string
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
func (o *CreateEmailTemplateRequest) GetReplyToOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *CreateEmailTemplateRequest) HasReplyTo() bool {
	return o != nil && !IsNil(o.ReplyTo)
}

// SetReplyTo sets field value
func (o *CreateEmailTemplateRequest) SetReplyTo(v string) {
	o.ReplyTo = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CreateEmailTemplateRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
func (o *CreateEmailTemplateRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CreateEmailTemplateRequest) HasSubject() bool {
	return o != nil && !IsNil(o.Subject)
}

// SetSubject sets field value
func (o *CreateEmailTemplateRequest) SetSubject(v string) {
	o.Subject = &v
}

func (o CreateEmailTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEmailTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	toSerialize["html"] = o.Html
	if !IsNil(o.LandingPage) {
		toSerialize["landingPage"] = o.LandingPage
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Preheader) {
		toSerialize["preheader"] = o.Preheader
	}
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	return toSerialize, nil
}

type NullableCreateEmailTemplateRequest struct {
	value *CreateEmailTemplateRequest
	isSet bool
}

func (v NullableCreateEmailTemplateRequest) Get() *CreateEmailTemplateRequest {
	return v.value
}

func (v *NullableCreateEmailTemplateRequest) Set(val *CreateEmailTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEmailTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEmailTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEmailTemplateRequest(val *CreateEmailTemplateRequest) *NullableCreateEmailTemplateRequest {
	return &NullableCreateEmailTemplateRequest{value: val, isSet: true}
}

func (v NullableCreateEmailTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEmailTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// CreateEmailTemplateMultipartRequestBody defines body for CreateEmailTemplate for multipart/form-data ContentType.
type CreateEmailTemplateMultipartRequestBody = CreateEmailTemplateRequest

// PatchEmailTemplateMultipartRequestBody defines body for PatchEmailTemplate for multipart/form-data ContentType.
type PatchEmailTemplateMultipartRequestBody = CreateEmailTemplateRequest

// UpdateEmailTemplateMultipartRequestBody defines body for UpdateEmailTemplate for multipart/form-data ContentType.
type UpdateEmailTemplateMultipartRequestBody = CreateEmailTemplateRequest
