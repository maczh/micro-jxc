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
	STOCK_SERVICE     = "jxc-stock"
	URI_GET_STOCK     = "/stock/get"
	URI_LIST_STOCK    = "/stock/list"
	URI_INOUT_STOCK   = "/stock/inout"
	URI_GET_STOCK_NUM = "/stock/num"
)

func GetSkuStock(shopId, skuId, storageId string) (*model.SkuStock, error) {
	params := make(map[string]string)
	params["storageId"] = storageId
	params["shopId"] = shopId
	params["skuId"] = skuId
	res, err := utils.CallNacos(STOCK_SERVICE, URI_GET_STOCK, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", STOCK_SERVICE, URI_GET_STOCK, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 && result.Data != nil {
		var skuStock model.SkuStock
		utils.FromJSON(utils.ToJSON(result.Data), &skuStock)
		return &skuStock, nil
	} else if result.Data == nil {
		return nil, errors.New("查无数据")
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListSkuStock(shopId, skuId, nameKeyWord, productId, storageId string, page, size int) ([]model.SkuStock, *common.ResultPage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	if skuId != "" {
		params["skuId"] = skuId
	}
	if nameKeyWord != "" {
		params["name"] = nameKeyWord
	}
	if productId != "" {
		params["productId"] = productId
	}
	if storageId != "" {
		params["storageId"] = storageId
	}
	if page > -1 {
		params["page"] = strconv.Itoa(page)
	}
	if size > -1 {
		params["size"] = strconv.Itoa(size)
	}
	res, err := utils.CallNacos(STOCK_SERVICE, URI_LIST_STOCK, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", STOCK_SERVICE, URI_LIST_STOCK, err.Error())
		return nil, nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var skuStockList []model.SkuStock
		utils.FromJSON(utils.ToJSON(result.Data), &skuStockList)
		return skuStockList, result.Page, nil
	} else {
		return nil, nil, errors.New(result.Msg)
	}
}

func InOutSkuStock(shopId, skuId, unit, storageId string, number, cost, price int) common.Result {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["skuId"] = skuId
	params["storageId"] = storageId
	params["unit"] = unit
	params["number"] = strconv.Itoa(number)
	params["price"] = strconv.Itoa(price)
	params["cost"] = strconv.Itoa(cost)
	res, err := utils.CallNacos(STOCK_SERVICE, URI_INOUT_STOCK, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", STOCK_SERVICE, URI_INOUT_STOCK, err.Error())
		return *common.Error(-1, err.Error())
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}

func GetSkuStockNumber(shopId, skuId, storageId, productId string) (int, error) {
	params := make(map[string]string)
	params["storageId"] = storageId
	params["shopId"] = shopId
	params["skuId"] = skuId
	params["productId"] = productId
	res, err := utils.CallNacos(STOCK_SERVICE, URI_GET_STOCK_NUM, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", STOCK_SERVICE, URI_GET_STOCK_NUM, err.Error())
		return 0, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		skuStock := result.Data.(map[string]interface{})
		return int(skuStock["stocks"].(float64)), nil
	} else {
		return 0, errors.New(result.Msg)
	}
}
