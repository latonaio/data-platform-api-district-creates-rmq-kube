package dpfm_api_input_reader

import (
	"data-platform-api-district-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToDistrict() *requests.District {
	data := sdc.District
	return &requests.District{
		District: data.District,
		Country:  data.Country,
	}
}

func (sdc *SDC) ConvertToDistrictText() *requests.DistrictText {
	dataDistrict := sdc.District
	data := sdc.District.DistrictText
	return &requests.DistrictText{
		District:     dataDistrict.District,
		Language:     data.Language,
		DistrictName: data.DistrictName,
	}
}
