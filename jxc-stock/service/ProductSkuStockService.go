package service

import (
	"ququ.im/common"
	"ququ.im/jxc-stock/dao"
	"ququ.im/jxc-stock/nacos"
)

func InOutSkuStock(shopId, skuId, unit, storageId string, number, cost, price int) common.Result {
	productSku, err := nacos.GetProductSku("", shopId, skuId)
	if err != nil {
		return *common.Error(-1, "查询货品异常:"+err.Error())
	}
	if productSku == nil {
		return *common.Error(-1, "无此货品")
	}
	shopStorage, err := nacos.GetStorage(shopId, storageId)
	if err != nil {
		return *common.Error(-1, "查询仓库异常:"+err.Error())
	}
	if shopStorage == nil || (shopStorage.ShopId != "" && shopStorage.ShopId != productSku.ShopId) {
		return *common.Error(-1, "无此仓库")
	}
	num := number
	if unit != productSku.BaseUnit {
		productUnit, err := nacos.GetProductUnit(shopId, productSku.ProductId, unit, productSku.BaseUnit)
		if err != nil {
			return *common.Error(-1, "查询产品单位换算异常:"+err.Error())
		}
		if productUnit == nil {
			return *common.Error(-1, "本产品单位不正确或请添加本产品的单位换算规则")
		}
		num = number * productUnit.Scale
	}
	if number > 0 {
		skuStock, err := dao.AddSkuStock(productSku, num, cost, price, storageId)
		if err != nil {
			return *common.Error(-1, "库存入库异常:"+err.Error())
		}
		return *common.Success(skuStock)
	} else {
		skuStock, err := dao.SubSkuStock(productSku, -num, price, storageId)
		if err != nil {
			return *common.Error(-1, "库存出库异常:"+err.Error())
		}
		return *common.Success(skuStock)
	}
}

func ListSkuStock(shopId, skuId, name, productId, storageId string, page, size int) common.Result {
	if page > 0 && size > 0 {
		count := dao.CountSkuStock(skuId, shopId, name, productId, storageId)
		skuStockList := dao.ListSkuStock(skuId, shopId, name, productId, storageId, page, size)
		if count > 0 {
			return *common.SuccessWithPage(skuStockList, count/size+1, page, size, count)
		} else {
			return *common.Error(-1, "查无此货品库存")
		}
	} else {
		skuStockList := dao.ListSkuStock(skuId, shopId, name, productId, storageId, 0, 0)
		if skuStockList != nil && len(skuStockList) > 0 {
			return *common.Success(skuStockList)
		} else {
			return *common.Error(-1, "查无此货品库存")
		}
	}
}

func GetSkuStockNumber(shopId, skuId, productId, storageId string) common.Result {
	stocks := 0
	if productId != "" {
		stocks = dao.GetProductStockNumber(shopId, productId, storageId)
		return *common.Success(map[string]int{"stocks": stocks})
	} else if skuId != "" {
		stocks = dao.GetSkuStockNumber(shopId, skuId, storageId)
	}
	return *common.Success(map[string]int{"stocks": stocks})
}

func GetSkuStock(shopId, skuId, storageId string) common.Result {
	return *common.Success(dao.GetSkuStock(shopId, skuId, storageId))
}
