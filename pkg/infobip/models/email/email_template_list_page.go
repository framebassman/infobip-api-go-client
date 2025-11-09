package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// EmailTemplateListPage struct for EmailTemplateListPage
type EmailTemplateListPage struct {
	// List of email template list items
	Templates *[]EmailTemplateListItem `json:"templates,omitempty"`

	// Page information (e.g. current page, total pages, etc)
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
}

// checks if the EmailTemplateListPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplateListPage{}

// NewEmailTemplateListPage instantiates a new EmailTemplateListPage object
func NewEmailTemplateListPage() *EmailTemplateListPage {
	this := EmailTemplateListPage{}
	return &this
}

// NewEmailTemplateListPageWithDefaults instantiates a new EmailTemplateListPage object with defaults
func NewEmailTemplateListPageWithDefaults() *EmailTemplateListPage {
	this := EmailTemplateListPage{}
	return &this
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *EmailTemplateListPage) GetTemplates() []EmailTemplateListItem {
	if o == nil || IsNil(o.Templates) {
		var ret []EmailTemplateListItem
		return ret
	}
	return *o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
func (o *EmailTemplateListPage) GetTemplatesOk() (*[]EmailTemplateListItem, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *EmailTemplateListPage) HasTemplates() bool {
	return o != nil && !IsNil(o.Templates)
}

// SetTemplates sets field value
func (o *EmailTemplateListPage) SetTemplates(v []EmailTemplateListItem) {
	o.Templates = &v
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *EmailTemplateListPage) GetPageInfo() PageInfo {
	if o == nil || IsNil(o.PageInfo) {
		var ret PageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
func (o *EmailTemplateListPage) GetPageInfoOk() (*PageInfo, bool) {
	if o == nil || IsNil(o.PageInfo) {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *EmailTemplateListPage) HasPageInfo() bool {
	return o != nil && !IsNil(o.PageInfo)
}

// SetPageInfo sets field value
func (o *EmailTemplateListPage) SetPageInfo(v PageInfo) {
	o.PageInfo = &v
}

func (o EmailTemplateListPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplateListPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Templates) {
		toSerialize["templates"] = o.Templates
	}
	if !IsNil(o.PageInfo) {
		toSerialize["pageInfo"] = o.PageInfo
	}
	return toSerialize, nil
}

type NullableEmailTemplateListPage struct {
	value *EmailTemplateListPage
	isSet bool
}

func (v NullableEmailTemplateListPage) Get() *EmailTemplateListPage {
	return v.value
}

func (v *NullableEmailTemplateListPage) Set(val *EmailTemplateListPage) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateListPage) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateListPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateListPage(val *EmailTemplateListPage) *NullableEmailTemplateListPage {
	return &NullableEmailTemplateListPage{value: val, isSet: true}
}

func (v NullableEmailTemplateListPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateListPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
