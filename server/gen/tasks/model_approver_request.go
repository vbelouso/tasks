/*
 * Tasks Hub
 *
 * This is a sample Tasks Hub Server based on the OpenAPI 3.0 specification.  Some useful links: - [The Tasks Hub repository](https://github.com/vbelouso/tasks) - [The source API definition](https://github.com/vbelouso/tasks/api/openapi-spec/v1/openapi.yaml)
 *
 * API version: 1.0.11
 * Contact: fake@msbang.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tasks

type ApproverRequest struct {

	Name string `json:"name"`

	Email string `json:"email"`
}

// AssertApproverRequestRequired checks if the required fields are not zero-ed
func AssertApproverRequestRequired(obj ApproverRequest) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"email": obj.Email,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseApproverRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApproverRequest (e.g. [][]ApproverRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApproverRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApproverRequest, ok := obj.(ApproverRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApproverRequestRequired(aApproverRequest)
	})
}
