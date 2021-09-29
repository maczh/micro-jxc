package model

type EntryType struct {
	Id     int    `json:"id" gorm:"column:id"`
	ShopId string `json:"shop_id" gorm:"column:shop_id"`
	Type   string `json:"type" gorm:"column:type"`
}
