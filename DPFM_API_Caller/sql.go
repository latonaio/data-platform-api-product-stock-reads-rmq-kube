package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Output_Formatter"
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
	var productStock *dpfm_api_output_formatter.ProductStock
	var productStockByBatch *dpfm_api_output_formatter.ProductStockByBatch
	var productStockByStorageBin *dpfm_api_output_formatter.ProductStockByStorageBin
	var productStockByStorageBinByBatch *dpfm_api_output_formatter.ProductStockByStorageBinByBatch
	var productStockAvailability *[]dpfm_api_output_formatter.ProductStockAvailability
	var productStockAvailabilityByBatch *[]dpfm_api_output_formatter.ProductStockAvailabilityByBatch
	var productStockAvailabilityByStorageBin *[]dpfm_api_output_formatter.ProductStockAvailabilityByStorageBin
	var productStockAvailabilityByStorageBinByBatch *[]dpfm_api_output_formatter.ProductStockAvailabilityByStorageBinByBatch
	for _, fn := range accepter {
		switch fn {
		case "ProductStock":
			func() {
				productStock = c.ProductStock(mtx, input, output, errs, log)
			}()
		case "ProductStockByBatch":
			func() {
				productStockByBatch = c.ProductStockByBatch(mtx, input, output, errs, log)
			}()
		case "ProductStockByStorageBin":
			func() {
				productStockByStorageBin = c.ProductStockByStorageBin(mtx, input, output, errs, log)
			}()
		case "ProductStockByStorageBinByBatch":
			func() {
				productStockByStorageBinByBatch = c.ProductStockByStorageBinByBatch(mtx, input, output, errs, log)
			}()
		case "ProductStockAvailability":
			func() {
				productStockAvailability = c.ProductStockAvailability(mtx, input, output, errs, log)
			}()
		case "ProductStockAvailabilityByBatch":
			func() {
				productStockAvailabilityByBatch = c.ProductStockAvailabilityByBatch(mtx, input, output, errs, log)
			}()
		case "ProductStockAvailabilityByStorageBin":
			func() {
				productStockAvailabilityByStorageBin = c.ProductStockAvailabilityByStorageBin(mtx, input, output, errs, log)
			}()
		case "ProductStockAvailabilityByStorageBinByBatch":
			func() {
				productStockAvailabilityByStorageBinByBatch = c.ProductStockAvailabilityByStorageBinByBatch(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		ProductStock:                                productStock,
		ProductStockByBatch:                         productStockByBatch,
		ProductStockByStorageBin:                    productStockByStorageBin,
		ProductStockByStorageBinByBatch:             productStockByStorageBinByBatch,
		ProductStockAvailability:                    productStockAvailability,
		ProductStockAvailabilityByBatch:             productStockAvailabilityByBatch,
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
) *dpfm_api_output_formatter.ProductStock {
	product := input.ProductStock.Product
	businessPartner := input.ProductStock.BusinessPartner
	plant := input.ProductStock.Plant

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_product_stock_data
		WHERE (Product, BusinessPartner, Plant) = (?, ?, ?);`, product, businessPartner, plant,
	)
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

func (c *DPFMAPICaller) ProductStockByBatch(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ProductStockByBatch {
	product := input.ProductStockByBatch.Product
	businessPartner := input.ProductStockByBatch.BusinessPartner
	plant := input.ProductStockByBatch.Plant
	batch := input.ProductStockByBatch.Batch

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_by_batch_data
		WHERE (Product, BusinessPartner, Plant, Batch) = (?, ?, ?, ?);`, product, businessPartner, plant, batch,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductStockByBatch(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductStockByStorageBin(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ProductStockByStorageBin {
	product := input.ProductStockByStorageBin.Product
	businessPartner := input.ProductStockByStorageBin.BusinessPartner
	plant := input.ProductStockByStorageBin.Plant
	storageLocation := input.ProductStockByStorageBin.StorageLocation
	storageBin := input.ProductStockByStorageBin.StorageBin

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_by_storage_bin_data
		WHERE (Product, BusinessPartner, Plant, StorageLocation, StorageBin) = (?, ?, ?, ?, ?);`, product, businessPartner, plant, storageBin, storageLocation,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductStockByStorageBin(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductStockByStorageBinByBatch(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ProductStockByStorageBinByBatch {
	product := input.ProductStockByStorageBinByBatch.Product
	businessPartner := input.ProductStockByStorageBinByBatch.BusinessPartner
	plant := input.ProductStockByStorageBinByBatch.Plant
	storageLocation := input.ProductStockByStorageBin.StorageLocation
	storageBin := input.ProductStockByStorageBin.StorageBin
	batch := input.ProductStockByStorageBinByBatch.Batch

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_by_storage_bin_by_batch_data
		WHERE (Product, BusinessPartner, Plant, StorageLocation, StorageBin, Batch) = (?, ?, ?, ?, ?, ?);`, product, businessPartner, plant, storageLocation, storageBin, batch,
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

func (c *DPFMAPICaller) ProductStockAvailability(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductStockAvailability {
	var args []interface{}
	product := input.ProductStock.Product
	businessPartner := input.ProductStock.BusinessPartner
	plant := input.ProductStock.Plant
	productStockAvailability := input.ProductStock.ProductStockAvailability

	cnt := 0
	for _, v := range productStockAvailability {
		args = append(args, product, businessPartner, plant, v.ProductStockAvailabilityDate)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_data
		WHERE (BusinessPartner, Product, Plant, ProductStockAvailabilityDate) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductStockAvailability(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductStockAvailabilityByBatch(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductStockAvailabilityByBatch {
	var args []interface{}
	product := input.ProductStockByBatch.Product
	businessPartner := input.ProductStockByBatch.BusinessPartner
	plant := input.ProductStockByBatch.Plant
	batch := input.ProductStockByBatch.Batch
	productStockAvailabilityByBatch := input.ProductStockByBatch.ProductStockAvailabilityByBatch

	cnt := 0
	for _, v := range productStockAvailabilityByBatch {
		args = append(args, product, businessPartner, plant, batch, v.ProductStockAvailabilityDate)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?),", cnt-1) + "(?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_by_batch_data
		WHERE (BusinessPartner, Product, Plant, Batch, ProductStockAvailabilityDate) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductStockAvailabilityByBatch(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductStockAvailabilityByStorageBin(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductStockAvailabilityByStorageBin {
	var args []interface{}
	product := input.ProductStockByStorageBin.Product
	businessPartner := input.ProductStockByStorageBin.BusinessPartner
	plant := input.ProductStockByStorageBin.Plant
	storageLocation := input.ProductStockByStorageBin.StorageLocation
	storageBin := input.ProductStockByStorageBin.StorageBin
	productStockAvailabilityByStorageBin := input.ProductStockByStorageBin.ProductStockAvailabilityByStorageBin

	cnt := 0
	for _, v := range productStockAvailabilityByStorageBin {
		args = append(args, product, businessPartner, plant, storageLocation, storageBin, v.ProductStockAvailabilityDate)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_by_storage_bin_data
		WHERE (BusinessPartner, Product, Plant, StorageLocation, StorageBin, ProductStockAvailabilityDate) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductStockAvailabilityByStorageBin(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductStockAvailabilityByStorageBinByBatch(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductStockAvailabilityByStorageBinByBatch {
	var args []interface{}
	product := input.ProductStockByStorageBinByBatch.Product
	businessPartner := input.ProductStockByStorageBinByBatch.BusinessPartner
	plant := input.ProductStockByStorageBinByBatch.Plant
	storageLocation := input.ProductStockByStorageBinByBatch.StorageLocation
	storageBin := input.ProductStockByStorageBinByBatch.StorageBin
	batch := input.ProductStockByStorageBinByBatch.Batch
	productStockAvailabilityByStorageBinByBatch := input.ProductStockByStorageBinByBatch.ProductStockAvailabilityByStorageBinByBatch

	cnt := 0
	for _, v := range productStockAvailabilityByStorageBinByBatch {
		args = append(args, product, businessPartner, plant, storageLocation, storageBin, batch, v.ProductStockAvailabilityDate)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?,?,?),", cnt-1) + "(?,?,?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_by_storage_bin_by_batch_data
			WHERE (BusinessPartner, Product, Plant, StorageLocation, StorageBin, Batch, ProductStockAvailabilityDate) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductStockAvailabilityByStorageBinByBatch(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
