package dao

import (
	"errors"
	"ququ.im/common/cache"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/jxc-base/model"
)

const TABLE_SHOP_STORAGE string = "shop_storage"

func ListShopStorageByShopId(shopId string) []model.ShopStorage {
	var shopStorageList []model.ShopStorage
	err := cache.GetCache("storage:shop", shopId, &shopStorageList)
	if err == nil && len(shopStorageList) > 0 {
		return shopStorageList
	}
	config.Mysql.Table(TABLE_SHOP_STORAGE).Where("shop_id = ?", shopId).Order("sort_number ASC").Find(&shopStorageList)
	if len(shopStorageList) == 0 {
		return nil
	} else {
		cache.SetCache("storage:shop", shopId, shopStorageList)
		return shopStorageList
	}
}

func GetShopStorageById(id int) *model.ShopStorage {
	var shopStorage model.ShopStorage
	config.Mysql.Table(TABLE_SHOP_STORAGE).Where("id = ? ", id).First(&shopStorage)
	if shopStorage.Id == 0 {
		return nil
	} else {
		return &shopStorage
	}
}

func GetShopStorage(shopId, storageId string) *model.ShopStorage {
	var shopStorage model.ShopStorage
	err := cache.GetCache("storage:id"+":"+shopId, storageId, &shopStorage)
	if err == nil && shopStorage.ShopId != "" {
		return &shopStorage
	}
	config.Mysql.Table(TABLE_SHOP_STORAGE).Where("shop_id = ? AND storage_id = ?", shopId, storageId).First(&shopStorage)
	if shopStorage.Id == 0 {
		return nil
	} else {
		cache.SetCache("storage:id"+":"+shopId, storageId, shopStorage)
		return &shopStorage
	}
}

func GetShopStorageByName(shopId, name string) *model.ShopStorage {
	var shopStorage model.ShopStorage
	config.Mysql.Table(TABLE_SHOP_STORAGE).Where("shop_id = ? AND name = ? ", shopId, name).First(&shopStorage)
	if shopStorage.Id == 0 {
		return nil
	} else {
		return &shopStorage
	}
}

func SaveShopStorage(shopStorage *model.ShopStorage) (*model.ShopStorage, error) {
	s := GetShopStorageByName(shopStorage.ShopId, shopStorage.Name)
	if s != nil {
		return s, nil
	}
	if shopStorage.SortNumber == 0 {
		shopStorageList := ListShopStorageByShopId(shopStorage.ShopId)
		if shopStorageList == nil {
			shopStorage.SortNumber = 101
		} else {
			shopStorage.SortNumber = shopStorageList[len(shopStorageList)-1].SortNumber + 1
		}
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_SHOP_STORAGE).Create(shopStorage).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("storage:shop", shopStorage.ShopId)
	return shopStorage, nil
}

func UpdateShopStorage(shopStorage *model.ShopStorage) (*model.ShopStorage, error) {
	if shopStorage.Id == 0 {
		return nil, errors.New("未指定Id")
	}
	data := make(map[string]interface{})
	if shopStorage.ShopId != "" {
		data["shop_id"] = shopStorage.ShopId
	}
	if shopStorage.Name != "" {
		data["name"] = shopStorage.Name
	}
	if shopStorage.SortNumber > 0 {
		data["sort_number"] = shopStorage.SortNumber
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_SHOP_STORAGE).Where("id = ?", shopStorage.Id).Update(data).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("storage:shop", shopStorage.ShopId)
	cache.DelCache("storage:id"+":"+shopStorage.ShopId, shopStorage.StorageId)
	return shopStorage, nil
}

//排序向后移一位
func IncrShopStorageSortNumber(shopId, storageId string) {
	shopStorage := GetShopStorage(shopId, storageId)
	if shopStorage == nil {
		return
	}
	var shopStorageNext model.ShopStorage
	config.Mysql.Table(TABLE_SHOP_STORAGE).Where("shop_id = ? AND sort_number > ?", shopStorage.ShopId, shopStorage.SortNumber).Order("sort_number ASC").First(&shopStorageNext)
	if shopStorageNext.Id == 0 {
		return
	}
	shopStorage.SortNumber, shopStorageNext.SortNumber = shopStorageNext.SortNumber, shopStorage.SortNumber
	UpdateShopStorage(shopStorage)
	UpdateShopStorage(&shopStorageNext)
}

//排序向前移一位
func DecrShopStorageSortNumber(shopId, storageId string) {
	shopStorage := GetShopStorage(shopId, storageId)
	if shopStorage == nil {
		return
	}
	var shopStorageBefore model.ShopStorage
	config.Mysql.Table(TABLE_SHOP_STORAGE).Where("shop_id = ? AND sort_number < ?", shopStorage.ShopId, shopStorage.SortNumber).Order("sort_number DESC").First(&shopStorageBefore)
	if shopStorageBefore.Id == 0 {
		return
	}
	shopStorage.SortNumber, shopStorageBefore.SortNumber = shopStorageBefore.SortNumber, shopStorage.SortNumber
	UpdateShopStorage(shopStorage)
	UpdateShopStorage(&shopStorageBefore)
}

func DeleteShopStorage(shopId, shopStorageId string) error {
	if shopStorageId == "" {
		return errors.New("未指定仓库编号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_SHOP_STORAGE).Where("shop_id = ? AND storage_id = ?", shopId, shopStorageId).Delete(model.ShopStorage{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	cache.DelCache("storage:shop", shopId)
	cache.DelCache("storage:id"+":"+shopId, shopStorageId)
	return nil
}
