package dao

import (
	"errors"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/jxc-base/model"
)

const TABLE_STOCK_MOVE string = "stock_move"

func ListStockMove(moveNo, shopId, skuId, fromStorageId, toStorageId, startTime, endTime string, page, size int) []model.StockMove {
	var stockMoveList []model.StockMove
	if page > 0 && size > 0 {
		config.Mysql.Table(TABLE_STOCK_MOVE).Where("shop_id = ? AND IF(? != '',move_no = ?,1=1) AND IF(? != '',sku_id = ?,1=1) AND IF(? != '',from_storage_id = ?,1=1) AND IF(? != '',to_storage_id = ?,1=1) AND IF(? != '',move_time > ?,1=1) AND IF(? != '',move_time < ?,1=1) ", shopId, moveNo, moveNo, skuId, skuId, fromStorageId, fromStorageId, toStorageId, toStorageId, startTime, startTime, endTime, endTime).Order("id DESC").Limit(size).Offset((page - 1) * size).Find(&stockMoveList)

	} else {
		config.Mysql.Table(TABLE_STOCK_MOVE).Where("shop_id = ? AND IF(? != '',move_no = ?,1=1) AND IF(? != '',sku_id = ?,1=1) AND IF(? != '',from_storage_id = ?,1=1) AND IF(? != '',to_storage_id = ?,1=1) AND IF(? != '',move_time > ?,1=1) AND IF(? != '',move_time < ?,1=1) ", shopId, moveNo, moveNo, skuId, skuId, fromStorageId, fromStorageId, toStorageId, toStorageId, startTime, startTime, endTime, endTime).Order("id DESC").Find(&stockMoveList)
	}
	if len(stockMoveList) == 0 {
		return nil
	} else {
		return stockMoveList
	}
}

func CountStockMove(moveNo, shopId, skuId, fromStorageId, toStorageId, startTime, endTime string) int {
	count := 0
	config.Mysql.Table(TABLE_STOCK_MOVE).Where("shop_id = ? AND IF(? != '',move_no = ?,1=1) AND IF(? != '',sku_id = ?,1=1) AND IF(? != '',from_storage_id = ?,1=1) AND IF(? != '',to_storage_id = ?,1=1) AND IF(? != '',move_time > ?,1=1) AND IF(? != '',move_time < ?,1=1) ", shopId, moveNo, moveNo, skuId, skuId, fromStorageId, fromStorageId, toStorageId, toStorageId, startTime, startTime, endTime, endTime).Count(&count)
	return count
}

func GetStockMove(moveNo, shopId, skuId string) *model.StockMove {
	var stockMove model.StockMove
	config.Mysql.Table(TABLE_STOCK_MOVE).Where("shop_id = ? AND move_no = ? AND sku_id = ?", shopId, moveNo, skuId).First(&stockMove)
	if stockMove.Id == 0 {
		return nil
	} else {
		return &stockMove
	}
}

func GetStockMoveById(id int) *model.StockMove {
	var stockMove model.StockMove
	config.Mysql.Table(TABLE_STOCK_MOVE).Where("id = ?", id).First(&stockMove)
	if stockMove.Id == 0 {
		return nil
	} else {
		return &stockMove
	}
}

func SaveStockMove(stockMove *model.StockMove) (*model.StockMove, error) {
	s := GetStockMove(stockMove.MoveNo, stockMove.ShopId, stockMove.SkuId)
	if s != nil {
		return s, nil
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_STOCK_MOVE).Create(stockMove).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	config.Redis.HSet("storage:last:"+stockMove.ShopId, stockMove.FromStorageId, stockMove.Operator+","+stockMove.MoveTime)
	config.Redis.HSet("storage:last:"+stockMove.ShopId, stockMove.ToStorageId, stockMove.Operator+","+stockMove.MoveTime)
	return stockMove, nil
}

func UpdateStockMove(stockMove *model.StockMove) (*model.StockMove, error) {
	if stockMove.Id == 0 {
		return nil, errors.New("未指定Id")
	}
	data := make(map[string]interface{})
	if stockMove.MoveNo != "" {
		data["move_no"] = stockMove.MoveNo
	}
	if stockMove.ShopId != "" {
		data["shop_id"] = stockMove.ShopId
	}
	if stockMove.SkuId != "" {
		data["sku_id"] = stockMove.SkuId
	}
	if stockMove.Unit != "" {
		data["unit"] = stockMove.Unit
	}
	if stockMove.FromStorageId != "" {
		data["from_storage_id"] = stockMove.FromStorageId
	}
	if stockMove.ToStorageId != "" {
		data["to_storage_id"] = stockMove.ToStorageId
	}
	if stockMove.Number > 0 {
		data["number"] = stockMove.Number
	}
	if stockMove.Remark != "" {
		data["remark"] = stockMove.Remark
	}
	if stockMove.Operator != "" {
		data["operator"] = stockMove.Operator
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_STOCK_MOVE).Where("id = ?", stockMove.Id).Update(data).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	config.Redis.HSet("storage:last:"+stockMove.ShopId, stockMove.FromStorageId, stockMove.Operator+","+stockMove.MoveTime)
	config.Redis.HSet("storage:last:"+stockMove.ShopId, stockMove.ToStorageId, stockMove.Operator+","+stockMove.MoveTime)
	return GetStockMoveById(stockMove.Id), nil
}

func DeleteStockMove(id int) error {
	if id == 0 {
		return errors.New("未指定移库单记录号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_STOCK_MOVE).Where("id = ?", id).Delete(model.StockMove{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
