package dpfm_api_output_formatter

type District struct {
	District     string              `json:"District"`
	Country      string              `json:"Country"`
	DistrictText DistrictText `json:"DistrictText"`
}

type DistrictText struct {
	District     string `json:"District"`
	Language     string `json:"Language"`
	DistrictName string `json:"DistrictName"`
}
