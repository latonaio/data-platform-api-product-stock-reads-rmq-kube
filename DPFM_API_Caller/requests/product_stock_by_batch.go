package requests

type ProductStockByBatch struct {
	Product                   string  `json:"Product"`
	BusinessPartner           int     `json:"BusinessPartner"`
	Plant                     string  `json:"Plant"`
	Batch                     string  `json:"Batch"`
	InventoryStockType        *string `json:"InventoryStockType"`
	InventorySpecialStockType *string `json:"InventorySpecialStockType"`
	ProductStock              float32 `json:"ProductStock"`
}
