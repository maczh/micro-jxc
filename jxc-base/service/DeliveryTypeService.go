package service

import (
	"ququ.im/common"
	"ququ.im/jxc-base/dao"
	"ququ.im/jxc-base/model"
)

func ListDeliveryTypeByShopId(shopId string) common.Result {
	deliveryTypeList := dao.ListDeliveryTypeByShopId(shopId)
	if deliveryTypeList == nil {
		return *common.Error(-1, "未定义出库类型")
	}
	return *common.Success(deliveryTypeList)
}

func SaveDeliveryType(shopId, deliveryType string) common.Result {
	u := new(model.DeliveryType)
	u.ShopId = shopId
	u.Type = deliveryType
	u, err := dao.SaveDeliveryType(u)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(u)
}

func UpdateDeliveryType(deliveryTypeId int, deliveryType string) common.Result {
	u := dao.GetDeliveryTypeById(deliveryTypeId)
	if u == nil {
		return *common.Error(-1, "无此出库类型记录")
	}
	u.Type = deliveryType
	u, err := dao.UpdateDeliveryType(u)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(u)
}

func DeleteDeliveryType(deliveryTypeId int) common.Result {
	err := dao.DeleteDeliveryType(deliveryTypeId)
	if err != nil {
		return *common.Error(-1, "无此出库类型记录")
	}
	return *common.Success(nil)
}
