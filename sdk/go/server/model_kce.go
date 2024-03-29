/*
 * Kowabunga API documentation
 *
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * API version: 0.30.0
 * Contact: csops@dalet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server




// Kce - A Kowabunga Compute Engine (KCE) is a wrapper object for bare virtual machines. It consists of an instance, one to several attached volumes and 2 network adapters (a private one, a public one). This is the prefered way for creating virtual machines. IP addresses will be automatically assigned.
type Kce struct {

	// The KCE (Kowabunga Compute Engine) ID (auto-generated).
	Id string `json:"id,omitempty"`

	// The KCE (Kowabunga Compute Engine) name.
	Name string `json:"name"`

	// The KCE (Kowabunga Compute Engine) description.
	Description string `json:"description,omitempty"`

	// The KCE (Kowabunga Compute Engine) memory size (in bytes).
	Memory int64 `json:"memory"`

	// The KCE (Kowabunga Compute Engine) number of vCPUs.
	Vcpus int64 `json:"vcpus"`

	// The KCE (Kowabunga Compute Engine) OS disk size (in bytes).
	Disk int64 `json:"disk"`

	// The KCE (Kowabunga Compute Engine) extra data disk size (in bytes). If unspecified, no extra data disk will be assigned.
	DataDisk int64 `json:"data_disk,omitempty"`

	// The KCE (Kowabunga Compute Engine) assigned private IPv4 address (read-only).
	Ip string `json:"ip,omitempty"`
}

// AssertKceRequired checks if the required fields are not zero-ed
func AssertKceRequired(obj Kce) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"memory": obj.Memory,
		"vcpus": obj.Vcpus,
		"disk": obj.Disk,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertKceConstraints checks if the values respects the defined constraints
func AssertKceConstraints(obj Kce) error {
	return nil
}
