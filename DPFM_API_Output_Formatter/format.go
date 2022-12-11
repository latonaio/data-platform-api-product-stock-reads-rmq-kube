package dpfm_api_output_formatter

import (
	"data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToProductStock(sdc *api_input_reader.SDC, rows *sql.Rows) (*ProductStock, error) {
	pm := &requests.ProductStock{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Product,
			&pm.Plant,
			&pm.StorageLocation,
			&pm.Batch,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.Project,
			&pm.InventoryStockType,
			&pm.InventorySpecialStockType,
			&pm.ProductStock,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	productStock := &ProductStock{
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
	return productStock, nil
}

func ConvertToProductStockAvailability(sdc *api_input_reader.SDC, rows *sql.Rows) (*ProductStockAvailability, error) {
	pm := &requests.ProductStockAvailability{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Product,
			&pm.Plant,
			&pm.Batch,
			&pm.BatchValidityEndDate,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.Project,
			&pm.InventoryStockType,
			&pm.InventorySpecialStockType,
			&pm.ProductStockAvailabilityDate,
			&pm.AvailableProductStock,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	productStockAvailability := &ProductStockAvailability{
		BusinessPartner:              data.BusinessPartner,
		Product:                      data.Product,
		Plant:                        data.Plant,
		Batch:                        data.Batch,
		BatchValidityEndDate:         data.BatchValidityEndDate,
		OrderID:                      data.OrderID,
		OrderItem:                    data.OrderItem,
		Project:                      data.Project,
		InventoryStockType:           data.InventoryStockType,
		InventorySpecialStockType:    data.InventorySpecialStockType,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
		AvailableProductStock:        data.AvailableProductStock,
	}
	return productStockAvailability, nil
}
