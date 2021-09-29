package service

import (
	"ququ.im/common"
	"ququ.im/jxc-api/nacos"
)

func ListShopStorageByShopId(shopId string) common.Result {
	shopStorageList, err := nacos.ListShopStorage(shopId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if shopStorageList == nil {
		return *common.Error(-1, "未定义仓库")
	}
	return *common.Success(shopStorageList)
}

func GetShopStorage(storageId, shopId, storageName string) common.Result {
	shopStorage, err := nacos.GetShopStorage(shopId, storageId, storageName)
	if shopStorage == nil {
		return *common.Error(-1, "未定义此仓库或参数为空")
	}
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(shopStorage)
}

func SaveShopStorage(shopId, storageName, storageId, remark string) common.Result {
	shopStorage, err := nacos.SaveShopStorage(shopId, storageId, storageName, remark)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(shopStorage)
}

func UpdateShopStorage(id int, shopId, storageId, storageName, remark string) common.Result {
	shopStorage, err := nacos.UpdateShopStorage(id, shopId, storageId, storageName, remark)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(shopStorage)
}

func IncrShopStorageSortNumber(shopId, storageId string) common.Result {
	shopStorage, err := nacos.IncrShopStorageSortNumber(shopId, storageId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(shopStorage)
}

func DecrShopStorageSortNumber(shopId, storageId string) common.Result {
	shopStorage, err := nacos.DecrShopStorageSortNumber(shopId, storageId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(shopStorage)
}

func DeleteShopStorage(shopId, shopStorageId string) common.Result {
	return nacos.DeleteShopStorage(shopId, shopStorageId)
}

func GetStorageLastOperator(shopId, storageId string) common.Result {
	result := nacos.GetStorageLastOperator(shopId, storageId)
	return *common.Success(result)
}
