/*
 * Trusted Issuers Registry
 *
 * The subset of the [Trusted Issuers Registryas defined by EBSI](https://api-pilot.ebsi.eu/docs/apis/trusted-issuers-registry/v4#/) as currently required by the [VCVerifier](https://github.com/FIWARE/VCVerifier).  
 *
 * API version: v3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package generated

type ProblemDetails struct {

	// An absolute URI that identifies the problem type. When dereferenced, it SHOULD provide human-readable documentation for the problem type.
	Type string `json:"type,omitempty"`

	// A short summary of the problem type.
	Title string `json:"title,omitempty"`

	// The HTTP status code generated by the origin server for this occurrence of the problem.
	Status float32 `json:"status,omitempty"`

	// A human readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`

	// An absolute URI that identifies the specific occurrence of the problem. It may or may not yield further information if dereferenced.
	Instance string `json:"instance,omitempty"`
}

// AssertProblemDetailsRequired checks if the required fields are not zero-ed
func AssertProblemDetailsRequired(obj ProblemDetails) error {
	return nil
}

// AssertRecurseProblemDetailsRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ProblemDetails (e.g. [][]ProblemDetails), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseProblemDetailsRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aProblemDetails, ok := obj.(ProblemDetails)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertProblemDetailsRequired(aProblemDetails)
	})
}
