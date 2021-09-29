package service

import (
	"ququ.im/common"
	"ququ.im/jxc-api/nacos"
)

func ListProductSku(shopId, keyword, skusMap, productId, categoryId string, status, page, size int) common.Result {
	productSkuList, resultPage, err := nacos.ListProductSku(shopId, keyword, skusMap, productId, categoryId, status, page, size)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if page > 0 && size > 0 {
		return *common.SuccessWithPage(productSkuList, resultPage.Count, page, size, resultPage.Total)
	} else {
		return *common.Success(productSkuList)
	}
}

func GetProductSku(skuGuid, shopId, skuId string) common.Result {
	productSku, err := nacos.GetProductSku(shopId, skuId, skuGuid)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productSku == nil {
		return *common.Error(-1, "无此货品")
	}
	return *common.Success(productSku)
}

func SaveProductSku(shopId, productId, skusMap, skuId, skuGuid, skuName, priceList string) common.Result {
	productInfo, err := nacos.GetProductInfo(0, shopId, productId, "")
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productInfo == nil {
		return *common.Error(-1, "商品编码不存在")
	}
	productSku, err := nacos.SaveProductSku(shopId, productId, skusMap, skuId, skuGuid, skuName, priceList)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productSku)
}

func DeleteProductSku(skuGuid, shopId, skuId string) common.Result {
	return nacos.DeleteProductSku(skuGuid, shopId, skuId)
}

func UpdateProductSku(skuGuid, skuId, name, skuName, barCode, skusMap, priceList string, status int) common.Result {
	productSku, err := nacos.UpdateProductSku(skuGuid, skuId, name, skuName, barCode, skusMap, priceList, status)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productSku)
}

func IncrProductSkuSortNumber(skuGuid string) common.Result {
	productSkuList, err := nacos.IncrProductSkuSortNumber(skuGuid)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productSkuList)
}

func DecrProductSkuSortNumber(skuGuid string) common.Result {
	productSkuList, err := nacos.DecrProductSkuSortNumber(skuGuid)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productSkuList)
}
