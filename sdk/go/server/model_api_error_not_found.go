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




type ApiErrorNotFound struct {

	Status int32 `json:"status"`

	Error string `json:"error"`

	Code string `json:"code"`
}

// AssertApiErrorNotFoundRequired checks if the required fields are not zero-ed
func AssertApiErrorNotFoundRequired(obj ApiErrorNotFound) error {
	elements := map[string]interface{}{
		"status": obj.Status,
		"error": obj.Error,
		"code": obj.Code,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertApiErrorNotFoundConstraints checks if the values respects the defined constraints
func AssertApiErrorNotFoundConstraints(obj ApiErrorNotFound) error {
	return nil
}