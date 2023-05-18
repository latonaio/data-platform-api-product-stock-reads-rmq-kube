package requests

type ProductStockAvailabilityByStorageBinByBatch struct {
	Product                      string  `json:"Product"`
	BusinessPartner              int     `json:"BusinessPartner"`
	Plant                        string  `json:"Plant"`
	StorageLocation              string  `json:"StorageLocation"`
	StorageBin                   string  `json:"StorageBin"`
	Batch                        string  `json:"Batch"`
	ProductStockAvailabilityDate string  `json:"ProductStockAvailabilityDate"`
	InventoryStockType           *string `json:"InventoryStockType"`
	InventorySpecialStockType    *string `json:"InventorySpecialStockType"`
	AvailableProductStock        float32 `json:"AvailableProductStock"`
}
