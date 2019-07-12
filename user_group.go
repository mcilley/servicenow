package servicenow

import (
	"encoding/json"
	"net/url"
)

// TableUserGroup defines the name of the table withing the JSONv2 web service to interface with
// SNOW TASKs
const TableUserGroup = "sys_user_group"

// GetUserGroupRequests method will take a url.Value type argument and call the GetRecordsFor method with
// the user_group table and query as the arguments, then format the response into the UserGroupRequest type
func (c Client) GetUserGroupRequests(query url.Values) ([]UserGroupRequest, error) {
	var res struct {
		Records []UserGroupRequest `json:"records"`
	}
	err := c.GetRecordsFor(TableUserGroup, query, &res)
	return res.Records, err
}

// UserGroupRequest is a struct type to define the formatted response received from the JSONv2
// ServiceNow Web Service
type UserGroupRequest struct {
	Active          bool        `json:"active,string"`
	Company         string      `json:"company"`
	Description     string      `json:"description"`
	DefaultAssignee string      `json:"default_assignee"`
	Email           string      `json:"email"`
	ExcludeManager  bool        `json:"exclude_manager,string"`
	HourlyRate      json.Number `json:"hourly_rate"`
	IncludeMembers  bool        `json:"include_members,string"`
	Manager         string      `json:"manager"`
	ManagedDomain   bool        `json:"managed_domain,string"`
	Name            string      `json:"name"`
	Parent          string      `json:"parent"`
	Points          json.Number `json:"points"`
	Source          string      `json:"source"`
	SysCreatedBy    string      `json:"sys_created_by"`
	SysCreatedOn    SNTime      `json:"sys_created_on"`
	SysDomain       string      `json:"sys_domain"`
	SysDomainPath   string      `json:"sys_domain_path"`
	SysID           string      `json:"sys_id"`
	SysModCount     json.Number `json:"sys_mod_count"`
	SysTags         string      `json:"sys_tags"`
	SysUpdatedBy    string      `json:"sys_updated_by"`
	SysUpdatedOn    SNTime      `json:"sys_updated_on"`
	Type            string      `json:"type"`
	UDepartment     string      `json:"u_department"`
	UCorrelationID  string      `json:"u_correlation_id"`
}
