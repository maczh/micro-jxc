package nacos

import (
	"errors"
	"ququ.im/common"
	"ququ.im/common/logs"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
	"strconv"
	"strings"
)

const (
	URI_GET_SHOP_STORAGE    = "/storage/get"
	URI_LIST_SHOP_STORAGE   = "/storage/list"
	URI_SAVE_SHOP_STORAGE   = "/storage/save"
	URI_UPDATE_SHOP_STORAGE = "/storage/update"
	URI_DEL_SHOP_STORAGE    = "/storage/del"
	URI_UP_SHOP_STORAGE     = "/storage/up"
	URI_DOWN_SHOP_STORAGE   = "/storage/down"
	URI_LAST_SHOP_STORAGE   = "/storage/last"
)

func GetShopStorage(shopId, storageId, storageName string) (*model.ShopStorage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	if storageId != "" {
		params["storageId"] = storageId
	}
	if storageName != "" {
		params["name"] = storageName
	}
	res, err := utils.CallNacos(BASE_SERVICE, URI_GET_SHOP_STORAGE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_GET_SHOP_STORAGE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var shopStorage model.ShopStorage
		utils.FromJSON(utils.ToJSON(result.Data), &shopStorage)
		return &shopStorage, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListShopStorage(shopId string) ([]model.ShopStorage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	res, err := utils.CallNacos(BASE_SERVICE, URI_LIST_SHOP_STORAGE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_LIST_SHOP_STORAGE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var shopStorageList []model.ShopStorage
		utils.FromJSON(utils.ToJSON(result.Data), &shopStorageList)
		return shopStorageList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func SaveShopStorage(shopId, storageId, storageName, remark string) (*model.ShopStorage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["name"] = storageName
	if storageId != "" {
		params["storageId"] = storageId
	}
	if remark != "" {
		params["remark"] = remark
	}
	res, err := utils.CallNacos(BASE_SERVICE, URI_SAVE_SHOP_STORAGE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_SAVE_SHOP_STORAGE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var shopStorage model.ShopStorage
		utils.FromJSON(utils.ToJSON(result.Data), &shopStorage)
		return &shopStorage, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func UpdateShopStorage(id int, shopId, storageId, storageName, remark string) (*model.ShopStorage, error) {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if storageId != "" {
		params["storageId"] = storageId
	}
	if storageName != "" {
		params["name"] = storageName
	}
	if remark != "" {
		params["remark"] = remark
	}
	res, err := utils.CallNacos(BASE_SERVICE, URI_UPDATE_SHOP_STORAGE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_UPDATE_SHOP_STORAGE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var shopStorage model.ShopStorage
		utils.FromJSON(utils.ToJSON(result.Data), &shopStorage)
		return &shopStorage, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func DeleteShopStorage(shopId, storageId string) common.Result {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["storageId"] = storageId
	res, err := utils.CallNacos(BASE_SERVICE, URI_DEL_SHOP_STORAGE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_DEL_SHOP_STORAGE, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}

func DecrShopStorageSortNumber(shopId, storageId string) ([]model.ShopStorage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["storageId"] = storageId
	res, err := utils.CallNacos(BASE_SERVICE, URI_UP_SHOP_STORAGE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_UP_SHOP_STORAGE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var shopStorageList []model.ShopStorage
		utils.FromJSON(utils.ToJSON(result.Data), &shopStorageList)
		return shopStorageList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func IncrShopStorageSortNumber(shopId, storageId string) ([]model.ShopStorage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["storageId"] = storageId
	res, err := utils.CallNacos(BASE_SERVICE, URI_DOWN_SHOP_STORAGE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_DOWN_SHOP_STORAGE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var shopStorageList []model.ShopStorage
		utils.FromJSON(utils.ToJSON(result.Data), &shopStorageList)
		return shopStorageList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func GetStorageLastOperator(shopId, storageId string) map[string]string {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["storageId"] = storageId
	res, err := utils.CallNacos(BASE_SERVICE, URI_LAST_SHOP_STORAGE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_LAST_SHOP_STORAGE, err.Error())
		return nil
	}
	var result common.Result
	utils.FromJSON(res, &result)
	resultMap := result.Data.(map[string]string)
	s := strings.Split(resultMap["lastInfo"], ",")
	m := make(map[string]string)
	m["operator"] = s[0]
	m["time"] = utils.GormTimeFormat(s[1])
	return m
}
