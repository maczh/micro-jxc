package model

type ProductCategory struct {
	Id         int    `json:"id" gorm:"column:id"`
	CategoryId string `json:"category_id" gorm:"column:category_id"`
	ShopId     string `json:"shop_id" gorm:"column:shop_id"`
	Name       string `json:"name" gorm:"column:name"`
	ParentId   string `json:"parent_id" gorm:"column:parent_id"`
	Level      int    `json:"level" gorm:"column:level"`
	SortNumber int    `json:"sort_number" gorm:"column:sort_number"`
}
