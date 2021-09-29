package model

type SkuStock struct {
	Id         int    `json:"id" gorm:"column:id"`
	ShopId     string `json:"shop_id" gorm:"column:shop_id"`
	ProductId  string `json:"product_id" gorm:"column:product_id"`
	SkuId      string `json:"sku_id" gorm:"column:sku_id"`
	Name       string `json:"name" gorm:"column:name"`
	Stocks     int    `json:"stocks" gorm:"column:stocks"`
	BaseUnit   string `json:"base_unit" gorm:"column:base_unit"`
	StorageId  string `json:"storage_id" gorm:"column:storage_id"`
	CostPrice  int    `json:"cost_price" gorm:"column:cost_price"`
	LastPrice  int    `json:"last_price" gorm:"column:last_price"`
	UpdateTime string `json:"update_time" gorm:"column:update_time"`
}
