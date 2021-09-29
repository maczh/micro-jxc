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
	BASE_SERVICE             = "jxc-base"
	URI_GET_DELIVERY_TYPE    = "/deliverytype/get"
	URI_LIST_DELIVERY_TYPE   = "/deliverytype/list"
	URI_SAVE_DELIVERY_TYPE   = "/deliverytype/save"
	URI_UPDATE_DELIVERY_TYPE = "/deliverytype/update"
	URI_DEL_DELIVERY_TYPE    = "/deliverytype/del"
)

func GetDeliveryType(shopId, deliveryType string) (*model.DeliveryType, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["deliveryType"] = deliveryType
	res, err := utils.CallNacos(BASE_SERVICE, URI_GET_DELIVERY_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_GET_DELIVERY_TYPE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var deliveryType model.DeliveryType
		utils.FromJSON(utils.ToJSON(result.Data), &deliveryType)
		return &deliveryType, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListDeliveryType(shopId string) ([]model.DeliveryType, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	res, err := utils.CallNacos(BASE_SERVICE, URI_LIST_DELIVERY_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_LIST_DELIVERY_TYPE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var deliveryTypeList []model.DeliveryType
		utils.FromJSON(utils.ToJSON(result.Data), &deliveryTypeList)
		return deliveryTypeList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func SaveDeliveryType(shopId, deliveryType string) (*model.DeliveryType, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["deliveryType"] = deliveryType
	res, err := utils.CallNacos(BASE_SERVICE, URI_SAVE_DELIVERY_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_SAVE_DELIVERY_TYPE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var deliveryType model.DeliveryType
		utils.FromJSON(utils.ToJSON(result.Data), &deliveryType)
		return &deliveryType, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func UpdateDeliveryType(id int, deliveryType string) (*model.DeliveryType, error) {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	params["deliveryType"] = deliveryType
	res, err := utils.CallNacos(BASE_SERVICE, URI_UPDATE_DELIVERY_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_UPDATE_DELIVERY_TYPE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var deliveryType model.DeliveryType
		utils.FromJSON(utils.ToJSON(result.Data), &deliveryType)
		return &deliveryType, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func DeleteDeliveryType(id int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(BASE_SERVICE, URI_DEL_DELIVERY_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_DEL_DELIVERY_TYPE, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}
