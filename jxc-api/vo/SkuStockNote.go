package vo

import "ququ.im/jxc-base/model"

type SkuStockNote struct {
	model.SkuStock
	SkuName      string            `json:"sku_name"`
	CategoryId   string            `json:"category_id"`
	BarCode      string            `json:"bar_code"`
	Sku          map[string]string `json:"sku_specs"`
	PriceList    map[string]string `json:"price_list"`
	CategoryName string            `json:"category_name"`
	StorageName  string            `json:"storage_name"`
}
