/*
Kowabunga API documentation

Kvm Orchestrator With A BUNch of Goods Added

API version: 0.34.0
Contact: csops@dalet.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KGWZoneSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KGWZoneSettings{}

// KGWZoneSettings A KGW Zone network settings.
type KGWZoneSettings struct {
	// The KGW (Kowabunga Network Gateway) zone ID (read-only).
	Zone string `json:"zone"`
	// The KGW (Kowabunga Network Gateway) public virtual IP (read-only).
	PublicIp string `json:"public_ip"`
	// The KGW (Kowabunga Network Gateway) private virtual IP (read-only).
	PrivateIp string `json:"private_ip"`
}

type _KGWZoneSettings KGWZoneSettings

// NewKGWZoneSettings instantiates a new KGWZoneSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKGWZoneSettings(zone string, publicIp string, privateIp string) *KGWZoneSettings {
	this := KGWZoneSettings{}
	this.Zone = zone
	this.PublicIp = publicIp
	this.PrivateIp = privateIp
	return &this
}

// NewKGWZoneSettingsWithDefaults instantiates a new KGWZoneSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKGWZoneSettingsWithDefaults() *KGWZoneSettings {
	this := KGWZoneSettings{}
	return &this
}

// GetZone returns the Zone field value
func (o *KGWZoneSettings) GetZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *KGWZoneSettings) GetZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zone, true
}

// SetZone sets field value
func (o *KGWZoneSettings) SetZone(v string) {
	o.Zone = v
}

// GetPublicIp returns the PublicIp field value
func (o *KGWZoneSettings) GetPublicIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value
// and a boolean to check if the value has been set.
func (o *KGWZoneSettings) GetPublicIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicIp, true
}

// SetPublicIp sets field value
func (o *KGWZoneSettings) SetPublicIp(v string) {
	o.PublicIp = v
}

// GetPrivateIp returns the PrivateIp field value
func (o *KGWZoneSettings) GetPrivateIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value
// and a boolean to check if the value has been set.
func (o *KGWZoneSettings) GetPrivateIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateIp, true
}

// SetPrivateIp sets field value
func (o *KGWZoneSettings) SetPrivateIp(v string) {
	o.PrivateIp = v
}

func (o KGWZoneSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KGWZoneSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["zone"] = o.Zone
	toSerialize["public_ip"] = o.PublicIp
	toSerialize["private_ip"] = o.PrivateIp
	return toSerialize, nil
}

func (o *KGWZoneSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"zone",
		"public_ip",
		"private_ip",
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

	varKGWZoneSettings := _KGWZoneSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKGWZoneSettings)

	if err != nil {
		return err
	}

	*o = KGWZoneSettings(varKGWZoneSettings)

	return err
}

type NullableKGWZoneSettings struct {
	value *KGWZoneSettings
	isSet bool
}

func (v NullableKGWZoneSettings) Get() *KGWZoneSettings {
	return v.value
}

func (v *NullableKGWZoneSettings) Set(val *KGWZoneSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableKGWZoneSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableKGWZoneSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKGWZoneSettings(val *KGWZoneSettings) *NullableKGWZoneSettings {
	return &NullableKGWZoneSettings{value: val, isSet: true}
}

func (v NullableKGWZoneSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKGWZoneSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

