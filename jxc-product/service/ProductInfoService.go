package service

import (
	"ququ.im/common"
	"ququ.im/jxc-base/model"
	"ququ.im/jxc-base/util"
	"ququ.im/jxc-product/dao"
)

func ListProductInfo(shopId, keyword, categoryId string, page, size int) common.Result {
	if shopId == "" {
		return *common.Error(-1, "商户编号不可为空")
	}
	if page > 0 && size > 0 {
		count := dao.CountProductInfoByShopId(shopId, keyword, categoryId)
		productInfoList := dao.ListProductInfoByShopId(shopId, keyword, categoryId, page, size)
		if count > 0 {
			return *common.SuccessWithPage(productInfoList, count/size+1, page, size, count)
		} else {
			return *common.Error(-1, "查无此商品")
		}
	} else {
		productInfoList := dao.ListProductInfoByShopId(shopId, keyword, categoryId, 0, 0)
		if productInfoList == nil {
			return *common.Error(-1, "未定义产品")
		}
		return *common.Success(productInfoList)
	}
}

func GetProductInfo(id int, shopId, productId, barCode string) common.Result {
	var productInfo *model.ProductInfo
	if id > 0 {
		productInfo = dao.GetProductInfoById(id)
	} else if shopId != "" && productId != "" {
		productInfo = dao.GetProductInfo(shopId, productId)
	} else if shopId != "" && barCode != "" {
		productInfo = dao.GetProductInfoByBarCode(shopId, barCode)
	} else {
		return *common.Error(-1, "参数不全")
	}
	if productInfo == nil {
		return *common.Error(-1, "未定义产品")
	}
	return *common.Success(productInfo)
}

func SaveProductInfo(productId, shopId, productName, baseUnit, barCode, categoryId string) common.Result {
	if shopId == "" {
		return *common.Error(-1, "商户编号不可为空")
	}
	if productName == "" {
		return *common.Error(-1, "商品名称不可为空")
	}
	if productId != "" {
		p := dao.GetProductInfo(shopId, productId)
		if p != nil {
			return *common.Error(-1, "商品自定义编号已经存在")
		}
	}
	productInfo := new(model.ProductInfo)
	productInfo.ShopId = shopId
	productInfo.Name = productName
	productInfo.BaseUnit = baseUnit
	productInfo.BarCode = barCode
	productInfo.CategoryId = categoryId
	if productId == "" {
		productInfo.ProductId = util.GenerateId("CP", "product", productInfo.ShopId, 8)
	}
	productInfo, err := dao.SaveProductInfo(productInfo)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productInfo)
}

func UpdateProductInfo(id int, productId, categoryId, productName, baseUnit, barCode string) common.Result {
	productInfo := dao.GetProductInfoById(id)
	if productInfo == nil {
		return *common.Error(-1, "无此商品记录")
	}
	if categoryId != "" {
		productInfo.CategoryId = categoryId
	}
	if productId != "" {
		p := dao.GetProductInfo(productInfo.ShopId, productId)
		if p != nil && p.Id != productInfo.Id {
			return *common.Error(-1, "商品自定义编号已存在")
		}
		productInfo.ProductId = productId
	}
	if productName != "" {
		productInfo.Name = productName
	}
	if baseUnit != "" {
		productInfo.BaseUnit = baseUnit
	}
	if barCode != "" {
		productInfo.BarCode = barCode
	}
	productInfo, err := dao.UpdateProductInfo(productInfo)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productInfo)
}

func DeleteProductInfo(id int) common.Result {
	err := dao.DeleteProductInfo(id)
	if err != nil {
		return *common.Error(-1, "无此商品记录")
	}
	return *common.Success(nil)
}
