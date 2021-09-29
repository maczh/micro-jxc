package service

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
	"ququ.im/jxc-order/dao"
	"time"
)

func ListStockMove(shopId, moveNo, skuId, fromStorageId, toStorageId, startTime, endTime string, page, size int) common.Result {
	if shopId == "" {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if endTime != "" {
		endTime = endTime + " 23:59:59"
	}
	stockMoveList := dao.ListStockMove(moveNo, shopId, skuId, fromStorageId, toStorageId, startTime, endTime, page, size)
	if stockMoveList == nil || len(stockMoveList) == 0 {
		return *common.Error(-1, "查无数据")
	} else {
		if page > 0 && size > 0 {
			count := dao.CountStockMove(moveNo, shopId, skuId, fromStorageId, toStorageId, startTime, endTime)
			return *common.SuccessWithPage(stockMoveList, count/size+1, page, size, count)
		} else {
			return *common.Success(stockMoveList)
		}
	}
}

func GetStockMove(id int, shopId, moveNo, skuId string) common.Result {
	var stockMove *model.StockMove
	if id > 0 {
		stockMove = dao.GetStockMoveById(id)
	} else {
		if shopId == "" {
			return *common.Error(-1, "未传入商户账号参数")
		}
		if moveNo == "" {
			return *common.Error(-1, "未传入移库单号参数")
		}
		if skuId == "" {
			return *common.Error(-1, "未传入移库商品编号参数")
		}
		stockMove = dao.GetStockMove(shopId, moveNo, skuId)
	}
	if stockMove == nil || stockMove.Id == 0 {
		return *common.Error(-1, "查无此数据")
	} else {
		return *common.Success(stockMove)
	}
}

func SaveStockMove(shopId, moveNo, skuId, unit, fromStorageId, toStorageId, operator, remark string, number int) common.Result {
	if shopId == "" {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if moveNo == "" {
		return *common.Error(-1, "未传入移库单号参数")
	}
	if skuId == "" {
		return *common.Error(-1, "未传入移库商品编号参数")
	}
	if fromStorageId == "" {
		return *common.Error(-1, "未传入移出的仓库编号参数")
	}
	if toStorageId == "" {
		return *common.Error(-1, "未传入移入的仓库编号参数")
	}
	if unit == "" {
		return *common.Error(-1, "未传入货品单位参数")
	}
	if number == 0 {
		return *common.Error(-1, "未传入移库数量参数")
	}
	stockMove := new(model.StockMove)
	stockMove.ShopId = shopId
	stockMove.MoveNo = moveNo
	stockMove.SkuId = skuId
	stockMove.Unit = unit
	stockMove.Number = number
	stockMove.FromStorageId = fromStorageId
	stockMove.ToStorageId = toStorageId
	stockMove.Remark = remark
	stockMove.Operator = operator
	stockMove.MoveTime = utils.ToDateTimeString(time.Now())
	stockMove, _ = dao.SaveStockMove(stockMove)
	return *common.Success(stockMove)
}

func UpdateStockMove(id int, skuId, unit, fromStorageId, toStorageId, operator, remark string, number int) common.Result {
	stockMove := dao.GetStockMoveById(id)
	if stockMove == nil {
		return *common.Error(-1, "无此移库单条目")
	}
	stockMove.Id = id
	stockMove.SkuId = skuId
	stockMove.Unit = unit
	stockMove.Number = number
	stockMove.FromStorageId = fromStorageId
	stockMove.ToStorageId = toStorageId
	stockMove.Remark = remark
	stockMove.Operator = operator
	stockMove, _ = dao.UpdateStockMove(stockMove)
	return *common.Success(stockMove)
}

func DeleteStockMove(id int) common.Result {
	err := dao.DeleteStockMove(id)
	if err != nil {
		return *common.Error(-1, "无此移库单记录")
	}
	return *common.Success(nil)
}
