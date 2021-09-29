package mongo

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"ququ.im/common/cache"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
	"ququ.im/jxc-base/util"
)

const COLLECTION_PRODUCT_SKU string = "ProductSku"

func ListProductSku(shopId, keyword, productId, categoryId string, sku map[string]string, status, page, size int) ([]model.ProductSku, error) {
	if shopId == "" {
		return nil, errors.New("缺少商户编号")
	}
	var productSkuList []model.ProductSku
	query := bson.M{}
	query["shopId"] = shopId
	if productId != "" {
		query["productId"] = productId
	}
	if categoryId != "" {
		query["categoryId"] = categoryId
	}
	if status > -1 {
		query["status"] = status
	}
	if len(sku) > 0 {
		for k, v := range sku {
			query["sku."+k] = v
		}
	}
	if keyword != "" {
		nameRegex := bson.M{"skuName": bson.M{"$regex": keyword}}
		pyfRegex := bson.M{"pinyinFull": bson.M{"$regex": keyword}}
		pyjRegex := bson.M{"pinyinFirst": bson.M{"$regex": keyword}}
		//regexs := bson.M{"$or":[]bson.M{nameRegex,pyfRegex,pyjRegex}}
		query["$or"] = []bson.M{nameRegex, pyfRegex, pyjRegex}
	}
	var err error
	if page > 0 && size > 0 {
		err = config.Mgo.C(COLLECTION_PRODUCT_SKU).Find(query).Sort("productId", "sortNumber").Skip(page * size).Limit(size).All(&productSkuList)
	} else {
		err = config.Mgo.C(COLLECTION_PRODUCT_SKU).Find(query).Sort("productId", "sortNumber").All(&productSkuList)
	}
	if err != nil {
		logs.Error("MongoDB组合查询异常:{}", err.Error())
		return nil, err
	}
	return productSkuList, nil
}

func CountProductSku(shopId, keyword, productId, categoryId string, sku map[string]string, status int) (int, error) {
	if shopId == "" {
		return 0, errors.New("缺少商户编号")
	}
	query := bson.M{}
	query["shopId"] = shopId
	if productId != "" {
		query["productId"] = productId
	}
	if categoryId != "" {
		query["categoryId"] = categoryId
	}
	if status > -1 {
		query["status"] = status
	}
	if len(sku) > 0 {
		for k, v := range sku {
			query["sku."+k] = v
		}
	}
	if keyword != "" {
		nameRegex := bson.M{"skuName": bson.M{"$regex": keyword}}
		pyfRegex := bson.M{"pinyinFull": bson.M{"$regex": keyword}}
		pyjRegex := bson.M{"pinyinFirst": bson.M{"$regex": keyword}}
		//regexs := bson.M{"$or":[]bson.M{nameRegex,pyfRegex,pyjRegex}}
		query["$or"] = []bson.M{nameRegex, pyfRegex, pyjRegex}
	}
	count, err := config.Mgo.C(COLLECTION_PRODUCT_SKU).Find(query).Count()
	if err != nil {
		logs.Error("MongoDB组合查询异常:{}", err.Error())
		return 0, err
	}
	return count, nil
}

func GetProductSkuBySkuGuid(skuGuid string) *model.ProductSku {
	var productSku model.ProductSku
	err := cache.GetCache("sku", skuGuid, &productSku)
	if err == nil && productSku.ShopId != "" {
		return &productSku
	}
	err = config.Mgo.C(COLLECTION_PRODUCT_SKU).Find(&bson.M{"skuGuid": skuGuid}).One(&productSku)
	if err != nil {
		logs.Error("MongoDB查询货品库异常:{}", err.Error())
		return nil
	}
	if productSku.SkuGuid == "" {
		return nil
	} else {
		cache.SetCache("sku", skuGuid, productSku)
		return &productSku
	}
}

func GetProductSkuBySkuid(shopId, skuId string) *model.ProductSku {
	var productSku model.ProductSku
	err := cache.GetCache("sku", shopId+":"+skuId, &productSku)
	if err == nil && productSku.ShopId != "" {
		return &productSku
	}
	err = config.Mgo.C(COLLECTION_PRODUCT_SKU).Find(&bson.M{"skuId": skuId, "shopId": shopId}).One(&productSku)
	if err != nil {
		logs.Error("MongoDB查询货品库异常:{}", err.Error())
		return nil
	}
	if productSku.SkuGuid == "" {
		return nil
	} else {
		cache.SetCache("sku", shopId+":"+skuId, productSku)
		return &productSku
	}
}

func SaveProductSku(productSku *model.ProductSku) *model.ProductSku {
	if productSku.Id == bson.ObjectId("") {
		productSku.Id = bson.NewObjectId()
	}
	if productSku.SkuGuid == "" {
		productSku.SkuGuid = utils.GetRandomHexString(32)
	}
	if productSku.SkuId == "" {
		productSku.SkuId = util.GenerateId("SP", "sku", productSku.ShopId, 8)
	}
	if productSku.SkuName == "" {
		productSku.SkuName = productSku.Name
		for _, v := range productSku.Sku {
			productSku.SkuName = productSku.SkuName + " " + v
		}
	}
	productSku.PyFull = utils.ToPinYin(productSku.Name, true, false)
	productSku.PyFirst = utils.ToPinYin(productSku.Name, false, false)
	productSku.Status = 1
	err := config.Mgo.C(COLLECTION_PRODUCT_SKU).Insert(productSku)
	if err != nil {
		logs.Error("MongoDB插入货品库异常:{}", err.Error())
		return nil
	}
	return productSku
}

func UpdateProductSku(productSku *model.ProductSku) *model.ProductSku {
	productSku.PyFull = utils.ToPinYin(productSku.Name, true, false)
	productSku.PyFirst = utils.ToPinYin(productSku.Name, false, false)
	err := config.Mgo.C(COLLECTION_PRODUCT_SKU).UpdateId(productSku.Id, productSku)
	if err != nil {
		logs.Error("MongoDB更新货品库异常:{}", err.Error())
		return nil
	}
	cache.DelCache("sku", productSku.SkuGuid)
	cache.DelCache("sku", productSku.ShopId+":"+productSku.SkuId)
	return productSku
}

func DeleteProductSku(skuGuid string) error {
	productSku := GetProductSkuBySkuGuid(skuGuid)
	err := config.Mgo.C(COLLECTION_PRODUCT_SKU).Remove(&bson.M{"skuGuid": skuGuid})
	if err != nil {
		logs.Error("MongoDB删除货品库异常:{}", err.Error())
		return err
	}
	cache.DelCache("sku", productSku.SkuGuid)
	cache.DelCache("sku", productSku.ShopId+":"+productSku.SkuId)
	return nil
}

func IncrProductSkuSortNumber(skuGuid string) []model.ProductSku {
	productSku := GetProductSkuBySkuGuid(skuGuid)
	if productSku == nil {
		return nil
	}
	productSkuList, _ := ListProductSku(productSku.ShopId, "", productSku.ProductId, "", nil, -1, 0, 0)
	for i, pSku := range productSkuList {
		if pSku.SkuGuid == skuGuid && len(productSkuList) > i+1 {
			nextSku := productSkuList[i+1]
			productSku.SortNumber, nextSku.SortNumber = nextSku.SortNumber, productSku.SortNumber
			UpdateProductSku(productSku)
			UpdateProductSku(&nextSku)
			break
		}
	}
	productSkuList, _ = ListProductSku(productSku.ShopId, "", productSku.ProductId, "", nil, -1, 0, 0)
	return productSkuList
}

func DecrProductSkuSortNumber(skuGuid string) []model.ProductSku {
	productSku := GetProductSkuBySkuGuid(skuGuid)
	if productSku == nil {
		return nil
	}
	productSkuList, _ := ListProductSku(productSku.ShopId, "", productSku.ProductId, "", nil, -1, 0, 0)
	for i, pSku := range productSkuList {
		if pSku.SkuGuid == skuGuid && i >= 1 {
			beforeSku := productSkuList[i-1]
			productSku.SortNumber, beforeSku.SortNumber = beforeSku.SortNumber, productSku.SortNumber
			UpdateProductSku(productSku)
			UpdateProductSku(&beforeSku)
			break
		}
	}
	productSkuList, _ = ListProductSku(productSku.ShopId, "", productSku.ProductId, "", nil, -1, 0, 0)
	return productSkuList
}
