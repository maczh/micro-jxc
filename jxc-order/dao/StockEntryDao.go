package dao

import (
	"errors"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
)

const TABLE_STOCK_ENTRY string = "stock_entry"

func ListStockEntry(entryNo, shopId, skuId, storageId, orderNo, supplierId, startTime, endTime string, page, size int) []model.StockEntry {
	var stockEntryList []model.StockEntry
	if page > 0 && size > 0 {
		config.Mysql.Table(TABLE_STOCK_ENTRY).Where("IF(? != '',entry_no = ?,1=1) AND IF(? != '',shop_id = ?,1=1) AND IF(? != '',sku_id = ?,1=1) AND IF(? != '',storage_id = ?,1=1) AND IF(? != '',order_no = ?,1=1) AND IF(? != '',supplier_id = ?,1=1) AND IF(? != '',entry_time > ?,1=1) AND IF(? != '',entry_time < ?,1=1) ", entryNo, entryNo, shopId, shopId, skuId, skuId, storageId, storageId, orderNo, orderNo, supplierId, supplierId, startTime, startTime, endTime, endTime).Order("id DESC").Limit(size).Offset((page - 1) * size).Find(&stockEntryList)

	} else {
		config.Mysql.Table(TABLE_STOCK_ENTRY).Where("IF(? != '',entry_no = ?,1=1) AND IF(? != '',shop_id = ?,1=1) AND IF(? != '',sku_id = ?,1=1) AND IF(? != '',storage_id = ?,1=1) AND IF(? != '',order_no = ?,1=1) AND IF(? != '',supplier_id = ?,1=1) AND IF(? != '',entry_time > ?,1=1) AND IF(? != '',entry_time < ?,1=1) ", entryNo, entryNo, shopId, shopId, skuId, skuId, storageId, storageId, orderNo, orderNo, supplierId, supplierId, startTime, startTime, endTime, endTime).Order("id DESC").Find(&stockEntryList)
	}
	if len(stockEntryList) == 0 {
		return nil
	} else {
		for i := 0; i < len(stockEntryList); i++ {
			stockEntryList[i].EntryTime = utils.GormTimeFormat(stockEntryList[i].EntryTime)
		}
		return stockEntryList
	}
}

func CountStockEntry(entryNo, shopId, skuId, storageId, orderNo, supplierId, startTime, endTime string) int {
	count := 0
	config.Mysql.Table(TABLE_STOCK_ENTRY).Where("IF(? != '',entry_no = ?,1=1) AND IF(? != '',shop_id = ?,1=1) AND IF(? != '',sku_id = ?,1=1) AND IF(? != '',storage_id = ?,1=1) AND IF(? != '',order_no = ?,1=1) AND IF(? != '',supplier_id = ?,1=1) AND IF(? != '',entry_time > ?,1=1) AND IF(? != '',entry_time < ?,1=1) ", entryNo, entryNo, shopId, shopId, skuId, skuId, storageId, storageId, orderNo, orderNo, supplierId, supplierId, startTime, startTime, endTime, endTime).Count(&count)
	return count
}

func GetStockEntry(shopId, entryNo, skuId string) *model.StockEntry {
	var stockEntry model.StockEntry
	config.Mysql.Table(TABLE_STOCK_ENTRY).Where("shop_id = ? AND entry_no = ? AND sku_id = ?", shopId, entryNo, skuId).First(&stockEntry)
	if stockEntry.Id == 0 {
		return nil
	} else {
		stockEntry.EntryTime = utils.GormTimeFormat(stockEntry.EntryTime)
		return &stockEntry
	}
}

func GetStockEntryById(id int) *model.StockEntry {
	var stockEntry model.StockEntry
	config.Mysql.Table(TABLE_STOCK_ENTRY).Where("id = ?", id).First(&stockEntry)
	if stockEntry.Id == 0 {
		return nil
	} else {
		stockEntry.EntryTime = utils.GormTimeFormat(stockEntry.EntryTime)
		return &stockEntry
	}
}

func SaveStockEntry(stockEntry *model.StockEntry) (*model.StockEntry, error) {
	s := GetStockEntry(stockEntry.ShopId, stockEntry.EntryNo, stockEntry.SkuId)
	if s != nil {
		return s, nil
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_STOCK_ENTRY).Create(stockEntry).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	config.Redis.HSet("storage:last:"+stockEntry.ShopId, stockEntry.StorageId, stockEntry.Operator+","+stockEntry.EntryTime)
	return stockEntry, nil
}

func UpdateStockEntry(stockEntry *model.StockEntry) (*model.StockEntry, error) {
	if stockEntry.Id == 0 {
		return nil, errors.New("未指定Id")
	}
	data := make(map[string]interface{})
	if stockEntry.EntryNo != "" {
		data["entry_no"] = stockEntry.EntryNo
	}
	if stockEntry.ShopId != "" {
		data["shop_id"] = stockEntry.ShopId
	}
	if stockEntry.Type != "" {
		data["type"] = stockEntry.Type
	}
	if stockEntry.SkuId != "" {
		data["sku_id"] = stockEntry.SkuId
	}
	if stockEntry.Unit != "" {
		data["unit"] = stockEntry.Unit
	}
	if stockEntry.OrderNo != "" {
		data["order_no"] = stockEntry.OrderNo
	}
	if stockEntry.Price > 0 {
		data["price"] = stockEntry.Price
	}
	if stockEntry.StorageId != "" {
		data["storage_id"] = stockEntry.StorageId
	}
	if stockEntry.Number > 0 {
		data["number"] = stockEntry.Number
	}
	if stockEntry.SupplierId != "" {
		data["supplier_id"] = stockEntry.SupplierId
	}
	if stockEntry.Remark != "" {
		data["remark"] = stockEntry.Remark
	}
	if stockEntry.Operator != "" {
		data["operator"] = stockEntry.Operator
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_STOCK_ENTRY).Where("id = ?", stockEntry.Id).Update(data).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	config.Redis.HSet("storage:last:"+stockEntry.ShopId, stockEntry.StorageId, stockEntry.Operator+","+stockEntry.EntryTime)
	return GetStockEntryById(stockEntry.Id), nil
}

func DeleteStockEntry(id int) error {
	if id == 0 {
		return errors.New("未指定入库单记录号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_STOCK_ENTRY).Where("id = ?", id).Delete(model.StockEntry{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
