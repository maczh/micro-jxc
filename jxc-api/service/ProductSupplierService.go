package service

import (
	"ququ.im/common"
	"ququ.im/jxc-api/nacos"
)

func ListProductSupplier(shopId, productId string) common.Result {
	productSupplierList, err := nacos.ListProductSupplier(shopId, productId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productSupplierList == nil {
		return *common.Error(-1, "未定义产品供应商")
	}
	return *common.Success(productSupplierList)
}

func GetProductSupplier(id int, shopId, productId, supplierId string) common.Result {
	productSupplier, err := nacos.GetProductSupplier(id, shopId, productId, supplierId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productSupplier == nil {
		return *common.Error(-1, "未定义此产品供应商")
	}
	return *common.Success(productSupplier)
}

func SaveProductSupplier(shopId, productId, supplierId, supplierName string) common.Result {
	productInfo, err := nacos.GetProductInfo(0, shopId, productId, "")
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productInfo == nil {
		return *common.Error(-1, "商品编码不存在")
	}
	productSupplier, err := nacos.SaveProductSupplier(shopId, productId, supplierId, supplierName)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productSupplier)
}

func UpdateProductSupplier(id int, shopId, supplierId, supplierName, productId string) common.Result {
	productInfo, err := nacos.GetProductInfo(0, shopId, productId, "")
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productInfo == nil {
		return *common.Error(-1, "商品编码不存在")
	}
	productSupplier, err := nacos.UpdateProductSupplier(id, shopId, productId, supplierId, supplierName)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productSupplier)
}

func DeleteProductSupplier(id int) common.Result {
	return nacos.DeleteProductSupplier(id)
}
