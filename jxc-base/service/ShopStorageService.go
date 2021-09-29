package service

import (
	"ququ.im/common"
	"ququ.im/common/config"
	"ququ.im/jxc-base/dao"
	"ququ.im/jxc-base/model"
	"ququ.im/jxc-base/util"
)

func ListShopStorageByShopId(shopId string) common.Result {
	shopStorageList := dao.ListShopStorageByShopId(shopId)
	if shopStorageList == nil {
		return *common.Error(-1, "未定义仓库")
	}
	return *common.Success(shopStorageList)
}

func GetShopStorage(storageId, shopId, storageName string) common.Result {
	var shopStorage *model.ShopStorage
	if shopId == "" {
		return *common.Error(-1, "缺少商户账号参数")
	}
	if storageId != "" {
		shopStorage = dao.GetShopStorage(shopId, storageId)
	} else if storageName != "" {
		shopStorage = dao.GetShopStorageByName(shopId, storageName)
	} else {
		return *common.Error(-1, "缺乏参数")
	}
	if shopStorage == nil {
		return *common.Error(-1, "未定义此仓库或参数为空")
	}
	return *common.Success(shopStorage)
}

func SaveShopStorage(shopId, storageName, storageId, remark string) common.Result {
	shopStorage := new(model.ShopStorage)
	shopStorage.ShopId = shopId
	shopStorage.Name = storageName
	if storageId == "" {
		shopStorage.StorageId = util.GenerateId("CK", "storage", shopId, 6)
	}
	shopStorage.Remark = remark
	shopStorage, err := dao.SaveShopStorage(shopStorage)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(shopStorage)
}

func UpdateShopStorage(id int, shopId, storageId, storageName, remark string) common.Result {
	var shopStorage *model.ShopStorage
	if id > 0 {
		shopStorage = dao.GetShopStorageById(id)
	} else {
		shopStorage = dao.GetShopStorage(shopId, storageId)
	}
	if shopStorage == nil {
		return *common.Error(-1, "无此仓库ID")
	}
	shopStorage.Name = storageName
	shopStorage.StorageId = storageId
	shopStorage.Remark = remark
	shopStorage, err := dao.UpdateShopStorage(shopStorage)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(shopStorage)
}

func IncrShopStorageSortNumber(shopId, storageId string) common.Result {
	shopStorage := dao.GetShopStorage(shopId, storageId)
	dao.IncrShopStorageSortNumber(shopId, storageId)
	return ListShopStorageByShopId(shopStorage.ShopId)
}

func DecrShopStorageSortNumber(shopId, storageId string) common.Result {
	shopStorage := dao.GetShopStorage(shopId, storageId)
	dao.DecrShopStorageSortNumber(shopId, storageId)
	return ListShopStorageByShopId(shopStorage.ShopId)
}

func DeleteShopStorage(shopId, shopStorageId string) common.Result {
	err := dao.DeleteShopStorage(shopId, shopStorageId)
	if err != nil {
		return *common.Error(-1, "无此仓库记录")
	}
	return *common.Success(nil)
}

func GetStorageLastOperator(shopId, storageId string) common.Result {
	result := make(map[string]string)
	last := config.Redis.HGet("storage:last:"+shopId, storageId).Val()
	if last == "" {
		last = "新仓库"
	}
	result["lastInfo"] = last
	return *common.Success(result)
}
