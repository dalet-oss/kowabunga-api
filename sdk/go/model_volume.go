/*
Kowabunga API documentation

Kvm Orchestrator With A BUNch of Goods Added

API version: 0.50.0
Contact: csops@dalet.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Volume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Volume{}

// Volume A storage volume.
type Volume struct {
	// The storage volume ID (auto-generated).
	Id *string `json:"id,omitempty"`
	// The storage volume name.
	Name string `json:"name"`
	// The storage volume description.
	Description *string `json:"description,omitempty"`
	// The storage volume type.
	Type string `json:"type"`
	// The storage volume size (in bytes).
	Size int64 `json:"size"`
}

type _Volume Volume

// NewVolume instantiates a new Volume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolume(name string, type_ string, size int64) *Volume {
	this := Volume{}
	this.Name = name
	this.Type = type_
	this.Size = size
	return &this
}

// NewVolumeWithDefaults instantiates a new Volume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeWithDefaults() *Volume {
	this := Volume{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Volume) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Volume) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Volume) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Volume) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Volume) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Volume) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Volume) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Volume) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Volume) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *Volume) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Volume) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Volume) SetType(v string) {
	o.Type = v
}

// GetSize returns the Size field value
func (o *Volume) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *Volume) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *Volume) SetSize(v int64) {
	o.Size = v
}

func (o Volume) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Volume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	toSerialize["size"] = o.Size
	return toSerialize, nil
}

func (o *Volume) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"size",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVolume := _Volume{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVolume)

	if err != nil {
		return err
	}

	*o = Volume(varVolume)

	return err
}

type NullableVolume struct {
	value *Volume
	isSet bool
}

func (v NullableVolume) Get() *Volume {
	return v.value
}

func (v *NullableVolume) Set(val *Volume) {
	v.value = val
	v.isSet = true
}

func (v NullableVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolume(val *Volume) *NullableVolume {
	return &NullableVolume{value: val, isSet: true}
}

func (v NullableVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

