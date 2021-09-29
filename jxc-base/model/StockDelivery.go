package model

type StockDelivery struct {
	Id           int    `json:"id" gorm:"column:id"`
	ShopId       string `json:"shop_id" gorm:"column:shop_id"`
	DeliveryNo   string `json:"delivery_no" gorm:"column:delivery_no"`
	Type         string `json:"type" gorm:"column:type"`
	SkuId        string `json:"sku_id" gorm:"column:sku_id"`
	Unit         string `json:"unit" gorm:"column:unit"`
	OrderNo      string `json:"order_no" gorm:"column:order_no"`
	Cost         int    `json:"cost" gorm:"column:cost"`
	Price        int    `json:"price" gorm:"column:price"`
	StorageId    string `json:"storage_id" gorm:"column:storage_id"`
	MultiStorage string `json:"multi_storage" gorm:"column:multi_storage"`
	Number       int    `json:"number" gorm:"column:number"`
	DeliveryTime string `json:"delivery_time" gorm:"column:delivery_time"`
	CustomerId   string `json:"customer_id" gorm:"column:customer_id"`
	CustomerName string `json:"customer_name" gorm:"column:customer_name"`
	Remark       string `json:"remark" gorm:"column:remark"`
	Operator     string `json:"operator" gorm:"column:operator"`
}
