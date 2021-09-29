package nacos

import (
	"errors"
	"ququ.im/common"
	"ququ.im/common/logs"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
	"strconv"
)

const (
	URI_GET_STOCK_MOVE    = "/move/get"
	URI_LIST_STOCK_MOVE   = "/move/list"
	URI_SAVE_STOCK_MOVE   = "/move/save"
	URI_UPDATE_STOCK_MOVE = "/move/update"
	URI_DEL_STOCK_MOVE    = "/move/del"
)

func GetStockMove(id int, shopId, skuId, moveNo string) (*model.StockMove, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if skuId != "" {
		params["skuId"] = skuId
	}
	if moveNo != "" {
		params["moveNo"] = moveNo
	}
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_GET_STOCK_MOVE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_GET_STOCK_MOVE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var stockMove model.StockMove
		utils.FromJSON(utils.ToJSON(result.Data), &stockMove)
		return &stockMove, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListStockMove(shopId, skuId, moveNo, fromStorageId, toStorageId, startTime, endTime string, page, size int) ([]model.StockMove, *common.ResultPage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["skuId"] = skuId
	params["moveNo"] = moveNo
	params["fromStorageId"] = fromStorageId
	params["toStorageId"] = toStorageId
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["page"] = strconv.Itoa(page)
	params["size"] = strconv.Itoa(size)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_LIST_STOCK_MOVE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_LIST_STOCK_MOVE, err.Error())
		return nil, nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var stockMoveList []model.StockMove
		utils.FromJSON(utils.ToJSON(result.Data), &stockMoveList)
		return stockMoveList, result.Page, nil
	} else {
		return nil, nil, errors.New(result.Msg)
	}
}

func SaveStockMove(shopId, skuId, moveNo, unit, fromStorageId, toStorageId, operator, remark string, number int) common.Result {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["skuId"] = skuId
	params["moveNo"] = moveNo
	params["unit"] = unit
	params["fromStorageId"] = fromStorageId
	params["toStorageId"] = toStorageId
	params["operator"] = operator
	params["remark"] = remark
	params["number"] = strconv.Itoa(number)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_SAVE_STOCK_MOVE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_SAVE_STOCK_MOVE, err.Error())
		return *common.Error(-1, err.Error())
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}

func UpdateStockMove(id int, skuId, unit, fromStorageId, toStorageId, operator, remark string, number int) (*model.StockMove, error) {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	params["skuId"] = skuId
	params["unit"] = unit
	params["fromStorageId"] = fromStorageId
	params["toStorageId"] = toStorageId
	params["operator"] = operator
	params["remark"] = remark
	params["number"] = strconv.Itoa(number)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_UPDATE_STOCK_MOVE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_UPDATE_STOCK_MOVE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var stockMove model.StockMove
		utils.FromJSON(utils.ToJSON(result.Data), &stockMove)
		return &stockMove, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func DeleteStockMove(id int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_DEL_STOCK_MOVE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_DEL_STOCK_MOVE, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}
