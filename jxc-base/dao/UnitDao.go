package dao

import (
	"errors"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/jxc-base/model"
)

const TABLE_UNIT string = "units"

func ListUnitByShopId(shopId string) []model.Unit {
	var unitList []model.Unit
	config.Mysql.Table(TABLE_UNIT).Where("shop_id = ? OR shop_id = '' OR shop_id is null", shopId).Find(&unitList)
	if len(unitList) == 0 {
		return nil
	} else {
		return unitList
	}
}

func GetUnit(shopId, u string) *model.Unit {
	var unit model.Unit
	config.Mysql.Table(TABLE_UNIT).Where("IF(? != '',shop_id = ?,1=1) AND unit = ?", shopId, shopId, u).First(&unit)
	if unit.Id == 0 {
		return nil
	} else {
		return &unit
	}
}

func GetUnitById(id int) *model.Unit {
	var unit model.Unit
	config.Mysql.Table(TABLE_UNIT).Where("id = ?", id).First(&unit)
	if unit.Id == 0 {
		return nil
	} else {
		return &unit
	}
}

func SaveUnit(unit *model.Unit) (*model.Unit, error) {
	d := GetUnit(unit.ShopId, unit.Unit)
	if d != nil {
		return d, nil
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_UNIT).Create(unit).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return GetUnit(unit.ShopId, unit.Unit), nil
}

func UpdateUnit(unit *model.Unit) (*model.Unit, error) {
	if unit.Id == 0 {
		return nil, errors.New("未指定单位编号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_UNIT).Where("id = ?", unit.Id).Update("unit", unit.Unit).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return unit, nil
}

func DeleteUnit(unitId int) error {
	if unitId == 0 {
		return errors.New("未指定单位编号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_UNIT).Where("id = ?", unitId).Delete(model.Unit{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
