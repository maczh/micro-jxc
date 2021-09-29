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
	URI_GET_SPECS    = "/specs/get"
	URI_LIST_SPECS   = "/specs/list"
	URI_SAVE_SPECS   = "/specs/save"
	URI_UPDATE_SPECS = "/specs/update"
	URI_DEL_SPECS    = "/specs/del"
	URI_ADD_SPECS    = "/specs/add"
	URI_REMOVE_SPECS = "/specs/remove"
)

func GetProductSpecs(id int, shopId, productId, specs string) (*model.ProductSpecs, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if productId != "" {
		params["productId"] = productId
	}
	if specs != "" {
		params["specs"] = specs
	}
	if id != 0 {
		params["id"] = strconv.Itoa(id)
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_GET_SPECS, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_GET_SPECS, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSpecs model.ProductSpecs
		utils.FromJSON(utils.ToJSON(result.Data), &productSpecs)
		return &productSpecs, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListProductSpecs(shopId, productId string) ([]model.ProductSpecs, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if productId != "" {
		params["productId"] = productId
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_LIST_SPECS, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_LIST_SPECS, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSpecsList []model.ProductSpecs
		utils.FromJSON(utils.ToJSON(result.Data), &productSpecsList)
		return productSpecsList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func SaveProductSpecs(shopId, productId, specs, values string) (*model.ProductSpecs, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["specs"] = specs
	params["productId"] = productId
	params["values"] = values
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_SAVE_SPECS, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_SAVE_SPECS, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSpecs model.ProductSpecs
		utils.FromJSON(utils.ToJSON(result.Data), &productSpecs)
		return &productSpecs, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func UpdateProductSpecs(id int, shopId, productId, specs, values string) (*model.ProductSpecs, error) {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if productId != "" {
		params["productId"] = productId
	}
	if specs != "" {
		params["specs"] = specs
	}
	if values != "" {
		params["values"] = values
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_UPDATE_SPECS, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_UPDATE_SPECS, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSpecs model.ProductSpecs
		utils.FromJSON(utils.ToJSON(result.Data), &productSpecs)
		return &productSpecs, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func DeleteProductSpecs(id int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_DEL_SPECS, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_DEL_SPECS, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}

func AddSpecsValue(id int, value string) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	params["value"] = value
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_ADD_SPECS, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_ADD_SPECS, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}

func RemoveSpecsValue(id int, value string) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	params["value"] = value
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_REMOVE_SPECS, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_REMOVE_SPECS, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}
