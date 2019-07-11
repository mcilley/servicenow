package servicenow

import "net/url"

// TableDepartment defines the name of the table withing the JSONv2 web service to interface with
// SNOW CMDB
const TableDepartment = "cmn_department.do"

// GetDepartmentItems method will take a url.Value type argument and call the GetRecordsFor method with
// the cmn_department table and query as the arguments, then format the response into a list of Department types
func (c Client) GetDepartmentItems(query url.Values) ([]Department, error) {
	var res struct {
		Records []Department
	}
	err := c.GetRecordsFor(TableDepartment, query, &res)
	return res.Records, err
}

type Department struct {
	Parent         string `json:"parent"`
	SysModCount    string `json:"sys_mod_count"`
	Description    string `json:"description"`
	HeadCount      string `json:"head_count"`
	SysUpdatedOn   string `json:"sys_updated_on"`
	SysTags        string `json:"sys_tags"`
	BusinessUnit   string `json:"business_unit"`
	SysID          string `json:"sys_id"`
	DeptHead       string `json:"dept_head"`
	SysUpdatedBy   string `json:"sys_updated_by"`
	CostCenter     string `json:"cost_center"`
	SysCreatedOn   string `json:"sys_created_on"`
	Name           string `json:"name"`
	Company        string `json:"company"`
	ID             string `json:"id"`
	PrimaryContact string `json:"primary_contact"`
	SysCreatedBy   string `json:"sys_created_by"`
	ResponseStatus string `json:"__status"`
}