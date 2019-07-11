package servicenow

import "net/url"

// TableLocation defines the name of the table withing the JSONv2 web service to interface with
// SNOW CMDB
const TableLocation = "cmn_location.do"

// GetLocationItems method will take a url.Value type argument and call the GetRecordsFor method with
// the cmn_location table and query as the arguments, then format the response into a list of Location types
func (c Client) GetLocationItems(query url.Values) ([]Location, error) {
	var res struct {
		Records []Location
	}
	err := c.GetRecordsFor(TableLocation, query, &res)
	return res.Records, err
}

type Location struct {
	Country        string `json:"country"`
	Parent         string `json:"parent"`
	City           string `json:"city"`
	Latitude       string `json:"latitude"`
	SysUpdatedOn   string `json:"sys_updated_on"`
	SysID          string `json:"sys_id"`
	SysUpdatedBy   string `json:"sys_updated_by"`
	StockRoom      string `json:"stock_room"`
	Street         string `json:"street"`
	SysCreatedOn   string `json:"sys_created_on"`
	Contact        string `json:"contact"`
	PhoneTerritory string `json:"phone_territory"`
	Company        string `json:"company"`
	LatLongError   string `json:"lat_long_error"`
	State          string `json:"state"`
	SysCreatedBy   string `json:"sys_created_by"`
	Longitude      string `json:"longitude"`
	ResponseStatus string `json:"___status"`
	Zip            string `json:"zip"`
	SysModCount    string `json:"sys_mod_count"`
	SysTags        string `json:"sys_tags"`
	TimeZone       string `json:"time_zone"`
	FullName       string `json:"full_name"`
	FaxPhone       string `json:"fax_phone"`
	Phone          string `json:"phone"`
	Name           string `json:"name"`
}