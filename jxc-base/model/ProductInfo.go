package model

type ProductInfo struct {
	Id         int    `json:"id" gorm:"column:id"`
	ShopId     string `json:"shop_id" gorm:"column:shop_id"`
	ProductId  string `json:"product_id" gorm:"column:product_id"`
	Name       string `json:"name" gorm:"column:name"`
	CategoryId string `json:"category_id" gorm:"column:category_id"`
	BaseUnit   string `json:"base_unit" gorm:"column:base_unit"`
	BarCode    string `json:"bar_code" gorm:"column:bar_code"`
	PyFull     string `json:"pinyin_full" gorm:"column:pinyin_full"`
	PyFirst    string `json:"pinyin_first" gorm:"column:pinyin_first"`
}
