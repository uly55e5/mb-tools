/*
 * Title
 *
 * Title
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mzserver

type GetSource200Response struct {
	Source string `json:"source,omitempty"`
}

// AssertGetSource200ResponseRequired checks if the required fields are not zero-ed
func AssertGetSource200ResponseRequired(obj GetSource200Response) error {
	return nil
}

// AssertRecurseGetSource200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetSource200Response (e.g. [][]GetSource200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetSource200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetSource200Response, ok := obj.(GetSource200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetSource200ResponseRequired(aGetSource200Response)
	})
}