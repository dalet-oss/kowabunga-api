/*
Kowabunga API documentation

Kvm Orchestrator With A BUNch of Goods Added

API version: 0.42.0
Contact: csops@dalet.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KawaiiVpcForwardRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KawaiiVpcForwardRule{}

// KawaiiVpcForwardRule A Kawaii VPC firewall forwarding rule.
type KawaiiVpcForwardRule struct {
	// The transport layer protocol to forward public traffic to.
	Protocol *string `json:"protocol,omitempty"`
	// The port (or list of ports) to forward public traffic from. Ranges are accepted. Format is a-b,c-d (e.g. 443; 22,80,443; 80,443,3000-3005).
	Ports string `json:"ports"`
}

type _KawaiiVpcForwardRule KawaiiVpcForwardRule

// NewKawaiiVpcForwardRule instantiates a new KawaiiVpcForwardRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKawaiiVpcForwardRule(ports string) *KawaiiVpcForwardRule {
	this := KawaiiVpcForwardRule{}
	var protocol string = "tcp"
	this.Protocol = &protocol
	this.Ports = ports
	return &this
}

// NewKawaiiVpcForwardRuleWithDefaults instantiates a new KawaiiVpcForwardRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKawaiiVpcForwardRuleWithDefaults() *KawaiiVpcForwardRule {
	this := KawaiiVpcForwardRule{}
	var protocol string = "tcp"
	this.Protocol = &protocol
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *KawaiiVpcForwardRule) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KawaiiVpcForwardRule) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *KawaiiVpcForwardRule) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *KawaiiVpcForwardRule) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPorts returns the Ports field value
func (o *KawaiiVpcForwardRule) GetPorts() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *KawaiiVpcForwardRule) GetPortsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ports, true
}

// SetPorts sets field value
func (o *KawaiiVpcForwardRule) SetPorts(v string) {
	o.Ports = v
}

func (o KawaiiVpcForwardRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KawaiiVpcForwardRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	toSerialize["ports"] = o.Ports
	return toSerialize, nil
}

func (o *KawaiiVpcForwardRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ports",
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

	varKawaiiVpcForwardRule := _KawaiiVpcForwardRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKawaiiVpcForwardRule)

	if err != nil {
		return err
	}

	*o = KawaiiVpcForwardRule(varKawaiiVpcForwardRule)

	return err
}

type NullableKawaiiVpcForwardRule struct {
	value *KawaiiVpcForwardRule
	isSet bool
}

func (v NullableKawaiiVpcForwardRule) Get() *KawaiiVpcForwardRule {
	return v.value
}

func (v *NullableKawaiiVpcForwardRule) Set(val *KawaiiVpcForwardRule) {
	v.value = val
	v.isSet = true
}

func (v NullableKawaiiVpcForwardRule) IsSet() bool {
	return v.isSet
}

func (v *NullableKawaiiVpcForwardRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKawaiiVpcForwardRule(val *KawaiiVpcForwardRule) *NullableKawaiiVpcForwardRule {
	return &NullableKawaiiVpcForwardRule{value: val, isSet: true}
}

func (v NullableKawaiiVpcForwardRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKawaiiVpcForwardRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

