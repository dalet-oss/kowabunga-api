/*
Kowabunga API documentation

Kvm Orchestrator With A BUNch of Goods Added

API version: 0.37.0
Contact: csops@dalet.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KGWNetIpZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KGWNetIpZone{}

// KGWNetIpZone A KGW Network IP zone settings.
type KGWNetIpZone struct {
	// The KGW (Kowabunga Network Gateway) zone name (read-only).
	Zone string `json:"zone"`
	// The KGW (Kowabunga Network Gateway) zone gateway public virtual IP (read-only).
	Public string `json:"public"`
	// The KGW (Kowabunga Network Gateway) zone gateway private virtual IP (read-only).
	Private string `json:"private"`
}

type _KGWNetIpZone KGWNetIpZone

// NewKGWNetIpZone instantiates a new KGWNetIpZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKGWNetIpZone(zone string, public string, private string) *KGWNetIpZone {
	this := KGWNetIpZone{}
	this.Zone = zone
	this.Public = public
	this.Private = private
	return &this
}

// NewKGWNetIpZoneWithDefaults instantiates a new KGWNetIpZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKGWNetIpZoneWithDefaults() *KGWNetIpZone {
	this := KGWNetIpZone{}
	return &this
}

// GetZone returns the Zone field value
func (o *KGWNetIpZone) GetZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *KGWNetIpZone) GetZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zone, true
}

// SetZone sets field value
func (o *KGWNetIpZone) SetZone(v string) {
	o.Zone = v
}

// GetPublic returns the Public field value
func (o *KGWNetIpZone) GetPublic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *KGWNetIpZone) GetPublicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *KGWNetIpZone) SetPublic(v string) {
	o.Public = v
}

// GetPrivate returns the Private field value
func (o *KGWNetIpZone) GetPrivate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Private
}

// GetPrivateOk returns a tuple with the Private field value
// and a boolean to check if the value has been set.
func (o *KGWNetIpZone) GetPrivateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Private, true
}

// SetPrivate sets field value
func (o *KGWNetIpZone) SetPrivate(v string) {
	o.Private = v
}

func (o KGWNetIpZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KGWNetIpZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["zone"] = o.Zone
	toSerialize["public"] = o.Public
	toSerialize["private"] = o.Private
	return toSerialize, nil
}

func (o *KGWNetIpZone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"zone",
		"public",
		"private",
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

	varKGWNetIpZone := _KGWNetIpZone{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKGWNetIpZone)

	if err != nil {
		return err
	}

	*o = KGWNetIpZone(varKGWNetIpZone)

	return err
}

type NullableKGWNetIpZone struct {
	value *KGWNetIpZone
	isSet bool
}

func (v NullableKGWNetIpZone) Get() *KGWNetIpZone {
	return v.value
}

func (v *NullableKGWNetIpZone) Set(val *KGWNetIpZone) {
	v.value = val
	v.isSet = true
}

func (v NullableKGWNetIpZone) IsSet() bool {
	return v.isSet
}

func (v *NullableKGWNetIpZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKGWNetIpZone(val *KGWNetIpZone) *NullableKGWNetIpZone {
	return &NullableKGWNetIpZone{value: val, isSet: true}
}

func (v NullableKGWNetIpZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKGWNetIpZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

