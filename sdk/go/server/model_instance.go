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




type Instance struct {

	// The virtual machine instance ID  (auto-generated).
	Id string `json:"id,omitempty"`

	// The virtual machine name
	Name string `json:"name"`

	// The virtual machine description.
	Description string `json:"description,omitempty"`

	// the virtual machine's memory size (in bytes).
	Memory int32 `json:"memory"`

	// the virtual machine's number of vCPUs.
	Vcpus int32 `json:"vcpus"`

	// a list of existing network adapters to be connected to the instance.
	Adapters []string `json:"adapters,omitempty"`

	// a list of existing storage volumes (i.e. disks) to be connected to the instance.
	Volumes []string `json:"volumes,omitempty"`
}

// AssertInstanceRequired checks if the required fields are not zero-ed
func AssertInstanceRequired(obj Instance) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"memory": obj.Memory,
		"vcpus": obj.Vcpus,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertInstanceConstraints checks if the values respects the defined constraints
func AssertInstanceConstraints(obj Instance) error {
	return nil
}