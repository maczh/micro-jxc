package service

import (
	"ququ.im/common"
	"ququ.im/jxc-api/nacos"
)

func ListProductSpecs(shopId, productId string) common.Result {
	productSpecsList, err := nacos.ListProductSpecs(shopId, productId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productSpecsList == nil {
		return *common.Error(-1, "未定义产品规格")
	}
	return *common.Success(productSpecsList)
}

func GetProductSpecs(id int, shopId, productId, specsName string) common.Result {
	productSpecs, err := nacos.GetProductSpecs(id, shopId, productId, specsName)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productSpecs == nil {
		return *common.Error(-1, "未定义产品规格")
	}
	return *common.Success(productSpecs)
}

func SaveProductSpecs(shopId, productId, specsName, specsValues string) common.Result {
	productInfo, err := nacos.GetProductInfo(0, shopId, productId, "")
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productInfo == nil {
		return *common.Error(-1, "商品编码不存在")
	}
	productSpecs, err := nacos.SaveProductSpecs(shopId, productId, specsName, specsValues)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productSpecs)
}

func UpdateProductSpecs(id int, shopId, productId, specsName, specsValues string) common.Result {
	productInfo, err := nacos.GetProductInfo(0, shopId, productId, "")
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productInfo == nil {
		return *common.Error(-1, "商品编码不存在")
	}
	productSpecs, err := nacos.UpdateProductSpecs(id, shopId, productId, specsName, specsValues)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productSpecs)
}

func DeleteProductSpecs(id int) common.Result {
	return nacos.DeleteProductSpecs(id)
}

func AddSpecsValue(id int, value string) common.Result {
	return nacos.AddSpecsValue(id, value)
}

func RemoveSpecsValue(id int, value string) common.Result {
	return nacos.RemoveSpecsValue(id, value)
}
