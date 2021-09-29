package vo

import "ququ.im/jxc-base/model"

type StockMoveRowDetail struct {
	model.StockMove
	ProductId       string            `json:"product_id"`
	SkuGuid         string            `json:"sku_guid"`
	SkuName         string            `json:"sku_name"`
	BarCode         string            `json:"bar_code"`
	Sku             map[string]string `json:"sku_specs"`
	PriceList       map[string]string `json:"price_list"`
	FromStorageName string            `json:"from_storage_name"`
	ToStorageName   string            `json:"to_storage_name"`
}
