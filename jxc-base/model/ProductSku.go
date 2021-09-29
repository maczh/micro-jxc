package model

import "gopkg.in/mgo.v2/bson"

type ProductSku struct {
	Id         bson.ObjectId     `bson:"_id"`
	ShopId     string            `json:"shopId" bson:"shopId"`
	ProductId  string            `json:"productId" bson:"productId"`
	SkuGuid    string            `json:"skuGuid" bson:"skuGuid"`
	SkuId      string            `json:"skuId" bson:"skuId"`
	Name       string            `json:"name" bson:"name"`
	SkuName    string            `json:"skuName" bson:"skuName"`
	CategoryId string            `json:"categoryId" bson:"categoryId"`
	BaseUnit   string            `json:"baseUnit" bson:"baseUnit"`
	BarCode    string            `json:"barCode" bson:"barCode"`
	Sku        map[string]string `json:"sku" bson:"sku"`
	PriceList  map[string]string `json:"priceList" bson:"priceList"`
	SortNumber int               `json:"sortNumber" bson:"sortNumber"`
	Status     int               `json:"status" bson:"status"`
	PyFull     string            `json:"pinyinFull" bson:"pinyinFull"`
	PyFirst    string            `json:"pinyinFirst" bson:"pinyinFirst"`
}
