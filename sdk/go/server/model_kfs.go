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




// Kfs - Kowabunga File System (KFS) provides an elastic NFS-like remote storage volume.
type Kfs struct {

	// The KFS ID (auto-generated).
	Id string `json:"id,omitempty"`

	// The KFS storage volume name.
	Name string `json:"name"`

	// The KFS storage volume description.
	Description string `json:"description,omitempty"`

	// The KFS storage volume access type.
	Access string `json:"access,omitempty"`

	// The KFS storage volume's NFS protocol versions to be supported.
	Protocols []int32 `json:"protocols,omitempty"`

	// The KFS endpoint FQDN (read-only).
	Endpoint string `json:"endpoint,omitempty"`

	// The KFS storage volume bytes used (read-only).
	Size int32 `json:"size,omitempty"`
}

// AssertKfsRequired checks if the required fields are not zero-ed
func AssertKfsRequired(obj Kfs) error {
	elements := map[string]interface{}{
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertKfsConstraints checks if the values respects the defined constraints
func AssertKfsConstraints(obj Kfs) error {
	return nil
}