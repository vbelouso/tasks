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

type GetApprovers200Response struct {

	Tasks []ApproverResponse `json:"tasks,omitempty"`
}

// AssertGetApprovers200ResponseRequired checks if the required fields are not zero-ed
func AssertGetApprovers200ResponseRequired(obj GetApprovers200Response) error {
	for _, el := range obj.Tasks {
		if err := AssertApproverResponseRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseGetApprovers200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetApprovers200Response (e.g. [][]GetApprovers200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetApprovers200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetApprovers200Response, ok := obj.(GetApprovers200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetApprovers200ResponseRequired(aGetApprovers200Response)
	})
}
