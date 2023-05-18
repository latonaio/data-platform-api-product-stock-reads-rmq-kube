package requests

type ProductStockAvailabilityByBatch struct {
	Product                      string  `json:"Product"`
	BusinessPartner              int     `json:"BusinessPartner"`
	Plant                        string  `json:"Plant"`
	Batch                        string  `json:"Batch"`
	ProductStockAvailabilityDate string  `json:"ProductStockAvailabilityDate"`
	InventoryStockType           *string `json:"InventoryStockType"`
	InventorySpecialStockType    *string `json:"InventorySpecialStockType"`
	AvailableProductStock        float32 `json:"AvailableProductStock"`
}
