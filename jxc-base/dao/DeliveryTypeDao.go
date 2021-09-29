package dao

import (
	"errors"
	"ququ.im/common/cache"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/jxc-base/model"
)

const TABLE_DELIVERY_TYPE string = "delivery_type"

func ListDeliveryTypeByShopId(shopId string) []model.DeliveryType {
	var deliveryTypeList []model.DeliveryType
	err := cache.GetCache("deliverytype:shop", shopId, &deliveryTypeList)
	if err == nil && len(deliveryTypeList) > 0 {
		return deliveryTypeList
	}
	config.Mysql.Table(TABLE_DELIVERY_TYPE).Where("shop_id = ? OR shop_id = '' OR shop_id is null", shopId).Find(&deliveryTypeList)
	if len(deliveryTypeList) == 0 {
		return nil
	} else {
		cache.SetCache("deliverytype:shop", shopId, deliveryTypeList)
		return deliveryTypeList
	}
}

func GetDeliveryType(shopId, dType string) *model.DeliveryType {
	var deliveryType model.DeliveryType
	config.Mysql.Table(TABLE_DELIVERY_TYPE).Where("IF(? != '',shop_id = ?,1=1) AND type = ?", shopId, shopId, dType).First(&deliveryType)
	if deliveryType.Id == 0 {
		return nil
	} else {
		return &deliveryType
	}
}

func GetDeliveryTypeById(deliveryTypeId int) *model.DeliveryType {
	var deliveryType model.DeliveryType
	config.Mysql.Table(TABLE_DELIVERY_TYPE).Where("id = ?", deliveryTypeId).First(&deliveryType)
	if deliveryType.Id == 0 {
		return nil
	} else {
		return &deliveryType
	}
}

func SaveDeliveryType(deliveryType *model.DeliveryType) (*model.DeliveryType, error) {
	d := GetDeliveryType(deliveryType.ShopId, deliveryType.Type)
	if d != nil {
		return d, nil
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_DELIVERY_TYPE).Create(deliveryType).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("deliverytype:shop", deliveryType.ShopId)
	return GetDeliveryType(deliveryType.ShopId, deliveryType.Type), nil
}

func UpdateDeliveryType(deliveryType *model.DeliveryType) (*model.DeliveryType, error) {
	if deliveryType.Id == 0 {
		return nil, errors.New("未指定出库类型编号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_DELIVERY_TYPE).Where("id = ?", deliveryType.Id).Update("type", deliveryType.Type).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("deliverytype:shop", deliveryType.ShopId)
	return deliveryType, nil
}

func DeleteDeliveryType(deliveryTypeId int) error {
	deliveryType := GetDeliveryTypeById(deliveryTypeId)
	if deliveryTypeId == 0 {
		return errors.New("未指定出库类型编号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_DELIVERY_TYPE).Where("id = ?", deliveryTypeId).Delete(model.DeliveryType{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	cache.DelCache("deliverytype:shop", deliveryType.ShopId)
	return nil
}
