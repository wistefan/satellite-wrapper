/*
 * Trusted Issuers Registry
 *
 * The subset of the [Trusted Issuers Registryas defined by EBSI](https://api-pilot.ebsi.eu/docs/apis/trusted-issuers-registry/v4#/) as currently required by the [VCVerifier](https://github.com/FIWARE/VCVerifier).  
 *
 * API version: v3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package generated

type IssuersResponse struct {

	// URI to issuers
	Self string `json:"self"`

	// list of issuers with their decentralized identifier
	Items []IssuerEntry `json:"items"`

	// Total number of items in a collection
	Total float32 `json:"total"`

	// Number of items to be returned per page
	PageSize float32 `json:"pageSize"`

	Links Links `json:"links"`
}

// AssertIssuersResponseRequired checks if the required fields are not zero-ed
func AssertIssuersResponseRequired(obj IssuersResponse) error {
	elements := map[string]interface{}{
		"self": obj.Self,
		"items": obj.Items,
		"total": obj.Total,
		"pageSize": obj.PageSize,
		"links": obj.Links,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Items {
		if err := AssertIssuerEntryRequired(el); err != nil {
			return err
		}
	}
	if err := AssertLinksRequired(obj.Links); err != nil {
		return err
	}
	return nil
}

// AssertRecurseIssuersResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IssuersResponse (e.g. [][]IssuersResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIssuersResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIssuersResponse, ok := obj.(IssuersResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIssuersResponseRequired(aIssuersResponse)
	})
}