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




type Host struct {

	// The host ID (auto-generated).
	Id string `json:"id,omitempty"`

	// The host name.
	Name string `json:"name"`

	// The host description.
	Description string `json:"description,omitempty"`

	// The protocol to use to issue libvirt connection.
	Protocol string `json:"protocol"`

	// The host libvirt's IPv4 address.
	Address string `json:"address"`

	// The host libvirt's port.
	Port int32 `json:"port,omitempty"`

	Tls HostTls `json:"tls,omitempty"`

	// Global cost associated to the host (deprecated, will be removed).
	Cost Cost `json:"cost,omitempty"`

	// Cost associated to the host's CPU resources.
	CpuCost Cost `json:"cpu_cost,omitempty"`

	// Cost associated to the host's memory resources.
	MemoryCost Cost `json:"memory_cost,omitempty"`

	// The host CPU resource over-commit ratio. Overcommitting CPU resources for VMs means allocating more virtual CPUs (vCPUs) to the virtual machines (VMs) than the physical cores available on the host. This can help optimize the utilization of the host CPU and increase the density of VMs per host.
	OvercommitCpuRatio int32 `json:"overcommit_cpu_ratio,omitempty"`

	// The host memory resource over-commit ratio. Memory overcommitment is a concept in computing that covers the assignment of more memory to virtual computing devices (or processes) than the physical machine they are hosted, or running on, actually has.
	OvercommitMemoryRatio int32 `json:"overcommit_memory_ratio,omitempty"`
}

// AssertHostRequired checks if the required fields are not zero-ed
func AssertHostRequired(obj Host) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"protocol": obj.Protocol,
		"address": obj.Address,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertHostTlsRequired(obj.Tls); err != nil {
		return err
	}
	return nil
}

// AssertHostConstraints checks if the values respects the defined constraints
func AssertHostConstraints(obj Host) error {
	return nil
}