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




// Kce - Kowabunga Compute Engine (KCE) is a wrapper object for bare virtual machines. It consists of an instance, one to several attached volumes and 2 network adapters (a private one, a public one). This is the prefered way for creating virtual machines. IP addresses will be automatically assigned.
type Kce struct {

	// The KCE ID  (auto-generated).
	Id string `json:"id,omitempty"`

	// The KCE virtual machine name
	Name string `json:"name"`

	// The KCE virtual machine description.
	Description string `json:"description,omitempty"`

	// The KCE virtual machine's memory size (in bytes).
	Memory int32 `json:"memory"`

	// The KCE virtual machine's number of vCPUs.
	Vcpus int32 `json:"vcpus"`

	// The KCE virtual machine's OS disk size (in bytes).
	Disk int32 `json:"disk"`

	// The KCE virtual machine's extra data disk size (in bytes). If unspecified, no extra data disk will be assigned.
	DataDisk int32 `json:"data_disk,omitempty"`

	// The KCE virtual machine's assigned private IPv4 address (read-only).
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