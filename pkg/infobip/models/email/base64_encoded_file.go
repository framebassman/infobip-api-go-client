package email

import (
	"encoding/json"

	. "github.com/framebassman/infobip-api-go-client/v3/pkg/infobip"
)

// Base64EncodedFile struct for Base64EncodedFile
type Base64EncodedFile struct {
	// Base64-encoded file content
	Content *string `json:"content,omitempty"`

	// File name
	FileName *string `json:"fileName,omitempty"`

	// Content type of file
	ContentType *string `json:"contentType,omitempty"`
}

// checks if the Base64EncodedFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Base64EncodedFile{}

// NewBase64EncodedFile instantiates a new Base64EncodedFile object
func NewBase64EncodedFile() *Base64EncodedFile {
	this := Base64EncodedFile{}
	return &this
}

// NewBase64EncodedFileWithDefaults instantiates a new Base64EncodedFile object with defaults
func NewBase64EncodedFileWithDefaults() *Base64EncodedFile {
	this := Base64EncodedFile{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Base64EncodedFile) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Base64EncodedFile) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Base64EncodedFile) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}
	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *Base64EncodedFile) SetContent(v string) {
	o.Content = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *Base64EncodedFile) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Base64EncodedFile) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *Base64EncodedFile) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}
	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *Base64EncodedFile) SetFileName(v string) {
	o.FileName = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *Base64EncodedFile) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Base64EncodedFile) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *Base64EncodedFile) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}
	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *Base64EncodedFile) SetContentType(v string) {
	o.ContentType = &v
}

func (o Base64EncodedFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Base64EncodedFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	return toSerialize, nil
}

type NullableBase64EncodedFile struct {
	value *Base64EncodedFile
	isSet bool
}

func (v NullableBase64EncodedFile) Get() *Base64EncodedFile {
	return v.value
}

func (v *NullableBase64EncodedFile) Set(val *Base64EncodedFile) {
	v.value = val
	v.isSet = true
}

func (v NullableBase64EncodedFile) IsSet() bool {
	return v.isSet
}

func (v *NullableBase64EncodedFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBase64EncodedFile(val *Base64EncodedFile) *NullableBase64EncodedFile {
	return &NullableBase64EncodedFile{value: val, isSet: true}
}

func (v NullableBase64EncodedFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBase64EncodedFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
