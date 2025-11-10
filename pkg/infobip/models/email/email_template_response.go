package email

import (
	"encoding/json"

	. "github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
)

// checks if the EmailTemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplateResponse{}

// EmailTemplateResponse defines model for EmailTemplateResponse.
type EmailTemplateResponse struct {
	ID              int64  `json:"id"`
	Name            string `json:"name,omitempty"`
	ImagePreviewURL string `json:"imagePreviewUrl,omitempty"`
	CreatedAt       string `json:"createdAt,omitempty"`
	UpdatedAt       string `json:"updatedAt,omitempty"`
}

func EmailTemplateResponseWithDefaults() *EmailTemplateResponse {
	this := EmailTemplateResponse{}
	return &this
}

// GetID returns the ID field value
func (o *EmailTemplateResponse) GetID() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ID
}

// GetIDOk returns a tuple with the ID field value and a boolean to check if the value has been set.
func (o *EmailTemplateResponse) GetIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *EmailTemplateResponse) SetID(v int64) {
	o.ID = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmailTemplateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmailTemplateResponse) HasName() bool {
	return o != nil && o.Name != ""
}

// SetName sets field value
func (o *EmailTemplateResponse) SetName(v string) {
	o.Name = v
}

// GetImagePreviewURL returns the ImagePreviewURL field value if set, zero value otherwise.
func (o *EmailTemplateResponse) GetImagePreviewURL() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ImagePreviewURL
}

func (o *EmailTemplateResponse) GetImagePreviewURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImagePreviewURL, true
}

func (o *EmailTemplateResponse) HasImagePreviewURL() bool {
	return o != nil && o.ImagePreviewURL != ""
}

func (o *EmailTemplateResponse) SetImagePreviewURL(v string) {
	o.ImagePreviewURL = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EmailTemplateResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedAt
}

func (o *EmailTemplateResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

func (o *EmailTemplateResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != ""
}

func (o *EmailTemplateResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EmailTemplateResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedAt
}

func (o *EmailTemplateResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

func (o *EmailTemplateResponse) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != ""
}

func (o *EmailTemplateResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o EmailTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.ID
	if o.Name != "" {
		toSerialize["name"] = o.Name
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

type NullableEmailTemplateResponse struct {
	value *EmailTemplateResponse
	isSet bool
}

func (v NullableEmailTemplateResponse) Get() *EmailTemplateResponse {
	return v.value
}

func (v *NullableEmailTemplateResponse) Set(val *EmailTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateResponse(val *EmailTemplateResponse) *NullableEmailTemplateResponse {
	return &NullableEmailTemplateResponse{value: val, isSet: true}
}

func (v NullableEmailTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
