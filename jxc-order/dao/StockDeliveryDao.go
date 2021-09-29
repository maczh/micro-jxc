package dao

import (
	"errors"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
)

const TABLE_STOCK_DELIVERY string = "stock_delivery"

func ListStockDelivery(deliveryNo, shopId, skuId, storageId, orderNo, customerId, startTime, endTime string, page, size int) []model.StockDelivery {
	var stockDeliveryList []model.StockDelivery
	if page > 0 && size > 0 {
		config.Mysql.Table(TABLE_STOCK_DELIVERY).Where("shop_id = ? AND IF(? != '',delivery_no = ?,1=1) AND IF(? != '',sku_id = ?,1=1) AND IF(? != '',storage_id = ?,1=1) AND IF(? != '',order_no = ?,1=1) AND IF(? != '',customer_id = ?,1=1) AND IF(? != '',delivery_time > ?,1=1) AND IF(? != '',delivery_time < ?,1=1) ", shopId, deliveryNo, deliveryNo, skuId, skuId, storageId, storageId, orderNo, orderNo, customerId, customerId, startTime, startTime, endTime, endTime).Order("id DESC").Limit(size).Offset((page - 1) * size).Find(&stockDeliveryList)

	} else {
		config.Mysql.Table(TABLE_STOCK_DELIVERY).Where("shop_id = ? AND IF(? != '',delivery_no = ?,1=1) AND IF(? != '',sku_id = ?,1=1) AND IF(? != '',storage_id = ?,1=1) AND IF(? != '',order_no = ?,1=1) AND IF(? != '',customer_id = ?,1=1) AND IF(? != '',delivery_time > ?,1=1) AND IF(? != '',delivery_time < ?,1=1) ", shopId, deliveryNo, deliveryNo, skuId, skuId, storageId, storageId, orderNo, orderNo, customerId, customerId, startTime, startTime, endTime, endTime).Order("id DESC").Find(&stockDeliveryList)
	}
	if len(stockDeliveryList) == 0 {
		return nil
	} else {
		for i := 0; i < len(stockDeliveryList); i++ {
			stockDeliveryList[i].DeliveryTime = utils.GormTimeFormat(stockDeliveryList[i].DeliveryTime)
		}
		return stockDeliveryList
	}
}

func CountStockDelivery(deliveryNo, shopId, skuGuid, storageId, orderNo, customerId, startTime, endTime string) int {
	count := 0
	config.Mysql.Table(TABLE_STOCK_DELIVERY).Where("shop_id = ? AND IF(? != '',delivery_no = ?,1=1) AND IF(? != '',sku_id = ?,1=1) AND IF(? != '',storage_id = ?,1=1) AND IF(? != '',order_no = ?,1=1) AND IF(? != '',customer_id = ?,1=1) AND IF(? != '',delivery_time > ?,1=1) AND IF(? != '',delivery_time < ?,1=1) ", shopId, deliveryNo, deliveryNo, skuGuid, skuGuid, storageId, storageId, orderNo, orderNo, customerId, customerId, startTime, startTime, endTime, endTime).Count(&count)
	return count
}

func GetStockDelivery(shopId, deliveryNo, skuId string) *model.StockDelivery {
	var stockDelivery model.StockDelivery
	config.Mysql.Table(TABLE_STOCK_DELIVERY).Where("shop_id = ? AND delivery_no = ? AND sku_id = ?", shopId, deliveryNo, skuId).First(&stockDelivery)
	if stockDelivery.Id == 0 {
		return nil
	} else {
		stockDelivery.DeliveryTime = utils.GormTimeFormat(stockDelivery.DeliveryTime)
		return &stockDelivery
	}
}

func GetStockDeliveryById(id int) *model.StockDelivery {
	var stockDelivery model.StockDelivery
	config.Mysql.Table(TABLE_STOCK_DELIVERY).Where("id = ?", id).First(&stockDelivery)
	if stockDelivery.Id == 0 {
		return nil
	} else {
		stockDelivery.DeliveryTime = utils.GormTimeFormat(stockDelivery.DeliveryTime)
		return &stockDelivery
	}
}

func SaveStockDelivery(stockDelivery *model.StockDelivery) (*model.StockDelivery, error) {
	s := GetStockDelivery(stockDelivery.ShopId, stockDelivery.DeliveryNo, stockDelivery.SkuId)
	if s != nil {
		return s, nil
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_STOCK_DELIVERY).Create(stockDelivery).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	config.Redis.HSet("storage:last:"+stockDelivery.ShopId, stockDelivery.StorageId, stockDelivery.Operator+","+stockDelivery.DeliveryTime)
	return GetStockDelivery(stockDelivery.ShopId, stockDelivery.DeliveryNo, stockDelivery.SkuId), nil
}

func UpdateStockDelivery(stockDelivery *model.StockDelivery) (*model.StockDelivery, error) {
	if stockDelivery.Id == 0 {
		return nil, errors.New("未指定Id")
	}
	data := make(map[string]interface{})
	if stockDelivery.DeliveryNo != "" {
		data["delivery_no"] = stockDelivery.DeliveryNo
	}
	if stockDelivery.ShopId != "" {
		data["shop_id"] = stockDelivery.ShopId
	}
	if stockDelivery.Type != "" {
		data["type"] = stockDelivery.Type
	}
	if stockDelivery.SkuId != "" {
		data["sku_id"] = stockDelivery.SkuId
	}
	if stockDelivery.Unit != "" {
		data["unit"] = stockDelivery.Unit
	}
	if stockDelivery.OrderNo != "" {
		data["order_no"] = stockDelivery.OrderNo
	}
	if stockDelivery.Cost > 0 {
		data["cost"] = stockDelivery.Cost
	}
	if stockDelivery.Price > 0 {
		data["price"] = stockDelivery.Price
	}
	if stockDelivery.StorageId != "" {
		data["storage_id"] = stockDelivery.StorageId
	}
	if stockDelivery.MultiStorage != "" {
		data["multi_storage"] = stockDelivery.MultiStorage
	}
	if stockDelivery.Number > 0 {
		data["number"] = stockDelivery.Number
	}
	if stockDelivery.CustomerId != "" {
		data["customer_id"] = stockDelivery.CustomerId
	}
	if stockDelivery.Remark != "" {
		data["remark"] = stockDelivery.Remark
	}
	if stockDelivery.Operator != "" {
		data["operator"] = stockDelivery.Operator
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_STOCK_DELIVERY).Where("id = ?", stockDelivery.Id).Update(data).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	config.Redis.HSet("storage:last:"+stockDelivery.ShopId, stockDelivery.StorageId, stockDelivery.Operator+","+stockDelivery.DeliveryTime)
	return GetStockDeliveryById(stockDelivery.Id), nil
}

func DeleteStockDelivery(id int) error {
	if id == 0 {
		return errors.New("未指定出库单记录号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_STOCK_DELIVERY).Where("id = ?", id).Delete(model.StockDelivery{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
