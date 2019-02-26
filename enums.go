package moogle

import (
	"bytes"
	"encoding/json"
)

// QueryStatus is top level query status
type QueryStatus int

const (
	// QueryOk indicates the response contains a valid result.
	QueryOk QueryStatus = iota
	// QueryZeroResults indicates that the geocode was successful but returned no results.
	// This may occur if the geocoder was passed a non-existent address.
	QueryZeroResults
	// OverDailyLimit indicates any of the following:
	// * The API key is missing or invalid.
	// * Billing has not been enabled on your account.
	// * A self-imposed usage cap has been exceeded.
	// * The provided method of payment is no longer valid (for example, a credit card has expired).
	OverDailyLimit
	// OverQueryLimit indicates the service has received too many requests from your application within the allowed time period.
	OverQueryLimit
	// MaxElementsExceeded indicates that the product of origins and destinations exceeds the per-query limit.
	MaxElementsExceeded
	// RequestDenied indicates that the service denied use of the Distance Matrix service by your application.
	RequestDenied
	// InvalidRequest indicates that the provided request was invalid.
	InvalidRequest
	// UnknownError indicates the request could not be processed due to a server error.
	// The request may succeed if you try again.
	UnknownError
)

var queryStatusToString = map[QueryStatus]string{
	QueryOk:             "OK",
	QueryZeroResults:    "ZERO_RESULTS",
	OverDailyLimit:      "OVER_DAILY_LIMIT",
	OverQueryLimit:      "OVER_QUERY_LIMIT",
	MaxElementsExceeded: "MAX_ELEMENTS_EXCEEDED",
	RequestDenied:       "REQUEST_DENIED",
	InvalidRequest:      "INVALID_REQUEST",
	UnknownError:        "UNKNOWN_ERROR",
}

var queryStatusToEnum = map[string]QueryStatus{
	"OK":                    QueryOk,
	"ZERO_RESULTS":          QueryZeroResults,
	"OVER_DAILY_LIMIT":      OverDailyLimit,
	"OVER_QUERY_LIMIT":      OverQueryLimit,
	"MAX_ELEMENTS_EXCEEDED": MaxElementsExceeded,
	"REQUEST_DENIED":        RequestDenied,
	"INVALID_REQUEST":       InvalidRequest,
	"UNKNOWN_ERROR":         UnknownError,
}

func (s QueryStatus) String() string {
	return queryStatusToString[s]
}

// MarshalJSON QueryStatus to JSON
func (s QueryStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(queryStatusToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON QueryStatus from JSON
func (s *QueryStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*s = queryStatusToEnum[j]
	return nil
}

// ElmStatus is Element level distance matrix status
type ElmStatus int

const (
	// ElmZeroResults indicates that the origin and/or destination of this pairing could not be geocoded.
	ElmZeroResults ElmStatus = iota
	// ElmOk indicates the response contains a valid result.
	ElmOk
	// NotFound indicates no route could be found between the origin and destination.
	NotFound
	// MaxRouteLengthExceeded indicates the requested route is too long and cannot be processed.
	MaxRouteLengthExceeded
)

var elmStatusToString = map[ElmStatus]string{
	ElmOk:                  "OK",
	ElmZeroResults:         "ZERO_RESULTS",
	NotFound:               "NOT_FOUND",
	MaxRouteLengthExceeded: "MAX_ROUTE_LENGTH_EXCEEDED",
}

var elmStatusToEnum = map[string]ElmStatus{
	"OK":                        ElmOk,
	"ZERO_RESULTS":              ElmZeroResults,
	"NOT_FOUND":                 NotFound,
	"MAX_ROUTE_LENGTH_EXCEEDED": MaxRouteLengthExceeded,
}

func (s ElmStatus) String() string {
	return elmStatusToString[s]
}

func (s ElmStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(elmStatusToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *ElmStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*s = elmStatusToEnum[j]
	return nil
}
