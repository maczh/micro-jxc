package model

type ShopStorage struct {
	Id         int    `json:"id" gorm:"column:id"`
	ShopId     string `json:"shop_id" gorm:"column:shop_id"`
	StorageId  string `json:"storage_id" gorm:"column:storage_id"`
	Name       string `json:"name" gorm:"column:name"`
	Remark     string `json:"remark" gorm:"column:remark"`
	SortNumber int    `json:"sort_number" gorm:"column:sort_number"`
}
