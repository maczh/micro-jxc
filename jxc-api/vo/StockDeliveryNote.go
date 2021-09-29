package vo

type StockDeliveryRow struct {
	Id           int               `json:"id" gorm:"column:id"` //表单内顺序号，从1开始递增
	SkuId        string            `json:"sku_id" gorm:"column:sku_id"`
	Unit         string            `json:"unit" gorm:"column:unit"`
	Cost         int               `json:"cost"`
	Price        int               `json:"price" gorm:"column:price"`
	StorageId    string            `json:"storage_id" gorm:"column:storage_id"`
	CustomerId   string            `json:"customer_id"`
	ProductId    string            `json:"product_id"`
	SkuGuid      string            `json:"sku_guid"`
	SkuName      string            `json:"sku_name"`
	BarCode      string            `json:"bar_code"`
	Sku          map[string]string `json:"sku_specs"`
	PriceList    map[string]string `json:"price_list"`
	CategoryName string            `json:"category_name"`
	StorageName  string            `json:"storage_name"`
	CustomerName string            `json:"customer_name"`
	Number       int               `json:"number" gorm:"column:number"`
	Remark       string            `json:"remark" gorm:"column:remark"`
}

type StockDeliveryNote struct {
	Id           int                `json:"id"` //第一条记录的ID
	ShopId       string             `json:"shop_id"`
	DeliveryNo   string             `json:"delivery_no"`
	Type         string             `json:"type"`
	OrderNo      string             `json:"order_no"`
	DeliveryTime string             `json:"delivery_time"`
	Operator     string             `json:"operator"`
	SkuList      []StockDeliveryRow `json:"sku_list"`
}

type StockDeliveryRowDetail struct {
	StockDeliveryRow
	ShopId       string `json:"shop_id"`
	DeliveryNo   string `json:"delivery_no"`
	Type         string `json:"type"`
	OrderNo      string `json:"order_no"`
	DeliveryTime string `json:"delivery_time"`
	SupplierId   string `json:"supplier_id"`
	Operator     string `json:"operator"`
}
