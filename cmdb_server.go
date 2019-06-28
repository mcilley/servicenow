package servicenow

import "net/url"

// TableCMDB defines the name of the table withing the JSONv2 web service to interface with
// SNOW CMDB
const TableCMDBServer = "cmdb_ci_server"


// GetCMDBItems method will take a url.Value type argument and call the GetRecordsFor method with
// the cmdb_ci_server table and query as the arguments, then format the response into a list of CMDBServer types
func (c Client) GetCMDBServers(query url.Values) ([]CMDBServer, error) {
	var res struct {
		Records []CMDBServer
	}
	err := c.GetRecordsFor(TableCMDBServer, query, &res)
	return res.Records, err
}

// CMDBServer is a struct type to define the formatted response received from the JSONv2
// ServiceNow Web Service
type CMDBServer struct {
	ResponseStatus      string `json:"__status"`
	Asset               string `json:"asset"`
	AssetTag            string `json:"asset_tag"`
	Assigned            string `json:"assigned"`
	AssignedTo          string `json:"assigned_to"`
	AssignmentGroup     string `json:"assignment_group"`
	Attributes          string `json:"attributes"`
	CanPrint            string `json:"can_print"`
	Category            string `json:"category"`
	CdRom               string `json:"cd_rom"`
	CdSpeed             string `json:"cd_speed"`
	ChangeControl       string `json:"change_control"`
	ChassisType         string `json:"chassis_type"`
	CheckedIn           string `json:"checked_in"`
	CheckedOut          string `json:"checked_out"`
	Classification      string `json:"classification"`
	Comments            string `json:"comments"`
	Company             string `json:"company"`
	CorrelationID       string `json:"correlation_id"`
	Cost                string `json:"cost"`
	CostCc              string `json:"cost_cc"`
	CostCenter          string `json:"cost_center"`
	CpuCoreCount        string `json:"cpu_core_count"`
	CpuCoreThread       string `json:"cpu_core_thread"`
	CpuCount            string `json:"cpu_count"`
	CpuManufacturer     string `json:"cpu_manufacturer"`
	CpuName             string `json:"cpu_name"`
	CpuSpeed            string `json:"cpu_speed"`
	CpuType             string `json:"cpu_type"`
	DefaultGateway      string `json:"default_gateway"`
	DeliveryDate        string `json:"delivery_date"`
	Department          string `json:"department"`
	DiscoverySource     string `json:"discovery_source"`
	DiskSpace           string `json:"disk_space"`
	DnsDomain           string `json:"dns_domain"`
	DrBackup            string `json:"dr_backup"`
	Due                 string `json:"due"`
	DueIn               string `json:"due_in"`
	FaultCount          string `json:"fault_count"`
	FirewallStatus      string `json:"firewall_status"`
	FirstDiscovered     string `json:"first_discovered"`
	Floppy              string `json:"floppy"`
	FormFactor          string `json:"form_factor"`
	Fqdn                string `json:"fqdn"`
	GlAccount           string `json:"gl_account"`
	HardwareStatus      string `json:"hardware_status"`
	HardwareSubstatus   string `json:"hardware_substatus"`
	HostName            string `json:"host_name"`
	InstallDate         string `json:"install_date"`
	InstallStatus       string `json:"install_status"`
	InvoiceNumber       string `json:"invoice_number"`
	IPAddress           string `json:"ip_address"`
	Justification       string `json:"justification"`
	LastDiscovered      string `json:"last_discovered"`
	LeaseID             string `json:"lease_id"`
	Location            string `json:"location"`
	MacAddress          string `json:"mac_address"`
	MaintenanceSchedule string `json:"maintenance_schedule"`
	ManagedBy           string `json:"managed_by"`
	Manufacturer        string `json:"manufacturer"`
	ModelID             string `json:"model_id"`
	ModelNumber         string `json:"model_number"`
	Monitor             string `json:"monitor"`
	Name                string `json:"name"`
	ObjectID            string `json:"object_id"`
	OperationalStatus   string `json:"operational_status"`
	OrderDate           string `json:"order_date"`
	Os                  string `json:"os"`
	OsAddressWidth      string `json:"os_address_width"`
	OsDomain            string `json:"os_domain"`
	OsServicePack       string `json:"os_service_pack"`
	OsVersion           string `json:"os_version"`
	OwnedBy             string `json:"owned_by"`
	PoNumber            string `json:"po_number"`
	PurchaseDate        string `json:"purchase_date"`
	Ram                 string `json:"ram"`
	Schedule            string `json:"schedule"`
	SerialNumber        string `json:"serial_number"`
	ShortDescription    string `json:"short_description"`
	SkipSync            string `json:"skip_sync"`
	StartDate           string `json:"start_date"`
	Subcategory         string `json:"subcategory"`
	SupportGroup        string `json:"support_group"`
	SupportedBy         string `json:"supported_by"`
	SysClassName        string `json:"sys_class_name"`
	SysClassPath        string `json:"sys_class_path"`
	SysCreatedBy        string `json:"sys_created_by"`
	SysCreatedOn        string `json:"sys_created_on"`
	SysDomain           string `json:"sys_domain"`
	SysDomainPath       string `json:"sys_domain_path"`
	SysID               string `json:"sys_id"`
	SysModCount         string `json:"sys_mod_count"`
	SysTags             string `json:"sys_tags"`
	SysUpdatedBy        string `json:"sys_updated_by"`
	SysUpdatedOn        string `json:"sys_updated_on"`
	Unverified          string `json:"unverified"`
	UsedFor             string `json:"used_for"`
	Vendor              string `json:"vendor"`
	Virtual             string `json:"virtual"`
	WarrantyExpiration  string `json:"warranty_expiration"`
}
