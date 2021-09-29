package vo

type StockEntryRow struct {
	Id           int               `json:"id"` ////表单内顺序号，从1开始递增
	SkuId        string            `json:"sku_id"`
	Unit         string            `json:"unit"`
	Price        int               `json:"price"`
	StorageId    string            `json:"storage_id"`
	SupplierId   string            `json:"supplier_id"`
	ProductId    string            `json:"product_id"`
	SkuGuid      string            `json:"sku_guid"`
	SkuName      string            `json:"sku_name"`
	BarCode      string            `json:"bar_code"`
	Sku          map[string]string `json:"sku_specs"`
	PriceList    map[string]string `json:"price_list"`
	CategoryName string            `json:"category_name"`
	StorageName  string            `json:"storage_name"`
	SupplierName string            `json:"supplier_name"`
	Number       int               `json:"number"`
	Remark       string            `json:"remark"`
}

type StockEntryNote struct {
	Id        int             `json:"id"` //第一条记录的ID
	ShopId    string          `json:"shop_id"`
	EntryNo   string          `json:"entry_no"`
	Type      string          `json:"type"`
	OrderNo   string          `json:"order_no"`
	EntryTime string          `json:"entry_time"`
	Operator  string          `json:"operator"`
	SkuList   []StockEntryRow `json:"sku_list"`
}

type StockEntryRowDetail struct {
	StockEntryRow
	ShopId     string `json:"shop_id"`
	EntryNo    string `json:"entry_no"`
	Type       string `json:"type"`
	OrderNo    string `json:"order_no"`
	EntryTime  string `json:"entry_time"`
	SupplierId string `json:"supplier_id"`
	Operator   string `json:"operator"`
}
