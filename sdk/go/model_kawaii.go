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
)

// checks if the Kawaii type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Kawaii{}

// Kawaii A Kawaii (Kowabunga Adapative WAn Intelligent Interface) is a network gateway used for your Internet inbound and outbound traffic.
type Kawaii struct {
	// The Kawaii ID (auto-generated).
	Id *string `json:"id,omitempty"`
	// The Kawaii name.
	Name *string `json:"name,omitempty"`
	// The Kawaii description.
	Description *string `json:"description,omitempty"`
	// The Kawaii list of assigned virtual IPs per-zone addresses (read-only).
	Netip KawaiiNetIp `json:"netip,omitempty"`
	// The Kawaii firewall settings from/to public Internet).
	Firewall KawaiiFirewall `json:"firewall,omitempty"`
	// The Kawaii list of NAT forwarding entries. Kawaii will forward public Internet traffic from all public virtual IPs to requested private subnet IP addresses.
	Dnat []KawaiiDNatRule `json:"dnat,omitempty"`
	// The Kawaii list of Kowabunga private VPC subnet peering entries.
	VpcPeerings []KawaiiVpcPeering `json:"vpc_peerings,omitempty"`
}

// NewKawaii instantiates a new Kawaii object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKawaii() *Kawaii {
	this := Kawaii{}
	return &this
}

// NewKawaiiWithDefaults instantiates a new Kawaii object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKawaiiWithDefaults() *Kawaii {
	this := Kawaii{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Kawaii) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kawaii) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Kawaii) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Kawaii) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Kawaii) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kawaii) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Kawaii) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Kawaii) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Kawaii) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kawaii) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Kawaii) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Kawaii) SetDescription(v string) {
	o.Description = &v
}

// GetNetip returns the Netip field value if set, zero value otherwise.
func (o *Kawaii) GetNetip() KawaiiNetIp {
	if o == nil || IsNil(o.Netip) {
		var ret KawaiiNetIp
		return ret
	}
	return o.Netip
}

// GetNetipOk returns a tuple with the Netip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kawaii) GetNetipOk() (KawaiiNetIp, bool) {
	if o == nil || IsNil(o.Netip) {
		return KawaiiNetIp{}, false
	}
	return o.Netip, true
}

// HasNetip returns a boolean if a field has been set.
func (o *Kawaii) HasNetip() bool {
	if o != nil && !IsNil(o.Netip) {
		return true
	}

	return false
}

// SetNetip gets a reference to the given KawaiiNetIp and assigns it to the Netip field.
func (o *Kawaii) SetNetip(v KawaiiNetIp) {
	o.Netip = v
}

// GetFirewall returns the Firewall field value if set, zero value otherwise.
func (o *Kawaii) GetFirewall() KawaiiFirewall {
	if o == nil || IsNil(o.Firewall) {
		var ret KawaiiFirewall
		return ret
	}
	return o.Firewall
}

// GetFirewallOk returns a tuple with the Firewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kawaii) GetFirewallOk() (KawaiiFirewall, bool) {
	if o == nil || IsNil(o.Firewall) {
		return KawaiiFirewall{}, false
	}
	return o.Firewall, true
}

// HasFirewall returns a boolean if a field has been set.
func (o *Kawaii) HasFirewall() bool {
	if o != nil && !IsNil(o.Firewall) {
		return true
	}

	return false
}

// SetFirewall gets a reference to the given KawaiiFirewall and assigns it to the Firewall field.
func (o *Kawaii) SetFirewall(v KawaiiFirewall) {
	o.Firewall = v
}

// GetDnat returns the Dnat field value if set, zero value otherwise.
func (o *Kawaii) GetDnat() []KawaiiDNatRule {
	if o == nil || IsNil(o.Dnat) {
		var ret []KawaiiDNatRule
		return ret
	}
	return o.Dnat
}

// GetDnatOk returns a tuple with the Dnat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kawaii) GetDnatOk() ([]KawaiiDNatRule, bool) {
	if o == nil || IsNil(o.Dnat) {
		return nil, false
	}
	return o.Dnat, true
}

// HasDnat returns a boolean if a field has been set.
func (o *Kawaii) HasDnat() bool {
	if o != nil && !IsNil(o.Dnat) {
		return true
	}

	return false
}

// SetDnat gets a reference to the given []KawaiiDNatRule and assigns it to the Dnat field.
func (o *Kawaii) SetDnat(v []KawaiiDNatRule) {
	o.Dnat = v
}

// GetVpcPeerings returns the VpcPeerings field value if set, zero value otherwise.
func (o *Kawaii) GetVpcPeerings() []KawaiiVpcPeering {
	if o == nil || IsNil(o.VpcPeerings) {
		var ret []KawaiiVpcPeering
		return ret
	}
	return o.VpcPeerings
}

// GetVpcPeeringsOk returns a tuple with the VpcPeerings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kawaii) GetVpcPeeringsOk() ([]KawaiiVpcPeering, bool) {
	if o == nil || IsNil(o.VpcPeerings) {
		return nil, false
	}
	return o.VpcPeerings, true
}

// HasVpcPeerings returns a boolean if a field has been set.
func (o *Kawaii) HasVpcPeerings() bool {
	if o != nil && !IsNil(o.VpcPeerings) {
		return true
	}

	return false
}

// SetVpcPeerings gets a reference to the given []KawaiiVpcPeering and assigns it to the VpcPeerings field.
func (o *Kawaii) SetVpcPeerings(v []KawaiiVpcPeering) {
	o.VpcPeerings = v
}

func (o Kawaii) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Kawaii) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Netip) {
		toSerialize["netip"] = o.Netip
	}
	if !IsNil(o.Firewall) {
		toSerialize["firewall"] = o.Firewall
	}
	if !IsNil(o.Dnat) {
		toSerialize["dnat"] = o.Dnat
	}
	if !IsNil(o.VpcPeerings) {
		toSerialize["vpc_peerings"] = o.VpcPeerings
	}
	return toSerialize, nil
}

type NullableKawaii struct {
	value *Kawaii
	isSet bool
}

func (v NullableKawaii) Get() *Kawaii {
	return v.value
}

func (v *NullableKawaii) Set(val *Kawaii) {
	v.value = val
	v.isSet = true
}

func (v NullableKawaii) IsSet() bool {
	return v.isSet
}

func (v *NullableKawaii) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKawaii(val *Kawaii) *NullableKawaii {
	return &NullableKawaii{value: val, isSet: true}
}

func (v NullableKawaii) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKawaii) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

