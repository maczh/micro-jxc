package service

import (
	"ququ.im/common"
	"ququ.im/jxc-api/nacos"
)

func ListProductInfo(shopId, keyword, categoryId string, page, size int) common.Result {
	productInfoList, resultPage, err := nacos.ListProductInfo(shopId, categoryId, keyword, page, size)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if page > 0 && size > 0 {
		return *common.SuccessWithPage(productInfoList, resultPage.Count, page, size, resultPage.Total)
	} else {
		return *common.Success(productInfoList)
	}
}

func GetProductInfo(id int, shopId, productId, barCode string) common.Result {
	productInfo, err := nacos.GetProductInfo(id, shopId, productId, barCode)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productInfo == nil {
		return *common.Error(-1, "未定义产品")
	}
	return *common.Success(productInfo)
}

func SaveProductInfo(productId, shopId, productName, baseUnit, barCode, categoryId string) common.Result {
	productCategory, err := nacos.GetProductCategory(0, shopId, categoryId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productCategory == nil {
		return *common.Error(-1, "商品分类编码不存在")
	}
	return nacos.SaveProductInfo(shopId, categoryId, productName, productId, barCode, baseUnit)
}

func UpdateProductInfo(id int, productId, categoryId, productName, baseUnit, barCode string) common.Result {
	p, _ := nacos.GetProductInfo(id, "", "", "")
	productCategory, err := nacos.GetProductCategory(0, p.ShopId, categoryId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productCategory == nil {
		return *common.Error(-1, "商品分类编码不存在")
	}
	return nacos.UpdateProductInfo(id, categoryId, productName, productId, baseUnit, barCode)
}

func DeleteProductInfo(id int) common.Result {
	return nacos.DeleteProductInfo(id)
}
