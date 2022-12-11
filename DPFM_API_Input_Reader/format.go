package dpfm_api_input_reader

import (
	"data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToProductStock() *requests.ProductStock {
	data := sdc.ProductStock
	return &requests.ProductStock{
		BusinessPartner:           data.BusinessPartner,
		Product:                   data.Product,
		Plant:                     data.Plant,
		StorageLocation:           data.StorageLocation,
		Batch:                     data.Batch,
		OrderID:                   data.OrderID,
		OrderItem:                 data.OrderItem,
		Project:                   data.Project,
		InventoryStockType:        data.InventoryStockType,
		InventorySpecialStockType: data.InventorySpecialStockType,
		ProductStock:              data.ProductStock,
	}
}

func (sdc *SDC) ConvertToProductStockAvailability() *requests.ProductStockAvailability {
	dataProductStock := sdc.ProductStock
	data := sdc.ProductStock.ProductStockAvailability
	return &requests.ProductStockAvailability{
		BusinessPartner:              dataProductStock.BusinessPartner,
		Product:                      dataProductStock.Product,
		Plant:                        dataProductStock.Plant,
		Batch:                        dataProductStock.Batch,
		BatchValidityEndDate:         data.BatchValidityEndDate,
		OrderID:                      dataProductStock.OrderID,
		OrderItem:                    dataProductStock.OrderItem,
		Project:                      dataProductStock.Project,
		InventoryStockType:           dataProductStock.InventoryStockType,
		InventorySpecialStockType:    dataProductStock.InventorySpecialStockType,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
		AvailableProductStock:        data.AvailableProductStock,
	}
}
