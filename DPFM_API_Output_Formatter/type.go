package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	ProductStock             *ProductStock             `json:"ProductStock"`
	ProductStockAvailability *ProductStockAvailability `json:"ProductStockAvailability"`
}

type ProductStockReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Deleted       string `json:"deleted"`
}

type ProductStock struct {
	BusinessPartner           int                      `json:"BusinessPartner"`
	Product                   string                   `json:"Product"`
	Plant                     string                   `json:"Plant"`
	StorageLocation           *string                  `json:"StorageLocation"`
	Batch                     *string                  `json:"Batch"`
	OrderID                   *int                     `json:"OrderID"`
	OrderItem                 *int                     `json:"OrderItem"`
	Project                   *string                  `json:"Project"`
	InventoryStockType        *string                  `json:"InventoryStockType"`
	InventorySpecialStockType *string                  `json:"InventorySpecialStockType"`
	ProductStock              *float32                 `json:"ProductStock"`
	ProductStockAvailability  ProductStockAvailability `json:"ProductStockAvailability"`
}

type ProductStockAvailability struct {
	BusinessPartner              int      `json:"BusinessPartner"`
	Product                      string   `json:"Product"`
	Plant                        string   `json:"Plant"`
	Batch                        *string  `json:"Batch"`
	BatchValidityEndDate         *string  `json:"BatchValidityEndDate"`
	OrderID                      *int     `json:"OrderID"`
	OrderItem                    *int     `json:"OrderItem"`
	Project                      *string  `json:"Project"`
	InventoryStockType           *string  `json:"InventoryStockType"`
	InventorySpecialStockType    *string  `json:"InventorySpecialStockType"`
	ProductStockAvailabilityDate string   `json:"ProductStockAvailabilityDate"`
	AvailableProductStock        *float32 `json:"AvailableProductStock"`
}
