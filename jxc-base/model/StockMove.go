package model

type StockMove struct {
	Id            int    `json:"id" gorm:"column:id"`
	ShopId        string `json:"shop_id" gorm:"column:shop_id"`
	MoveNo        string `json:"move_no" gorm:"column:move_no"`
	SkuId         string `json:"sku_id" gorm:"column:sku_id"`
	Number        int    `json:"number" gorm:"column:number"`
	Unit          string `json:"unit" gorm:"column:unit"`
	FromStorageId string `json:"from_storage_id" gorm:"column:from_storage_id"`
	ToStorageId   string `json:"to_storage_id" gorm:"column:to_storage_id"`
	MoveTime      string `json:"move_time" gorm:"column:move_time"`
	Remark        string `json:"remark" gorm:"column:remark"`
	Operator      string `json:"operator" gorm:"column:operator"`
}
