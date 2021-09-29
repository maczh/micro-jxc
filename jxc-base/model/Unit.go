package model

type Unit struct {
	Id     int    `json:"id" gorm:"column:id"`
	ShopId string `json:"shop_id" gorm:"column:shop_id"`
	Unit   string `json:"unit" gorm:"column:unit"`
}
