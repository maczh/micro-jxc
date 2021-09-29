package service

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
	"ququ.im/jxc-order/dao"
	"time"
)

func ListStockDelivery(shopId, deliveryNo, skuId, storageId, orderNo, customerId, startTime, endTime string, page, size int) common.Result {
	if shopId == "" {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if endTime != "" {
		endTime = endTime + " 23:59:59"
	}
	stockDeliveryList := dao.ListStockDelivery(deliveryNo, shopId, skuId, storageId, orderNo, customerId, startTime, endTime, page, size)
	if stockDeliveryList == nil || len(stockDeliveryList) == 0 {
		return *common.Error(-1, "查无数据")
	} else {
		if page > 0 && size > 0 {
			count := dao.CountStockDelivery(deliveryNo, shopId, skuId, storageId, orderNo, customerId, startTime, endTime)
			return *common.SuccessWithPage(stockDeliveryList, count/size+1, page, size, count)
		} else {
			return *common.Success(stockDeliveryList)
		}
	}
}

func GetStockDelivery(shopId, deliveryNo, skuId string, id int) common.Result {
	var stockDelivery *model.StockDelivery
	if id > 0 {
		stockDelivery = dao.GetStockDeliveryById(id)
	} else {
		if shopId == "" {
			return *common.Error(-1, "未传入商户账号参数")
		}
		if deliveryNo == "" {
			return *common.Error(-1, "未传入出库单号参数")
		}
		if skuId == "" {
			return *common.Error(-1, "未传入出库商品编号参数")
		}
		stockDelivery = dao.GetStockDelivery(shopId, deliveryNo, skuId)
	}
	if stockDelivery == nil || stockDelivery.Id == 0 {
		return *common.Error(-1, "查无此数据")
	} else {
		return *common.Success(stockDelivery)
	}
}

func SaveStockDelivery(shopId, deliveryNo, deliveryType, skuId, unit, orderNo, storageId, customerId, customerName, operator, remark string, number, cost, price int) common.Result {
	if shopId == "" {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if deliveryNo == "" {
		return *common.Error(-1, "未传入出库单号参数")
	}
	if skuId == "" {
		return *common.Error(-1, "未传入出库商品编号参数")
	}
	if deliveryType == "" {
		return *common.Error(-1, "未传入出库类型参数")
	}
	if unit == "" {
		return *common.Error(-1, "未传入货品单位参数")
	}
	if number == 0 {
		return *common.Error(-1, "未传入出库数量参数")
	}
	stockDelivery := new(model.StockDelivery)
	stockDelivery.ShopId = shopId
	stockDelivery.DeliveryNo = deliveryNo
	stockDelivery.SkuId = skuId
	stockDelivery.Type = deliveryType
	stockDelivery.Unit = unit
	stockDelivery.Number = number
	stockDelivery.OrderNo = orderNo
	stockDelivery.Cost = cost
	stockDelivery.Price = price
	stockDelivery.StorageId = storageId
	stockDelivery.CustomerId = customerId
	stockDelivery.CustomerName = customerName
	stockDelivery.Remark = remark
	stockDelivery.Operator = operator
	stockDelivery.DeliveryTime = utils.ToDateTimeString(time.Now())
	stockDelivery, _ = dao.SaveStockDelivery(stockDelivery)
	return *common.Success(dao.GetStockDelivery(shopId, deliveryNo, skuId))
}

func UpdateStockDelivery(id int, skuId, unit, orderNo, storageId, customerId, customerName, operator, remark string, number, cost, price int) common.Result {
	stockDelivery := dao.GetStockDeliveryById(id)
	if stockDelivery == nil {
		return *common.Error(-1, "无此出库单条目")
	}
	stockDelivery.Id = id
	stockDelivery.SkuId = skuId
	stockDelivery.Unit = unit
	stockDelivery.Number = number
	stockDelivery.OrderNo = orderNo
	stockDelivery.Cost = cost
	stockDelivery.Price = price
	stockDelivery.StorageId = storageId
	stockDelivery.CustomerId = customerId
	stockDelivery.CustomerName = customerName
	stockDelivery.Remark = remark
	stockDelivery.Operator = operator
	stockDelivery, _ = dao.UpdateStockDelivery(stockDelivery)
	return *common.Success(stockDelivery)
}

func DeleteStockDelivery(id int) common.Result {
	err := dao.DeleteStockDelivery(id)
	if err != nil {
		return *common.Error(-1, "无此出库单记录")
	}
	return *common.Success(nil)
}
