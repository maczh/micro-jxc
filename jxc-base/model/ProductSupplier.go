package model

type ProductSupplier struct {
	Id           int    `json:"id" gorm:"column:id"`
	ShopId       string `json:"shop_id" gorm:"column:shop_id"`
	ProductId    string `json:"product_id" gorm:"column:product_id"`
	SupplierId   string `json:"supplier_id" gorm:"column:supplier_id"`
	SupplierName string `json:"supplier_name" gorm:"column:supplier_name"`
}
