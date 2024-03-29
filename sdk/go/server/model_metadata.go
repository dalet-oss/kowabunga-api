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




// Metadata - A key/value metadata.
type Metadata struct {

	// The metadata key.
	Key string `json:"key,omitempty"`

	// The metadata value.
	Value string `json:"value,omitempty"`
}

// AssertMetadataRequired checks if the required fields are not zero-ed
func AssertMetadataRequired(obj Metadata) error {
	return nil
}

// AssertMetadataConstraints checks if the values respects the defined constraints
func AssertMetadataConstraints(obj Metadata) error {
	return nil
}
