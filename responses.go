package moogle

// Strings
const missingAPIKey = "You must use an API key to authenticate each request to Google Maps Platform APIs. For additional information, please refer to http://g.co/dev/maps-no-account"
const geocodeInvalidRequest = "Invalid request. Missing the 'address', 'components', 'latlng' or 'place_id' parameter."

var MATRIX_QUERY_LIMIT = MatrixResponse{
	DestinationAddresses: []string{},
	OriginAddresses:      []string{},
	Rows:                 []DistanceRow{},
	Status:               OverQueryLimit,
}

var MATRIX_DENIED = MatrixResponse{
	DestinationAddresses: []string{},
	OriginAddresses:      []string{},
	ErrorMessage:         missingAPIKey,
	Rows:                 []DistanceRow{},
	Status:               RequestDenied,
}

var GEOCODE_QUERY_LIMIT = GeocodeResponse{
	ErrorMessage: missingAPIKey,
	Results:      []AddressResult{},
	Status:       OverQueryLimit,
}

var GEOCODE_DENIED = GeocodeResponse{
	ErrorMessage: missingAPIKey,
	Results:      []AddressResult{},
	Status:       RequestDenied,
}

var GEOCODE_INVALID_REQUEST = GeocodeResponse{
	ErrorMessage: geocodeInvalidRequest,
	Results:      []AddressResult{},
	Status:       InvalidRequest,
}

var GEOCODE_ZERO_RESULTS = GeocodeResponse{
	Results: []AddressResult{},
	Status:  QueryZeroResults,
}
