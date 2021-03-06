package servicenow

import "net/url"

// TableChoice defines the name of the table withing the JSONv2 web service to interface with
// SNOW CMDB
const TableChoice = "sys_choice.do"

// GetChoiceItems method will take a url.Value type argument and call the GetRecordsFor method with
// the sys_choice table and query as the arguments, then format the response into a list of Choice type
func (c Client) GetChoiceItems(query url.Values) ([]Choice, error) {
	var res struct {
		Records []Choice
	}
	err := c.GetRecordsFor(TableChoice, query, &res)
	return res.Records, err
}

type Choice struct {
	DependentValue string `json:"dependent_value"`
	SysModCount    string `json:"sys_mod_count"`
	Language       string `json:"language"`
	Label          string `json:"label"`
	SysUpdatedOn   string `json:"sys_updated_on"`
	SysDomainPath  string `json:"sys_domain_path"`
	SysTags        string `json:"sys_tags"`
	Sequence       string `json:"sequence"`
	SysID          string `json:"sys_id"`
	Inactive       string `json:"inactive"`
	SysUpdatedBy   string `json:"sys_updated_by"`
	SysCreatedOn   string `json:"sys_created_on"`
	Hint           string `json:"hint"`
	SysDomain      string `json:"sys_domain"`
	Name           string `json:"name"`
	Value          string `json:"value"`
	SysCreatedBy   string `json:"sys_created_by"`
	Element        string `json:"element"`
	ResponseStatus string `json:"__status"`
}