/*
Kowabunga API documentation

Kvm Orchestrator With A BUNch of Goods Added

API version: 0.10.0
Contact: csops@dalet.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KCE type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KCE{}

// KCE Kowabunga Compute Engine (KCE) is a wrapper object for bare virtual machines. It consists of an instance, one to several attached volumes and 2 network adapters (a private one, a public one). This is the prefered way for creating virtual machines. IP addresses will be automatically assigned.
type KCE struct {
	// The KCE ID  (auto-generated).
	Id *string `json:"id,omitempty"`
	// The KCE virtual machine name
	Name string `json:"name"`
	// The KCE virtual machine description.
	Description *string `json:"description,omitempty"`
	// The KCE virtual machine's memory size (in bytes).
	Memory int32 `json:"memory"`
	// The KCE virtual machine's number of vCPUs.
	Vcpus int32 `json:"vcpus"`
	// The KCE virtual machine's OS disk size (in bytes).
	Disk int32 `json:"disk"`
	// The KCE virtual machine's extra data disk size (in bytes). If unspecified, no extra data disk will be assigned.
	DataDisk *int32 `json:"data_disk,omitempty"`
	// The KCE virtual machine's assigned private IPv4 address (read-only).
	Ip *string `json:"ip,omitempty"`
}

type _KCE KCE

// NewKCE instantiates a new KCE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKCE(name string, memory int32, vcpus int32, disk int32) *KCE {
	this := KCE{}
	this.Name = name
	this.Memory = memory
	this.Vcpus = vcpus
	this.Disk = disk
	var dataDisk int32 = 0
	this.DataDisk = &dataDisk
	return &this
}

// NewKCEWithDefaults instantiates a new KCE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKCEWithDefaults() *KCE {
	this := KCE{}
	var dataDisk int32 = 0
	this.DataDisk = &dataDisk
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KCE) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KCE) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KCE) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KCE) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *KCE) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KCE) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KCE) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KCE) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KCE) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KCE) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KCE) SetDescription(v string) {
	o.Description = &v
}

// GetMemory returns the Memory field value
func (o *KCE) GetMemory() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *KCE) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *KCE) SetMemory(v int32) {
	o.Memory = v
}

// GetVcpus returns the Vcpus field value
func (o *KCE) GetVcpus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vcpus
}

// GetVcpusOk returns a tuple with the Vcpus field value
// and a boolean to check if the value has been set.
func (o *KCE) GetVcpusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vcpus, true
}

// SetVcpus sets field value
func (o *KCE) SetVcpus(v int32) {
	o.Vcpus = v
}

// GetDisk returns the Disk field value
func (o *KCE) GetDisk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Disk
}

// GetDiskOk returns a tuple with the Disk field value
// and a boolean to check if the value has been set.
func (o *KCE) GetDiskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disk, true
}

// SetDisk sets field value
func (o *KCE) SetDisk(v int32) {
	o.Disk = v
}

// GetDataDisk returns the DataDisk field value if set, zero value otherwise.
func (o *KCE) GetDataDisk() int32 {
	if o == nil || IsNil(o.DataDisk) {
		var ret int32
		return ret
	}
	return *o.DataDisk
}

// GetDataDiskOk returns a tuple with the DataDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KCE) GetDataDiskOk() (*int32, bool) {
	if o == nil || IsNil(o.DataDisk) {
		return nil, false
	}
	return o.DataDisk, true
}

// HasDataDisk returns a boolean if a field has been set.
func (o *KCE) HasDataDisk() bool {
	if o != nil && !IsNil(o.DataDisk) {
		return true
	}

	return false
}

// SetDataDisk gets a reference to the given int32 and assigns it to the DataDisk field.
func (o *KCE) SetDataDisk(v int32) {
	o.DataDisk = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *KCE) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KCE) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *KCE) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *KCE) SetIp(v string) {
	o.Ip = &v
}

func (o KCE) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KCE) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["memory"] = o.Memory
	toSerialize["vcpus"] = o.Vcpus
	toSerialize["disk"] = o.Disk
	if !IsNil(o.DataDisk) {
		toSerialize["data_disk"] = o.DataDisk
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	return toSerialize, nil
}

func (o *KCE) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"memory",
		"vcpus",
		"disk",
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

	varKCE := _KCE{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKCE)

	if err != nil {
		return err
	}

	*o = KCE(varKCE)

	return err
}

type NullableKCE struct {
	value *KCE
	isSet bool
}

func (v NullableKCE) Get() *KCE {
	return v.value
}

func (v *NullableKCE) Set(val *KCE) {
	v.value = val
	v.isSet = true
}

func (v NullableKCE) IsSet() bool {
	return v.isSet
}

func (v *NullableKCE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKCE(val *KCE) *NullableKCE {
	return &NullableKCE{value: val, isSet: true}
}

func (v NullableKCE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKCE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

