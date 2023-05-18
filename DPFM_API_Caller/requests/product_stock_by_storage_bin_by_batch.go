package requests

type ProductStockByStorageBinByBatch struct {
	Product                   string  `json:"Product"`
	BusinessPartner           int     `json:"BusinessPartner"`
	Plant                     string  `json:"Plant"`
	StorageLocation           string  `json:"StorageLocation"`
	StorageBin                string  `json:"StorageBin"`
	Batch                     string  `json:"Batch"`
	InventoryStockType        *string `json:"InventoryStockType"`
	InventorySpecialStockType *string `json:"InventorySpecialStockType"`
	ProductStock              float32 `json:"ProductStock"`
}
