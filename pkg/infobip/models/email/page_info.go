package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// PageInfo struct for PageInfo
type PageInfo struct {
	// Current page number
	Page *int32 `json:"page,omitempty"`

	// Total number of pages
	PageCount *int32 `json:"pageCount,omitempty"`

	// Total number of items
	TotalItems *int32 `json:"totalItems,omitempty"`
}

// checks if the PageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageInfo{}

// NewPageInfo instantiates a new PageInfo object
func NewPageInfo() *PageInfo {
	this := PageInfo{}
	return &this
}

// NewPageInfoWithDefaults instantiates a new PageInfo object with defaults
func NewPageInfoWithDefaults() *PageInfo {
	this := PageInfo{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PageInfo) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
func (o *PageInfo) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PageInfo) HasPage() bool {
	return o != nil && !IsNil(o.Page)
}

// SetPage sets field value
func (o *PageInfo) SetPage(v int32) {
	o.Page = &v
}

// GetPageCount returns the PageCount field value if set, zero value otherwise.
func (o *PageInfo) GetPageCount() int32 {
	if o == nil || IsNil(o.PageCount) {
		var ret int32
		return ret
	}
	return *o.PageCount
}

// GetPageCountOk returns a tuple with the PageCount field value if set, nil otherwise
func (o *PageInfo) GetPageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PageCount) {
		return nil, false
	}
	return o.PageCount, true
}

// HasPageCount returns a boolean if a field has been set.
func (o *PageInfo) HasPageCount() bool {
	return o != nil && !IsNil(o.PageCount)
}

// SetPageCount sets field value
func (o *PageInfo) SetPageCount(v int32) {
	o.PageCount = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *PageInfo) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
func (o *PageInfo) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *PageInfo) HasTotalItems() bool {
	return o != nil && !IsNil(o.TotalItems)
}

// SetTotalItems sets field value
func (o *PageInfo) SetTotalItems(v int32) {
	o.TotalItems = &v
}

func (o PageInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageCount) {
		toSerialize["pageCount"] = o.PageCount
	}
	if !IsNil(o.TotalItems) {
		toSerialize["totalItems"] = o.TotalItems
	}
	return toSerialize, nil
}

type NullablePageInfo struct {
	value *PageInfo
	isSet bool
}

func (v NullablePageInfo) Get() *PageInfo {
	return v.value
}

func (v *NullablePageInfo) Set(val *PageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageInfo(val *PageInfo) *NullablePageInfo {
	return &NullablePageInfo{value: val, isSet: true}
}

func (v NullablePageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
