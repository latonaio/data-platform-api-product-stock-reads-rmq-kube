package dpfm_api_output_formatter

import (
	"data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToProductStock(rows *sql.Rows) (*[]ProductStock, error) {
	defer rows.Close()
	pm := &requests.ProductStock{}
	productStock := make([]ProductStock, 0)

	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.InventoryStockType,
			&pm.ProductStock,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
		data := pm
		productStock = append(productStock, ProductStock{
			Product:                                data.Product,
			BusinessPartner:                        data.BusinessPartner,
			Plant:                                  data.Plant,
			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID: data.BusinessPartner,
			Buyer:                                  data.Buyer,
			Seller:                                 data.Seller,
			DeliverToParty:                         data.DeliverToParty,
			DeliverFromParty:                       data.DeliverFromParty,
			DeliverToPlant:                         data.DeliverToPlant,
			DeliverFromPlant:                       data.DeliverFromPlant,
			InventoryStockType:                     data.InventoryStockType,
			ProductStock:                           data.ProductStock,
			CreationDate:                           data.CreationDate,
			CreationTime:                           data.CreationTime,
			LastChangeDate:                         data.LastChangeDate,
			LastChangeTime:                         data.LastChangeTime,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &productStock, nil
}

//func ConvertToProductStockByBatch(rows *sql.Rows) (*ProductStockByBatch, error) {
//	defer rows.Close()
//	pm := &requests.ProductStockByBatch{}
//
//	i := 0
//	for rows.Next() {
//		i++
//		err := rows.Scan(
//			&pm.Product,
//			&pm.BusinessPartner,
//			&pm.Plant,
//			&pm.Batch,
//			&pm.SupplyChainRelationshipID,
//			&pm.SupplyChainRelationshipDeliveryID,
//			&pm.SupplyChainRelationshipDeliveryPlantID,
//			&pm.Buyer,
//			&pm.Seller,
//			&pm.DeliverToParty,
//			&pm.DeliverFromParty,
//			&pm.DeliverToPlant,
//			&pm.DeliverFromPlant,
//			&pm.InventoryStockType,
//			&pm.ProductStock,
//			&pm.CreationDate,
//			&pm.CreationTime,
//			&pm.LastChangeDate,
//			&pm.LastChangeTime,
//		)
//		if err != nil {
//			fmt.Printf("err = %+v \n", err)
//			return nil, err
//		}
//	}
//	if i == 0 {
//		fmt.Printf("DBに対象のレコードが存在しません。")
//		return nil, nil
//	}
//
//	data := pm
//	productStockByBatch := &ProductStockByBatch{
//		Product:                                data.Product,
//		BusinessPartner:                        data.BusinessPartner,
//		Plant:                                  data.Plant,
//		Batch:                                  data.Batch,
//		SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
//		SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
//		SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
//		Buyer:                                  data.Buyer,
//		Seller:                                 data.Seller,
//		DeliverToParty:                         data.DeliverToParty,
//		DeliverFromParty:                       data.DeliverFromParty,
//		DeliverToPlant:                         data.DeliverToPlant,
//		DeliverFromPlant:                       data.DeliverFromPlant,
//		InventoryStockType:                     data.InventoryStockType,
//		ProductStock:                           data.ProductStock,
//		CreationDate:                           data.CreationDate,
//		CreationTime:                           data.CreationTime,
//		LastChangeDate:                         data.LastChangeDate,
//		LastChangeTime:                         data.LastChangeTime,
//	}
//
//	return productStockByBatch, nil
//}
//func ConvertToProductStockByOrder(rows *sql.Rows) (*ProductStockByOrder, error) {
//	defer rows.Close()
//	pm := &requests.ProductStockByOrder{}
//
//	i := 0
//	for rows.Next() {
//		i++
//		err := rows.Scan(
//			&pm.Product,
//			&pm.OrderID,
//			&pm.OrderItem,
//			&pm.SupplyChainRelationshipID,
//			&pm.SupplyChainRelationshipDeliveryID,
//			&pm.SupplyChainRelationshipDeliveryPlantID,
//			&pm.Buyer,
//			&pm.Seller,
//			&pm.DeliverToParty,
//			&pm.DeliverFromParty,
//			&pm.DeliverToPlant,
//			&pm.DeliverFromPlant,
//			&pm.InventoryStockType,
//			&pm.ProductStock,
//			&pm.CreationDate,
//			&pm.CreationTime,
//			&pm.LastChangeDate,
//			&pm.LastChangeTime,
//		)
//		if err != nil {
//			fmt.Printf("err = %+v \n", err)
//			return nil, err
//		}
//	}
//	if i == 0 {
//		fmt.Printf("DBに対象のレコードが存在しません。")
//		return nil, nil
//	}
//
//	data := pm
//	productStockByOrder := &ProductStockByOrder{
//		Product:                                data.Product,
//		OrderID:                                data.OrderID,
//		OrderItem:                              data.OrderItem,
//		SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
//		SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
//		SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
//		Buyer:                                  data.Buyer,
//		Seller:                                 data.Seller,
//		DeliverToParty:                         data.DeliverToParty,
//		DeliverFromParty:                       data.DeliverFromParty,
//		DeliverToPlant:                         data.DeliverToPlant,
//		DeliverFromPlant:                       data.DeliverFromPlant,
//		InventoryStockType:                     data.InventoryStockType,
//		ProductStock:                           data.ProductStock,
//		CreationDate:                           data.CreationDate,
//		CreationTime:                           data.CreationTime,
//		LastChangeDate:                         data.LastChangeDate,
//		LastChangeTime:                         data.LastChangeTime,
//	}
//
//	return productStockByOrder, nil
//}
//func ConvertToProductStockByProject(rows *sql.Rows) (*ProductStockByProject, error) {
//	defer rows.Close()
//	pm := &requests.ProductStockByProject{}
//
//	i := 0
//	for rows.Next() {
//		i++
//		err := rows.Scan(
//			&pm.Product,
//			&pm.Project,
//			&pm.WBSElement,
//			&pm.SupplyChainRelationshipID,
//			&pm.SupplyChainRelationshipDeliveryID,
//			&pm.SupplyChainRelationshipDeliveryPlantID,
//			&pm.Buyer,
//			&pm.Seller,
//			&pm.DeliverToParty,
//			&pm.DeliverFromParty,
//			&pm.DeliverToPlant,
//			&pm.DeliverFromPlant,
//			&pm.InventoryStockType,
//			&pm.ProductStock,
//			&pm.CreationDate,
//			&pm.CreationTime,
//			&pm.LastChangeDate,
//			&pm.LastChangeTime,
//		)
//		if err != nil {
//			fmt.Printf("err = %+v \n", err)
//			return nil, err
//		}
//	}
//	if i == 0 {
//		fmt.Printf("DBに対象のレコードが存在しません。")
//		return nil, nil
//	}
//
//	data := pm
//	productStockByProject := &ProductStockByProject{
//		Product:                                data.Product,
//		Project:                                data.Project,
//		WBSElement:                             data.WBSElement,
//		SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
//		SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
//		SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
//		Buyer:                                  data.Buyer,
//		Seller:                                 data.Seller,
//		DeliverToParty:                         data.DeliverToParty,
//		DeliverFromParty:                       data.DeliverFromParty,
//		DeliverToPlant:                         data.DeliverToPlant,
//		DeliverFromPlant:                       data.DeliverFromPlant,
//		InventoryStockType:                     data.InventoryStockType,
//		ProductStock:                           data.ProductStock,
//		CreationDate:                           data.CreationDate,
//		CreationTime:                           data.CreationTime,
//		LastChangeDate:                         data.LastChangeDate,
//		LastChangeTime:                         data.LastChangeTime,
//	}
//
//	return productStockByProject, nil
//}
//
//func ConvertToProductStockByStorageBin(rows *sql.Rows) (*ProductStockByStorageBin, error) {
//	defer rows.Close()
//	pm := &requests.ProductStockByStorageBin{}
//
//	i := 0
//	for rows.Next() {
//		i++
//		err := rows.Scan(
//			&pm.Product,
//			&pm.BusinessPartner,
//			&pm.Plant,
//			&pm.StorageLocation,
//			&pm.StorageBin,
//			&pm.SupplyChainRelationshipID,
//			&pm.SupplyChainRelationshipDeliveryID,
//			&pm.SupplyChainRelationshipDeliveryPlantID,
//			&pm.Buyer,
//			&pm.Seller,
//			&pm.DeliverToParty,
//			&pm.DeliverFromParty,
//			&pm.DeliverToPlant,
//			&pm.DeliverFromPlant,
//			&pm.InventoryStockType,
//			&pm.ProductStock,
//			&pm.CreationDate,
//			&pm.CreationTime,
//			&pm.LastChangeDate,
//			&pm.LastChangeTime,
//		)
//		if err != nil {
//			fmt.Printf("err = %+v \n", err)
//			return nil, err
//		}
//	}
//	if i == 0 {
//		fmt.Printf("DBに対象のレコードが存在しません。")
//		return nil, nil
//	}
//
//	data := pm
//	productStockByStorageBin := &ProductStockByStorageBin{
//		Product:                                data.Product,
//		BusinessPartner:                        data.BusinessPartner,
//		Plant:                                  data.Plant,
//		StorageLocation:                        data.StorageLocation,
//		StorageBin:                             data.StorageBin,
//		SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
//		SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
//		SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
//		Buyer:                                  data.Buyer,
//		Seller:                                 data.Seller,
//		DeliverToParty:                         data.DeliverToParty,
//		DeliverFromParty:                       data.DeliverFromParty,
//		DeliverToPlant:                         data.DeliverToPlant,
//		DeliverFromPlant:                       data.DeliverFromPlant,
//		InventoryStockType:                     data.InventoryStockType,
//		ProductStock:                           data.ProductStock,
//		CreationDate:                           data.CreationDate,
//		CreationTime:                           data.CreationTime,
//		LastChangeDate:                         data.LastChangeDate,
//		LastChangeTime:                         data.LastChangeTime,
//	}
//
//	return productStockByStorageBin, nil
//}

func ConvertToProductStockByStorageBinByBatch(rows *sql.Rows) (*[]ProductStockByStorageBinByBatch, error) {
	defer rows.Close()
	pm := &requests.ProductStockByStorageBinByBatch{}
	productStockByStorageBinByBatch := make([]ProductStockByStorageBinByBatch, 0)

	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.StorageLocation,
			&pm.StorageBin,
			&pm.Batch,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.InventoryStockType,
			&pm.ProductStock,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
		data := pm
		productStockByStorageBinByBatch = append(productStockByStorageBinByBatch, ProductStockByStorageBinByBatch{
			Product:                                data.Product,
			BusinessPartner:                        data.BusinessPartner,
			Plant:                                  data.Plant,
			StorageLocation:                        data.StorageLocation,
			StorageBin:                             data.StorageBin,
			Batch:                                  data.Batch,
			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
			Buyer:                                  data.Buyer,
			Seller:                                 data.Seller,
			DeliverToParty:                         data.DeliverToParty,
			DeliverFromParty:                       data.DeliverFromParty,
			DeliverToPlant:                         data.DeliverToPlant,
			DeliverFromPlant:                       data.DeliverFromPlant,
			InventoryStockType:                     data.InventoryStockType,
			ProductStock:                           data.ProductStock,
			CreationDate:                           data.CreationDate,
			CreationTime:                           data.CreationTime,
			LastChangeDate:                         data.LastChangeDate,
			LastChangeTime:                         data.LastChangeTime,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &productStockByStorageBinByBatch, nil
}

//func ConvertToProductStockAvailability(rows *sql.Rows) (*[]ProductStockAvailability, error) {
//	defer rows.Close()
//	productStockAvailability := make([]ProductStockAvailability, 0)
//
//	i := 0
//	for rows.Next() {
//		i++
//		pm := &requests.ProductStockAvailability{}
//
//		err := rows.Scan(
//			&pm.Product,
//			&pm.BusinessPartner,
//			&pm.Plant,
//			&pm.SupplyChainRelationshipID,
//			&pm.SupplyChainRelationshipDeliveryID,
//			&pm.SupplyChainRelationshipDeliveryPlantID,
//			&pm.Buyer,
//			&pm.Seller,
//			&pm.DeliverToParty,
//			&pm.DeliverFromParty,
//			&pm.DeliverToPlant,
//			&pm.DeliverFromPlant,
//			&pm.ProductStockAvailabilityDate,
//			&pm.AvailableProductStock,
//			&pm.CreationDate,
//			&pm.CreationTime,
//			&pm.LastChangeDate,
//			&pm.LastChangeTime,
//		)
//		if err != nil {
//			fmt.Printf("err = %+v \n", err)
//			return &productStockAvailability, err
//		}
//
//		data := pm
//		productStockAvailability = append(productStockAvailability, ProductStockAvailability{
//			Product:                                data.Product,
//			BusinessPartner:                        data.BusinessPartner,
//			Plant:                                  data.Plant,
//			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
//			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
//			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
//			Buyer:                                  data.Buyer,
//			Seller:                                 data.Seller,
//			DeliverToParty:                         data.DeliverToParty,
//			DeliverFromParty:                       data.DeliverFromParty,
//			DeliverToPlant:                         data.DeliverToPlant,
//			DeliverFromPlant:                       data.DeliverFromPlant,
//			ProductStockAvailabilityDate:           data.ProductStockAvailabilityDate,
//			AvailableProductStock:                  data.AvailableProductStock,
//			CreationDate:                           data.CreationDate,
//			CreationTime:                           data.CreationTime,
//			LastChangeDate:                         data.LastChangeDate,
//			LastChangeTime:                         data.LastChangeTime,
//		})
//	}
//	if i == 0 {
//		fmt.Printf("DBに対象のレコードが存在しません。")
//		return &productStockAvailability, nil
//	}
//
//	return &productStockAvailability, nil
//}
//
//func ConvertToProductStockAvailabilityByBatch(rows *sql.Rows) (*[]ProductStockAvailabilityByBatch, error) {
//	defer rows.Close()
//	productStockAvailabilityByBatch := make([]ProductStockAvailabilityByBatch, 0)
//
//	i := 0
//	for rows.Next() {
//		i++
//		pm := &requests.ProductStockAvailabilityByBatch{}
//
//		err := rows.Scan(
//			&pm.Product,
//			&pm.BusinessPartner,
//			&pm.Plant,
//			&pm.Batch,
//			&pm.ProductStockAvailabilityDate,
//			&pm.InventoryStockType,
//			&pm.InventorySpecialStockType,
//			&pm.AvailableProductStock,
//		)
//		if err != nil {
//			fmt.Printf("err = %+v \n", err)
//			return &productStockAvailabilityByBatch, err
//		}
//		data := pm
//		productStockAvailabilityByBatch = append(productStockAvailabilityByBatch, ProductStockAvailabilityByBatch{
//			Product:                      data.Product,
//			BusinessPartner:              data.BusinessPartner,
//			Plant:                        data.Plant,
//			Batch:                        data.Batch,
//			ProductStockAvailabilityDate: data.ProductStockAvailabilityDate,
//			InventoryStockType:           data.InventoryStockType,
//			InventorySpecialStockType:    data.InventorySpecialStockType,
//			AvailableProductStock:        data.AvailableProductStock,
//		})
//	}
//	if i == 0 {
//		fmt.Printf("DBに対象のレコードが存在しません。")
//		return &productStockAvailabilityByBatch, nil
//	}
//
//	return &productStockAvailabilityByBatch, nil
//}
//func ConvertToProductStockAvailabilityByOrder(rows *sql.Rows) (*[]ProductStockAvailabilityByOrder, error) {
//	defer rows.Close()
//	productStockAvailabilityByOrder := make([]ProductStockAvailabilityByOrder, 0)
//
//	i := 0
//	for rows.Next() {
//		i++
//		pm := &requests.ProductStockAvailabilityByOrder{}
//
//		err := rows.Scan(
//			&pm.Product,
//			&pm.OrderID,
//			&pm.OrderItem,
//			&pm.SupplyChainRelationshipID,
//			&pm.SupplyChainRelationshipDeliveryID,
//			&pm.SupplyChainRelationshipDeliveryPlantID,
//			&pm.Buyer,
//			&pm.Seller,
//			&pm.DeliverToParty,
//			&pm.DeliverFromParty,
//			&pm.DeliverToPlant,
//			&pm.DeliverFromPlant,
//			&pm.ProductStockAvailabilityDate,
//			&pm.AvailableProductStock,
//			&pm.CreationDate,
//			&pm.CreationTime,
//			&pm.LastChangeDate,
//			&pm.LastChangeTime,
//		)
//		if err != nil {
//			fmt.Printf("err = %+v \n", err)
//			return &productStockAvailabilityByOrder, err
//		}
//		data := pm
//		productStockAvailabilityByOrder = append(productStockAvailabilityByOrder, ProductStockAvailabilityByOrder{
//			Product:                                data.Product,
//			OrderID:                                data.OrderID,
//			OrderItem:                              data.OrderItem,
//			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
//			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
//			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
//			Buyer:                                  data.Buyer,
//			Seller:                                 data.Seller,
//			DeliverToParty:                         data.DeliverToParty,
//			DeliverFromParty:                       data.DeliverFromParty,
//			DeliverToPlant:                         data.DeliverToPlant,
//			DeliverFromPlant:                       data.DeliverFromPlant,
//			ProductStockAvailabilityDate:           data.ProductStockAvailabilityDate,
//			AvailableProductStock:                  data.AvailableProductStock,
//			CreationDate:                           data.CreationDate,
//			CreationTime:                           data.CreationTime,
//			LastChangeDate:                         data.LastChangeDate,
//			LastChangeTime:                         data.LastChangeTime,
//		})
//	}
//	if i == 0 {
//		fmt.Printf("DBに対象のレコードが存在しません。")
//		return &productStockAvailabilityByOrder, nil
//	}
//
//	return &productStockAvailabilityByOrder, nil
//}
//func ConvertToProductStockAvailabilityByProject(rows *sql.Rows) (*[]ProductStockAvailabilityByProject, error) {
//	defer rows.Close()
//	productStockAvailabilityByProject := make([]ProductStockAvailabilityByProject, 0)
//
//	i := 0
//	for rows.Next() {
//		i++
//		pm := &requests.ProductStockAvailabilityByProject{}
//
//		err := rows.Scan(
//			&pm.Product,
//			&pm.Project,
//			&pm.WBSElement,
//			&pm.SupplyChainRelationshipID,
//			&pm.SupplyChainRelationshipDeliveryID,
//			&pm.SupplyChainRelationshipDeliveryPlantID,
//			&pm.Buyer,
//			&pm.Seller,
//			&pm.DeliverToParty,
//			&pm.DeliverFromParty,
//			&pm.DeliverToPlant,
//			&pm.DeliverFromPlant,
//			&pm.ProductStockAvailabilityDate,
//			&pm.AvailableProductStock,
//			&pm.CreationDate,
//			&pm.CreationTime,
//			&pm.LastChangeDate,
//			&pm.LastChangeTime,
//		)
//		if err != nil {
//			fmt.Printf("err = %+v \n", err)
//			return &productStockAvailabilityByProject, err
//		}
//		data := pm
//		productStockAvailabilityByProject = append(productStockAvailabilityByProject, ProductStockAvailabilityByProject{
//			Product:                                data.Product,
//			Project:                                data.Project,
//			WBSElement:                             data.WBSElement,
//			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
//			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
//			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
//			Buyer:                                  data.Buyer,
//			Seller:                                 data.Seller,
//			DeliverToParty:                         data.DeliverToParty,
//			DeliverFromParty:                       data.DeliverFromParty,
//			DeliverToPlant:                         data.DeliverToPlant,
//			DeliverFromPlant:                       data.DeliverFromPlant,
//			ProductStockAvailabilityDate:           data.ProductStockAvailabilityDate,
//			AvailableProductStock:                  data.AvailableProductStock,
//			CreationDate:                           data.CreationDate,
//			CreationTime:                           data.CreationTime,
//			LastChangeDate:                         data.LastChangeDate,
//			LastChangeTime:                         data.LastChangeTime,
//		})
//	}
//	if i == 0 {
//		fmt.Printf("DBに対象のレコードが存在しません。")
//		return &productStockAvailabilityByProject, nil
//	}
//
//	return &productStockAvailabilityByProject, nil
//}
//func ConvertToProductStockAvailabilityByStorageBin(rows *sql.Rows) (*[]ProductStockAvailabilityByStorageBin, error) {
//	defer rows.Close()
//	productStockAvailabilityByStorageBin := make([]ProductStockAvailabilityByStorageBin, 0)
//
//	i := 0
//	for rows.Next() {
//		i++
//		pm := &requests.ProductStockAvailabilityByStorageBin{}
//
//		err := rows.Scan(
//			&pm.Product,
//			&pm.BusinessPartner,
//			&pm.Plant,
//			&pm.StorageLocation,
//			&pm.StorageBin,
//			&pm.SupplyChainRelationshipID,
//			&pm.SupplyChainRelationshipDeliveryID,
//			&pm.SupplyChainRelationshipDeliveryPlantID,
//			&pm.Buyer,
//			&pm.Seller,
//			&pm.DeliverToParty,
//			&pm.DeliverFromParty,
//			&pm.DeliverToPlant,
//			&pm.DeliverFromPlant,
//			&pm.ProductStockAvailabilityDate,
//			&pm.AvailableProductStock,
//			&pm.CreationDate,
//			&pm.CreationTime,
//			&pm.LastChangeDate,
//			&pm.LastChangeTime,
//		)
//		if err != nil {
//			fmt.Printf("err = %+v \n", err)
//			return &productStockAvailabilityByStorageBin, err
//		}
//		data := pm
//		productStockAvailabilityByStorageBin = append(productStockAvailabilityByStorageBin, ProductStockAvailabilityByStorageBin{
//			Product:                                data.Product,
//			BusinessPartner:                        data.BusinessPartner,
//			Plant:                                  data.Plant,
//			StorageLocation:                        data.StorageLocation,
//			StorageBin:                             data.StorageBin,
//			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
//			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
//			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
//			Buyer:                                  data.Buyer,
//			Seller:                                 data.Seller,
//			DeliverToParty:                         data.DeliverToParty,
//			DeliverFromParty:                       data.DeliverFromParty,
//			DeliverToPlant:                         data.DeliverToPlant,
//			DeliverFromPlant:                       data.DeliverFromPlant,
//			ProductStockAvailabilityDate:           data.ProductStockAvailabilityDate,
//			AvailableProductStock:                  data.AvailableProductStock,
//			CreationDate:                           data.CreationDate,
//			CreationTime:                           data.CreationTime,
//			LastChangeDate:                         data.LastChangeDate,
//			LastChangeTime:                         data.LastChangeTime,
//		})
//	}
//	if i == 0 {
//		fmt.Printf("DBに対象のレコードが存在しません。")
//		return &productStockAvailabilityByStorageBin, nil
//	}
//
//	return &productStockAvailabilityByStorageBin, nil
//}
//
//func ConvertToProductStockAvailabilityByStorageBinByBatch(rows *sql.Rows) (*[]ProductStockAvailabilityByStorageBinByBatch, error) {
//	defer rows.Close()
//	productStockAvailabilityByStorageBinByBatch := make([]ProductStockAvailabilityByStorageBinByBatch, 0)
//
//	i := 0
//	for rows.Next() {
//		i++
//		pm := &requests.ProductStockAvailabilityByStorageBinByBatch{}
//		err := rows.Scan(
//			&pm.Product,
//			&pm.BusinessPartner,
//			&pm.Plant,
//			&pm.StorageLocation,
//			&pm.StorageBin,
//			&pm.Batch,
//			&pm.SupplyChainRelationshipID,
//			&pm.SupplyChainRelationshipDeliveryID,
//			&pm.SupplyChainRelationshipDeliveryPlantID,
//			&pm.Buyer,
//			&pm.Seller,
//			&pm.DeliverToParty,
//			&pm.DeliverFromParty,
//			&pm.DeliverToPlant,
//			&pm.DeliverFromPlant,
//			&pm.ProductStockAvailabilityDate,
//			&pm.AvailableProductStock,
//			&pm.CreationDate,
//			&pm.CreationTime,
//			&pm.LastChangeDate,
//			&pm.LastChangeTime,
//		)
//		if err != nil {
//			fmt.Printf("err = %+v \n", err)
//			return &productStockAvailabilityByStorageBinByBatch, err
//		}
//		data := pm
//		productStockAvailabilityByStorageBinByBatch = append(productStockAvailabilityByStorageBinByBatch, ProductStockAvailabilityByStorageBinByBatch{
//			Product:                                data.Product,
//			BusinessPartner:                        data.BusinessPartner,
//			Plant:                                  data.Plant,
//			StorageLocation:                        data.StorageLocation,
//			StorageBin:                             data.StorageBin,
//			Batch:                                  data.Batch,
//			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
//			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
//			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
//			Buyer:                                  data.Buyer,
//			Seller:                                 data.Seller,
//			DeliverToParty:                         data.DeliverToParty,
//			DeliverFromParty:                       data.DeliverFromParty,
//			DeliverToPlant:                         data.DeliverToPlant,
//			DeliverFromPlant:                       data.DeliverFromPlant,
//			ProductStockAvailabilityDate:           data.ProductStockAvailabilityDate,
//			AvailableProductStock:                  data.AvailableProductStock,
//			CreationDate:                           data.CreationDate,
//			CreationTime:                           data.CreationTime,
//			LastChangeDate:                         data.LastChangeDate,
//			LastChangeTime:                         data.LastChangeTime,
//		})
//	}
//	if i == 0 {
//		fmt.Printf("DBに対象のレコードが存在しません。")
//		return &productStockAvailabilityByStorageBinByBatch, nil
//	}
//
//	return &productStockAvailabilityByStorageBinByBatch, nil
//}
