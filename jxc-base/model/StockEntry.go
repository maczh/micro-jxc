package model

type StockEntry struct {
	Id         int    `json:"id" gorm:"column:id"`
	ShopId     string `json:"shop_id" gorm:"column:shop_id"`
	EntryNo    string `json:"entry_no" gorm:"column:entry_no"`
	Type       string `json:"type" gorm:"column:type"`
	SkuId      string `json:"sku_id" gorm:"column:sku_id"`
	Unit       string `json:"unit" gorm:"column:unit"`
	OrderNo    string `json:"order_no" gorm:"column:order_no"`
	Price      int    `json:"price" gorm:"column:price"`
	StorageId  string `json:"storage_id" gorm:"column:storage_id"`
	Number     int    `json:"number" gorm:"column:number"`
	EntryTime  string `json:"entry_time" gorm:"column:entry_time"`
	SupplierId string `json:"supplier_id" gorm:"column:supplier_id"`
	Remark     string `json:"remark" gorm:"column:remark"`
	Operator   string `json:"operator" gorm:"column:operator"`
}
