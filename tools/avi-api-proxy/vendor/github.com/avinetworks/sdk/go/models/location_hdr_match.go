package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// LocationHdrMatch location hdr match
// swagger:model LocationHdrMatch
type LocationHdrMatch struct {

	// Case sensitivity to use for the match. Enum options - SENSITIVE, INSENSITIVE.
	MatchCase *string `json:"match_case,omitempty"`

	// Criterion to use for matching location header value in the HTTP response. Enum options - HDR_EXISTS, HDR_DOES_NOT_EXIST, HDR_BEGINS_WITH, HDR_DOES_NOT_BEGIN_WITH, HDR_CONTAINS, HDR_DOES_NOT_CONTAIN, HDR_ENDS_WITH, HDR_DOES_NOT_END_WITH, HDR_EQUALS, HDR_DOES_NOT_EQUAL.
	// Required: true
	MatchCriteria *string `json:"match_criteria"`

	// String value(s) in the location header.
	Value []string `json:"value,omitempty"`
}
