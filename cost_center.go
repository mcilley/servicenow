package servicenow

import "net/url"

// TableCostCenter defines the name of the table withing the JSONv2 web service to interface with
// SNOW CMDB
const TableCostCenter = "cmn_cost_center.do"

// GetCostCenterItems method will take a url.Value type argument and call the GetRecordsFor method with
// the cmn_cost_center table and query as the arguments, then format the response into a list of CostCenter types
func (c Client) GetCostCenterItems(query url.Values) ([]CostCenter, error) {
	var res struct {
		Records []CostCenter
	}
	err := c.GetRecordsFor(TableCostCenter, query, &res)
	return res.Records, err
}

type CostCenter struct {
	AccountNumber  string `json:"account_number"`
	Parent         string `json:"parent"`
	Code           string `json:"code"`
	Manager        string `json:"manager"`
	SysModCount    string `json:"sys_mod_count"`
	ValidFrom      string `json:"valid_from"`
	SysUpdatedOn   string `json:"sys_updated_on"`
	SysTags        string `json:"sys_tags"`
	SysID          string `json:"sys_id"`
	SysUpdatedBy   string `json:"sys_updated_by"`
	SysCreatedOn   string `json:"sys_created_on"`
	ValidTo        string `json:"valid_to"`
	Name           string `json:"name"`
	Location       string `json:"location"`
	SysCreatedBy   string `json:"sys_created_by"`
	ResponseStatus string `json:"__status"`
}