package service

import (
	"ququ.im/common"
	"ququ.im/jxc-api/nacos"
)

func ListProductUnit(shopId, productId string) common.Result {
	productUnitList, err := nacos.ListProductUnit(shopId, productId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productUnitList == nil {
		return *common.Error(-1, "未定义产品单位换算")
	}
	return *common.Success(productUnitList)
}

func GetProductUnit(id int, shopId, productId, unit, baseUnit string) common.Result {
	productUnit, err := nacos.GetProductUnit(id, shopId, productId, unit, baseUnit)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productUnit == nil {
		return *common.Error(-1, "未定义此产品单位换算")
	}
	return *common.Success(productUnit)
}

func SaveProductUnit(shopId, productId, unit, baseUnit string, scale int) common.Result {
	productInfo, err := nacos.GetProductInfo(0, shopId, productId, "")
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productInfo == nil {
		return *common.Error(-1, "商品编码不存在")
	}
	productUnit, err := nacos.SaveProductUnit(shopId, productId, unit, baseUnit, scale)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productUnit)
}

func UpdateProductUnit(shopId, productId, unit, baseUnit string, scale int) common.Result {
	productInfo, err := nacos.GetProductInfo(0, shopId, productId, "")
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productInfo == nil {
		return *common.Error(-1, "商品编码不存在")
	}
	productUnit, err := nacos.UpdateProductUnit(shopId, productId, unit, baseUnit, scale)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productUnit)
}

func DeleteProductUnit(productUnitId int) common.Result {
	return nacos.DeleteProductUnit(productUnitId)
}
