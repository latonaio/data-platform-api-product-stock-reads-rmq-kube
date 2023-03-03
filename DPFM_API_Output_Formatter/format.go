package dpfm_api_output_formatter

import (
	"data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToProductStock(rows *sql.Rows) (*ProductStock, error) {
	defer rows.Close()
	pm := &requests.ProductStock{}

	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Product,
			&pm.Plant,
			// &pm.Batch,
			&pm.InventoryStockType,
			&pm.InventorySpecialStockType,
			&pm.ProductStock,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	data := pm
	productStock := &ProductStock{
		BusinessPartner:           data.BusinessPartner,
		Product:                   data.Product,
		Plant:                     data.Plant,
		Batch:                     data.Batch,
		InventoryStockType:        data.InventoryStockType,
		InventorySpecialStockType: data.InventorySpecialStockType,
		ProductStock:              data.ProductStock,
	}

	return productStock, nil
}
func ConvertToProductStockByBatch(rows *sql.Rows) (*ProductStock, error) {
	defer rows.Close()
	pm := &requests.ProductStock{}

	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Product,
			&pm.Plant,
			&pm.Batch,
			&pm.InventoryStockType,
			&pm.InventorySpecialStockType,
			&pm.ProductStock,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	data := pm
	productStock := &ProductStock{
		BusinessPartner:           data.BusinessPartner,
		Product:                   data.Product,
		Plant:                     data.Plant,
		Batch:                     data.Batch,
		InventoryStockType:        data.InventoryStockType,
		InventorySpecialStockType: data.InventorySpecialStockType,
		ProductStock:              data.ProductStock,
	}

	return productStock, nil
}

func ConvertToProductStockAvailability(rows *sql.Rows) (*ProductStockAvailability, error) {
	defer rows.Close()
	pm := &requests.ProductStockAvailability{}

	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Product,
			&pm.Plant,
			&pm.Batch,
			&pm.ProductStockAvailabilityDate,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.Project,
			&pm.InventoryStockType,
			&pm.InventorySpecialStockType,
			&pm.AvailableProductStock,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	data := pm
	productStockAvailability := &ProductStockAvailability{
		BusinessPartner:              data.BusinessPartner,
		Product:                      data.Product,
		Plant:                        data.Plant,
		Batch:                        data.Batch,
		ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
		OrderID:                      data.OrderID,
		OrderItem:                    data.OrderItem,
		Project:                      data.Project,
		InventoryStockType:           data.InventoryStockType,
		InventorySpecialStockType:    data.InventorySpecialStockType,
		AvailableProductStock:        data.AvailableProductStock,
	}

	return productStockAvailability, nil
}
