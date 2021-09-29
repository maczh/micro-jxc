package service

import (
	"ququ.im/common"
	"ququ.im/jxc-base/model"
	"ququ.im/jxc-product/dao"
)

func ListProductSupplier(shopId, productId string) common.Result {
	var productSupplierList []model.ProductSupplier
	productSupplierList = dao.ListProductSupplierByProductId(shopId, productId)
	if productSupplierList == nil {
		return *common.Error(-1, "未定义产品供应商")
	}
	return *common.Success(productSupplierList)
}

func GetProductSupplier(id int, shopId, productId, supplierId string) common.Result {
	var productSupplier *model.ProductSupplier
	if id > 0 {
		productSupplier = dao.GetProductSupplierById(id)
	} else if shopId != "" && productId != "" && supplierId != "" {
		productSupplier = dao.GetProductSupplier(shopId, productId, supplierId)
	} else {
		return *common.Error(-1, "参数不全")
	}
	if productSupplier == nil {
		return *common.Error(-1, "未定义此产品供应商")
	}
	return *common.Success(productSupplier)
}

func SaveProductSupplier(shopId, productId, supplierId, supplierName string) common.Result {
	productSupplier := dao.GetProductSupplier(shopId, productId, supplierId)
	if productSupplier != nil {
		return *common.Error(-1, "此服务商已经存在")
	}
	productSupplier = new(model.ProductSupplier)
	productSupplier.ShopId = shopId
	productSupplier.ProductId = productId
	productSupplier.SupplierId = supplierId
	productSupplier.SupplierName = supplierName
	_, err := dao.SaveProductSupplier(productSupplier)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(dao.GetProductSupplier(shopId, productId, supplierId))
}

func UpdateProductSupplier(id int, shopId, supplierId, supplierName, productId string) common.Result {
	productSupplier := dao.GetProductSupplierById(id)
	if productSupplier == nil {
		return *common.Error(-1, "无此商品供应商ID")
	}
	productSupplier.ShopId = shopId
	productSupplier.ProductId = productId
	productSupplier.SupplierId = supplierId
	productSupplier.SupplierName = supplierName
	_, err := dao.UpdateProductSupplier(productSupplier)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(dao.GetProductSupplier(shopId, productId, supplierId))
}

func DeleteProductSupplier(id int) common.Result {
	err := dao.DeleteProductSupplier(id)
	if err != nil {
		return *common.Error(-1, "无此商品分类记录")
	}
	return *common.Success(nil)
}
