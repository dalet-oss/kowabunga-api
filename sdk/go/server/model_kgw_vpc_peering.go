/*
 * Kowabunga API documentation
 *
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * API version: 0.34.0
 * Contact: csops@dalet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// KgwVpcPeering - A KGW internal VPC subnet peering.
type KgwVpcPeering struct {

	// Kowabunga Subnet ID to be peered with (IP addresses will be automatically assigned into)..
	Subnet string `json:"subnet"`

	// Ports to be reachable from peered subnet. Accept Ranges. If specified, traffic will be filtered..
	Ports string `json:"ports,omitempty"`

	// The KGW (Kowabunga Network Gateway) auto-assigned private IPs in peered subnet (read-only).
	Ips []string `json:"ips,omitempty"`
}

// AssertKgwVpcPeeringRequired checks if the required fields are not zero-ed
func AssertKgwVpcPeeringRequired(obj KgwVpcPeering) error {
	elements := map[string]interface{}{
		"subnet": obj.Subnet,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertKgwVpcPeeringConstraints checks if the values respects the defined constraints
func AssertKgwVpcPeeringConstraints(obj KgwVpcPeering) error {
	return nil
}