/*
 * Kowabunga API documentation
 *
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * API version: 0.10.0
 * Contact: csops@dalet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// KgwNat - KGW Nat definition
type KgwNat struct {

	// Target Private IP. Leave blank for a new generated one
	PrivateIp string `json:"private_ip"`

	// Public IP from created Adapter. Leave empty to use the default Public IP
	PublicIp string `json:"public_ip,omitempty"`

	// Ports being forwarded from the public to the private IP. Accept Ranges
	Ports string `json:"ports"`
}

// AssertKgwNatRequired checks if the required fields are not zero-ed
func AssertKgwNatRequired(obj KgwNat) error {
	elements := map[string]interface{}{
		"private_ip": obj.PrivateIp,
		"ports": obj.Ports,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertKgwNatConstraints checks if the values respects the defined constraints
func AssertKgwNatConstraints(obj KgwNat) error {
	return nil
}