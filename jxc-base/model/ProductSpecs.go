package model

type ProductSpecs struct {
	Id        int    `json:"id" gorm:"column:id"`
	ShopId    string `json:"shop_id" gorm:"column:shop_id"`
	ProductId string `json:"product_id" gorm:"column:product_id"`
	Name      string `json:"name" gorm:"column:name"`
	Values    string `json:"values" gorm:"column:values"`
}
