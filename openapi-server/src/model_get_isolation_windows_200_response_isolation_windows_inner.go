/*
 * Title
 *
 * Title
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mzserver

type GetIsolationWindows200ResponseIsolationWindowsInner struct {

	Low float32 `json:"low,omitempty"`

	High float32 `json:"high,omitempty"`
}

// AssertGetIsolationWindows200ResponseIsolationWindowsInnerRequired checks if the required fields are not zero-ed
func AssertGetIsolationWindows200ResponseIsolationWindowsInnerRequired(obj GetIsolationWindows200ResponseIsolationWindowsInner) error {
	return nil
}

// AssertRecurseGetIsolationWindows200ResponseIsolationWindowsInnerRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GetIsolationWindows200ResponseIsolationWindowsInner (e.g. [][]GetIsolationWindows200ResponseIsolationWindowsInner), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGetIsolationWindows200ResponseIsolationWindowsInnerRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGetIsolationWindows200ResponseIsolationWindowsInner, ok := obj.(GetIsolationWindows200ResponseIsolationWindowsInner)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGetIsolationWindows200ResponseIsolationWindowsInnerRequired(aGetIsolationWindows200ResponseIsolationWindowsInner)
	})
}