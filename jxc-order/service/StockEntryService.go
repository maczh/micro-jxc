package service

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
	"ququ.im/jxc-order/dao"
	"time"
)

func ListStockEntry(shopId, entryNo, skuId, storageId, orderNo, supplierId, startTime, endTime string, page, size int) common.Result {
	if shopId == "" {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if endTime != "" {
		endTime = endTime + " 23:59:59"
	}
	stockEntryList := dao.ListStockEntry(entryNo, shopId, skuId, storageId, orderNo, supplierId, startTime, endTime, page, size)
	if stockEntryList == nil || len(stockEntryList) == 0 {
		return *common.Error(-1, "查无数据")
	} else {
		if page > 0 && size > 0 {
			count := dao.CountStockEntry(entryNo, shopId, skuId, storageId, orderNo, supplierId, startTime, endTime)
			return *common.SuccessWithPage(stockEntryList, count/size+1, page, size, count)
		} else {
			return *common.Success(stockEntryList)
		}
	}
}

func GetStockEntry(shopId, entryNo, skuId string, id int) common.Result {
	var stockEntry *model.StockEntry
	if id > 0 {
		stockEntry = dao.GetStockEntryById(id)
	} else {
		if shopId == "" {
			return *common.Error(-1, "未传入商户账号参数")
		}
		if entryNo == "" {
			return *common.Error(-1, "未传入入库单号参数")
		}
		if skuId == "" {
			return *common.Error(-1, "未传入入库商品编号参数")
		}
		stockEntry = dao.GetStockEntry(shopId, entryNo, skuId)
	}
	if stockEntry == nil || stockEntry.Id == 0 {
		return *common.Error(-1, "查无此数据")
	} else {
		return *common.Success(stockEntry)
	}
}

func SaveStockEntry(shopId, entryNo, entryType, skuId, unit, orderNo, storageId, supplierId, operator, remark string, number, price int) common.Result {
	if shopId == "" {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if entryNo == "" {
		return *common.Error(-1, "未传入入库单号参数")
	}
	if skuId == "" {
		return *common.Error(-1, "未传入入库商品编号参数")
	}
	if entryType == "" {
		return *common.Error(-1, "未传入入库类型参数")
	}
	if unit == "" {
		return *common.Error(-1, "未传入货品单位参数")
	}
	if number == 0 {
		return *common.Error(-1, "未传入入库数量参数")
	}
	stockEntry := new(model.StockEntry)
	stockEntry.ShopId = shopId
	stockEntry.EntryNo = entryNo
	stockEntry.SkuId = skuId
	stockEntry.Type = entryType
	stockEntry.Unit = unit
	stockEntry.Number = number
	stockEntry.OrderNo = orderNo
	stockEntry.Price = price
	stockEntry.StorageId = storageId
	stockEntry.SupplierId = supplierId
	stockEntry.Remark = remark
	stockEntry.Operator = operator
	stockEntry.EntryTime = utils.ToDateTimeString(time.Now())
	stockEntry, _ = dao.SaveStockEntry(stockEntry)
	return *common.Success(dao.GetStockEntry(shopId, entryNo, skuId))
}

func UpdateStockEntry(id int, skuId, unit, orderNo, storageId, supplierId, operator, remark string, number, price int) common.Result {
	stockEntry := dao.GetStockEntryById(id)
	if stockEntry == nil {
		return *common.Error(-1, "无此入库单条目")
	}
	stockEntry.Id = id
	stockEntry.SkuId = skuId
	stockEntry.Unit = unit
	stockEntry.Number = number
	stockEntry.OrderNo = orderNo
	stockEntry.Price = price
	stockEntry.StorageId = storageId
	stockEntry.SupplierId = supplierId
	stockEntry.Remark = remark
	stockEntry.Operator = operator
	stockEntry, _ = dao.UpdateStockEntry(stockEntry)
	return *common.Success(stockEntry)
}

func DeleteStockEntry(id int) common.Result {
	err := dao.DeleteStockEntry(id)
	if err != nil {
		return *common.Error(-1, "无此入库单记录")
	}
	return *common.Success(nil)
}
