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

// checks if the Subnet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subnet{}

// Subnet struct for Subnet
type Subnet struct {
	// The subnet ID (auto-generated).
	Id *string `json:"id,omitempty"`
	// The subnet name.
	Name string `json:"name"`
	// The subnet description.
	Description *string `json:"description,omitempty"`
	// The subnet CIDR (e.g. 192.168.0.0/24).
	Cidr string `json:"cidr"`
	// The subnet router/gateway IP address (e.g. 192.168.0.254).
	Gateway string `json:"gateway"`
	// The subnet DNS server IP address (gateway value if unspecified).
	Dns *string `json:"dns,omitempty"`
	// The list of extra routes to be access through designated gateway (format is 10.0.0.0/8).
	ExtraRoutes []string `json:"extra_routes,omitempty"`
	// The subnet list of reserved IPv4 ranges (i.e. no IP address can be assigned from there).
	Reserved []IpRange `json:"reserved,omitempty"`
}

type _Subnet Subnet

// NewSubnet instantiates a new Subnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubnet(name string, cidr string, gateway string) *Subnet {
	this := Subnet{}
	this.Name = name
	this.Cidr = cidr
	this.Gateway = gateway
	return &this
}

// NewSubnetWithDefaults instantiates a new Subnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubnetWithDefaults() *Subnet {
	this := Subnet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Subnet) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Subnet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Subnet) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Subnet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Subnet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Subnet) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Subnet) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Subnet) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Subnet) SetDescription(v string) {
	o.Description = &v
}

// GetCidr returns the Cidr field value
func (o *Subnet) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *Subnet) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *Subnet) SetCidr(v string) {
	o.Cidr = v
}

// GetGateway returns the Gateway field value
func (o *Subnet) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *Subnet) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *Subnet) SetGateway(v string) {
	o.Gateway = v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *Subnet) GetDns() string {
	if o == nil || IsNil(o.Dns) {
		var ret string
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetDnsOk() (*string, bool) {
	if o == nil || IsNil(o.Dns) {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *Subnet) HasDns() bool {
	if o != nil && !IsNil(o.Dns) {
		return true
	}

	return false
}

// SetDns gets a reference to the given string and assigns it to the Dns field.
func (o *Subnet) SetDns(v string) {
	o.Dns = &v
}

// GetExtraRoutes returns the ExtraRoutes field value if set, zero value otherwise.
func (o *Subnet) GetExtraRoutes() []string {
	if o == nil || IsNil(o.ExtraRoutes) {
		var ret []string
		return ret
	}
	return o.ExtraRoutes
}

// GetExtraRoutesOk returns a tuple with the ExtraRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetExtraRoutesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtraRoutes) {
		return nil, false
	}
	return o.ExtraRoutes, true
}

// HasExtraRoutes returns a boolean if a field has been set.
func (o *Subnet) HasExtraRoutes() bool {
	if o != nil && !IsNil(o.ExtraRoutes) {
		return true
	}

	return false
}

// SetExtraRoutes gets a reference to the given []string and assigns it to the ExtraRoutes field.
func (o *Subnet) SetExtraRoutes(v []string) {
	o.ExtraRoutes = v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *Subnet) GetReserved() []IpRange {
	if o == nil || IsNil(o.Reserved) {
		var ret []IpRange
		return ret
	}
	return o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetReservedOk() ([]IpRange, bool) {
	if o == nil || IsNil(o.Reserved) {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *Subnet) HasReserved() bool {
	if o != nil && !IsNil(o.Reserved) {
		return true
	}

	return false
}

// SetReserved gets a reference to the given []IpRange and assigns it to the Reserved field.
func (o *Subnet) SetReserved(v []IpRange) {
	o.Reserved = v
}

func (o Subnet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subnet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["cidr"] = o.Cidr
	toSerialize["gateway"] = o.Gateway
	if !IsNil(o.Dns) {
		toSerialize["dns"] = o.Dns
	}
	if !IsNil(o.ExtraRoutes) {
		toSerialize["extra_routes"] = o.ExtraRoutes
	}
	if !IsNil(o.Reserved) {
		toSerialize["reserved"] = o.Reserved
	}
	return toSerialize, nil
}

func (o *Subnet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"cidr",
		"gateway",
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

	varSubnet := _Subnet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubnet)

	if err != nil {
		return err
	}

	*o = Subnet(varSubnet)

	return err
}

type NullableSubnet struct {
	value *Subnet
	isSet bool
}

func (v NullableSubnet) Get() *Subnet {
	return v.value
}

func (v *NullableSubnet) Set(val *Subnet) {
	v.value = val
	v.isSet = true
}

func (v NullableSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubnet(val *Subnet) *NullableSubnet {
	return &NullableSubnet{value: val, isSet: true}
}

func (v NullableSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

