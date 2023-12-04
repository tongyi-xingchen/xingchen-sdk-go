/*
XingChen 开放接口定义

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xingchen

import (
	"encoding/json"
)

// checks if the Repository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Repository{}

// Repository 知识库文件信息
type Repository struct {
	FileSavePath         *string `json:"fileSavePath,omitempty"`
	Filename             *string `json:"filename,omitempty"`
	FileUrl              *string `json:"fileUrl,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Repository Repository

// NewRepository instantiates a new Repository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepository() *Repository {
	this := Repository{}
	return &this
}

// NewRepositoryWithDefaults instantiates a new Repository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryWithDefaults() *Repository {
	this := Repository{}
	return &this
}

// GetFileSavePath returns the FileSavePath field value if set, zero value otherwise.
func (o *Repository) GetFileSavePath() string {
	if o == nil || IsNil(o.FileSavePath) {
		var ret string
		return ret
	}
	return *o.FileSavePath
}

// GetFileSavePathOk returns a tuple with the FileSavePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetFileSavePathOk() (*string, bool) {
	if o == nil || IsNil(o.FileSavePath) {
		return nil, false
	}
	return o.FileSavePath, true
}

// HasFileSavePath returns a boolean if a field has been set.
func (o *Repository) HasFileSavePath() bool {
	if o != nil && !IsNil(o.FileSavePath) {
		return true
	}

	return false
}

// SetFileSavePath gets a reference to the given string and assigns it to the FileSavePath field.
func (o *Repository) SetFileSavePath(v string) {
	o.FileSavePath = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *Repository) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *Repository) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *Repository) SetFilename(v string) {
	o.Filename = &v
}

// GetFileUrl returns the FileUrl field value if set, zero value otherwise.
func (o *Repository) GetFileUrl() string {
	if o == nil || IsNil(o.FileUrl) {
		var ret string
		return ret
	}
	return *o.FileUrl
}

// GetFileUrlOk returns a tuple with the FileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetFileUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FileUrl) {
		return nil, false
	}
	return o.FileUrl, true
}

// HasFileUrl returns a boolean if a field has been set.
func (o *Repository) HasFileUrl() bool {
	if o != nil && !IsNil(o.FileUrl) {
		return true
	}

	return false
}

// SetFileUrl gets a reference to the given string and assigns it to the FileUrl field.
func (o *Repository) SetFileUrl(v string) {
	o.FileUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Repository) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Repository) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Repository) SetStatus(v string) {
	o.Status = &v
}

func (o Repository) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Repository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileSavePath) {
		toSerialize["fileSavePath"] = o.FileSavePath
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.FileUrl) {
		toSerialize["fileUrl"] = o.FileUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Repository) UnmarshalJSON(bytes []byte) (err error) {
	varRepository := _Repository{}

	err = json.Unmarshal(bytes, &varRepository)

	if err != nil {
		return err
	}

	*o = Repository(varRepository)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "fileSavePath")
		delete(additionalProperties, "filename")
		delete(additionalProperties, "fileUrl")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepository struct {
	value *Repository
	isSet bool
}

func (v NullableRepository) Get() *Repository {
	return v.value
}

func (v *NullableRepository) Set(val *Repository) {
	v.value = val
	v.isSet = true
}

func (v NullableRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepository(val *Repository) *NullableRepository {
	return &NullableRepository{value: val, isSet: true}
}

func (v NullableRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
