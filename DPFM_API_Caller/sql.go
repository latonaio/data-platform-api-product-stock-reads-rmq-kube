package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var productStock *[]dpfm_api_output_formatter.ProductStock
	var productStockByBatch *[]dpfm_api_output_formatter.ProductStockByBatch
	var productStockByOrder *[]dpfm_api_output_formatter.ProductStockByOrder
	var productStockByProject *[]dpfm_api_output_formatter.ProductStockByProject
	var productStockByStorageBin *[]dpfm_api_output_formatter.ProductStockByStorageBin
	var productStockByStorageBinByBatch *[]dpfm_api_output_formatter.ProductStockByStorageBinByBatch
	var productStockAvailability *[]dpfm_api_output_formatter.ProductStockAvailability
	var productStockAvailabilityByBatch *[]dpfm_api_output_formatter.ProductStockAvailabilityByBatch
	var productStockAvailabilityByOrder *[]dpfm_api_output_formatter.ProductStockAvailabilityByOrder
	var productStockAvailabilityByProject *[]dpfm_api_output_formatter.ProductStockAvailabilityByProject
	var productStockAvailabilityByStorageBin *[]dpfm_api_output_formatter.ProductStockAvailabilityByStorageBin
	var productStockAvailabilityByStorageBinByBatch *[]dpfm_api_output_formatter.ProductStockAvailabilityByStorageBinByBatch
	for _, fn := range accepter {
		switch fn {
		case "ProductStock":
			func() {
				productStock = c.ProductStock(mtx, input, output, errs, log)
			}()
		//case "ProductStockByBatch":
		//	func() {
		//		productStockByBatch = c.ProductStockByBatch(mtx, input, output, errs, log)
		//	}()
		//case "ProductStockByOrder":
		//	func() {
		//		productStockByOrder = c.ProductStockByOrder(mtx, input, output, errs, log)
		//	}()
		//case "ProductStockByProject":
		//	func() {
		//		productStockByProject = c.ProductStockByProject(mtx, input, output, errs, log)
		//	}()
		//case "ProductStockByStorageBin":
		//	func() {
		//		productStockByStorageBin = c.ProductStockByStorageBin(mtx, input, output, errs, log)
		//	}()
		case "ProductStocksByStorageBinByBatch":
			func() {
				productStockByStorageBinByBatch = c.ProductStocksByStorageBinByBatch(mtx, input, output, errs, log)
			}()
		//case "ProductStockAvailability":
		//	func() {
		//		productStockAvailability = c.ProductStockAvailability(mtx, input, output, errs, log)
		//	}()w
		//case "ProductStockAvailabilityByBatch":
		//	func() {
		//		productStockAvailabilityByBatch = c.ProductStockAvailabilityByBatch(mtx, input, output, errs, log)
		//	}()
		//case "ProductStockAvailabilityByOrder":
		//	func() {
		//		productStockAvailabilityByOrder = c.ProductStockAvailabilityByOrder(mtx, input, output, errs, log)
		//	}()
		//case "ProductStockAvailabilityByProject":
		//	func() {
		//		productStockAvailabilityByProject = c.ProductStockAvailabilityByProject(mtx, input, output, errs, log)
		//	}()
		//case "ProductStockAvailabilityByStorageBin":
		//	func() {
		//		productStockAvailabilityByStorageBin = c.ProductStockAvailabilityByStorageBin(mtx, input, output, errs, log)
		//	}()
		//case "ProductStockAvailabilityByStorageBinByBatch":
		//	func() {
		//		productStockAvailabilityByStorageBinByBatch = c.ProductStockAvailabilityByStorageBinByBatch(mtx, input, output, errs, log)
		//	}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		ProductStock:                                productStock,
		ProductStockByBatch:                         productStockByBatch,
		ProductStockByOrder:                         productStockByOrder,
		ProductStockByProject:                       productStockByProject,
		ProductStockByStorageBin:                    productStockByStorageBin,
		ProductStockByStorageBinByBatch:             productStockByStorageBinByBatch,
		ProductStockAvailability:                    productStockAvailability,
		ProductStockAvailabilityByBatch:             productStockAvailabilityByBatch,
		ProductStockAvailabilityByOrder:             productStockAvailabilityByOrder,
		ProductStockAvailabilityByProject:           productStockAvailabilityByProject,
		ProductStockAvailabilityByStorageBin:        productStockAvailabilityByStorageBin,
		ProductStockAvailabilityByStorageBinByBatch: productStockAvailabilityByStorageBinByBatch,
	}

	return data
}

func (c *DPFMAPICaller) ProductStock(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductStock {
	where := "WHERE 1 = 1"

	if input.ProductStock.Product != nil {
		where = fmt.Sprintf("%s\nAND Product =  \"%s\"", where, *input.ProductStock.Product)
	}
	if input.ProductStock.BusinessPartner != nil && *input.ProductStock.BusinessPartner != 0 {
		where = fmt.Sprintf("%s\nAND BusinessPartner = %d", where, *input.ProductStock.BusinessPartner)
	}
	if input.ProductStock.Plant != nil {
		where = fmt.Sprintf("%s\nAND Plant = \"%s\"", where, *input.ProductStock.Plant)
	}
	if input.ProductStock.SupplyChainRelationshipID != nil && *input.ProductStock.SupplyChainRelationshipID != 0 {
		where = fmt.Sprintf("%s\nAND SupplyChainRelationshipID = %d", where, *input.ProductStock.SupplyChainRelationshipID)
	}
	if input.ProductStock.SupplyChainRelationshipDeliveryID != nil && *input.ProductStock.SupplyChainRelationshipDeliveryID != 0 {
		where = fmt.Sprintf("%s\nAND SupplyChainRelationshipDeliveryID = %d", where, *input.ProductStock.SupplyChainRelationshipDeliveryID)
	}
	if input.ProductStock.SupplyChainRelationshipDeliveryPlantID != nil && *input.ProductStock.SupplyChainRelationshipDeliveryPlantID != 0 {
		where = fmt.Sprintf("%s\nAND SupplyChainRelationshipDeliveryPlantID = %d", where, *input.ProductStock.SupplyChainRelationshipDeliveryPlantID)
	}
	if input.ProductStock.Buyer != nil && *input.ProductStock.Buyer != 0 {
		where = fmt.Sprintf("%s\nAND Buyer = %d", where, *input.ProductStock.Buyer)
	}
	if input.ProductStock.Seller != nil && *input.ProductStock.Seller != 0 {
		where = fmt.Sprintf("%s\nAND Seller = %d", where, *input.ProductStock.Seller)
	}
	if input.ProductStock.DeliverToParty != nil && *input.ProductStock.DeliverToParty != 0 {
		where = fmt.Sprintf("%s\nAND DeliverToParty = %d", where, *input.ProductStock.DeliverToParty)
	}
	if input.ProductStock.DeliverFromParty != nil && *input.ProductStock.DeliverFromParty != 0 {
		where = fmt.Sprintf("%s\nAND DeliverFromParty = %d", where, *input.ProductStock.DeliverFromParty)
	}
	if input.ProductStock.DeliverToPlant != nil {
		where = fmt.Sprintf("%s\nAND DeliverToPlant = \"%s\"", where, *input.ProductStock.DeliverToPlant)
	}
	if input.ProductStock.DeliverFromPlant != nil {
		where = fmt.Sprintf("%s\nAND DeliverFromPlant = \"%s\"", where, *input.ProductStock.DeliverFromPlant)
	}
	if input.ProductStock.InventoryStockType != nil {
		where = fmt.Sprintf("%s\nAND InventoryStockType = \"%s\"", where, *input.ProductStock.InventoryStockType)
	}
	groupBy := "\nGROUP BY Product, BusinessPartner, Plant "

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_product_stock_data
		` + where + groupBy + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductStock(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

//func (c *DPFMAPICaller) ProductStockAvailability(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *dpfm_api_output_formatter.ProductStockAvailability {
//	where := "WHERE 1 = 1"
//
//	if input.ProductStock.ProductStockAvailability.Product != nil {
//		where = fmt.Sprintf("%s\nAND Product =  \"%s\"", where, *input.ProductStock.ProductStockAvailability.Product)
//	}
//	if input.ProductStock.ProductStockAvailability.BusinessPartner != nil && *input.ProductStock.ProductStockAvailability.BusinessPartner != 0 {
//		where = fmt.Sprintf("%s\nAND BusinessPartner = %d", where, *input.ProductStock.ProductStockAvailability.BusinessPartner)
//	}
//	if input.ProductStock.ProductStockAvailability.Plant != nil {
//		where = fmt.Sprintf("%s\nAND Plant = \"%s\"", where, *input.ProductStock.ProductStockAvailability.Plant)
//	}
//	if input.ProductStock.ProductStockAvailability.SupplyChainRelationshipID != nil && *input.ProductStock.ProductStockAvailability.SupplyChainRelationshipID != 0 {
//		where = fmt.Sprintf("%s\nAND SupplyChainRelationshipID = %d", where, *input.ProductStock.ProductStockAvailability.SupplyChainRelationshipID)
//	}
//	if input.ProductStock.ProductStockAvailability.SupplyChainRelationshipDeliveryID != nil && *input.ProductStock.ProductStockAvailability.SupplyChainRelationshipDeliveryID != 0 {
//		where = fmt.Sprintf("%s\nAND SupplyChainRelationshipDeliveryID = %d", where, *input.ProductStock.ProductStockAvailability.SupplyChainRelationshipDeliveryID)
//	}
//	if input.ProductStock.ProductStockAvailability.SupplyChainRelationshipDeliveryPlantID != nil && *input.ProductStock.ProductStockAvailability.SupplyChainRelationshipDeliveryPlantID != 0 {
//		where = fmt.Sprintf("%s\nAND SupplyChainRelationshipDeliveryPlantID = %d", where, *input.ProductStock.ProductStockAvailability.SupplyChainRelationshipDeliveryPlantID)
//	}
//	if input.ProductStock.ProductStockAvailability.Buyer != nil && *input.ProductStock.ProductStockAvailability.Buyer != 0 {
//		where = fmt.Sprintf("%s\nAND Buyer = %d", where, *input.ProductStock.ProductStockAvailability.Buyer)
//	}
//	if input.ProductStock.ProductStockAvailability.Seller != nil && *input.ProductStock.ProductStockAvailability.Seller != 0 {
//		where = fmt.Sprintf("%s\nAND Seller = %d", where, *input.ProductStock.ProductStockAvailability.Seller)
//	}
//	if input.ProductStock.DeliverToParty != nil && *input.ProductStock.ProductStockAvailability.DeliverToParty != 0 {
//		where = fmt.Sprintf("%s\nAND DeliverToParty = %d", where, *input.ProductStock.ProductStockAvailability.DeliverToParty)
//	}
//	if input.ProductStock.ProductStockAvailability.DeliverFromParty != nil && *input.ProductStock.ProductStockAvailability.DeliverFromParty != 0 {
//		where = fmt.Sprintf("%s\nAND DeliverFromParty = %d", where, *input.ProductStock.ProductStockAvailability.DeliverFromParty)
//	}
//	if input.ProductStock.ProductStockAvailability.DeliverToPlant != nil {
//		where = fmt.Sprintf("%s\nAND DeliverToPlant = \"%s\"", where, *input.ProductStock.ProductStockAvailability.DeliverToPlant)
//	}
//	if input.ProductStock.ProductStockAvailability.DeliverFromPlant != nil {
//		where = fmt.Sprintf("%s\nAND DeliverFromPlant = \"%s\"", where, *input.ProductStock.ProductStockAvailability.DeliverFromPlant)
//	}
//	if input.ProductStock.ProductStockAvailability.ProductStockAvailabilityDate != nil {
//		where = fmt.Sprintf("%s\nAND ProductStockAvailabilityDate = \"%s\"", where, *input.ProductStock.ProductStockAvailability.ProductStockAvailabilityDate)
//	}
//	groupBy := "\nGROUP BY Product, BusinessPartner, Plant "
//
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_product_stock_availability_data
//		` + where + groupBy + `;`)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//	defer rows.Close()
//
//	data, err := dpfm_api_output_formatter.ConvertToProductStockAvailability(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}
//
//func (c *DPFMAPICaller) ProductStockByBatch(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *dpfm_api_output_formatter.ProductStockByBatch {
//	product := input.ProductStock.ProductStockByBatch.Product
//	businessPartner := input.ProductStock.ProductStockByBatch.BusinessPartner
//	plant := input.ProductStock.ProductStockByBatch.Plant
//	batch := input.ProductStock.ProductStockByBatch.Batch
//
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_by_batch_data
//		WHERE (Product, BusinessPartner, Plant, Batch) = (?, ?, ?, ?);`, product, businessPartner, plant, batch,
//	)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//	defer rows.Close()
//
//	data, err := dpfm_api_output_formatter.ConvertToProductStockByBatch(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}
//func (c *DPFMAPICaller) ProductStockByOrder(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *dpfm_api_output_formatter.ProductStockByOrder {
//	product := input.ProductStockByOrder.Product
//	businessPartner := input.ProductStockByOrder.BusinessPartner
//	plant := input.ProductStockByOrder.Plant
//	batch := input.ProductStockByOrder.Batch
//
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_by_Order_data
//		WHERE (Product, BusinessPartner, Plant, Batch) = (?, ?, ?, ?);`, product, businessPartner, plant, batch,
//	)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//	defer rows.Close()
//
//	data, err := dpfm_api_output_formatter.ConvertToProductStockByOrder(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}
//
//func (c *DPFMAPICaller) ProductStockByProject(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *dpfm_api_output_formatter.ProductStockByProject {
//	product := input.ProductStockByProject.Product
//	businessPartner := input.ProductStockByProject.BusinessPartner
//	plant := input.ProductStockByProject.Plant
//	batch := input.ProductStockByProject.Batch
//
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_by_Project_data
//		WHERE (Product, BusinessPartner, Plant, Batch) = (?, ?, ?, ?);`, product, businessPartner, plant, batch,
//	)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//	defer rows.Close()
//
//	data, err := dpfm_api_output_formatter.ConvertToProductStockByProject(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}
//
//func (c *DPFMAPICaller) ProductStockByStorageBin(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *dpfm_api_output_formatter.ProductStockByStorageBin {
//	product := input.ProductStockByStorageBin.Product
//	businessPartner := input.ProductStockByStorageBin.BusinessPartner
//	plant := input.ProductStockByStorageBin.Plant
//	storageLocation := input.ProductStockByStorageBin.StorageLocation
//	storageBin := input.ProductStockByStorageBin.StorageBin
//
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_by_storage_bin_data
//		WHERE (Product, BusinessPartner, Plant, StorageLocation, StorageBin) = (?, ?, ?, ?, ?);`, product, businessPartner, plant, storageBin, storageLocation,
//	)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//	defer rows.Close()
//
//	data, err := dpfm_api_output_formatter.ConvertToProductStockByStorageBin(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}

func (c *DPFMAPICaller) ProductStocksByStorageBinByBatch(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductStockByStorageBinByBatch {
	var args []interface{}

	cnt := 0

	for _, v := range input.ProductStock.ProductStockByStorageBinByBatch {
		product := v.Product
		businessPartner := v.BusinessPartner
		plant := v.Plant

		args = append(args, product, businessPartner, plant)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_by_storage_bin_by_batch_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` )`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductStockByStorageBinByBatch(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

//func (c *DPFMAPICaller) ProductStockAvailabilityByBatch(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *[]dpfm_api_output_formatter.ProductStockAvailabilityByBatch {
//	var args []interface{}
//	product := input.ProductStockByBatch.Product
//	businessPartner := input.ProductStockByBatch.BusinessPartner
//	plant := input.ProductStockByBatch.Plant
//	batch := input.ProductStockByBatch.Batch
//	productStockAvailabilityByBatch := input.ProductStockByBatch.ProductStockAvailabilityByBatch
//
//	cnt := 0
//	for _, v := range productStockAvailabilityByBatch {
//		args = append(args, product, businessPartner, plant, batch, v.ProductStockAvailabilityDate)
//		cnt++
//	}
//	repeat := strings.Repeat("(?,?,?,?,?),", cnt-1) + "(?,?,?,?,?)"
//
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_by_batch_data
//		WHERE (BusinessPartner, Product, Plant, Batch, ProductStockAvailabilityDate) IN ( `+repeat+` );`, args...,
//	)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//	defer rows.Close()
//
//	data, err := dpfm_api_output_formatter.ConvertToProductStockAvailabilityByBatch(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}
//func (c *DPFMAPICaller) ProductStockAvailabilityByOrder(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *[]dpfm_api_output_formatter.ProductStockAvailabilityByOrder {
//	var args []interface{}
//	product := input.ProductStockByOrder.Product
//	businessPartner := input.ProductStockByOrder.BusinessPartner
//	plant := input.ProductStockByOrder.Plant
//	batch := input.ProductStockByOrder.Batch
//	productStockAvailabilityByOrder := input.ProductStockByOrder.ProductStockAvailabilityByOrder
//
//	cnt := 0
//	for _, v := range productStockAvailabilityByOrder {
//		args = append(args, product, businessPartner, plant, batch, v.ProductStockAvailabilityDate)
//		cnt++
//	}
//	repeat := strings.Repeat("(?,?,?,?,?),", cnt-1) + "(?,?,?,?,?)"
//
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_by_Order_data
//		WHERE (BusinessPartner, Product, Plant, Batch, ProductStockAvailabilityDate) IN ( `+repeat+` );`, args...,
//	)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//	defer rows.Close()
//
//	data, err := dpfm_api_output_formatter.ConvertToProductStockAvailabilityByOrder(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}
//func (c *DPFMAPICaller) ProductStockAvailabilityByProject(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *[]dpfm_api_output_formatter.ProductStockAvailabilityByProject {
//	var args []interface{}
//	product := input.ProductStockByProject.Product
//	businessPartner := input.ProductStockByProject.BusinessPartner
//	plant := input.ProductStockByProject.Plant
//	batch := input.ProductStockByProject.Batch
//	productStockAvailabilityByProject := input.ProductStockByProject.ProductStockAvailabilityByProject
//
//	cnt := 0
//	for _, v := range productStockAvailabilityByProject {
//		args = append(args, product, businessPartner, plant, batch, v.ProductStockAvailabilityDate)
//		cnt++
//	}
//	repeat := strings.Repeat("(?,?,?,?,?),", cnt-1) + "(?,?,?,?,?)"
//
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_by_Project_data
//		WHERE (BusinessPartner, Product, Plant, Batch, ProductStockAvailabilityDate) IN ( `+repeat+` );`, args...,
//	)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//	defer rows.Close()
//
//	data, err := dpfm_api_output_formatter.ConvertToProductStockAvailabilityByProject(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}
//func (c *DPFMAPICaller) ProductStockAvailabilityByStorageBin(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *[]dpfm_api_output_formatter.ProductStockAvailabilityByStorageBin {
//	var args []interface{}
//	product := input.ProductStockByStorageBin.Product
//	businessPartner := input.ProductStockByStorageBin.BusinessPartner
//	plant := input.ProductStockByStorageBin.Plant
//	storageLocation := input.ProductStockByStorageBin.StorageLocation
//	storageBin := input.ProductStockByStorageBin.StorageBin
//	productStockAvailabilityByStorageBin := input.ProductStockByStorageBin.ProductStockAvailabilityByStorageBin
//
//	cnt := 0
//	for _, v := range productStockAvailabilityByStorageBin {
//		args = append(args, product, businessPartner, plant, storageLocation, storageBin, v.ProductStockAvailabilityDate)
//		cnt++
//	}
//	repeat := strings.Repeat("(?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?)"
//
//	rows, err := c.db.Query(
//		`SELECT *
//		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_by_storage_bin_data
//		WHERE (BusinessPartner, Product, Plant, StorageLocation, StorageBin, ProductStockAvailabilityDate) IN ( `+repeat+` );`, args...,
//	)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//	defer rows.Close()
//
//	data, err := dpfm_api_output_formatter.ConvertToProductStockAvailabilityByStorageBin(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}
//
//func (c *DPFMAPICaller) ProductStockAvailabilityByStorageBinByBatch(
//	mtx *sync.Mutex,
//	input *dpfm_api_input_reader.SDC,
//	output *dpfm_api_output_formatter.SDC,
//	errs *[]error,
//	log *logger.Logger,
//) *[]dpfm_api_output_formatter.ProductStockAvailabilityByStorageBinByBatch {
//	var args []interface{}
//	product := input.ProductStockByStorageBinByBatch.Product
//	businessPartner := input.ProductStockByStorageBinByBatch.BusinessPartner
//	plant := input.ProductStockByStorageBinByBatch.Plant
//	storageLocation := input.ProductStockByStorageBinByBatch.StorageLocation
//	storageBin := input.ProductStockByStorageBinByBatch.StorageBin
//	batch := input.ProductStockByStorageBinByBatch.Batch
//	productStockAvailabilityByStorageBinByBatch := input.ProductStockByStorageBinByBatch.ProductStockAvailabilityByStorageBinByBatch
//
//	cnt := 0
//	for _, v := range productStockAvailabilityByStorageBinByBatch {
//		args = append(args, product, businessPartner, plant, storageLocation, storageBin, batch, v.ProductStockAvailabilityDate)
//		cnt++
//	}
//	repeat := strings.Repeat("(?,?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?,?)"
//
//	rows, err := c.db.Query(
//		`SELECT *
//			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_by_storage_bin_by_batch_data
//			WHERE (BusinessPartner, Product, Plant, StorageLocation, StorageBin, Batch, ProductStockAvailabilityDate) IN ( `+repeat+` );`, args...,
//	)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//	defer rows.Close()
//
//	data, err := dpfm_api_output_formatter.ConvertToProductStockAvailabilityByStorageBinByBatch(rows)
//	if err != nil {
//		*errs = append(*errs, err)
//		return nil
//	}
//
//	return data
//}
