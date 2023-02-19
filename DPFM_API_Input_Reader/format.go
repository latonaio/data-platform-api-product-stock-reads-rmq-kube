package dpfm_api_input_reader

import (
	"data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToProductStock() *requests.ProductStock {
	data := sdc.ProductStock
	return &requests.ProductStock{
        BusinessPartner:                     data.BusinessPartner,
        Product:                             data.Product,
        Plant:                               data.Plant,
        Batch:                               data.Batch,
        ProductStockAvailabilityDate:        data.ProductStockAvailabilityDate,
        OrderID:                             data.OrderID,
        OrderItem:                           data.OrderItem,
        Project:                             data.Project,
        InventoryStockType:                  data.InventoryStockType,
        InventorySpecialStockType:           data.InventorySpecialStockType,
        AvailableProductStock:               data.AvailableProductStock,
	}
}

func (sdc *SDC) ConvertToProductStockAvailability() *requests.ProductStockAvailability {
	dataProductStock := sdc.ProductStock
	data := sdc.ProductStock.ProductStockAvailability
	return &requests.ProductStockAvailability{
		BusinessPartner:                    data.BusinessPartner,
		Product:                            data.Product,
		Plant:                              data.Plant,
		Batch:                              data.Batch,
		ProductStockAvailabilityDate:       data.ProductStockAvailabilityDate,
		OrderID:                            data.OrderID,
		OrderItem:                          data.OrderItem,
		Project:                            data.Project,
		InventoryStockType:                 data.InventoryStockType,
		InventorySpecialStockType:          data.InventorySpecialStockType,
		AvailableProductStock:              data.AvailableProductStock,
	}
}
