package servicenow

import "net/url"

// TableAsset defines the name of the table withing the JSONv2 web service to interface with
// SNOW CMDB
const TableAsset = "alm_asset"

// GetAssetItems method will take a url.Value type argument and call the GetRecordsFor method with
// the alm_asset table and query as the arguments, then format the response into a list of Asset types
func (c Client) GetAssetItems(query url.Values) ([]Model, error) {
	var res struct {
		Records []Model
	}
	err := c.GetRecordsFor(TableModel, query, &res)
	return res.Records, err
}

type Asset struct {
	Parent             string `json:"parent"`
	SkipSync           string `json:"skip_sync"`
	ResidualDate       string `json:"residual_date"`
	Residual           string `json:"residual"`
	SysUpdatedOn       string `json:"sys_updated_on"`
	RequestLine        string `json:"request_line"`
	SysUpdatedBy       string `json:"sys_updated_by"`
	DueIn              string `json:"due_in"`
	ModelCategory      string `json:"model_category"`
	SysCreatedOn       string `json:"sys_created_on"`
	SysDomain          string `json:"sys_domain"`
	DisposalReason     string `json:"disposal_reason"`
	Model              string `json:"model"`
	InstallDate        string `json:"install_date"`
	GlAccount          string `json:"gl_account"`
	InvoiceNumber      string `json:"invoice_number"`
	SysCreatedBy       string `json:"sys_created_by"`
	WarrantyExpiration string `json:"warranty_expiration"`
	ResponseStatus     string `json:"__status"`
	AssetTag           string `json:"asset_tag"`
	DepreciatedAmount  string `json:"depreciated_amount"`
	Substatus          string `json:"substatus"`
	PreAllocated       string `json:"pre_allocated"`
	OwnedBy            string `json:"owned_by"`
	CheckedOut         string `json:"checked_out"`
	DisplayName        string `json:"display_name"`
	SysDomainPath      string `json:"sys_domain_path"`
	DeliveryDate       string `json:"delivery_date"`
	RetirementDate     string `json:"retirement_date"`
	Beneficiary        string `json:"beneficiary"`
	InstallStatus      string `json:"install_status"`
	CostCenter         string `json:"cost_center"`
	SupportedBy        string `json:"supported_by"`
	Assigned           string `json:"assigned"`
	PurchaseDate       string `json:"purchase_date"`
	WorkNotes          string `json:"work_notes"`
	ManagedBy          string `json:"managed_by"`
	SysClassName       string `json:"sys_class_name"`
	SysID              string `json:"sys_id"`
	PoNumber           string `json:"po_number"`
	Stockroom          string `json:"stockroom"`
	CheckedIn          string `json:"checked_in"`
	ResalePrice        string `json:"resale_price"`
	Vendor             string `json:"vendor"`
	Company            string `json:"company"`
	Retired            string `json:"retired"`
	Justification      string `json:"justification"`
	Department         string `json:"department"`
	ExpenditureType    string `json:"expenditure_type"`
	Depreciation       string `json:"depreciation"`
	AssignedTo         string `json:"assigned_to"`
	DepreciationDate   string `json:"depreciation_date"`
	OldStatus          string `json:"old_status"`
	Comments           string `json:"comments"`
	Cost               string `json:"cost"`
	Quantity           string `json:"quantity"`
	AcquisitionMethod  string `json:"acquisition_method"`
	Ci                 string `json:"ci"`
	SysModCount        string `json:"sys_mod_count"`
	OldSubstatus       string `json:"old_substatus"`
	SerialNumber       string `json:"serial_number"`
	SysTags            string `json:"sys_tags"`
	OrderDate          string `json:"order_date"`
	SupportGroup       string `json:"support_group"`
	ReservedFor        string `json:"reserved_for"`
	Due                string `json:"due"`
	Location           string `json:"location"`
	LeaseID            string `json:"lease_id"`
	SalvageValue       string `json:"salvage_value"`
}