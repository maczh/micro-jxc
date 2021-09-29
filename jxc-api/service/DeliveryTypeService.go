package service

import (
	"ququ.im/common"
	"ququ.im/jxc-api/nacos"
)

func ListDeliveryTypeByShopId(shopId string) common.Result {
	deliveryTypeList, err := nacos.ListDeliveryType(shopId)
	if err != nil {
		return *common.Error(-1, "获取出库类型列表异常:"+err.Error())
	}
	if deliveryTypeList == nil {
		return *common.Error(-1, "未定义出库类型")
	}
	return *common.Success(deliveryTypeList)
}

func SaveDeliveryType(shopId, deliveryType string) common.Result {
	dType, err := nacos.SaveDeliveryType(shopId, deliveryType)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(dType)
}

func UpdateDeliveryType(deliveryTypeId int, deliveryType string) common.Result {
	u, err := nacos.UpdateDeliveryType(deliveryTypeId, deliveryType)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(u)
}

func DeleteDeliveryType(deliveryTypeId int) common.Result {
	return nacos.DeleteDeliveryType(deliveryTypeId)
}
