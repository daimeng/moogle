package moogle

import (
	"bytes"
	"encoding/json"
)

type QueryStatus int

const (
	QueryOk QueryStatus = iota
	QueryZeroResults
	OverDailyLimit
	OverQueryLimit
	MaxElementsExceeded
	RequestDenied
	InvalidRequest
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

func (s QueryStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(queryStatusToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *QueryStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*s = queryStatusToEnum[j]
	return nil
}

type ElmStatus int

const (
	ElmZeroResults ElmStatus = iota
	ElmOk
	NotFound
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
