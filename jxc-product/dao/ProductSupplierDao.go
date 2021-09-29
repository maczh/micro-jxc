package dao

import (
	"errors"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/jxc-base/model"
)

const TABLE_PRODUCT_SUPPLIER string = "product_supplier"

func ListProductSupplierByProductId(shopId, productId string) []model.ProductSupplier {
	var productSupplierList []model.ProductSupplier
	config.Mysql.Table(TABLE_PRODUCT_SUPPLIER).Where("shop_id = ? AND product_id = ? ", shopId, productId).Order("id ASC").Find(&productSupplierList)
	if len(productSupplierList) == 0 {
		return nil
	} else {
		return productSupplierList
	}
}

func GetProductSupplier(shopId, productId, supplierId string) *model.ProductSupplier {
	var productSupplier model.ProductSupplier
	config.Mysql.Table(TABLE_PRODUCT_SUPPLIER).Where("shop_id = ? AND product_id = ? AND supplier_id = ?", shopId, productId, supplierId).First(&productSupplier)
	if productSupplier.Id == 0 {
		return nil
	} else {
		return &productSupplier
	}
}

func GetProductSupplierById(productSupplierId int) *model.ProductSupplier {
	var productSupplier model.ProductSupplier
	config.Mysql.Table(TABLE_PRODUCT_SUPPLIER).Where("id = ?", productSupplierId).First(&productSupplier)
	if productSupplier.Id == 0 {
		return nil
	} else {
		return &productSupplier
	}
}

func SaveProductSupplier(productSupplier *model.ProductSupplier) (*model.ProductSupplier, error) {
	s := GetProductSupplier(productSupplier.ShopId, productSupplier.ProductId, productSupplier.SupplierId)
	if s != nil {
		return s, nil
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_SUPPLIER).Create(productSupplier).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return productSupplier, nil
}

func UpdateProductSupplier(productSupplier *model.ProductSupplier) (*model.ProductSupplier, error) {
	if productSupplier.Id == 0 {
		return nil, errors.New("未指定Id")
	}
	data := make(map[string]interface{})
	if productSupplier.ShopId != "" {
		data["shop_id"] = productSupplier.ShopId
	}
	if productSupplier.ProductId != "" {
		data["product_id"] = productSupplier.ProductId
	}
	if productSupplier.SupplierId != "" {
		data["supplier_id"] = productSupplier.SupplierId
	}
	if productSupplier.SupplierName != "" {
		data["supplier_name"] = productSupplier.SupplierName
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_SUPPLIER).Where("id = ?", productSupplier.Id).Update(data).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return productSupplier, nil
}

func DeleteProductSupplier(productSupplierId int) error {
	if productSupplierId == 0 {
		return errors.New("未指定商品供货商编号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_SUPPLIER).Where("id = ?", productSupplierId).Delete(model.ProductSupplier{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
