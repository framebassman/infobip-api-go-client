package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// GetEmailTemplatesParams struct for GetEmailTemplatesParams
type GetEmailTemplatesParams struct {
	// Page number (optional)
	Page *int32 `json:"page,omitempty"`

	// Page size (optional)
	PageSize *int32 `json:"pageSize,omitempty"`

	// Template name filter (optional)
	Name *string `json:"name,omitempty"`
}

// checks if the GetEmailTemplatesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEmailTemplatesParams{}

// NewGetEmailTemplatesParams instantiates a new GetEmailTemplatesParams object
func NewGetEmailTemplatesParams() *GetEmailTemplatesParams {
	this := GetEmailTemplatesParams{}
	return &this
}

// NewGetEmailTemplatesParamsWithDefaults instantiates a new GetEmailTemplatesParams object with defaults
func NewGetEmailTemplatesParamsWithDefaults() *GetEmailTemplatesParams {
	this := GetEmailTemplatesParams{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetEmailTemplatesParams) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
func (o *GetEmailTemplatesParams) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetEmailTemplatesParams) HasPage() bool {
	return o != nil && !IsNil(o.Page)
}

// SetPage sets field value
func (o *GetEmailTemplatesParams) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *GetEmailTemplatesParams) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
func (o *GetEmailTemplatesParams) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *GetEmailTemplatesParams) HasPageSize() bool {
	return o != nil && !IsNil(o.PageSize)
}

// SetPageSize sets field value
func (o *GetEmailTemplatesParams) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetEmailTemplatesParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
func (o *GetEmailTemplatesParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetEmailTemplatesParams) HasName() bool {
	return o != nil && !IsNil(o.Name)
}

// SetName sets field value
func (o *GetEmailTemplatesParams) SetName(v string) {
	o.Name = &v
}

func (o GetEmailTemplatesParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEmailTemplatesParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetEmailTemplatesParams struct {
	value *GetEmailTemplatesParams
	isSet bool
}

func (v NullableGetEmailTemplatesParams) Get() *GetEmailTemplatesParams {
	return v.value
}

func (v *NullableGetEmailTemplatesParams) Set(val *GetEmailTemplatesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEmailTemplatesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEmailTemplatesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEmailTemplatesParams(val *GetEmailTemplatesParams) *NullableGetEmailTemplatesParams {
	return &NullableGetEmailTemplatesParams{value: val, isSet: true}
}

func (v NullableGetEmailTemplatesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEmailTemplatesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
