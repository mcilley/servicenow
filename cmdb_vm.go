package servicenow

import "net/url"

// TableCMDBVm defines the name of the table withing the JSONv2 web service to interface with
// SNOW CMDB
const TableCMDBVm = "cmdb_ci_vmware_instance"

// GetCMDBVmItems method will take a url.Value type argument and call the GetRecordsFor method with
// the cmdb_ci_vmware_instance table and query as the arguments, then format the response into a list of CMDBVm types
func (c Client) GetCMDBVmItems(query url.Values) ([]CMDBVm, error) {
	var res struct {
		Records []CMDBVm
	}
	err := c.GetRecordsFor(TableCMDBVm, query, &res)
	return res.Records, err
}

// CMDBServer is a struct type to define the formatted response received from the JSONv2
// ServiceNow Web Service
type CMDBVm struct {
	ResponseStatus        string `json:"__status"`
	Asset                 string `json:"asset"`
	AssetTag              string `json:"asset_tag"`
	Assigned              string `json:"assigned"`
	AssignedTo            string `json:"assigned_to"`
	AssignmentGroup       string `json:"assignment_group"`
	Attributes            string `json:"attributes"`
	BiosUuid              string `json:"bios_uuid"`
	CanPrint              string `json:"can_print"`
	Category              string `json:"category"`
	ChangeControl         string `json:"change_control"`
	CheckedIn             string `json:"checked_in"`
	CheckedOut            string `json:"checked_out"`
	Comments              string `json:"comments"`
	Company               string `json:"company"`
	CorrelationID         string `json:"correlation_id"`
	Cost                  string `json:"cost"`
	CostCc                string `json:"cost_cc"`
	Cpus                  string `json:"cpus"`
	DeliveryDate          string `json:"delivery_date"`
	Department            string `json:"department"`
	DiscoverySource       string `json:"discovery_source"`
	Disks                 string `json:"disks"`
	DisksSize             string `json:"disks_size"`
	DnsDomain             string `json:"dns_domain"`
	Due                   string `json:"due"`
	DueIn                 string `json:"due_in"`
	FaultCount            string `json:"fault_count"`
	FirstDiscovered       string `json:"first_discovered"`
	Fqdn                  string `json:"fqdn"`
	GlAccount             string `json:"gl_account"`
	GuestID               string `json:"guest_id"`
	GuestReconciled       string `json:"guest_reconciled"`
	ImagePath             string `json:"image_path"`
	InstallDate           string `json:"install_date"`
	InstallStatus         string `json:"install_status"`
	InvoiceNumber         string `json:"invoice_number"`
	IPAddress             string `json:"ip_address"`
	Justification         string `json:"justification"`
	LastDiscovered        string `json:"last_discovered"`
	LeaseID               string `json:"lease_id"`
	Location              string `json:"location"`
	MacAddress            string `json:"mac_address"`
	MaintenanceSchedule   string `json:"maintenance_schedule"`
	ManagedBy             string `json:"managed_by"`
	Manufacturer          string `json:"manufacturer"`
	Memory                string `json:"memory"`
	ModelID               string `json:"model_id"`
	ModuleNumber          string `json:"model_number"`
	Monitor               string `json:"monitor"`
	Name                  string `json:"name"`
	Nics                  string `json:"nics"`
	ObectID               string `json:"object_id"`
	OperationalStatus     string `json:"operational_status"`
	OrderDate             string `json:"order_date"`
	OwnedBy               string `json:"owned_by"`
	PoNumber              string `json:"po_number"`
	PurchaseDate          string `json:"purchase_date"`
	Schedule              string `json:"schedule"`
	SerialNumber          string `json:"serial_number"`
	Server                string `json:"server"`
	ShortDescription      string `json:"short_description"`
	SkipSync              string `json:"skip_sync"`
	StartDate             string `json:"start_date"`
	State                 string `json:"state"`
	SubCategory           string `json:"subcategory"`
	SupportGroup          string `json:"support_group"`
	SupportedBy           string `json:"supported_by"`
	SysClassName          string `json:"sys_class_name"`
	SysClassPath          string `json:"sys_class_path"`
	SysCreatedBy          string `json:"sys_created_by"`
	SysCreatedOn          string `json:"sys_created_on"`
	SysDomain             string `json:"sys_domain"`
	SysDomainPath         string `json:"sys_domain_path"`
	SysID                 string `json:"sys_id"`
	SysModCount           string `json:"sys_mod_count"`
	SysTags               string `json:"sys_tags"`
	SysUpdatedBy          string `json:"sys_updated_by"`
	SysUpdatedOn          string `json:"sys_updated_on"`
	Template              string `json:"template"`
	TerminatedOn          string `json:"terminated_on"`
	TerminationProtection string `json:"termination_protection"`
	ULastDrfttDate        string `json:"u_last_drftt_date"`
	ULastDrfttResult      string `json:"u_last_drftt_result"`
	USoxInScope           string `json:"u_sox_in_scope"`
	USubtype              string `json:"u_subtype"`
	UnVerified            string `json:"unverified"`
	VcenterRef            string `json:"vcenter_ref"`
	VcenterUuid           string `json:"vcenter_uuid"`
	Vendor                string `json:"vendor"`
	VmInstID              string `json:"vm_inst_id"`
	VmInstanceUuid        string `json:"vm_instance_uuid"`
	WarrantyExpiration    string `json:"warranty_expiration"`
}
