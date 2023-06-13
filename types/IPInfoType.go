package types

type IPInfoType struct {
	IP          string `json:"ip"`
	FullIP      string `json:"full_ip"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Province    string `json:"province"`
	City        string `json:"city"`
	Distinct    string `json:"distinct"`
	Isp         string `json:"isp"`
	Operator    string `json:"operator"`
	Lon         string `json:"lon"`
	Lat         string `json:"lat"`
	NetStr      string `json:"net_str"`
}
