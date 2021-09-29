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
	URI_GET_STOCK_ENTRY    = "/entry/get"
	URI_LIST_STOCK_ENTRY   = "/entry/list"
	URI_SAVE_STOCK_ENTRY   = "/entry/save"
	URI_UPDATE_STOCK_ENTRY = "/entry/update"
	URI_DEL_STOCK_ENTRY    = "/entry/del"
)

func GetStockEntry(shopId, skuId, entryNo string, id int) (*model.StockEntry, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if skuId != "" {
		params["skuId"] = skuId
	}
	if entryNo != "" {
		params["entryNo"] = entryNo
	}
	if id > 0 {
		params["id"] = strconv.Itoa(id)
	}
	res, err := utils.CallNacos(ORDER_SERVICE, URI_GET_STOCK_ENTRY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_GET_STOCK_ENTRY, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var stockEntry model.StockEntry
		utils.FromJSON(utils.ToJSON(result.Data), &stockEntry)
		return &stockEntry, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListStockEntry(shopId, skuId, entryNo, orderNo, supplierId, storageId, startTime, endTime string, page, size int) ([]model.StockEntry, *common.ResultPage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["skuId"] = skuId
	params["entryNo"] = entryNo
	params["orderNo"] = orderNo
	params["supplierId"] = supplierId
	params["storageId"] = storageId
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["page"] = strconv.Itoa(page)
	params["size"] = strconv.Itoa(size)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_LIST_STOCK_ENTRY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_LIST_STOCK_ENTRY, err.Error())
		return nil, nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var stockEntryList []model.StockEntry
		utils.FromJSON(utils.ToJSON(result.Data), &stockEntryList)
		return stockEntryList, result.Page, nil
	} else {
		return nil, nil, errors.New(result.Msg)
	}
}

func SaveStockEntry(shopId, skuId, entryNo, entryType, unit, orderNo, storageId, supplierId, operator, remark string, number, price int) (*model.StockEntry, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["skuId"] = skuId
	params["entryNo"] = entryNo
	params["type"] = entryType
	params["unit"] = unit
	params["orderNo"] = orderNo
	params["storageId"] = storageId
	params["supplierId"] = supplierId
	params["operator"] = operator
	params["remark"] = remark
	params["number"] = strconv.Itoa(number)
	params["price"] = strconv.Itoa(price)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_SAVE_STOCK_ENTRY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_SAVE_STOCK_ENTRY, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var stockEntry model.StockEntry
		utils.FromJSON(utils.ToJSON(result.Data), &stockEntry)
		return &stockEntry, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func UpdateStockEntry(id int, skuId, unit, orderNo, storageId, supplierId, operator, remark string, number, price int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	params["skuId"] = skuId
	params["unit"] = unit
	params["orderNo"] = orderNo
	params["storageId"] = storageId
	params["supplierId"] = supplierId
	params["operator"] = operator
	params["remark"] = remark
	params["number"] = strconv.Itoa(number)
	params["price"] = strconv.Itoa(price)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_UPDATE_STOCK_ENTRY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_UPDATE_STOCK_ENTRY, err.Error())
		return *common.Error(-1, err.Error())
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}

func DeleteStockEntry(id int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_DEL_STOCK_ENTRY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_DEL_STOCK_ENTRY, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}
