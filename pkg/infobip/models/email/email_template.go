package email

import (
	"encoding/json"

	. "github.com/framebassman/infobip-api-go-client/v3/pkg/infobip"
)

// EmailTemplate struct for EmailTemplate
type EmailTemplate struct {
	// Template ID
	Id *string `json:"id,omitempty"`

	// Template name
	Name *string `json:"name,omitempty"`

	// Template subject
	Subject *string `json:"subject,omitempty"`

	// Template HTML body
	Body *string `json:"body,omitempty"`

	// Language of the template (optional)
	Language *string `json:"language,omitempty"`
}

// checks if the EmailTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplate{}

// NewEmailTemplate instantiates a new EmailTemplate object
func NewEmailTemplate() *EmailTemplate {
	this := EmailTemplate{}
	return &this
}

// NewEmailTemplateWithDefaults instantiates a new EmailTemplate object with defaults
func NewEmailTemplateWithDefaults() *EmailTemplate {
	this := EmailTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
func (o *EmailTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailTemplate) HasId() bool {
	return o != nil && !IsNil(o.Id)
}

// SetId sets field value
func (o *EmailTemplate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmailTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
func (o *EmailTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmailTemplate) HasName() bool {
	return o != nil && !IsNil(o.Name)
}

// SetName sets field value
func (o *EmailTemplate) SetName(v string) {
	o.Name = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *EmailTemplate) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
func (o *EmailTemplate) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailTemplate) HasSubject() bool {
	return o != nil && !IsNil(o.Subject)
}

// SetSubject sets field value
func (o *EmailTemplate) SetSubject(v string) {
	o.Subject = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *EmailTemplate) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
func (o *EmailTemplate) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *EmailTemplate) HasBody() bool {
	return o != nil && !IsNil(o.Body)
}

// SetBody sets field value
func (o *EmailTemplate) SetBody(v string) {
	o.Body = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *EmailTemplate) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
func (o *EmailTemplate) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *EmailTemplate) HasLanguage() bool {
	return o != nil && !IsNil(o.Language)
}

// SetLanguage sets field value
func (o *EmailTemplate) SetLanguage(v string) {
	o.Language = &v
}

func (o EmailTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplate) ToMap() (map[string]interface{}, error) {
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

type NullableEmailTemplate struct {
	value *EmailTemplate
	isSet bool
}

func (v NullableEmailTemplate) Get() *EmailTemplate {
	return v.value
}

func (v *NullableEmailTemplate) Set(val *EmailTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplate(val *EmailTemplate) *NullableEmailTemplate {
	return &NullableEmailTemplate{value: val, isSet: true}
}

func (v NullableEmailTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
