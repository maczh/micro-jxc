package service

import (
	"ququ.im/common"
	"ququ.im/jxc-base/model"
	"ququ.im/jxc-product/dao"
)

func ListProductUnit(shopId, productId string) common.Result {
	var productUnitList []model.ProductUnit
	productUnitList = dao.ListProductUnitByProductId(shopId, productId)
	if productUnitList == nil {
		return *common.Error(-1, "未定义产品单位换算")
	}
	return *common.Success(productUnitList)
}

func GetProductUnit(shopId, productId, unit, baseUnit string) common.Result {
	productUnit := dao.GetProductUnit(shopId, productId, unit, baseUnit)
	if productUnit == nil {
		return *common.Error(-1, "未定义此产品单位换算")
	}
	return *common.Success(productUnit)
}

func SaveProductUnit(shopId, productId, unit, baseUnit string, scale int) common.Result {
	productUnit := dao.GetProductUnit(shopId, productId, unit, baseUnit)
	if productUnit != nil {
		return *common.Success(productUnit)
	}
	productUnit = new(model.ProductUnit)
	productUnit.ShopId = shopId
	productUnit.Unit = unit
	productUnit.ProductId = productId
	productUnit.BaseUnit = baseUnit
	productUnit.Scale = scale
	_, err := dao.SaveProductUnit(productUnit)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(dao.GetProductUnit(shopId, productId, unit, baseUnit))
}

func UpdateProductUnit(shopId, productId, unit, baseUnit string, scale int) common.Result {
	productUnit := dao.GetProductUnit(shopId, productId, unit, baseUnit)
	if productUnit == nil {
		return *common.Error(-1, "无此商品单位换算")
	}
	productUnit.Unit = unit
	productUnit.ProductId = productId
	productUnit.BaseUnit = baseUnit
	productUnit.Scale = scale
	_, err := dao.UpdateProductUnit(productUnit)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(dao.GetProductUnit(shopId, productId, unit, baseUnit))
}

func DeleteProductUnit(productUnitId int) common.Result {
	err := dao.DeleteProductUnit(productUnitId)
	if err != nil {
		return *common.Error(-1, "无此商品单位换算")
	}
	return *common.Success(nil)
}
