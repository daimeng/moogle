package moogle

var OVER_QUERY_LIMIT = MatrixResponse{
	Status: OverQueryLimit,
}

var MATRIX_NO_KEY = MatrixResponse{
	ErrorMessage: "You must use an API key to authenticate each request to Google Maps Platform APIs. For additional information, please refer to http://g.co/dev/maps-no-account",
	Status:       RequestDenied,
}

var GEOCODE_INVALID_REQUEST = GeocodeResponse{
	ErrorMessage: "Invalid request. Missing the 'address', 'components', 'latlng' or 'place_id' parameter.",
	Results:      []AddressResult{},
	Status:       InvalidRequest,
}

var GEOCODE_ZERO_RESULTS = GeocodeResponse{
	Results: []AddressResult{},
	Status:  QueryZeroResults,
}
