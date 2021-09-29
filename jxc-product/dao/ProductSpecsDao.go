package dao

import (
	"errors"
	"ququ.im/common/cache"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/jxc-base/model"
)

const TABLE_PRODUCT_SPECS string = "product_specs"

func ListProductSpecsByProductId(shopId, productId string) []model.ProductSpecs {
	var productSpecsList []model.ProductSpecs
	err := cache.GetCache("specs:productid:"+shopId, productId, &productSpecsList)
	if err == nil && len(productSpecsList) > 0 {
		return productSpecsList
	}
	config.Mysql.Table(TABLE_PRODUCT_SPECS).Where("shop_id = ? AND product_id = ? ", shopId, productId).Order("id ASC").Find(&productSpecsList)
	if len(productSpecsList) == 0 {
		return nil
	} else {
		cache.SetCache("specs:productid:"+shopId, productId, productSpecsList)
		return productSpecsList
	}
}

func GetProductSpecs(shopId, productId, specs string) *model.ProductSpecs {
	var productSpecs model.ProductSpecs
	err := cache.GetCache("specs:product:"+shopId+":"+productId, specs, &productSpecs)
	if err == nil && productSpecs.Id > 0 {
		return &productSpecs
	}
	config.Mysql.Table(TABLE_PRODUCT_SPECS).Where("shop_id = ? AND product_id = ? AND name = ?", shopId, productId, specs).First(&productSpecs)
	if productSpecs.Id == 0 {
		return nil
	} else {
		cache.SetCache("specs:product:"+shopId+":"+productId, specs, productSpecs)
		return &productSpecs
	}
}

func GetProductSpecsById(productSpecsId int) *model.ProductSpecs {
	var productSpecs model.ProductSpecs
	config.Mysql.Table(TABLE_PRODUCT_SPECS).Where("id = ? ", productSpecsId).First(&productSpecs)
	if productSpecs.Id == 0 {
		return nil
	} else {
		return &productSpecs
	}
}

func SaveProductSpecs(productSpecs *model.ProductSpecs) (*model.ProductSpecs, error) {
	s := GetProductSpecs(productSpecs.ShopId, productSpecs.ProductId, productSpecs.Name)
	if s != nil {
		return s, nil
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_SPECS).Create(productSpecs).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("specs:productid:"+productSpecs.ShopId, productSpecs.ProductId)
	cache.DelCache("specs:product:"+productSpecs.ShopId+":"+productSpecs.ProductId, productSpecs.Name)
	return productSpecs, nil
}

func UpdateProductSpecs(productSpecs *model.ProductSpecs) (*model.ProductSpecs, error) {
	if productSpecs.Id == 0 {
		return nil, errors.New("未指定Id")
	}
	data := make(map[string]interface{})
	if productSpecs.ShopId != "" {
		data["shop_id"] = productSpecs.ShopId
	}
	if productSpecs.ProductId != "" {
		data["product_id"] = productSpecs.ProductId
	}
	if productSpecs.Name != "" {
		data["name"] = productSpecs.Name
	}
	if productSpecs.Values != "" {
		data["values"] = productSpecs.Values
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_SPECS).Where("id = ?", productSpecs.Id).Update(data).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("specs:productid", productSpecs.ProductId)
	cache.DelCache("specs:product:"+productSpecs.ShopId+":"+productSpecs.ProductId, productSpecs.Name)
	return productSpecs, nil
}

func DeleteProductSpecs(productSpecsId int) error {
	if productSpecsId == 0 {
		return errors.New("未指定商品规格编号")
	}
	productSpecs := GetProductSpecsById(productSpecsId)
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_SPECS).Where("id = ?", productSpecsId).Delete(model.ProductSpecs{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	cache.DelCache("specs:productid", productSpecs.ProductId)
	cache.DelCache("specs:product:"+productSpecs.ShopId+":"+productSpecs.ProductId, productSpecs.Name)
	return nil
}
