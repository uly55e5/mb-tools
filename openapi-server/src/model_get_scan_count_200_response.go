/*
 * Title
 *
 * Title
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mzserver

type GetScanCount200Response struct {
	ScanCount int64 `json:"scanCount,omitempty"`
}

// AssertGetScanCount200ResponseRequired checks if the required fields are not zero-ed
func AssertGetScanCount200ResponseRequired(obj GetScanCount200Response) error {
	return nil
}

// AssertRecurseGetScanCount200ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetScanCount200Response (e.g. [][]GetScanCount200Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetScanCount200ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetScanCount200Response, ok := obj.(GetScanCount200Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetScanCount200ResponseRequired(aGetScanCount200Response)
	})
}
