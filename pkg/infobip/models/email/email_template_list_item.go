package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// EmailTemplateListItem struct for EmailTemplateListItem
type EmailTemplateListItem struct {
	// Template ID
	Id *int64 `json:"id,omitempty"`

	// Template name
	Name *string `json:"name,omitempty"`

	// Template subject
	Subject *string `json:"subject,omitempty"`

	// Template HTML body
	Body *string `json:"body,omitempty"`

	// Language of the template (optional)
	Language *string `json:"language,omitempty"`
}

// checks if the EmailTemplateListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplateListItem{}

// NewEmailTemplateListItem instantiates a new EmailTemplateListItem object
func NewEmailTemplateListItem() *EmailTemplateListItem {
	this := EmailTemplateListItem{}
	return &this
}

// NewEmailTemplateListItemWithDefaults instantiates a new EmailTemplateListItem object with defaults
func NewEmailTemplateListItemWithDefaults() *EmailTemplateListItem {
	this := EmailTemplateListItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailTemplateListItem) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
func (o *EmailTemplateListItem) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailTemplateListItem) HasId() bool {
	return o != nil && !IsNil(o.Id)
}

// SetId sets field value
func (o *EmailTemplateListItem) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmailTemplateListItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
func (o *EmailTemplateListItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmailTemplateListItem) HasName() bool {
	return o != nil && !IsNil(o.Name)
}

// SetName sets field value
func (o *EmailTemplateListItem) SetName(v string) {
	o.Name = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *EmailTemplateListItem) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
func (o *EmailTemplateListItem) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailTemplateListItem) HasSubject() bool {
	return o != nil && !IsNil(o.Subject)
}

// SetSubject sets field value
func (o *EmailTemplateListItem) SetSubject(v string) {
	o.Subject = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *EmailTemplateListItem) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
func (o *EmailTemplateListItem) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *EmailTemplateListItem) HasBody() bool {
	return o != nil && !IsNil(o.Body)
}

// SetBody sets field value
func (o *EmailTemplateListItem) SetBody(v string) {
	o.Body = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *EmailTemplateListItem) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
func (o *EmailTemplateListItem) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *EmailTemplateListItem) HasLanguage() bool {
	return o != nil && !IsNil(o.Language)
}

// SetLanguage sets field value
func (o *EmailTemplateListItem) SetLanguage(v string) {
	o.Language = &v
}

func (o EmailTemplateListItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplateListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullableEmailTemplateListItem struct {
	value *EmailTemplateListItem
	isSet bool
}

func (v NullableEmailTemplateListItem) Get() *EmailTemplateListItem {
	return v.value
}

func (v *NullableEmailTemplateListItem) Set(val *EmailTemplateListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateListItem(val *EmailTemplateListItem) *NullableEmailTemplateListItem {
	return &NullableEmailTemplateListItem{value: val, isSet: true}
}

func (v NullableEmailTemplateListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
