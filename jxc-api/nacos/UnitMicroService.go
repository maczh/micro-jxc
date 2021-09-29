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
	URI_LIST_UNIT   = "/unit/list"
	URI_SAVE_UNIT   = "/unit/save"
	URI_UPDATE_UNIT = "/unit/update"
	URI_DEL_UNIT    = "/unit/del"
)

func ListUnit(shopId string) ([]model.Unit, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	res, err := utils.CallNacos(BASE_SERVICE, URI_LIST_UNIT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_LIST_UNIT, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var unitList []model.Unit
		utils.FromJSON(utils.ToJSON(result.Data), &unitList)
		return unitList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func SaveUnit(shopId, unit string) (*model.Unit, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	params["unit"] = unit
	res, err := utils.CallNacos(BASE_SERVICE, URI_SAVE_UNIT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_SAVE_UNIT, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var unit model.Unit
		utils.FromJSON(utils.ToJSON(result.Data), &unit)
		return &unit, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func UpdateUnit(id int, unit string) (*model.Unit, error) {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	params["unit"] = unit
	res, err := utils.CallNacos(BASE_SERVICE, URI_UPDATE_UNIT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_UPDATE_UNIT, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var unit model.Unit
		utils.FromJSON(utils.ToJSON(result.Data), &unit)
		return &unit, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func DeleteUnit(id int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(BASE_SERVICE, URI_DEL_UNIT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_DEL_UNIT, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}
