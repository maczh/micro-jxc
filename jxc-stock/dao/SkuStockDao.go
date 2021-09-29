package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
	"time"
)

const TABLE_SKU_STOCK string = "sku_stock"

func ListSkuStock(skuId, shopId, name, productId, storageId string, page, size int) []model.SkuStock {
	var skuStockList []model.SkuStock
	if page > 0 && size > 0 {
		config.Mysql.Table(TABLE_SKU_STOCK).Where("IF(? != '',sku_id = ?,1=1) AND IF(? != '',shop_id = ?,1=1) AND IF(? != '',product_id = ?,1=1) AND IF(? != '', storage_id = ?,1=1) AND IF(? != '',name like ?,1=1)", skuId, skuId, shopId, shopId, productId, productId, storageId, storageId, name, "%"+name+"%").Order("update_time DESC").Limit(size).Offset((page - 1) * size).Find(&skuStockList)
	} else {
		config.Mysql.Table(TABLE_SKU_STOCK).Where("IF(? != '',sku_id = ?,1=1) AND IF(? != '',shop_id = ?,1=1) AND IF(? != '',product_id = ?,1=1) AND IF(? != '', storage_id = ?,1=1) AND IF(? != '',name like ?,1=1)", skuId, skuId, shopId, shopId, productId, productId, storageId, storageId, name, "%"+name+"%").Order("update_time DESC").Find(&skuStockList)
	}
	if len(skuStockList) == 0 {
		return nil
	} else {
		for i := 0; i < len(skuStockList); i++ {
			skuStockList[i].UpdateTime = utils.GormTimeFormat(skuStockList[i].UpdateTime)
		}
		return skuStockList
	}
}

func CountSkuStock(skuId, shopId, name, productId, storageId string) int {
	count := 0
	config.Mysql.Table(TABLE_SKU_STOCK).Where("IF(? != '',sku_id = ?,1=1) AND IF(? != '',shop_id = ?,1=1) AND IF(? != '',product_id = ?,1=1) AND IF(? != '', storage_id = ?,1=1) AND IF(? != '',name like ?,1=1)", skuId, skuId, shopId, shopId, productId, productId, storageId, storageId, name, "%"+name+"%").Count(&count)
	return count
}

func GetSkuStock(shopId, skuId, storageId string) *model.SkuStock {
	var skuStock model.SkuStock
	config.Mysql.Table(TABLE_SKU_STOCK).Where("shop_id = ? AND sku_id = ? AND storage_id = ?", shopId, skuId, storageId).First(&skuStock)
	if skuStock.Id == 0 {
		return nil
	} else {
		skuStock.UpdateTime = utils.GormTimeFormat(skuStock.UpdateTime)
		return &skuStock
	}
}

func GetSkuStockNumber(shopId, skuId, storageId string) int {
	if storageId != "" {
		skuStock := GetSkuStock(shopId, skuId, storageId)
		return skuStock.Stocks
	}
	var skuStock []int
	config.Mysql.Table(TABLE_SKU_STOCK).Where("shop_id = ? AND sku_id = ?", shopId, skuId).Pluck("SUM(stocks) as skuStock", &skuStock)
	return skuStock[0]
}

func GetProductStockNumber(shopId, productId, storageId string) int {
	var skuStock []int
	config.Mysql.Table(TABLE_SKU_STOCK).Where("shop_id = ? AND product_id = ? AND IF(? != '', storage_id = ?, 1=1)", shopId, productId, storageId, storageId).Pluck("SUM(stocks) as skuStock", &skuStock)
	return skuStock[0]
}

func SaveSkuStock(skuStock *model.SkuStock) (*model.SkuStock, error) {
	s := GetSkuStock(skuStock.ShopId, skuStock.SkuId, skuStock.StorageId)
	if s != nil {
		return s, nil
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_SKU_STOCK).Create(skuStock).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return skuStock, nil
}

func UpdateSkuStock(skuStock *model.SkuStock) (*model.SkuStock, error) {
	if skuStock.Id == 0 {
		return nil, errors.New("未指定Id")
	}
	data := make(map[string]interface{})
	if skuStock.SkuId != "" {
		data["sku_id"] = skuStock.SkuId
	}
	if skuStock.ProductId != "" {
		data["product_id"] = skuStock.ProductId
	}
	if skuStock.ShopId != "" {
		data["shop_id"] = skuStock.ShopId
	}
	if skuStock.Name != "" {
		data["name"] = skuStock.Name
	}
	if skuStock.Stocks > 0 {
		data["stocks"] = skuStock.Stocks
	}
	if skuStock.BaseUnit != "" {
		data["base_unit"] = skuStock.BaseUnit
	}
	if skuStock.StorageId != "" {
		data["storage_id"] = skuStock.StorageId
	}
	if skuStock.CostPrice > 0 {
		data["cost_price"] = skuStock.CostPrice
	}
	if skuStock.LastPrice > 0 {
		data["last_price"] = skuStock.LastPrice
	}
	data["update_time"] = utils.ToDateTimeString(time.Now())
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_SKU_STOCK).Where("id = ?", skuStock.Id).Update(data).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return GetSkuStock(skuStock.ShopId, skuStock.SkuId, skuStock.StorageId), nil
}

func AddSkuStock(productSku *model.ProductSku, number, costPrice, lastPrice int, storageId string) (*model.SkuStock, error) {
	//计算库存总金额
	amount := 0
	var amounts []int
	config.Mysql.Table(TABLE_SKU_STOCK).Where("shop_id = ? AND sku_id = ?", productSku.ShopId, productSku.SkuId).Pluck("SUM(stocks * cost_price) as amount", &amounts)
	amount = amounts[0] + number*lastPrice
	//计算库存总量
	total := GetSkuStockNumber(productSku.ShopId, productSku.SkuId, "") + number
	cost := amount / total
	if costPrice > 0 {
		cost = costPrice
	}
	skuStock := GetSkuStock(productSku.ShopId, productSku.SkuId, storageId)
	if skuStock == nil {
		skuStock = new(model.SkuStock)
		skuStock.SkuId = productSku.SkuId
		skuStock.ProductId = productSku.ProductId
		skuStock.ShopId = productSku.ShopId
		skuStock.Name = productSku.SkuName
		skuStock.Stocks = number
		skuStock.BaseUnit = productSku.BaseUnit
		skuStock.StorageId = storageId
		skuStock.CostPrice = cost
		skuStock.LastPrice = lastPrice
		skuStock.UpdateTime = utils.ToDateTimeString(time.Now())
		SaveSkuStock(skuStock)
	} else {
		tx := config.Mysql.Begin()
		err := tx.Table(TABLE_SKU_STOCK).Where("id = ?", skuStock.Id).Update("stocks", gorm.Expr("stocks + ?", number)).Error
		if err != nil {
			logs.Error("更新数据错误:{}", err.Error())
			tx.Rollback()
			return nil, err
		}
		tx.Commit()
	}
	data := make(map[string]interface{})
	data["cost_price"] = cost
	data["last_price"] = lastPrice
	data["update_time"] = utils.ToDateTimeString(time.Now())
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_SKU_STOCK).Where("shop_id = ? AND sku_id = ?", productSku.ShopId, productSku.SkuId).Update(data).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return GetSkuStock(productSku.ShopId, productSku.SkuId, storageId), nil
}

func SubSkuStock(productSku *model.ProductSku, number, price int, storageId string) (*model.SkuStock, error) {
	skuStock := GetSkuStock(productSku.ShopId, productSku.SkuId, storageId)
	if skuStock == nil {
		return nil, errors.New("本仓库无此库存货品")
	} else {
		//如果有价格，重新计算成本
		if price > 0 {
			total := GetSkuStockNumber(productSku.ShopId, productSku.SkuId, "")
			amount := skuStock.CostPrice * total
			costnew := (amount - price*number) / (total - number)
			tx := config.Mysql.Begin()
			err := tx.Table(TABLE_SKU_STOCK).Where("shop_id = ? AND sku_id = ?", productSku.ShopId, productSku.SkuId).Update("cost_price", costnew).Error
			if err != nil {
				logs.Error("更新数据错误:{}", err.Error())
				tx.Rollback()
				return nil, err
			}
			tx.Commit()
		}
		//减库存
		tx := config.Mysql.Begin()
		err := tx.Table(TABLE_SKU_STOCK).Where("id = ?", skuStock.Id).Update(map[string]interface{}{"stocks": gorm.Expr("stocks - ?", number), "update_time": utils.ToDateTimeString(time.Now())}).Error
		if err != nil {
			logs.Error("更新数据错误:{}", err.Error())
			tx.Rollback()
			return nil, err
		}
		tx.Commit()
	}
	return GetSkuStock(productSku.ShopId, productSku.SkuId, storageId), nil
}
