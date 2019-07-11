package servicenow

import "net/url"

// TableCompany defines the name of the table withing the JSONv2 web service to interface with
// SNOW CMDB
const TableCompany = "core_company.do"

// GetCompanyItems method will take a url.Value type argument and call the GetRecordsFor method with
// the core_company table and query as the arguments, then format the response into a list of Company types
func (c Client) GetCompanyItems(query url.Values) ([]Company, error) {
	var res struct {
		Records []Company
	}
	err := c.GetRecordsFor(TableCompany, query, &res)
	return res.Records, err
}

type Company struct {
	BannerImageLight string `json:"banner_image_light"`
	Country          string `json:"country"`
	Parent           string `json:"parent"`
	Notes            string `json:"notes"`
	City             string `json:"city"`
	StockSymbol      string `json:"stock_symbol"`
	Latitude         string `json:"latitude"`
	Discount         string `json:"discount"`
	SysUpdatedOn     string `json:"sys_updated_on"`
	SysClassName     string `json:"sys_class_name"`
	Manufacturer     string `json:"manufacturer"`
	AppleIcon        string `json:"apple_icon"`
	SysID            string `json:"sys_id"`
	MarketCap        string `json:"market_cap"`
	SysUpdatedBy     string `json:"sys_updated_by"`
	NumEmployees     string `json:"num_employees"`
	FiscalYear       string `json:"fiscal_year"`
	RankTier         string `json:"rank_tier"`
	Street           string `json:"street"`
	SysCreatedOn     string `json:"sys_created_on"`
	Vendor           string `json:"vendor"`
	Contact          string `json:"contact"`
	LatLongError     string `json:"lat_long_error"`
	StockPrice       string `json:"stock_price"`
	Theme            string `json:"theme"`
	BannerImage      string `json:"banner_image"`
	State            string `json:"state"`
	SysCreatedBy     string `json:"sys_created_by"`
	Longitude        string `json:"longitude"`
	VendorType       string `json:"vendor_type"`
	ResponseStatus   string `json:"__status"`
	Zip              string `json:"zip"`
	Profits          string `json:"profits"`
	RevenuePerYear   string `json:"revenue_per_year"`
	Website          string `json:"website"`
	PubliclyTraded   string `json:"publicly_traded"`
	SysModCount      string `json:"sys_mod_count"`
	SysTags          string `json:"sys_tags"`
	FaxPhone         string `json:"fax_phone"`
	Phone            string `json:"phone"`
	VendorManager    string `json:"vendor_manager"`
	BannerText       string `json:"banner_text"`
	Name             string `json:"name"`
	Customer         string `json:"customer"`
	Primary          string `json:"primary"`
}