package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-stock-reads-rmq-kube/DPFM_API_Output_Formatter"
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
	var productStockAvailability *dpfm_api_output_formatter.ProductStockAvailability
	for _, fn := range accepter {
		switch fn {
		case "ProductStock":
			func() {
				productStock = c.ProductStock(mtx, input, output, errs, log)
			}()
		case "ProductStockAvailability":
			func() {
				productStockAvailability = c.ProductStockAvailability(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		ProductStock:             productStock,
		ProductStockAvailability: productStockAvailability,
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
	businessPartner := input.ProductStock.BusinessPartner
	product := input.ProductStock.Product
	plant := input.ProductStock.Plant

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Product, Plant, StorageLocation, Batch, OrderID, OrderItem, Project, 
		InventoryStockType, InventorySpecialStockType, ProductStock
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_product_stock_data
		WHERE (BusinessPartner, Product, Plant) = (?, ?, ?);`, businessPartner, product, plant,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToProductStock(input, rows)
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
) *dpfm_api_output_formatter.ProductStockAvailability {
	businessPartner := input.ProductStock.BusinessPartner
	product := input.ProductStock.Product
	plant := input.ProductStock.Plant
	productStockAvailabilityDate := input.ProductStock.ProductStockAvailability.ProductStockAvailabilityDate

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Product, Plant, Batch, BatchValidityEndDate, OrderID, OrderItem, Project, 
		InventoryStockType, InventorySpecialStockType, ProductStockAvailabilityDate, AvailableProductStock
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_stock_availability_data
		WHERE (BusinessPartner, Product, Plant, ProductStockAvailabilityDate) = (?, ?, ?, ?);`, businessPartner, product, plant, productStockAvailabilityDate,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToProductStockAvailability(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
