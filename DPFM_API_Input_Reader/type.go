package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string       `json:"connection_key"`
	Result            bool         `json:"result"`
	RedisKey          string       `json:"redis_key"`
	Filepath          string       `json:"filepath"`
	APIStatusCode     int          `json:"api_status_code"`
	RuntimeSessionID  string       `json:"runtime_session_id"`
	BusinessPartnerID *int         `json:"business_partner"`
	ServiceLabel      string       `json:"service_label"`
	APIType           string       `json:"api_type"`
	ProductStock      ProductStock `json:"ProductStock"`
	APISchema         string       `json:"api_schema"`
	Accepter          []string     `json:"accepter"`
	Deleted           bool         `json:"deleted"`
}

type ProductStock struct {
	Product                                     *string                                       `json:"Product"`
	BusinessPartner                             *int                                          `json:"BusinessPartner"`
	Plant                                       *string                                       `json:"Plant"`
	SupplyChainRelationshipID                   *int                                          `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID           *int                                          `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID      *int                                          `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                       *int                                          `json:"Buyer"`
	Seller                                      *int                                          `json:"Seller"`
	DeliverToParty                              *int                                          `json:"DeliverToParty"`
	DeliverFromParty                            *int                                          `json:"DeliverFromParty"`
	DeliverToPlant                              *string                                       `json:"DeliverToPlant"`
	DeliverFromPlant                            *string                                       `json:"DeliverFromPlant"`
	InventoryStockType                          *string                                       `json:"InventoryStockType"`
	ProductStock                                *float32                                      `json:"ProductStock"`
	CreationDate                                *string                                       `json:"CreationDate"`
	CreationTime                                *string                                       `json:"CreationTime"`
	LastChangeDate                              *string                                       `json:"LastChangeDate"`
	LastChangeTime                              *string                                       `json:"LastChangeTime"`
	ProductStockAvailability                    []ProductStockAvailability                    `json:"ProductStockAvailability"`
	ProductStockByBatch                         []ProductStockByBatch                         `json:"ProductStockByBatch"`
	ProductStockByOrder                         []ProductStockByOrder                         `json:"ProductStockByOrder"`
	ProductStockByProject                       []ProductStockByProject                       `json:"ProductStockByProject"`
	ProductStockByStorageBin                    []ProductStockByStorageBin                    `json:"ProductStockByStorageBin"`
	ProductStockByStorageBinByBatch             []ProductStockByStorageBinByBatch             `json:"ProductStockByStorageBinByBatch"`
	ProductStockAvailabilityByBatch             []ProductStockAvailabilityByBatch             `json:"ProductStockAvailabilityByBatch"`
	ProductStockAvailabilityByOrder             []ProductStockAvailabilityByOrder             `json:"ProductStockAvailabilityByOrder"`
	ProductStockAvailabilityByProject           []ProductStockAvailabilityByProject           `json:"ProductStockAvailabilityByProject"`
	ProductStockAvailabilityByStorageBin        []ProductStockAvailabilityByStorageBin        `json:"ProductStockAvailabilityByStorageBin"`
	ProductStockAvailabilityByStorageBinByBatch []ProductStockAvailabilityByStorageBinByBatch `json:"ProductStockAvailabilityByStorageBinByBatch"`
}

type ProductStockDoc struct {
	Product                  string  `json:"Product"`
	BusinessPartner          int     `json:"BusinessPartner"`
	Plant                    string  `json:"Plant"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type ProductStockAvailability struct {
	Product                                string   `json:"Product"`
	BusinessPartner                        int      `json:"BusinessPartner"`
	Plant                                  string   `json:"Plant"`
	SupplyChainRelationshipID              int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int      `json:"Buyer"`
	Seller                                 int      `json:"Seller"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	ProductStockAvailabilityDate           string   `json:"ProductStockAvailabilityDate"`
	AvailableProductStock                  *float32 `json:"AvailableProductStock"`
	CreationDate                           *string  `json:"CreationDate"`
	CreationTime                           *string  `json:"CreationTime"`
	LastChangeDate                         *string  `json:"LastChangeDate"`
	LastChangeTime                         *string  `json:"LastChangeTime"`
}

type ProductStockByBatch struct {
	Product                                string                            `json:"Product"`
	BusinessPartner                        int                               `json:"BusinessPartner"`
	Plant                                  string                            `json:"Plant"`
	Batch                                  string                            `json:"Batch"`
	SupplyChainRelationshipID              int                               `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int                               `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int                               `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int                               `json:"Buyer"`
	Seller                                 int                               `json:"Seller"`
	DeliverToParty                         int                               `json:"DeliverToParty"`
	DeliverFromParty                       int                               `json:"DeliverFromParty"`
	DeliverToPlant                         string                            `json:"DeliverToPlant"`
	DeliverFromPlant                       string                            `json:"DeliverFromPlant"`
	InventoryStockType                     string                            `json:"InventoryStockType"`
	ProductStock                           *float32                          `json:"ProductStock"`
	CreationDate                           *string                           `json:"CreationDate"`
	CreationTime                           *string                           `json:"CreationTime"`
	LastChangeDate                         *string                           `json:"LastChangeDate"`
	LastChangeTime                         *string                           `json:"LastChangeTime"`
	ProductStockAvailabilityByBatch        []ProductStockAvailabilityByBatch `json:"ProductStockAvailabilityByBatch"`
}

type ProductStockByOrder struct {
	Product                                string   `json:Product`
	OrderID                                int      `json:OrderID`
	OrderItem                              int      `json:OrderItem`
	SupplyChainRelationshipID              int      `json:SupplyChainRelationshipID`
	SupplyChainRelationshipDeliveryID      int      `json:SupplyChainRelationshipDeliveryID`
	SupplyChainRelationshipDeliveryPlantID int      `json:SupplyChainRelationshipDeliveryPlantID`
	Buyer                                  int      `json:Buyer`
	Seller                                 int      `json:Seller`
	DeliverToParty                         int      `json:DeliverToParty`
	DeliverFromParty                       int      `json:DeliverFromParty`
	DeliverToPlant                         string   `json:DeliverToPlant`
	DeliverFromPlant                       string   `json:DeliverFromPlant`
	InventoryStockType                     string   `json:InventoryStockType`
	ProductStock                           *float32 `json:ProductStock`
	CreationDate                           *string  `json:CreationDate`
	CreationTime                           *string  `json:CreationTime`
	LastChangeDate                         *string  `json:LastChangeDate`
	LastChangeTime                         *string  `json:LastChangeTime`
}

type ProductStockByProject struct {
	Product                                string   `json:Product`
	Project                                int      `json:Project`
	WBSElement                             int      `json:WBSElement`
	SupplyChainRelationshipID              int      `json:SupplyChainRelationshipID`
	SupplyChainRelationshipDeliveryID      int      `json:SupplyChainRelationshipDeliveryID`
	SupplyChainRelationshipDeliveryPlantID int      `json:SupplyChainRelationshipDeliveryPlantID`
	Buyer                                  int      `json:Buyer`
	Seller                                 int      `json:Seller`
	DeliverToParty                         int      `json:DeliverToParty`
	DeliverFromParty                       int      `json:DeliverFromParty`
	DeliverToPlant                         string   `json:DeliverToPlant`
	DeliverFromPlant                       string   `json:DeliverFromPlant`
	InventoryStockType                     string   `json:InventoryStockType`
	ProductStock                           *float32 `json:ProductStock`
	CreationDate                           *string  `json:CreationDate`
	CreationTime                           *string  `json:CreationTime`
	LastChangeDate                         *string  `json:LastChangeDate`
	LastChangeTime                         *string  `json:LastChangeTime`
}

type ProductStockByStorageBin struct {
	Product                                string                                 `json:"Product"`
	BusinessPartner                        int                                    `json:"BusinessPartner"`
	Plant                                  string                                 `json:"Plant"`
	StorageLocation                        string                                 `json:"StorageLocation"`
	StorageBin                             string                                 `json:"StorageBin"`
	SupplyChainRelationshipID              int                                    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int                                    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int                                    `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int                                    `json:"Buyer"`
	Seller                                 int                                    `json:"Seller"`
	DeliverToParty                         int                                    `json:"DeliverToParty"`
	DeliverFromParty                       int                                    `json:"DeliverFromParty"`
	DeliverToPlant                         string                                 `json:"DeliverToPlant"`
	DeliverFromPlant                       string                                 `json:"DeliverFromPlant"`
	InventoryStockType                     string                                 `json:"InventoryStockType"`
	ProductStock                           *float32                               `json:"ProductStock"`
	CreationDate                           *string                                `json:"CreationDate"`
	CreationTime                           *string                                `json:"CreationTime"`
	LastChangeDate                         *string                                `json:"LastChangeDate"`
	LastChangeTime                         *string                                `json:"LastChangeTime"`
	ProductStockAvailabilityByStorageBin   []ProductStockAvailabilityByStorageBin `json:"ProductStockAvailabilityByStorageBin"`
}

type ProductStockByStorageBinByBatch struct {
	Product                                     string                                        `json:"Product"`
	BusinessPartner                             int                                           `json:"BusinessPartner"`
	Plant                                       string                                        `json:"Plant"`
	StorageLocation                             string                                        `json:"StorageLocation"`
	StorageBin                                  string                                        `json:"StorageBin"`
	Batch                                       string                                        `json:"Batch"`
	SupplyChainRelationshipID                   int                                           `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID           int                                           `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID      int                                           `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                       int                                           `json:"Buyer"`
	Seller                                      int                                           `json:"Seller"`
	DeliverToParty                              int                                           `json:"DeliverToParty"`
	DeliverFromParty                            int                                           `json:"DeliverFromParty"`
	DeliverToPlant                              string                                        `json:"DeliverToPlant"`
	DeliverFromPlant                            string                                        `json:"DeliverFromPlant"`
	InventoryStockType                          string                                        `json:"InventoryStockType"`
	ProductStock                                *float32                                      `json:"ProductStock"`
	CreationDate                                *string                                       `json:"CreationDate"`
	CreationTime                                *string                                       `json:"CreationTime"`
	LastChangeDate                              *string                                       `json:"LastChangeDate"`
	LastChangeTime                              *string                                       `json:"LastChangeTime"`
	ProductStockAvailabilityByStorageBinByBatch []ProductStockAvailabilityByStorageBinByBatch `json:"ProductStockAvailabilityByStorageBinByBatch"`
}

type ProductStockAvailabilityByBatch struct {
	Product                                string   `json:"Product"`
	BusinessPartner                        int      `json:"BusinessPartner"`
	Plant                                  string   `json:"Plant"`
	Batch                                  string   `json:"Batch"`
	SupplyChainRelationshipID              int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int      `json:"Buyer"`
	Seller                                 int      `json:"Seller"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	ProductStockAvailabilityDate           string   `json:"ProductStockAvailabilityDate"`
	AvailableProductStock                  *float32 `json:"AvailableProductStock"`
	CreationDate                           *string  `json:"CreationDate"`
	CreationTime                           *string  `json:"CreationTime"`
	LastChangeDate                         *string  `json:"LastChangeDate"`
	LastChangeTime                         *string  `json:"LastChangeTime"`
}

type ProductStockAvailabilityByOrder struct {
	Product                                string   `json:Product`
	OrderID                                int      `json:OrderID`
	OrderItem                              int      `json:OrderItem`
	SupplyChainRelationshipID              int      `json:SupplyChainRelationshipID`
	SupplyChainRelationshipDeliveryID      int      `json:SupplyChainRelationshipDeliveryID`
	SupplyChainRelationshipDeliveryPlantID int      `json:SupplyChainRelationshipDeliveryPlantID`
	Buyer                                  int      `json:Buyer`
	Seller                                 int      `json:Seller`
	DeliverToParty                         int      `json:DeliverToParty`
	DeliverFromParty                       int      `json:DeliverFromParty`
	DeliverToPlant                         string   `json:DeliverToPlant`
	DeliverFromPlant                       string   `json:DeliverFromPlant`
	ProductStockAvailabilityDate           string   `json:ProductStockAvailabilityDate`
	AvailableProductStock                  *float32 `json:AvailableProductStock`
	CreationDate                           *string  `json:CreationDate`
	CreationTime                           *string  `json:CreationTime`
	LastChangeDate                         *string  `json:LastChangeDate`
	LastChangeTime                         *string  `json:LastChangeTime`
}

type ProductStockAvailabilityByProject struct {
	Product                                string   `json:Product`
	Project                                int      `json:Project`
	WBSElement                             int      `json:WBSElement`
	SupplyChainRelationshipID              int      `json:SupplyChainRelationshipID`
	SupplyChainRelationshipDeliveryID      int      `json:SupplyChainRelationshipDeliveryID`
	SupplyChainRelationshipDeliveryPlantID int      `json:SupplyChainRelationshipDeliveryPlantID`
	Buyer                                  int      `json:Buyer`
	Seller                                 int      `json:Seller`
	DeliverToParty                         int      `json:DeliverToParty`
	DeliverFromParty                       int      `json:DeliverFromParty`
	DeliverToPlant                         string   `json:DeliverToPlant`
	DeliverFromPlant                       string   `json:DeliverFromPlant`
	ProductStockAvailabilityDate           string   `json:ProductStockAvailabilityDate`
	AvailableProductStock                  *float32 `json:AvailableProductStock`
	CreationDate                           *string  `json:CreationDate`
	CreationTime                           *string  `json:CreationTime`
	LastChangeDate                         *string  `json:LastChangeDate`
	LastChangeTime                         *string  `json:LastChangeTime`
}

type ProductStockAvailabilityByStorageBin struct {
	Product                                string   `json:"Product"`
	BusinessPartner                        int      `json:"BusinessPartner"`
	Plant                                  string   `json:"Plant"`
	StorageLocation                        string   `json:"StorageLocation"`
	StorageBin                             string   `json:"StorageBin"`
	SupplyChainRelationshipID              int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int      `json:"Buyer"`
	Seller                                 int      `json:"Seller"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	ProductStockAvailabilityDate           string   `json:"ProductStockAvailabilityDate"`
	AvailableProductStock                  *float32 `json:"AvailableProductStock"`
	CreationDate                           *string  `json:"CreationDate"`
	CreationTime                           *string  `json:"CreationTime"`
	LastChangeDate                         *string  `json:"LastChangeDate"`
	LastChangeTime                         *string  `json:"LastChangeTime"`
}

type ProductStockAvailabilityByStorageBinByBatch struct {
	Product                                string   `json:"Product"`
	BusinessPartner                        int      `json:"BusinessPartner"`
	Plant                                  string   `json:"Plant"`
	StorageLocation                        string   `json:"StorageLocation"`
	StorageBin                             string   `json:"StorageBin"`
	Batch                                  string   `json:"Batch"`
	SupplyChainRelationshipID              int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int      `json:"Buyer"`
	Seller                                 int      `json:"Seller"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	ProductStockAvailabilityDate           string   `json:"ProductStockAvailabilityDate"`
	AvailableProductStock                  *float32 `json:"AvailableProductStock"`
	CreationDate                           *string  `json:"CreationDate"`
	CreationTime                           *string  `json:"CreationTime"`
	LastChangeDate                         *string  `json:"LastChangeDate"`
	LastChangeTime                         *string  `json:"LastChangeTime"`
}
