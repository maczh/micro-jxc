package service

import (
	"ququ.im/common"
	"ququ.im/jxc-base/dao"
	"ququ.im/jxc-base/model"
)

func ListEntryTypeByShopId(shopId string) common.Result {
	entryTypeList := dao.ListEntryTypeByShopId(shopId)
	if entryTypeList == nil {
		return *common.Error(-1, "未定义入库类型")
	}
	return *common.Success(entryTypeList)
}

func SaveEntryType(shopId, entryType string) common.Result {
	u := new(model.EntryType)
	u.ShopId = shopId
	u.Type = entryType
	u, err := dao.SaveEntryType(u)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(u)
}

func UpdateEntryType(entryTypeId int, entryType string) common.Result {
	u := dao.GetEntryTypeById(entryTypeId)
	if u == nil {
		return *common.Error(-1, "无此入库类型记录")
	}
	u.Type = entryType
	u, err := dao.UpdateEntryType(u)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(u)
}

func DeleteEntryType(entryTypeId int) common.Result {
	err := dao.DeleteEntryType(entryTypeId)
	if err != nil {
		return *common.Error(-1, "无此入库类型记录")
	}
	return *common.Success(nil)
}
