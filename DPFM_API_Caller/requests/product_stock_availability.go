package requests

type ProductStockAvailability struct {
	BusinessPartner              int      `json:"BusinessPartner"`
	Product                      string   `json:"Product"`
	Plant                        string   `json:"Plant"`
	Batch                        *string  `json:"Batch"`
	ProductStockAvailabilityDate *string  `json:"ProductStockAvailabilityDate"`
	OrderID                      *int     `json:"OrderID"`
	OrderItem                    *int     `json:"OrderItem"`
	Project                      *string  `json:"Project"`
	InventoryStockType           *string  `json:"InventoryStockType"`
	InventorySpecialStockType    *string  `json:"InventorySpecialStockType"`
	AvailableProductStock        *float32 `json:"AvailableProductStock"`
}
