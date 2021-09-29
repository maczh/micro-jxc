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
	URI_GET_ENTRY_TYPE    = "/entrytype/get"
	URI_LIST_ENTRY_TYPE   = "/entrytype/list"
	URI_SAVE_ENTRY_TYPE   = "/entrytype/save"
	URI_UPDATE_ENTRY_TYPE = "/entrytype/update"
	URI_DEL_ENTRY_TYPE    = "/entrytype/del"
)

func GetEntryType(shopId, entryType string) (*model.EntryType, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["entryType"] = entryType
	res, err := utils.CallNacos(BASE_SERVICE, URI_GET_ENTRY_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_GET_ENTRY_TYPE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var entryType model.EntryType
		utils.FromJSON(utils.ToJSON(result.Data), &entryType)
		return &entryType, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListEntryType(shopId string) ([]model.EntryType, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	res, err := utils.CallNacos(BASE_SERVICE, URI_LIST_ENTRY_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_LIST_ENTRY_TYPE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var entryTypeList []model.EntryType
		utils.FromJSON(utils.ToJSON(result.Data), &entryTypeList)
		return entryTypeList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func SaveEntryType(shopId, entryType string) (*model.EntryType, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["entryType"] = entryType
	res, err := utils.CallNacos(BASE_SERVICE, URI_SAVE_ENTRY_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_SAVE_ENTRY_TYPE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var entryType model.EntryType
		utils.FromJSON(utils.ToJSON(result.Data), &entryType)
		return &entryType, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func UpdateEntryType(id int, entryType string) (*model.EntryType, error) {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	params["entryType"] = entryType
	res, err := utils.CallNacos(BASE_SERVICE, URI_UPDATE_ENTRY_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_UPDATE_ENTRY_TYPE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var entryType model.EntryType
		utils.FromJSON(utils.ToJSON(result.Data), &entryType)
		return &entryType, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func DeleteEntryType(id int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(BASE_SERVICE, URI_DEL_ENTRY_TYPE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_DEL_ENTRY_TYPE, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}
