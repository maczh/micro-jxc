package dao

import (
	"errors"
	"ququ.im/common/cache"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/jxc-base/model"
)

const TABLE_PRODUCT_UNIT string = "product_unit"

func ListProductUnitByProductId(shopId, productId string) []model.ProductUnit {
	var productUnitList []model.ProductUnit
	err := cache.GetCache("productunit:"+shopId, productId, &productUnitList)
	if err == nil && len(productUnitList) > 0 {
		return productUnitList
	}
	config.Mysql.Table(TABLE_PRODUCT_UNIT).Where("shop_id = ? AND product_id = ? ", shopId, productId).Order("id ASC").Find(&productUnitList)
	if len(productUnitList) == 0 {
		return nil
	} else {
		cache.SetCache("productunit:"+shopId, productId, productUnitList)
		return productUnitList
	}
}

func GetProductUnit(shopId, productId, unit, baseUnit string) *model.ProductUnit {
	var productUnit model.ProductUnit
	err := cache.GetCache("productunit:scale:"+shopId, productId+":"+unit+":"+baseUnit, &productUnit)
	if err == nil && productUnit.Id > 0 {
		return &productUnit
	}
	config.Mysql.Table(TABLE_PRODUCT_UNIT).Where("shop_id = ? AND product_id = ? AND unit = ? AND base_unit = ?", shopId, productId, unit, baseUnit).First(&productUnit)
	if productUnit.Id == 0 {
		return nil
	} else {
		cache.SetCache("productunit:scale:"+shopId, productId+":"+unit+":"+baseUnit, productUnit)
		return &productUnit
	}
}

func GetProductUnitById(productUnitId int) *model.ProductUnit {
	var productUnit model.ProductUnit
	config.Mysql.Table(TABLE_PRODUCT_UNIT).Where("id = ? ", productUnitId).First(&productUnit)
	if productUnit.Id == 0 {
		return nil
	} else {
		return &productUnit
	}
}

func SaveProductUnit(productUnit *model.ProductUnit) (*model.ProductUnit, error) {
	s := GetProductUnit(productUnit.ShopId, productUnit.ProductId, productUnit.Unit, productUnit.BaseUnit)
	if s != nil {
		return s, nil
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_UNIT).Create(productUnit).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("productunit:"+productUnit.ShopId, productUnit.ProductId)

	return productUnit, nil
}

func UpdateProductUnit(productUnit *model.ProductUnit) (*model.ProductUnit, error) {
	if productUnit.Id == 0 {
		return nil, errors.New("未指定Id")
	}
	data := make(map[string]interface{})
	if productUnit.ProductId != "" {
		data["product_id"] = productUnit.ProductId
	}
	if productUnit.Unit != "" {
		data["unit"] = productUnit.Unit
	}
	if productUnit.BaseUnit != "" {
		data["base_unit"] = productUnit.BaseUnit
	}
	if productUnit.Scale > 0 {
		data["scale"] = productUnit.Scale
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_UNIT).Where("id = ?", productUnit.Id).Update(data).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("productunit:"+productUnit.ShopId, productUnit.ProductId)
	cache.DelCache("productunit:scale:"+productUnit.ShopId, productUnit.ProductId+":"+productUnit.Unit+":"+productUnit.BaseUnit)
	return productUnit, nil
}

func DeleteProductUnit(productUnitId int) error {
	if productUnitId == 0 {
		return errors.New("未指定商品单位换算规则编号")
	}
	productUnit := GetProductUnitById(productUnitId)
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_UNIT).Where("id = ?", productUnitId).Delete(model.ProductUnit{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	cache.DelCache("productunit:"+productUnit.ShopId, productUnit.ProductId)
	cache.DelCache("productunit:scale:"+productUnit.ShopId, productUnit.ProductId+":"+productUnit.Unit+":"+productUnit.BaseUnit)
	return nil
}
