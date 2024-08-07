/*
 * Kowabunga API documentation
 *
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * API version: 0.42.0
 * Contact: csops@dalet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// KgwVpcNetIpZone - A KGW VPC Network IP zone settings.
type KgwVpcNetIpZone struct {

	// The KGW (Kowabunga Network Gateway) zone name (read-only).
	Zone string `json:"zone"`

	// The KGW (Kowabunga Network Gateway) zone gateway private IP address in VPC peered subnet  (read-only).
	Private string `json:"private"`
}

// AssertKgwVpcNetIpZoneRequired checks if the required fields are not zero-ed
func AssertKgwVpcNetIpZoneRequired(obj KgwVpcNetIpZone) error {
	elements := map[string]interface{}{
		"zone": obj.Zone,
		"private": obj.Private,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertKgwVpcNetIpZoneConstraints checks if the values respects the defined constraints
func AssertKgwVpcNetIpZoneConstraints(obj KgwVpcNetIpZone) error {
	return nil
}
