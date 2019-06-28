package servicenow

import "net/url"

const TableModel = "cmdb_model"
const TableHwModel = "cmdb_hardware_product_model"

func (c Client) GetModelItems(query url.Values) ([]Model, error) {
	var res struct {
		Records []Model
	}
	err := c.GetRecordsFor(TableModel, query, &res)
	return res.Records, err
}

//HW Model table returns the same format as cmdb_model table
func (c Client) GetHwModelItems(query url.Values) ([]Model, error) {
	var res struct {
		Records []Model
	}
	err := c.GetRecordsFor(TableHwModel, query, &res)
	return res.Records, err
}

type Model struct {
	ShortDescription      string `json:"short_description"`
	FlowRate              string `json:"flow_rate"`
	Description           string `json:"description"`
	SysUpdatedOn          string `json:"sys_updated_on"`
	Type                  string `json:"type"`
	SysClassName          string `json:"sys_class_name"`
	Manufacturer          string `json:"manufacturer"`
	SysID                 string `json:"sys_id"`
	SoundPower            string `json:"sound_power"`
	SysUpdatedBy          string `json:"sys_updated_by"`
	SysCreatedOn          string `json:"sys_created_on"`
	Certified             string `json:"certified"`
	ModelNumber           string `json:"model_number"`
	Barcode               string `json:"barcode"`
	Bundle                string `json:"bundle"`
	ExpenditureType       string `json:"expenditure_type"`
	Depreciation          string `json:"depreciation"`
	SysCreatedBy          string `json:"sys_created_by"`
	ResponseStatus        string `json:"__status"`
	Owner                 string `json:"owner"`
	Comments              string `json:"comments"`
	Cost                  string `json:"cost"`
	AcquisitionMethod     string `json:"acquisition_method"`
	PowerConsumption      string `json:"power_consumption"`
	SysModCount           string `json:"sys_mod_count"`
	SLA                   string `json:"sla"`
	Weight                string `json:"weight"`
	CmdbModelCategory     string `json:"cmdb_model_category"`
	DisplayName           string `json:"display_name"`
	SysTags               string `json:"sys_tags"`
	AssetTrackingStrategy string `json:"asset_tracking_strategy"`
	Picture               string `json:"picture"`
	MainComponent         string `json:"main_component"`
	FullName              string `json:"full_name"`
	RackUnits             string `json:"rack_units"`
	ProductCatalogItem    string `json:"product_catalog_item"`
	Name                  string `json:"name"`
	CmdbCiClass           string `json:"cmdb_ci_class"`
	SalvageValue          string `json:"salvage_value"`
	Status                string `json:"status"`
}
