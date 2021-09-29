package dao

import (
	"errors"
	"ququ.im/common/cache"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/jxc-base/model"
)

const TABLE_ENTRY_TYPE string = "entry_type"

func ListEntryTypeByShopId(shopId string) []model.EntryType {
	var entryTypeList []model.EntryType
	err := cache.GetCache("entrytype:shop", shopId, &entryTypeList)
	if err == nil && len(entryTypeList) > 0 {
		return entryTypeList
	}
	config.Mysql.Table(TABLE_ENTRY_TYPE).Where("shop_id = ? OR shop_id = '' OR shop_id is null", shopId).Find(&entryTypeList)
	if len(entryTypeList) == 0 {
		return nil
	} else {
		cache.SetCache("entrytype:shop", shopId, entryTypeList)
		return entryTypeList
	}
}

func GetEntryType(shopId, eType string) *model.EntryType {
	var entryType model.EntryType
	config.Mysql.Table(TABLE_ENTRY_TYPE).Where("IF(? != '',shop_id = ?,1=1) AND type = ?", shopId, shopId, eType).First(&entryType)
	if entryType.Id == 0 {
		return nil
	} else {
		return &entryType
	}
}

func GetEntryTypeById(entryTypeId int) *model.EntryType {
	var entryType model.EntryType
	config.Mysql.Table(TABLE_ENTRY_TYPE).Where("id = ?", entryTypeId).First(&entryType)
	if entryType.Id == 0 {
		return nil
	} else {
		return &entryType
	}
}

func SaveEntryType(entryType *model.EntryType) (*model.EntryType, error) {
	d := GetEntryType(entryType.ShopId, entryType.Type)
	if d != nil {
		return d, nil
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_ENTRY_TYPE).Create(entryType).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("entrytype:shop", entryType.ShopId)
	return GetEntryType(entryType.ShopId, entryType.Type), nil
}

func UpdateEntryType(entryType *model.EntryType) (*model.EntryType, error) {
	if entryType.Id == 0 {
		return nil, errors.New("未指定入库类型编号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_ENTRY_TYPE).Where("id = ?", entryType.Id).Update("type", entryType.Type).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("entrytype:shop", entryType.ShopId)
	return entryType, nil
}

func DeleteEntryType(entryTypeId int) error {
	entryType := GetEntryTypeById(entryTypeId)
	if entryTypeId == 0 {
		return errors.New("未指定入库类型编号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_ENTRY_TYPE).Where("id = ?", entryTypeId).Delete(model.EntryType{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	cache.DelCache("entrytype:shop", entryType.ShopId)
	return nil
}
