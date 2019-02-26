package moogle

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type TextedInt struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type DistanceElm struct {
	Distance *TextedInt `json:"distance,omitempty"`
	Duration *TextedInt `json:"duration,omitempty"`
	Status   ElmStatus  `json:"status"`
}

type DistanceRow struct {
	Elements []DistanceElm `json:"elements"`
}

type MatrixResponse struct {
	DestinationAddresses []string      `json:"destination_addresses"`
	ErrorMessage         string        `json:"error_message,omitempty"`
	OriginAddresses      []string      `json:"origin_addresses"`
	Rows                 []DistanceRow `json:"rows"`
	Status               QueryStatus   `json:"status"`
}

type AddressComponent struct {
}

type GeocodeGeometry struct {
}

type AddressPlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}

type AddressResult struct {
	AddressComponents []AddressComponent `json:"address_components"`
	FormattedAddress  string             `json:"formatted_address"`
	Geometry          *GeocodeGeometry   `json:"geometry"`
	PartialMatch      bool               `json:"partial_match,omitempty"`
	PlaceID           string             `json:"place_id,omitempty"`
	PlusCode          string             `json:"plus_code,omitempty"`
	Types             []string           `json:"types"`
}

type GeocodeResponse struct {
	ErrorMessage string          `json:"error_message,omitempty"`
	Results      []AddressResult `json:"results"`
	Status       QueryStatus     `json:"status"`
}

type Box struct {
	Northeast Point `json:"northeast"`
	Southwest Point `json:"southwest"`
}
