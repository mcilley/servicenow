package servicenow

import "net/url"

const TableScheduleMaint = "cmn_schedule_maintenance.do"

func (c Client) GetScheduleMaintItems(query url.Values) ([]ScheduleMaint, error) {
	var res struct {
		Records []ScheduleMaint
	}
	err := c.GetRecordsFor(TableScheduleMaint, query, &res)
	return res.Records, err
}

type ScheduleMaint struct {
	Parent         string `json:"parent"`
	Document       string `json:"document"`
	Description    string `json:"description"`
	Source         string `json:"source"`
	SysUpdatedOn   string `json:"sys_updated_on"`
	Type           string `json:"type"`
	DocumentKey    string `json:"document_key"`
	SysClassName   string `json:"sys_class_name"`
	SysID          string `json:"sys_id"`
	SysUpdatedBy   string `json:"sys_updated_by"`
	AppliesTo      string `json:"applies_to"`
	ReadOnly       string `json:"read_only"`
	SysCreatedOn   string `json:"sys_created_on"`
	SysName        string `json:"sys_name"`
	SysScope       string `json:"sys_scope"`
	SysCreatedBy   string `json:"sys_created_by"`
	ResponseStatus string `json:"__status"`
	SysModCount    string `json:"sys_mod_count"`
	SysTags        string `json:"sys_tags"`
	TimeZone       string `json:"time_zone"`
	SysPackage     string `json:"sys_package"`
	Condition      string `json:"condition"`
	SysUpdateName  string `json:"sys_update_name"`
	Name           string `json:"name"`
	SysPolicy      string `json:"sys_policy"`
}