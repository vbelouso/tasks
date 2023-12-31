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

type TaskResponse struct {

	Id int64 `json:"id,omitempty"`

	Title string `json:"title"`

	Description string `json:"description"`

	Approvers []ApproverRequest `json:"approvers,omitempty"`
}

// AssertTaskResponseRequired checks if the required fields are not zero-ed
func AssertTaskResponseRequired(obj TaskResponse) error {
	elements := map[string]interface{}{
		"title": obj.Title,
		"description": obj.Description,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Approvers {
		if err := AssertApproverRequestRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseTaskResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TaskResponse (e.g. [][]TaskResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTaskResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTaskResponse, ok := obj.(TaskResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTaskResponseRequired(aTaskResponse)
	})
}
