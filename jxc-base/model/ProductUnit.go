package model

type ProductUnit struct {
	Id        int    `json:"id" gorm:"column:id"`
	ShopId    string `json:"shop_id" gorm:"column:shop_id"`
	ProductId string `json:"product_id" gorm:"column:product_id"`
	Unit      string `json:"unit" gorm:"column:unit"`
	BaseUnit  string `json:"base_unit" gorm:"column:base_unit"`
	Scale     int    `json:"scale" gorm:"column:scale"`
}
