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




// ProjectResources - The global project resource quotas/usage (0 for unlimited).
type ProjectResources struct {

	// the maximum total number of vCPUs allowed to be consumed by sum of all instances.
	Vcpus int32 `json:"vcpus,omitempty"`

	// the maximum total memory (in bytes) allowed to be consumed by sum of all instances.
	Memory int32 `json:"memory,omitempty"`

	// the maximum total disk capacity allowed to be consumed by sum of all instances.
	Storage int32 `json:"storage,omitempty"`

	// the maximum number of instances allowed to be spawned.
	Instances int32 `json:"instances,omitempty"`
}

// AssertProjectResourcesRequired checks if the required fields are not zero-ed
func AssertProjectResourcesRequired(obj ProjectResources) error {
	return nil
}

// AssertProjectResourcesConstraints checks if the values respects the defined constraints
func AssertProjectResourcesConstraints(obj ProjectResources) error {
	return nil
}