package service

import (
	"errors"
	"github.com/araddon/gou"
	"ququ.im/common"
	"ququ.im/common/logs"
	"ququ.im/common/utils"
	"ququ.im/jxc-api/nacos"
	"ququ.im/jxc-api/vo"
	"ququ.im/jxc-base/model"
	"strconv"
	"time"
)

func ListStockEntryRow(shopId, entryNo, skuId, storageId, orderNo, supplierId, startTime, endTime string, page, size int) common.Result {
	var stockEntryList []model.StockEntry
	var stockEntryRowList []vo.StockEntryRowDetail
	var resultPage *common.ResultPage
	var err error
	if page > 0 && size > 0 {
		stockEntryList, resultPage, err = nacos.ListStockEntry(shopId, skuId, entryNo, orderNo, supplierId, storageId, startTime, endTime, page, size)
	} else {
		stockEntryList, _, err = nacos.ListStockEntry(shopId, skuId, entryNo, orderNo, supplierId, storageId, startTime, endTime, 0, 0)
	}
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if stockEntryList == nil || len(stockEntryList) == 0 {
		return *common.Error(-1, "查无数据")
	}
	utils.FromJSON(utils.ToJSON(stockEntryList), &stockEntryRowList)
	for i := 0; i < len(stockEntryRowList); i++ {
		productSku, _ := nacos.GetProductSku(shopId, stockEntryRowList[i].SkuId, "")
		stockEntryRowList[i].SkuName = productSku.SkuName
		stockEntryRowList[i].Sku = productSku.Sku
		stockEntryRowList[i].PriceList = productSku.PriceList
		stockEntryRowList[i].BarCode = productSku.BarCode
		stockEntryRowList[i].ProductId = productSku.ProductId
		stockEntryRowList[i].SkuGuid = productSku.SkuGuid
		shopStorage, _ := nacos.GetShopStorage(shopId, stockEntryRowList[i].StorageId, "")
		stockEntryRowList[i].StorageName = shopStorage.Name
		productSupplier, _ := nacos.GetProductSupplier(0, shopId, productSku.ProductId, stockEntryRowList[i].SupplierId)
		stockEntryRowList[i].SupplierName = productSupplier.SupplierName
	}
	if page > 0 && size > 0 {
		return *common.SuccessWithPage(stockEntryRowList, resultPage.Count, page, size, resultPage.Total)
	} else {
		return *common.Success(stockEntryRowList)
	}
}

func GetStockEntryNote(shopId, entryNo string) common.Result {
	stockEntryList, _, err := nacos.ListStockEntry(shopId, "", entryNo, "", "", "", "", "", 0, 0)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if stockEntryList == nil || len(stockEntryList) == 0 {
		return *common.Error(-1, "无此入库单")
	}
	stockEntryNote := new(vo.StockEntryNote)
	stockEntryNote.ShopId = shopId
	stockEntryNote.EntryNo = entryNo
	stockEntryNote.Id = stockEntryList[0].Id
	stockEntryNote.EntryTime = stockEntryList[0].EntryTime
	stockEntryNote.Type = stockEntryList[0].Type
	stockEntryNote.OrderNo = stockEntryList[0].OrderNo
	stockEntryNote.Operator = stockEntryList[0].Operator
	for _, stockEntry := range stockEntryList {
		stockEntryRow := new(vo.StockEntryRow)
		stockEntryRow.Id = stockEntry.Id
		stockEntryRow.SkuId = stockEntry.SkuId
		stockEntryRow.Unit = stockEntry.Unit
		stockEntryRow.StorageId = stockEntry.StorageId
		stockEntryRow.Number = stockEntry.Number
		stockEntryRow.Price = stockEntry.Price
		stockEntryRow.SupplierId = stockEntry.SupplierId
		stockEntryRow.Remark = stockEntry.Remark
		productSku, _ := nacos.GetProductSku(shopId, stockEntryRow.SkuId, "")
		stockEntryRow.SkuName = productSku.SkuName
		stockEntryRow.Sku = productSku.Sku
		stockEntryRow.PriceList = productSku.PriceList
		stockEntryRow.BarCode = productSku.BarCode
		stockEntryRow.ProductId = productSku.ProductId
		stockEntryRow.SkuGuid = productSku.SkuGuid
		shopStorage, _ := nacos.GetShopStorage(shopId, stockEntryRow.StorageId, "")
		stockEntryRow.StorageName = shopStorage.Name
		productSupplier, _ := nacos.GetProductSupplier(0, shopId, productSku.ProductId, stockEntryRow.SupplierId)
		stockEntryRow.SupplierName = productSupplier.SupplierName
		stockEntryNote.SkuList = append(stockEntryNote.SkuList, *stockEntryRow)
	}
	return *common.Success(stockEntryNote)
}

func SaveOneStockEntryRow(shopId, entryNo, entryType, skuId, unit, orderNo, storageId, supplierId, operator, remark string, number, price int) (*model.StockEntry, error) {
	if shopId == "" {
		return nil, errors.New("缺少商户账号参数")
	}
	if entryNo == "" {
		return nil, errors.New("缺少入库单号参数")
	}
	if skuId == "" {
		return nil, errors.New("缺少入库商品编号参数")
	}
	if entryType == "" {
		return nil, errors.New("缺少入库类型参数")
	}
	if unit == "" {
		return nil, errors.New("缺少货品单位参数")
	}
	if storageId == "" {
		return nil, errors.New("缺少商户仓库编码参数")
	}
	if number == 0 {
		return nil, errors.New("缺少入库数量参数")
	}
	if price == 0 {
		//如果已有库存的情况下，不传入入库价格，则默认价格为最近一次入库价，该货品首次入库则价格不可为空
		skuStockList, _, err := nacos.ListSkuStock(shopId, skuId, "", "", "", 0, 0)
		if err != nil || skuStockList == nil || len(skuStockList) == 0 {
			return nil, errors.New("缺少成本价格参数")
		} else {
			price = skuStockList[0].LastPrice
		}
	}
	//检查货品是否存在
	productSku, _ := nacos.GetProductSku(shopId, skuId, "")
	if productSku == nil {
		return nil, errors.New("货品编码不正确")
	}
	shopStorage, _ := nacos.GetShopStorage(shopId, storageId, "")
	if shopStorage == nil {
		return nil, errors.New("商户仓库编码不正确")
	}
	if supplierId != "" {
		productSupplier, _ := nacos.GetProductSupplier(0, shopId, productSku.ProductId, supplierId)
		if productSupplier == nil {
			return nil, errors.New("商品供应商不存在，请先添加此商品的供应商信息")
		}
	}
	if unit != productSku.BaseUnit {
		productUnit, _ := nacos.GetProductUnit(0, shopId, productSku.ProductId, unit, productSku.BaseUnit)
		if productUnit == nil {
			return nil, errors.New("商品单位不存在，请先添加此商品单位换算成" + productSku.BaseUnit + "的换算规则")
		}
	}
	stockEntry, err := nacos.SaveStockEntry(shopId, skuId, entryNo, entryType, unit, orderNo, storageId, supplierId, operator, remark, number, price)
	if err != nil {
		return nil, err
	}
	//增加库存
	result := nacos.InOutSkuStock(shopId, skuId, unit, storageId, number, 0, price)
	if result.Status != 1 {
		//增加库存失败，回滚，删除入库单记录
		nacos.DeleteStockEntry(stockEntry.Id)
		return nil, errors.New(result.Msg)
	}
	return stockEntry, nil
}

func SaveStockEntryNote(shopId, entryNo, entryType, orderNo, operator, stockEntryRowList string) common.Result {
	if shopId == "" {
		return *common.Error(-1, "缺少商户账号参数")
	}
	if entryNo == "" {
		entryNo = "IN" + shopId + time.Now().Format("20060102150405")
		//return *common.Error(-1, "缺少入库单号参数")
	}
	stockEntryList, _, _ := nacos.ListStockEntry(shopId, "", entryNo, "", "", "", "", "", 0, 0)
	if stockEntryList != nil && len(stockEntryList) > 0 {
		return *common.Error(-1, "此入库单号已经存在")
	}
	if operator == "" {
		return *common.Error(-1, "缺少操作员参数")
	}
	if entryType == "" {
		return *common.Error(-1, "缺少入库类型参数")
	}
	if stockEntryRowList == "" {
		return *common.Error(-1, "缺少入库单详细入库清单参数")
	}
	if !gou.IsJsonArray([]byte(stockEntryRowList)) {
		return *common.Error(-1, "入库清单参数不是JSON数组格式")
	}
	var stockEntryRows []vo.StockEntryRow
	utils.FromJSON(stockEntryRowList, &stockEntryRows)
	//先校验入库清单的正确性
	for i, entryRow := range stockEntryRows {
		if entryRow.SkuId == "" {
			return *common.Error(-1, "第"+strconv.Itoa(i+1)+"个货品编码不能为空")
		}
		if entryRow.Unit == "" {
			return *common.Error(-1, "货品"+entryRow.SkuId+"的单位不能为空")
		}
		if entryRow.Number == 0 {
			return *common.Error(-1, "货品"+entryRow.SkuId+"的入库数量不能为0")
		}
		if entryRow.StorageId == "" {
			return *common.Error(-1, "货品"+entryRow.SkuId+"的仓库编码不能为空")
		}
		//检查货品是否存在
		productSku, _ := nacos.GetProductSku(shopId, entryRow.SkuId, "")
		if productSku == nil {
			return *common.Error(-1, "货品"+entryRow.SkuId+"的货品编码不存在")
		}
		shopStorage, _ := nacos.GetShopStorage(shopId, entryRow.StorageId, "")
		if shopStorage == nil {
			return *common.Error(-1, "货品"+entryRow.SkuId+"的仓库编码不存在")
		}
		if entryRow.SupplierId != "" {
			productSupplier, _ := nacos.GetProductSupplier(0, shopId, productSku.ProductId, entryRow.SupplierId)
			if productSupplier == nil {
				return *common.Error(-1, "货品"+entryRow.SkuId+"的供应商账号不存在")
			}
		}
		if entryRow.Unit != productSku.BaseUnit {
			productUnit, _ := nacos.GetProductUnit(0, shopId, productSku.ProductId, entryRow.Unit, productSku.BaseUnit)
			if productUnit == nil {
				return *common.Error(-1, "货品"+entryRow.SkuId+"的单位换算规格不存在")
			}
		}
	}
	//按记录批量入库
	for _, entryRow := range stockEntryRows {
		stockEntry, err := SaveOneStockEntryRow(shopId, entryNo, entryType, entryRow.SkuId, entryRow.Unit, orderNo, entryRow.StorageId, entryRow.SupplierId, operator, entryRow.Remark, entryRow.Number, entryRow.Price)
		if err != nil {
			logs.Error("入库单批量入库错误:{}\n,入库参数{}", err.Error(), stockEntry)
			return *common.Error(-1, "入库单批量入库错误:"+err.Error())
		}
	}
	return GetStockEntryNote(shopId, entryNo)
}

func UpdateStockEntryRow(id int, unit, orderNo, storageId, supplierId, operator, remark string, number, price int) common.Result {
	stockEntry, err := nacos.GetStockEntry("", "", "", id)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	productSku, _ := nacos.GetProductSku(stockEntry.ShopId, stockEntry.SkuId, "")
	if supplierId != "" {
		productSupplier, _ := nacos.GetProductSupplier(0, stockEntry.ShopId, productSku.ProductId, supplierId)
		if productSupplier == nil {
			return *common.Error(-1, "商品供应商不存在，请先添加此商品的供应商信息")
		}
	}
	if storageId != "" {
		shopStorage, _ := nacos.GetShopStorage(stockEntry.ShopId, storageId, "")
		if shopStorage == nil {
			return *common.Error(-1, "商户仓库编码不正确")
		}
	}
	//如果改数量或改单位
	if (number > 0 && number != stockEntry.Number) || (unit != "" && unit != stockEntry.Unit) || (price > 0 && price != stockEntry.Price) {
		if unit != "" && unit != productSku.BaseUnit {
			productUnit, err := nacos.GetProductUnit(0, stockEntry.ShopId, productSku.ProductId, unit, productSku.BaseUnit)
			if err != nil || productUnit == nil {
				return *common.Error(-1, "新的货品单位换算规则不存在")
			}
		}
		//先减掉原库存
		nacos.InOutSkuStock(stockEntry.ShopId, stockEntry.SkuId, stockEntry.Unit, stockEntry.StorageId, -stockEntry.Number, 0, stockEntry.Price)
		//再加上新库存
		if unit == "" {
			unit = stockEntry.Unit
		}
		if storageId == "" {
			storageId = stockEntry.StorageId
		}
		if number == 0 {
			number = stockEntry.Number
		}
		if price == 0 {
			price = stockEntry.Price
		}
		nacos.InOutSkuStock(stockEntry.ShopId, stockEntry.SkuId, unit, storageId, number, 0, price)
	}
	return nacos.UpdateStockEntry(id, stockEntry.SkuId, unit, orderNo, storageId, supplierId, operator, remark, number, price)
}

func DeleteStockEntryRow(id int) common.Result {
	stockEntry, err := nacos.GetStockEntry("", "", "", id)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	//减掉相应的库存
	result := nacos.InOutSkuStock(stockEntry.ShopId, stockEntry.SkuId, stockEntry.Unit, stockEntry.StorageId, -stockEntry.Number, 0, stockEntry.Price)
	if result.Status != 1 {
		return result
	}
	return nacos.DeleteStockEntry(id)
}
