package requests

type ProductStock struct {
	BusinessPartner           int      `json:"BusinessPartner"`
	Product                   string   `json:"Product"`
	Plant                     string   `json:"Plant"`
	Batch                     string   `json:"Batch"`
	InventoryStockType        *string  `json:"InventoryStockType"`
	InventorySpecialStockType *string  `json:"InventorySpecialStockType"`
	ProductStock              *float32 `json:"ProductStock"`
}
