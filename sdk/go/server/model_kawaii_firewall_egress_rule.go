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




// KawaiiFirewallEgressRule - A Kawaii public firewall egress rule.
type KawaiiFirewallEgressRule struct {

	// The destination IP or CIDR to accept/drop public traffic to.
	Destination string `json:"destination,omitempty"`

	// The transport layer protocol to accept/drop public traffic to.
	Protocol string `json:"protocol,omitempty"`

	// The port (or list of ports) to accept/drop public traffic from. Ranges are accepted. Format is a-b,c-d (e.g. 443; 22,80,443; 80,443,3000-3005).
	Ports string `json:"ports"`
}

// AssertKawaiiFirewallEgressRuleRequired checks if the required fields are not zero-ed
func AssertKawaiiFirewallEgressRuleRequired(obj KawaiiFirewallEgressRule) error {
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

// AssertKawaiiFirewallEgressRuleConstraints checks if the values respects the defined constraints
func AssertKawaiiFirewallEgressRuleConstraints(obj KawaiiFirewallEgressRule) error {
	return nil
}