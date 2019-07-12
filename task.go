package servicenow

import (
	"encoding/json"
	"net/url"
)

// TableTasks defines the name of the table withing the JSONv2 web service to interface with
// SNOW TASKs
const TableTasks = "sc_task"

// GetTasksRequests method will take a url.Value type argument and call the GetRecordsFor method with
// the task table and query as the arguments, then format the response into the TaskRequest type
func (c Client) GetTasksRequests(query url.Values) ([]TaskRequest, error) {
	var res struct {
		Records []TaskRequest `json:"records"`
	}
	err := c.GetRecordsFor(TableTasks, query, &res)
	return res.Records, err
}

// CreateTask method will take a TaskRequest type argument and call the Insert method with
// the task table and task obj as the arguments, then format the response into the TaskRequest type
func (c Client) CreateTask(task TaskRequest) ([]TaskRequest, error) {
	var res struct {
		Records []TaskRequest `json:"records"`
	}
	err := c.Insert(TableTasks, task, &res)
	return res.Records, err
}

// UpdateTask method will take a url.Value type argument and call the GetRecordsFor method with
// the task table and query as the arguments, then format the response into the TaskRequest type
func (c Client) UpdateTask(query url.Values, update map[string]interface{}) ([]TaskRequest, error) {
	var res struct {
		Records []TaskRequest
	}
	err := c.Update(TableTasks, query, update, &res)
	return res.Records, err
}

// TaskRequest is a struct type to define the formatted response received from the JSONv2
// ServiceNow Web Service
type TaskRequest struct {
	Active                        bool        `json:"active,string"`
	ActivityDue                   string      `json:"activity_due"`
	AdditionalAssigneeList        string      `json:"additional_assignee_list"`
	Approval                      string      `json:"approval"`
	ApprovalHistory               string      `json:"approval_history"`
	ApprovalSet                   string      `json:"approval_set"`
	AssignedTo                    string      `json:"assigned_to"`
	AssignmentGroup               string      `json:"assignment_group"`
	BackoutPlan                   string      `json:"backout_plan"`
	BusinessDuration              string      `json:"business_duration"`
	BusinessService               string      `json:"business_service"`
	CabDate                       SNTime      `json:"cab_date"`
	CabDelegate                   string      `json:"cab_delegate"`
	CabRecommendation             string      `json:"cab_recommendation"`
	CabRequired                   bool        `json:"cab_required,string"`
	CalendarDuration              string      `json:"calendar_duration"`
	Category                      string      `json:"category"`
	ChangePlan                    string      `json:"change_plan"`
	ClosedAt                      SNTime      `json:"closed_at"`
	ClosedBy                      string      `json:"closed_by"`
	CloseNotes                    string      `json:"close_notes"`
	CMDBCi                        string      `json:"cmdb_ci"`
	Comments                      string      `json:"comments"`
	CommentsAndWorkNotes          string      `json:"comments_and_work_notes"`
	Company                       string      `json:"company"`
	ConflictLastRun               string      `json:"conflict_last_run"`
	ConflictStatus                string      `json:"conflict_status"`
	ContactType                   string      `json:"contact_type"`
	CorrelationDisplay            string      `json:"correlation_display"`
	CorrelationID                 string      `json:"correlation_id"`
	Description                   string      `json:"description"`
	DueDate                       SNTime      `json:"due_date"`
	EndDate                       SNTime      `json:"end_date"`
	Escalation                    json.Number `json:"escalation"`
	ExpectedStart                 SNTime      `json:"expected_start"`
	FollowUp                      string      `json:"follow_up"`
	GroupList                     string      `json:"group_list"`
	ImplementationPlan            string      `json:"implementation_plan"`
	Justification                 string      `json:"justification"`
	Knowledge                     bool        `json:"knowledge,string"`
	Location                      string      `json:"location"`
	MadeSLA                       bool        `json:"made_sla,string"`
	Number                        string      `json:"number"`
	OpenedAt                      SNTime      `json:"opened_at"`
	OpenedBy                      string      `json:"opened_by"`
	Order                         string      `json:"order"`
	OutsideMaintenanceSchedule    bool        `json:"outside_maintenance_schedule,string"`
	Parent                        string      `json:"parent"`
	Phase                         string      `json:"phase"`
	PhaseState                    string      `json:"phase_state"`
	Priority                      json.Number `json:"priority"`
	ProductionSystem              bool        `json:"production_system,string"`
	Reason                        string      `json:"reason"`
	ReassignmentCount             json.Number `json:"reassignment_count"`
	RequestedBy                   string      `json:"requested_by"`
	RequestedByDate               SNTime      `json:"requested_by_date"`
	ReviewComments                string      `json:"review_comments"`
	ReviewDate                    SNTime      `json:"review_date"`
	ReviewStatus                  string      `json:"review_status"`
	Risk                          json.Number `json:"risk"`
	Scope                         json.Number `json:"scope"`
	ShortDescription              string      `json:"short_description"`
	Skills                        string      `json:"skills"`
	SLADue                        string      `json:"sla_due"`
	StartDate                     SNTime      `json:"start_date"`
	State                         string      `json:"state"`
	Status                        string      `json:"__status"`
	SysClassName                  string      `json:"sys_class_name"`
	SysCreatedBy                  string      `json:"sys_created_by"`
	SysCreatedOn                  SNTime      `json:"sys_created_on"`
	SysDomain                     string      `json:"sys_domain"`
	SysDomainPath                 string      `json:"sys_domain_path"`
	SysID                         string      `json:"sys_id"`
	SysModCount                   json.Number `json:"sys_mod_count"`
	SysTags                       string      `json:"sys_tags"`
	SysUpdatedBy                  string      `json:"sys_updated_by"`
	SysUpdatedOn                  SNTime      `json:"sys_updated_on"`
	TaskFor                       string      `json:"task_for"`
	TestPlan                      string      `json:"test_plan"`
	TimeWorked                    string      `json:"time_worked"`
	Type                          string      `json:"type"`
	UAccountManager               string      `json:"u_account_manager"`
	UApproval                     string      `json:"u_approval"`
	UApprovalHistory              string      `json:"u_approval_history"`
	UBbrSalesRef                  string      `json:"u_bbr_sales_ref"`
	UBillable                     string      `json:"u_billable"`
	UBillableStartDate            string      `json:"u_billable_start_date"`
	UCallerID                     string      `json:"u_caller_id"`
	UCallNumber                   string      `json:"u_call_number"`
	UChangeCategory               string      `json:"u_change_category"`
	UChangeType                   string      `json:"u_change_type"`
	UCiContacts                   string      `json:"u_ci_contacts"`
	UCiNotAvailable               string      `json:"u_ci_not_available"`
	UCiSelected                   string      `json:"u_ci_selected"`
	UCiType                       string      `json:"u_ci_type"`
	UCloseCode                    string      `json:"u_close_code"`
	UCommercialApproval           string      `json:"u_commercial_approval"`
	UCommercialRequirementsDetail string      `json:"u_commercial_requirements_detail"`
	UCompleted                    string      `json:"u_completed"`
	UCreatedFromCall              string      `json:"u_created_from_call"`
	UCreateProject                string      `json:"u_create_project"`
	UCustomerApprover             string      `json:"u_customer_approver"`
	UCustomerInternal             string      `json:"u_customer_internal"`
	UEscalationComments           string      `json:"u_escalation_comments"`
	UFirstResponse                string      `json:"u_first_response"`
	UFollowUpDate                 string      `json:"u_follow_up_date"`
	UImpact                       string      `json:"u_impact"`
	UImpactRpn                    string      `json:"u_impact_rpn"`
	UIsClientVisible              string      `json:"u_is_client_visible"`
	UITILWatchList                string      `json:"u_itil_watch_list"`
	UJiraID                       string      `json:"u_jira_id"`
	ULikelihood                   string      `json:"u_likelihood"`
	ULoadApproverSelector         string      `json:"u_loadapproverselector"`
	UMonthlyRecurringRevenue      string      `json:"u_monthly_recurring_revenue"`
	UNonRecurringRevenue          string      `json:"u_non_recurring_revenue"`
	UOnBehalfOf                   string      `json:"u_on_behalf_of"`
	UOtherChangeTypeExplain       string      `json:"u_other_change_type_explain"`
	UPcTemplate                   string      `json:"u_pc_template"`
	UponApproval                  string      `json:"upon_approval"`
	UponReject                    string      `json:"upon_reject"`
	UPoNumber                     string      `json:"u_po_number"`
	UPreAgreementJustification    string      `json:"u_pre_agreement_justification"`
	UProvideDetailsCategory       string      `json:"u_provide_details_category"`
	URedundancy                   string      `json:"u_redundancy"`
	URefNumber                    string      `json:"u_refnumber"`
	UResolvedTimeStamp            string      `json:"u_resolved_time_stamp"`
	UReviewed                     string      `json:"u_reviewed"`
	URft                          string      `json:"u_rft"`
	URollback                     string      `json:"u_rollback"`
	URpn                          string      `json:"u_rpn"`
	UserInput                     string      `json:"user_input"`
	UStateReason                  string      `json:"u_state_reason"`
	UTaskEscalation               bool        `json:"u_task_escalation,string"`
	UTaskEscalationComments       string      `json:"u_task_escalation_comments"`
	UTemplate                     string      `json:"u_template"`
	UUnassignedReasonExplain      string      `json:"u_unassigned_reason_explain"`
	UWatchList                    string      `json:"u_watch_list"`
	UZtmpAssigneeBlank            string      `json:"u_ztmpassigneeblank"`
	WatchList                     string      `json:"watch_list"`
	WorkEnd                       string      `json:"work_end"`
	WorkNotes                     string      `json:"work_notes"`
	WorkNotesList                 string      `json:"work_notes_list"`
	WorkStart                     string      `json:"work_start"`
}
