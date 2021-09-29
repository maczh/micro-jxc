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
	ORDER_SERVICE             = "jxc-order"
	URI_GET_STOCK_DELIVERY    = "/delivery/get"
	URI_LIST_STOCK_DELIVERY   = "/delivery/list"
	URI_SAVE_STOCK_DELIVERY   = "/delivery/save"
	URI_UPDATE_STOCK_DELIVERY = "/delivery/update"
	URI_DEL_STOCK_DELIVERY    = "/delivery/del"
)

func GetStockDelivery(shopId, skuId, deliveryNo string, id int) (*model.StockDelivery, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if skuId != "" {
		params["skuId"] = skuId
	}
	if deliveryNo != "" {
		params["deliveryNo"] = deliveryNo
	}
	if id > 0 {
		params["id"] = strconv.Itoa(id)
	}
	res, err := utils.CallNacos(ORDER_SERVICE, URI_GET_STOCK_DELIVERY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_GET_STOCK_DELIVERY, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var stockDelivery model.StockDelivery
		utils.FromJSON(utils.ToJSON(result.Data), &stockDelivery)
		return &stockDelivery, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListStockDelivery(shopId, skuId, deliveryNo, orderNo, customerId, storageId, startTime, endTime string, page, size int) ([]model.StockDelivery, *common.ResultPage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["skuId"] = skuId
	params["deliveryNo"] = deliveryNo
	params["orderNo"] = orderNo
	params["customerId"] = customerId
	params["storageId"] = storageId
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["page"] = strconv.Itoa(page)
	params["size"] = strconv.Itoa(size)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_LIST_STOCK_DELIVERY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_LIST_STOCK_DELIVERY, err.Error())
		return nil, nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var stockDeliveryList []model.StockDelivery
		utils.FromJSON(utils.ToJSON(result.Data), &stockDeliveryList)
		return stockDeliveryList, result.Page, nil
	} else {
		return nil, nil, errors.New(result.Msg)
	}
}

func SaveStockDelivery(shopId, skuId, deliveryNo, deliveryType, unit, orderNo, storageId, customerId, customerName, operator, remark string, number, cost, price int) (*model.StockDelivery, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["skuId"] = skuId
	params["deliveryNo"] = deliveryNo
	params["type"] = deliveryType
	params["unit"] = unit
	params["orderNo"] = orderNo
	params["storageId"] = storageId
	params["customerId"] = customerId
	params["customerName"] = customerName
	params["operator"] = operator
	params["remark"] = remark
	params["number"] = strconv.Itoa(number)
	params["cost"] = strconv.Itoa(cost)
	params["price"] = strconv.Itoa(price)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_SAVE_STOCK_DELIVERY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_SAVE_STOCK_DELIVERY, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var stockDelivery model.StockDelivery
		utils.FromJSON(utils.ToJSON(result.Data), &stockDelivery)
		return &stockDelivery, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func UpdateStockDelivery(id int, skuId, unit, orderNo, storageId, customerId, customerName, operator, remark string, number, cost, price int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	params["skuId"] = skuId
	params["unit"] = unit
	params["orderNo"] = orderNo
	params["storageId"] = storageId
	params["customerId"] = customerId
	params["customerName"] = customerName
	params["operator"] = operator
	params["remark"] = remark
	params["number"] = strconv.Itoa(number)
	params["cost"] = strconv.Itoa(cost)
	params["price"] = strconv.Itoa(price)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_UPDATE_STOCK_DELIVERY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_UPDATE_STOCK_DELIVERY, err.Error())
		return *common.Error(-1, err.Error())
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}

func DeleteStockDelivery(id int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(ORDER_SERVICE, URI_DEL_STOCK_DELIVERY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", ORDER_SERVICE, URI_DEL_STOCK_DELIVERY, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}
