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
	URI_GET_PRODUCT_UNIT    = "/unit/get"
	URI_LIST_PRODUCT_UNIT   = "/unit/list"
	URI_SAVE_PRODUCT_UNIT   = "/unit/save"
	URI_UPDATE_PRODUCT_UNIT = "/unit/update"
	URI_DEL_PRODUCT_UNIT    = "/unit/del"
)

func GetProductUnit(id int, shopId, productId, unit, baseUnit string) (*model.ProductUnit, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if productId != "" {
		params["productId"] = productId
	}
	if unit != "" {
		params["unit"] = unit
	}
	if baseUnit != "" {
		params["baseUnit"] = baseUnit
	}
	if id != 0 {
		params["id"] = strconv.Itoa(id)
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_GET_PRODUCT_UNIT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_GET_PRODUCT_UNIT, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productUnit model.ProductUnit
		utils.FromJSON(utils.ToJSON(result.Data), &productUnit)
		return &productUnit, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListProductUnit(shopId, productId string) ([]model.ProductUnit, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if productId != "" {
		params["productId"] = productId
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_LIST_PRODUCT_UNIT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_LIST_PRODUCT_UNIT, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productUnitList []model.ProductUnit
		utils.FromJSON(utils.ToJSON(result.Data), &productUnitList)
		return productUnitList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func SaveProductUnit(shopId, productId, unit, baseUnit string, scale int) (*model.ProductUnit, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["productId"] = productId
	params["unit"] = unit
	params["baseUnit"] = baseUnit
	params["scale"] = strconv.Itoa(scale)
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_SAVE_PRODUCT_UNIT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_SAVE_PRODUCT_UNIT, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productUnit model.ProductUnit
		utils.FromJSON(utils.ToJSON(result.Data), &productUnit)
		return &productUnit, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func UpdateProductUnit(shopId, productId, unit, baseUnit string, scale int) (*model.ProductUnit, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["productId"] = productId
	params["unit"] = unit
	params["baseUnit"] = baseUnit
	params["scale"] = strconv.Itoa(scale)
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_UPDATE_PRODUCT_UNIT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_UPDATE_PRODUCT_UNIT, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productUnit model.ProductUnit
		utils.FromJSON(utils.ToJSON(result.Data), &productUnit)
		return &productUnit, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func DeleteProductUnit(id int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_DEL_PRODUCT_UNIT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_DEL_PRODUCT_UNIT, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}
