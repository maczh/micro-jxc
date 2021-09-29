package service

import (
	"ququ.im/common"
	"ququ.im/jxc-api/nacos"
	"ququ.im/jxc-api/vo"
	"ququ.im/jxc-base/model"
)

func ListSkuStock(shopId, skuId, name, productId, storageId string, page, size int) common.Result {
	var skuStockList []model.SkuStock
	var skuStockNoteList []vo.SkuStockNote
	var err error
	var resultPage *common.ResultPage
	if page > 0 && size > 0 {
		skuStockList, resultPage, err = nacos.ListSkuStock(shopId, skuId, name, productId, storageId, page, size)
	} else {
		skuStockList, _, err = nacos.ListSkuStock(shopId, skuId, name, productId, storageId, 0, 0)
	}
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if skuStockList == nil || len(skuStockList) == 0 {
		return *common.Error(-1, "查无数据")
	}
	for _, skuStock := range skuStockList {
		skuStockNote := new(vo.SkuStockNote)
		skuStockNote.Id = skuStock.Id
		skuStockNote.ShopId = skuStock.ShopId
		skuStockNote.ProductId = skuStock.ProductId
		skuStockNote.SkuId = skuStock.SkuId
		skuStockNote.Name = skuStock.Name
		skuStockNote.Stocks = skuStock.Stocks
		skuStockNote.BaseUnit = skuStock.BaseUnit
		skuStockNote.StorageId = skuStock.StorageId
		skuStockNote.CostPrice = skuStock.CostPrice
		skuStockNote.LastPrice = skuStock.LastPrice
		skuStockNote.UpdateTime = skuStock.UpdateTime
		productSku, _ := nacos.GetProductSku(shopId, skuStock.SkuId, "")
		skuStockNote.CategoryId = productSku.CategoryId
		skuStockNote.SkuName = productSku.SkuName
		skuStockNote.BarCode = productSku.BarCode
		skuStockNote.Sku = productSku.Sku
		skuStockNote.PriceList = productSku.PriceList
		productCategory, _ := nacos.GetProductCategory(0, shopId, productSku.CategoryId)
		skuStockNote.CategoryName = productCategory.Name
		shopStorage, _ := nacos.GetShopStorage(shopId, skuStock.StorageId, "")
		skuStockNote.StorageName = shopStorage.Name
		skuStockNoteList = append(skuStockNoteList, *skuStockNote)
	}
	if page > 0 && size > 0 {
		return *common.SuccessWithPage(skuStockNoteList, resultPage.Count, page, size, resultPage.Total)
	} else {
		return *common.Success(skuStockNoteList)
	}
}

func GetSkuStockDetail(shopId, skuId, storageId string) common.Result {
	skuStock, err := nacos.GetSkuStock(shopId, skuId, storageId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if skuStock == nil {
		return *common.Error(-1, "查无数据")
	}
	skuStockNote := new(vo.SkuStockNote)
	skuStockNote.Id = skuStock.Id
	skuStockNote.ShopId = skuStock.ShopId
	skuStockNote.ProductId = skuStock.ProductId
	skuStockNote.SkuId = skuStock.SkuId
	skuStockNote.Name = skuStock.Name
	skuStockNote.Stocks = skuStock.Stocks
	skuStockNote.BaseUnit = skuStock.BaseUnit
	skuStockNote.StorageId = skuStock.StorageId
	skuStockNote.CostPrice = skuStock.CostPrice
	skuStockNote.LastPrice = skuStock.LastPrice
	skuStockNote.UpdateTime = skuStock.UpdateTime
	productSku, _ := nacos.GetProductSku(shopId, skuId, "")
	skuStockNote.CategoryId = productSku.CategoryId
	skuStockNote.SkuName = productSku.SkuName
	skuStockNote.BarCode = productSku.BarCode
	skuStockNote.Sku = productSku.Sku
	skuStockNote.PriceList = productSku.PriceList
	productCategory, _ := nacos.GetProductCategory(0, shopId, productSku.CategoryId)
	skuStockNote.CategoryName = productCategory.Name
	shopStorage, _ := nacos.GetShopStorage(shopId, storageId, "")
	skuStockNote.StorageName = shopStorage.Name
	return *common.Success(skuStockNote)
}

func InOutSkuStock(shopId, skuId, unit, storageId string, number, cost, price int) common.Result {
	return nacos.InOutSkuStock(shopId, skuId, unit, storageId, number, cost, price)
}

func GetSkuStockNumber(shopId, skuId, productId, storageId string) common.Result {
	stocks, _ := nacos.GetSkuStockNumber(shopId, skuId, storageId, productId)
	return *common.Success(map[string]int{"stocks": stocks})
}
