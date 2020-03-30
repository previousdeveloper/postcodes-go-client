package model
type BulkPostCodeResponse struct {
	Status int `json:"status"`
	Result []struct {
		Query  string `json:"query"`
		Result struct {
			Postcode                  string  `json:"postcode"`
			Quality                   int     `json:"quality"`
			Eastings                  int     `json:"eastings"`
			Northings                 int     `json:"northings"`
			Country                   string  `json:"country"`
			NhsHa                     string  `json:"nhs_ha"`
			Longitude                 float64 `json:"longitude"`
			Latitude                  float64 `json:"latitude"`
			EuropeanElectoralRegion   string  `json:"european_electoral_region"`
			PrimaryCareTrust          string  `json:"primary_care_trust"`
			Region                    string  `json:"region"`
			Lsoa                      string  `json:"lsoa"`
			Msoa                      string  `json:"msoa"`
			Incode                    string  `json:"incode"`
			Outcode                   string  `json:"outcode"`
			ParliamentaryConstituency string  `json:"parliamentary_constituency"`
			AdminDistrict             string  `json:"admin_district"`
			Parish                    string  `json:"parish"`
			AdminCounty               string  `json:"admin_county"`
			AdminWard                 string  `json:"admin_ward"`
			Ced                       string  `json:"ced"`
			Ccg                       string  `json:"ccg"`
			Nuts                      string  `json:"nuts"`
			Codes                     struct {
				AdminDistrict             string `json:"admin_district"`
				AdminCounty               string `json:"admin_county"`
				AdminWard                 string `json:"admin_ward"`
				Parish                    string `json:"parish"`
				ParliamentaryConstituency string `json:"parliamentary_constituency"`
				Ccg                       string `json:"ccg"`
				CcgID                     string `json:"ccg_id"`
				Ced                       string `json:"ced"`
				Nuts                      string `json:"nuts"`
			} `json:"codes"`
		} `json:"result"`
	} `json:"result"`
}
