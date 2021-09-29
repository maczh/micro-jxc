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

func ListStockDeliveryRow(shopId, deliveryNo, skuId, storageId, orderNo, customerId, startTime, endTime string, page, size int) common.Result {
	var stockDeliveryList []model.StockDelivery
	var stockDeliveryRowList []vo.StockDeliveryRowDetail
	var resultPage *common.ResultPage
	var err error
	if page > 0 && size > 0 {
		stockDeliveryList, resultPage, err = nacos.ListStockDelivery(shopId, skuId, deliveryNo, orderNo, customerId, storageId, startTime, endTime, page, size)
	} else {
		stockDeliveryList, _, err = nacos.ListStockDelivery(shopId, skuId, deliveryNo, orderNo, customerId, storageId, startTime, endTime, 0, 0)
	}
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if stockDeliveryList == nil || len(stockDeliveryList) == 0 {
		return *common.Error(-1, "查无数据")
	}
	utils.FromJSON(utils.ToJSON(stockDeliveryList), &stockDeliveryRowList)
	for i := 0; i < len(stockDeliveryRowList); i++ {
		productSku, _ := nacos.GetProductSku(shopId, stockDeliveryRowList[i].SkuId, "")
		stockDeliveryRowList[i].SkuName = productSku.SkuName
		stockDeliveryRowList[i].Sku = productSku.Sku
		stockDeliveryRowList[i].PriceList = productSku.PriceList
		stockDeliveryRowList[i].BarCode = productSku.BarCode
		stockDeliveryRowList[i].ProductId = productSku.ProductId
		stockDeliveryRowList[i].SkuGuid = productSku.SkuGuid
		shopStorage, _ := nacos.GetShopStorage(shopId, stockDeliveryRowList[i].StorageId, "")
		stockDeliveryRowList[i].StorageName = shopStorage.Name
	}
	if page > 0 && size > 0 {
		return *common.SuccessWithPage(stockDeliveryRowList, resultPage.Count, page, size, resultPage.Total)
	} else {
		return *common.Success(stockDeliveryRowList)
	}
}

func GetStockDeliveryNote(shopId, deliveryNo string) common.Result {
	stockDeliveryList, _, err := nacos.ListStockDelivery(shopId, "", deliveryNo, "", "", "", "", "", 0, 0)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if stockDeliveryList == nil || len(stockDeliveryList) == 0 {
		return *common.Error(-1, "无此出库单")
	}
	stockDeliveryNote := new(vo.StockDeliveryNote)
	stockDeliveryNote.ShopId = shopId
	stockDeliveryNote.DeliveryNo = deliveryNo
	stockDeliveryNote.Id = stockDeliveryList[0].Id
	stockDeliveryNote.DeliveryTime = stockDeliveryList[0].DeliveryTime
	stockDeliveryNote.Type = stockDeliveryList[0].Type
	stockDeliveryNote.OrderNo = stockDeliveryList[0].OrderNo
	stockDeliveryNote.Operator = stockDeliveryList[0].Operator
	for _, stockDelivery := range stockDeliveryList {
		stockDeliveryRow := new(vo.StockDeliveryRow)
		stockDeliveryRow.Id = stockDelivery.Id
		stockDeliveryRow.SkuId = stockDelivery.SkuId
		stockDeliveryRow.Unit = stockDelivery.Unit
		stockDeliveryRow.StorageId = stockDelivery.StorageId
		stockDeliveryRow.Number = stockDelivery.Number
		stockDeliveryRow.Price = stockDelivery.Price
		stockDeliveryRow.Cost = stockDelivery.Cost
		stockDeliveryRow.CustomerId = stockDelivery.CustomerId
		stockDeliveryRow.CustomerName = stockDelivery.CustomerName
		stockDeliveryRow.Remark = stockDelivery.Remark
		productSku, _ := nacos.GetProductSku(shopId, stockDeliveryRow.SkuId, "")
		stockDeliveryRow.SkuName = productSku.SkuName
		stockDeliveryRow.Sku = productSku.Sku
		stockDeliveryRow.PriceList = productSku.PriceList
		stockDeliveryRow.BarCode = productSku.BarCode
		stockDeliveryRow.ProductId = productSku.ProductId
		stockDeliveryRow.SkuGuid = productSku.SkuGuid
		shopStorage, _ := nacos.GetShopStorage(shopId, stockDeliveryRow.StorageId, "")
		stockDeliveryRow.StorageName = shopStorage.Name
		stockDeliveryNote.SkuList = append(stockDeliveryNote.SkuList, *stockDeliveryRow)
	}
	return *common.Success(stockDeliveryNote)
}

func SaveOneStockDeliveryRow(shopId, deliveryNo, deliveryType, skuId, unit, orderNo, storageId, customerId, customerName, operator, remark string, number, price int) (*model.StockDelivery, error) {
	if shopId == "" {
		return nil, errors.New("缺少商户账号参数")
	}
	if deliveryNo == "" {
		return nil, errors.New("缺少出库单号参数")
	}
	if skuId == "" {
		return nil, errors.New("缺少出库商品编号参数")
	}
	if deliveryType == "" {
		return nil, errors.New("缺少出库类型参数")
	}
	if unit == "" {
		return nil, errors.New("缺少货品单位参数")
	}
	if storageId == "" {
		return nil, errors.New("缺少商户仓库编码参数")
	}
	if number == 0 {
		return nil, errors.New("缺少出库数量参数")
	}
	cost := 0
	//检查货品是否存在
	productSku, _ := nacos.GetProductSku(shopId, skuId, "")
	if productSku == nil {
		return nil, errors.New("货品编码不正确")
	}
	shopStorage, _ := nacos.GetShopStorage(shopId, storageId, "")
	if shopStorage == nil {
		return nil, errors.New("商户仓库编码不正确")
	}
	//检查是否有库存记录
	skuStock, _ := nacos.GetSkuStock(shopId, skuId, storageId)
	if skuStock == nil {
		return nil, errors.New(shopStorage.Name + "仓库中没有" + productSku.SkuName + "的库存记录")
	}
	cost = skuStock.CostPrice
	if unit != productSku.BaseUnit {
		productUnit, _ := nacos.GetProductUnit(0, shopId, productSku.ProductId, unit, productSku.BaseUnit)
		if productUnit == nil {
			return nil, errors.New("商品单位不存在，请先添加此商品单位换算成" + productSku.BaseUnit + "的换算规则")
		}
	}
	stockDelivery, err := nacos.SaveStockDelivery(shopId, skuId, deliveryNo, deliveryType, unit, orderNo, storageId, customerId, customerName, operator, remark, number, cost, price)
	if err != nil {
		return nil, err
	}
	//减少库存
	result := nacos.InOutSkuStock(shopId, skuId, unit, storageId, -number, 0, 0)
	if result.Status != 1 {
		//增加库存失败，回滚，删除出库单记录
		nacos.DeleteStockDelivery(stockDelivery.Id)
		return nil, errors.New(result.Msg)
	}
	return stockDelivery, nil
}

func SaveStockDeliveryNote(shopId, deliveryNo, deliveryType, orderNo, operator, stockDeliveryRowList string) common.Result {
	if shopId == "" {
		return *common.Error(-1, "缺少商户账号参数")
	}
	if deliveryNo == "" {
		deliveryNo = "OUT" + shopId + time.Now().Format("20060102150405")
		//return *common.Error(-1, "缺少出库单号参数")
	}
	stockDeliveryList, _, _ := nacos.ListStockDelivery(shopId, "", deliveryNo, "", "", "", "", "", 0, 0)
	if stockDeliveryList != nil && len(stockDeliveryList) > 0 {
		return *common.Error(-1, "此出库单号已经存在")
	}
	if operator == "" {
		return *common.Error(-1, "缺少操作员参数")
	}
	if deliveryType == "" {
		return *common.Error(-1, "缺少出库类型参数")
	}
	if stockDeliveryRowList == "" {
		return *common.Error(-1, "缺少出库单详细出库清单参数")
	}
	if !gou.IsJsonArray([]byte(stockDeliveryRowList)) {
		return *common.Error(-1, "出库清单参数不是JSON数组格式")
	}
	var stockDeliveryRows []vo.StockDeliveryRow
	utils.FromJSON(stockDeliveryRowList, &stockDeliveryRows)
	//先校验出库清单的正确性
	for i, deliveryRow := range stockDeliveryRows {
		if deliveryRow.SkuId == "" {
			return *common.Error(-1, "第"+strconv.Itoa(i+1)+"个货品编码不能为空")
		}
		if deliveryRow.Unit == "" {
			return *common.Error(-1, "货品"+deliveryRow.SkuId+"的单位不能为空")
		}
		if deliveryRow.Number == 0 {
			return *common.Error(-1, "货品"+deliveryRow.SkuId+"的出库数量不能为0")
		}
		if deliveryRow.StorageId == "" {
			return *common.Error(-1, "货品"+deliveryRow.SkuId+"的仓库编码不能为空")
		}
		//检查货品是否存在
		productSku, _ := nacos.GetProductSku(shopId, deliveryRow.SkuId, "")
		if productSku == nil {
			return *common.Error(-1, "货品"+deliveryRow.SkuId+"的货品编码不存在")
		}
		shopStorage, _ := nacos.GetShopStorage(shopId, deliveryRow.StorageId, "")
		if shopStorage == nil {
			return *common.Error(-1, "货品"+deliveryRow.SkuId+"的仓库编码不存在")
		}
		if deliveryRow.Unit != productSku.BaseUnit {
			productUnit, _ := nacos.GetProductUnit(0, shopId, productSku.ProductId, deliveryRow.Unit, productSku.BaseUnit)
			if productUnit == nil {
				return *common.Error(-1, "货品"+deliveryRow.SkuId+"的单位换算规格不存在")
			}
		}
	}
	//按记录批量出库
	for _, deliveryRow := range stockDeliveryRows {
		stockDelivery, err := SaveOneStockDeliveryRow(shopId, deliveryNo, deliveryType, deliveryRow.SkuId, deliveryRow.Unit, orderNo, deliveryRow.StorageId, deliveryRow.CustomerId, deliveryRow.CustomerName, operator, deliveryRow.Remark, deliveryRow.Number, deliveryRow.Price)
		if err != nil {
			logs.Error("出库单批量出库错误:{}\n,出库参数{}", err.Error(), stockDelivery)
			return *common.Error(-1, "出库单批量出库错误:"+err.Error())
		}
	}
	return GetStockDeliveryNote(shopId, deliveryNo)
}

func UpdateStockDeliveryRow(id int, unit, orderNo, storageId, customerId, customerName, operator, remark string, number, price int) common.Result {
	stockDelivery, err := nacos.GetStockDelivery("", "", "", id)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	productSku, _ := nacos.GetProductSku(stockDelivery.ShopId, stockDelivery.SkuId, "")
	skuStock, _ := nacos.GetSkuStock(stockDelivery.ShopId, stockDelivery.SkuId, stockDelivery.StorageId)
	if skuStock == nil {
		return *common.Error(-1, "仓库中没有此货品的库存记录")
	}
	if storageId != "" && storageId != stockDelivery.StorageId {
		shopStorage, _ := nacos.GetShopStorage(stockDelivery.ShopId, storageId, "")
		if shopStorage == nil {
			return *common.Error(-1, "新仓库编码不正确")
		}
		skuStockNew, _ := nacos.GetSkuStock(stockDelivery.ShopId, stockDelivery.SkuId, storageId)
		if skuStockNew == nil {
			return *common.Error(-1, "新仓库中没有此货品的库存记录")
		}
	}
	//如果改数量或改单位
	if (number > 0 && number != stockDelivery.Number) || (unit != "" && unit != stockDelivery.Unit) || (storageId != "" && storageId != stockDelivery.StorageId) {
		if unit != "" && unit != productSku.BaseUnit {
			productUnit, err := nacos.GetProductUnit(0, stockDelivery.ShopId, productSku.ProductId, unit, productSku.BaseUnit)
			if err != nil || productUnit == nil {
				return *common.Error(-1, "新的货品单位换算规则不存在")
			}
		}
		//先加上原出库的库存
		nacos.InOutSkuStock(stockDelivery.ShopId, stockDelivery.SkuId, stockDelivery.Unit, stockDelivery.StorageId, stockDelivery.Number, skuStock.CostPrice, skuStock.LastPrice)
		//再减去新库存
		if unit == "" {
			unit = stockDelivery.Unit
		}
		if storageId == "" {
			storageId = stockDelivery.StorageId
		}
		if number == 0 {
			number = stockDelivery.Number
		}
		if price == 0 {
			price = stockDelivery.Price
		}
		nacos.InOutSkuStock(stockDelivery.ShopId, stockDelivery.SkuId, unit, storageId, -number, 0, price)
	}
	return nacos.UpdateStockDelivery(id, stockDelivery.SkuId, unit, orderNo, storageId, customerId, customerName, operator, remark, number, skuStock.CostPrice, price)
}

func DeleteStockDeliveryRow(id int) common.Result {
	stockDelivery, err := nacos.GetStockDelivery("", "", "", id)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	skuStock, _ := nacos.GetSkuStock(stockDelivery.ShopId, stockDelivery.SkuId, stockDelivery.StorageId)
	if skuStock == nil {
		return *common.Error(-1, "仓库中没有此货品的库存记录")
	}
	//补上相应的库存
	result := nacos.InOutSkuStock(stockDelivery.ShopId, stockDelivery.SkuId, stockDelivery.Unit, stockDelivery.StorageId, stockDelivery.Number, skuStock.CostPrice, skuStock.LastPrice)
	if result.Status != 1 {
		return result
	}
	return nacos.DeleteStockDelivery(id)
}
