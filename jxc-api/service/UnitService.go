package service

import (
	"ququ.im/common"
	"ququ.im/jxc-api/nacos"
)

func ListUnitByShopId(shopId string) common.Result {
	unitList, err := nacos.ListUnit(shopId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if unitList == nil {
		return *common.Error(-1, "未定义单位")
	}
	return *common.Success(unitList)
}

func SaveUnit(shopId, unit string) common.Result {
	u, err := nacos.SaveUnit(shopId, unit)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(u)
}

func UpdateUnit(unitId int, unit string) common.Result {
	u, err := nacos.UpdateUnit(unitId, unit)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(u)
}

func DeleteUnit(unitId int) common.Result {
	return nacos.DeleteUnit(unitId)
}
