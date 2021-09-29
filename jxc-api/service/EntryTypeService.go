package service

import (
	"ququ.im/common"
	"ququ.im/jxc-api/nacos"
)

func ListEntryTypeByShopId(shopId string) common.Result {
	entryTypeList, err := nacos.ListEntryType(shopId)
	if err != nil {
		return *common.Error(-1, "获取入库类型错误:"+err.Error())
	}
	return *common.Success(entryTypeList)
}

func SaveEntryType(shopId, entryType string) common.Result {
	u, err := nacos.SaveEntryType(shopId, entryType)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(u)
}

func UpdateEntryType(entryTypeId int, entryType string) common.Result {
	u, err := nacos.UpdateEntryType(entryTypeId, entryType)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(u)
}

func DeleteEntryType(entryTypeId int) common.Result {
	return nacos.DeleteEntryType(entryTypeId)
}
