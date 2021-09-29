package service

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
	"ququ.im/jxc-product/dao"
	"ququ.im/jxc-product/mongo"
)

func ListProductSku(shopId, keyword, skusMap, productId, categoryId string, status, page, size int) common.Result {
	if shopId == "" {
		return *common.Error(-1, "商户编号不可为空")
	}
	sku := make(map[string]string)
	if skusMap != "" {
		utils.FromJSON(skusMap, &sku)
	}
	if page > 0 && size > 0 {
		count, err := mongo.CountProductSku(shopId, keyword, productId, categoryId, sku, status)
		if err != nil {
			return *common.Error(-1, "查询异常:"+err.Error())
		}
		productSkuList, err := mongo.ListProductSku(shopId, keyword, productId, categoryId, sku, status, page, size)
		if err != nil {
			return *common.Error(-1, "查询异常:"+err.Error())
		}
		return *common.SuccessWithPage(productSkuList, count/size+1, page, size, count)
	} else {
		productSkuList, err := mongo.ListProductSku(shopId, keyword, productId, categoryId, sku, status, 0, 0)
		if err != nil {
			return *common.Error(-1, "查询异常:"+err.Error())
		}
		return *common.Success(productSkuList)
	}
}

func GetProductSku(skuGuid, shopId, skuId string) common.Result {
	var productSku *model.ProductSku
	if skuGuid != "" {
		productSku = mongo.GetProductSkuBySkuGuid(skuGuid)
	}
	if skuId != "" {
		productSku = mongo.GetProductSkuBySkuid(shopId, skuId)
	}
	if productSku == nil {
		return *common.Error(-1, "无此货品")
	}
	return *common.Success(productSku)
}

func SaveProductSku(shopId, productId, skusMap, skuId, skuGuid, skuName, priceList string) common.Result {
	if shopId == "" {
		return *common.Error(-1, "商户账号不可为空")
	}
	if productId == "" {
		return *common.Error(-1, "商品编号不可为空")
	}
	if skusMap == "" {
		return *common.Error(-1, "货品规格参数不可为空")
	}
	productInfo := dao.GetProductInfo(shopId, productId)
	if productInfo == nil {
		return *common.Error(-1, "无此商品")
	}
	sku := make(map[string]string)
	prices := make(map[string]string)
	utils.FromJSON(skusMap, &sku)
	if priceList != "" {
		utils.FromJSON(priceList, &prices)
	}
	var productSku *model.ProductSku
	if skuGuid != "" {
		productSku = mongo.GetProductSkuBySkuGuid(skuGuid)
	}
	if skuId != "" {
		productSku = mongo.GetProductSkuBySkuid(productInfo.ShopId, skuId)
	}
	if productSku == nil {
		productSku = new(model.ProductSku)
	}
	productSku.ProductId = productId
	productSku.ShopId = shopId
	productSku.BaseUnit = productInfo.BaseUnit
	productSku.Name = productInfo.Name
	productSku.BarCode = productInfo.BarCode
	productSku.CategoryId = productInfo.CategoryId
	productSku.Sku = sku
	productSku.PriceList = prices
	productSku.SkuId = skuId
	productSku.SkuGuid = skuGuid
	productSku.SkuName = skuName
	productSku.Status = 1
	//获取排序号
	productSkuList, _ := mongo.ListProductSku(productInfo.ShopId, "", productId, "", nil, -1, 0, 0)
	if productSkuList != nil && len(productSkuList) > 0 {
		productSku.SortNumber = productSkuList[len(productSkuList)-1].SortNumber + 1
	} else {
		productSku.SortNumber = 101
	}
	productSku = mongo.SaveProductSku(productSku)
	return *common.Success(productSku)
}

func DeleteProductSku(skuGuid, shopId, skuId string) common.Result {
	if skuGuid == "" && (shopId != "" && skuId != "") {
		productSku := mongo.GetProductSkuBySkuid(shopId, skuId)
		if productSku == nil {
			return *common.Error(-1, "无此货品")
		}
		skuGuid = productSku.SkuGuid
	}
	err := mongo.DeleteProductSku(skuGuid)
	if err != nil {
		return *common.Error(-1, "删除货品失败:"+err.Error())
	}
	return *common.Success(nil)
}

func UpdateProductSku(skuGuid, skuId, name, skuName, barCode, skusMap, priceList string, status int) common.Result {
	productSku := mongo.GetProductSkuBySkuGuid(skuGuid)
	if productSku == nil {
		return *common.Error(-1, "无此货品")
	}
	if skuId != "" {
		s := mongo.GetProductSkuBySkuid(productSku.ShopId, skuId)
		if s != nil && skuId != productSku.SkuId {
			return *common.Error(-1, "货品编码skuId已存在，无法修改")
		}
		productSku.SkuId = skuId
	}
	if name != "" {
		productSku.Name = name
	}
	if skuName != "" {
		productSku.SkuName = skuName
	}
	if barCode != "" {
		productSku.BarCode = barCode
	}
	if status > -1 {
		productSku.Status = status
	}
	if skusMap != "" {
		sku := make(map[string]string)
		utils.FromJSON(skusMap, &sku)
		productSku.Sku = sku
	}
	if priceList != "" {
		prices := make(map[string]string)
		utils.FromJSON(priceList, &prices)
		productSku.PriceList = prices
	}
	productSku = mongo.UpdateProductSku(productSku)
	return *common.Success(productSku)
}

func IncrProductSkuSortNumber(skuGuid string) common.Result {
	if skuGuid == "" {
		return *common.Error(-1, "货品全局编码不可为空")
	}
	return *common.Success(mongo.IncrProductSkuSortNumber(skuGuid))
}

func DecrProductSkuSortNumber(skuGuid string) common.Result {
	if skuGuid == "" {
		return *common.Error(-1, "货品全局编码不可为空")
	}
	return *common.Success(mongo.DecrProductSkuSortNumber(skuGuid))
}
