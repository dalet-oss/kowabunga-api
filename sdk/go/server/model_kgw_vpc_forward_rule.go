/*
 * Kowabunga API documentation
 *
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * API version: 0.40.0
 * Contact: csops@dalet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// KgwVpcForwardRule - A KGW VPC firewall forwarding rule.
type KgwVpcForwardRule struct {

	// The transport layer protocol to forward public traffic to.
	Protocol string `json:"protocol,omitempty"`

	// The port (or list of ports) to forward public traffic from. Ranges are accepted. Format is a-b,c-d (e.g. 443; 22,80,443; 80,443,3000-3005).
	Ports string `json:"ports"`
}

// AssertKgwVpcForwardRuleRequired checks if the required fields are not zero-ed
func AssertKgwVpcForwardRuleRequired(obj KgwVpcForwardRule) error {
	elements := map[string]interface{}{
		"ports": obj.Ports,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertKgwVpcForwardRuleConstraints checks if the values respects the defined constraints
func AssertKgwVpcForwardRuleConstraints(obj KgwVpcForwardRule) error {
	return nil
}
