package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the CreateEmailTemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEmailTemplateResponse{}

// CreateEmailTemplateResponse defines model for CreateEmailTemplateResponse.
type CreateEmailTemplateResponse struct {
	ID              int64  `json:"id"`
	Name            string `json:"name,omitempty"`
	From            string `json:"from,omitempty"`
	ReplyTo         string `json:"replyTo,omitempty"`
	Subject         string `json:"subject,omitempty"`
	Preheader       string `json:"preheader,omitempty"`
	HTML            string `json:"html,omitempty"`
	IsHTMLEditable  bool   `json:"isHtmlEditable,omitempty"`
	LandingPageID   string `json:"landingPageId,omitempty"`
	ImagePreviewURL string `json:"imagePreviewUrl,omitempty"`
	CreatedAt       string `json:"createdAt,omitempty"`
	UpdatedAt       string `json:"updatedAt,omitempty"`
}

func NewCreateEmailTemplateResponse() *CreateEmailTemplateResponse {
	this := CreateEmailTemplateResponse{}
	return &this
}

func NewCreateEmailTemplateResponseWithDefaults() *CreateEmailTemplateResponse {
	this := CreateEmailTemplateResponse{}
	return &this
}

// GetID returns the ID field value
func (o *CreateEmailTemplateResponse) GetID() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ID
}

// GetIDOk returns a tuple with the ID field value and a boolean to check if the value has been set.
func (o *CreateEmailTemplateResponse) GetIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *CreateEmailTemplateResponse) SetID(v int64) {
	o.ID = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateEmailTemplateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEmailTemplateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateEmailTemplateResponse) HasName() bool {
	return o != nil && o.Name != ""
}

// SetName sets field value
func (o *CreateEmailTemplateResponse) SetName(v string) {
	o.Name = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *CreateEmailTemplateResponse) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.From
}

func (o *CreateEmailTemplateResponse) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

func (o *CreateEmailTemplateResponse) HasFrom() bool {
	return o != nil && o.From != ""
}

func (o *CreateEmailTemplateResponse) SetFrom(v string) {
	o.From = v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *CreateEmailTemplateResponse) GetReplyTo() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReplyTo
}

func (o *CreateEmailTemplateResponse) GetReplyToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplyTo, true
}

func (o *CreateEmailTemplateResponse) HasReplyTo() bool {
	return o != nil && o.ReplyTo != ""
}

func (o *CreateEmailTemplateResponse) SetReplyTo(v string) {
	o.ReplyTo = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CreateEmailTemplateResponse) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Subject
}

func (o *CreateEmailTemplateResponse) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

func (o *CreateEmailTemplateResponse) HasSubject() bool {
	return o != nil && o.Subject != ""
}

func (o *CreateEmailTemplateResponse) SetSubject(v string) {
	o.Subject = v
}

// GetPreheader returns the Preheader field value if set, zero value otherwise.
func (o *CreateEmailTemplateResponse) GetPreheader() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Preheader
}

func (o *CreateEmailTemplateResponse) GetPreheaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Preheader, true
}

func (o *CreateEmailTemplateResponse) HasPreheader() bool {
	return o != nil && o.Preheader != ""
}

func (o *CreateEmailTemplateResponse) SetPreheader(v string) {
	o.Preheader = v
}

// GetHTML returns the HTML field value if set, zero value otherwise.
func (o *CreateEmailTemplateResponse) GetHTML() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.HTML
}

func (o *CreateEmailTemplateResponse) GetHTMLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HTML, true
}

func (o *CreateEmailTemplateResponse) HasHTML() bool {
	return o != nil && o.HTML != ""
}

func (o *CreateEmailTemplateResponse) SetHTML(v string) {
	o.HTML = v
}

// GetIsHTMLEditable returns the IsHTMLEditable field value if set, zero value otherwise.
func (o *CreateEmailTemplateResponse) GetIsHTMLEditable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsHTMLEditable
}

func (o *CreateEmailTemplateResponse) GetIsHTMLEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsHTMLEditable, true
}

func (o *CreateEmailTemplateResponse) HasIsHTMLEditable() bool {
	return o != nil && o.IsHTMLEditable
}

func (o *CreateEmailTemplateResponse) SetIsHTMLEditable(v bool) {
	o.IsHTMLEditable = v
}

// GetLandingPageID returns the LandingPageID field value if set, zero value otherwise.
func (o *CreateEmailTemplateResponse) GetLandingPageID() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LandingPageID
}

func (o *CreateEmailTemplateResponse) GetLandingPageIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LandingPageID, true
}

func (o *CreateEmailTemplateResponse) HasLandingPageID() bool {
	return o != nil && o.LandingPageID != ""
}

func (o *CreateEmailTemplateResponse) SetLandingPageID(v string) {
	o.LandingPageID = v
}

// GetImagePreviewURL returns the ImagePreviewURL field value if set, zero value otherwise.
func (o *CreateEmailTemplateResponse) GetImagePreviewURL() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ImagePreviewURL
}

func (o *CreateEmailTemplateResponse) GetImagePreviewURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImagePreviewURL, true
}

func (o *CreateEmailTemplateResponse) HasImagePreviewURL() bool {
	return o != nil && o.ImagePreviewURL != ""
}

func (o *CreateEmailTemplateResponse) SetImagePreviewURL(v string) {
	o.ImagePreviewURL = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateEmailTemplateResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedAt
}

func (o *CreateEmailTemplateResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

func (o *CreateEmailTemplateResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != ""
}

func (o *CreateEmailTemplateResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CreateEmailTemplateResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedAt
}

func (o *CreateEmailTemplateResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

func (o *CreateEmailTemplateResponse) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != ""
}

func (o *CreateEmailTemplateResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o CreateEmailTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEmailTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.ID
	if o.Name != "" {
		toSerialize["name"] = o.Name
	}
	if o.From != "" {
		toSerialize["from"] = o.From
	}
	if o.ReplyTo != "" {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if o.Subject != "" {
		toSerialize["subject"] = o.Subject
	}
	if o.Preheader != "" {
		toSerialize["preheader"] = o.Preheader
	}
	if o.HTML != "" {
		toSerialize["html"] = o.HTML
	}
	if o.IsHTMLEditable {
		toSerialize["isHtmlEditable"] = o.IsHTMLEditable
	}
	if o.LandingPageID != "" {
		toSerialize["landingPageId"] = o.LandingPageID
	}
	if o.ImagePreviewURL != "" {
		toSerialize["imagePreviewUrl"] = o.ImagePreviewURL
	}
	if o.CreatedAt != "" {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != "" {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableCreateEmailTemplateResponse struct {
	value *CreateEmailTemplateResponse
	isSet bool
}

func (v NullableCreateEmailTemplateResponse) Get() *CreateEmailTemplateResponse {
	return v.value
}

func (v *NullableCreateEmailTemplateResponse) Set(val *CreateEmailTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEmailTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEmailTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEmailTemplateResponse(val *CreateEmailTemplateResponse) *NullableCreateEmailTemplateResponse {
	return &NullableCreateEmailTemplateResponse{value: val, isSet: true}
}

func (v NullableCreateEmailTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEmailTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
