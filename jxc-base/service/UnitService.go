package service

import (
	"ququ.im/common"
	"ququ.im/jxc-base/dao"
	"ququ.im/jxc-base/model"
)

func ListUnitByShopId(shopId string) common.Result {
	unitList := dao.ListUnitByShopId(shopId)
	if unitList == nil {
		return *common.Error(-1, "未定义单位")
	}
	return *common.Success(unitList)
}

func SaveUnit(shopId, unit string) common.Result {
	u := new(model.Unit)
	u.ShopId = shopId
	u.Unit = unit
	u, err := dao.SaveUnit(u)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(u)
}

func UpdateUnit(unitId int, unit string) common.Result {
	u := dao.GetUnitById(unitId)
	if u == nil {
		return *common.Error(-1, "无此单位记录")
	}
	u.Unit = unit
	u, err := dao.UpdateUnit(u)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(u)
}

func DeleteUnit(unitId int) common.Result {
	err := dao.DeleteUnit(unitId)
	if err != nil {
		return *common.Error(-1, "无此单位记录")
	}
	return *common.Success(nil)
}
