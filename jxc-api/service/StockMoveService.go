package service

import (
	"ququ.im/common"
	"ququ.im/common/logs"
	"ququ.im/common/utils"
	"ququ.im/jxc-api/nacos"
	"ququ.im/jxc-api/vo"
	"ququ.im/jxc-base/model"
	"time"
)

func ListStockMove(shopId, skuId, moveNo, fromStorageId, toStorageId, startTime, endTime string, page, size int) common.Result {
	var stockMoveList []model.StockMove
	var resultPage *common.ResultPage
	var stockMoveRowDetailList []vo.StockMoveRowDetail
	var err error
	if page > 0 && size > 0 {
		stockMoveList, resultPage, err = nacos.ListStockMove(shopId, skuId, moveNo, fromStorageId, toStorageId, startTime, endTime, page, size)
	} else {
		stockMoveList, _, err = nacos.ListStockMove(shopId, skuId, moveNo, fromStorageId, toStorageId, startTime, endTime, 0, 0)
	}
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if stockMoveList == nil || len(stockMoveList) == 0 {
		return *common.Error(-1, "查无数据")
	}
	for _, stockMove := range stockMoveList {
		stockMoveRowDetail := new(vo.StockMoveRowDetail)
		stockMoveRowDetail.Id = stockMove.Id
		stockMoveRowDetail.ShopId = stockMove.ShopId
		stockMoveRowDetail.MoveNo = stockMove.MoveNo
		stockMoveRowDetail.SkuId = stockMove.SkuId
		stockMoveRowDetail.Number = stockMove.Number
		stockMoveRowDetail.Unit = stockMove.Unit
		stockMoveRowDetail.FromStorageId = stockMove.FromStorageId
		stockMoveRowDetail.ToStorageId = stockMove.ToStorageId
		stockMoveRowDetail.MoveTime = utils.GormTimeFormat(stockMove.MoveTime)
		stockMoveRowDetail.Operator = stockMove.Operator
		stockMoveRowDetail.Remark = stockMove.Remark
		productSku, _ := nacos.GetProductSku(stockMove.ShopId, stockMove.SkuId, "")
		stockMoveRowDetail.SkuName = productSku.SkuName
		stockMoveRowDetail.Sku = productSku.Sku
		stockMoveRowDetail.ProductId = productSku.ProductId
		stockMoveRowDetail.PriceList = productSku.PriceList
		stockMoveRowDetail.SkuGuid = productSku.SkuGuid
		stockMoveRowDetail.BarCode = productSku.BarCode
		fromStorage, _ := nacos.GetShopStorage(stockMove.ShopId, stockMove.FromStorageId, "")
		stockMoveRowDetail.FromStorageName = fromStorage.Name
		toStorage, _ := nacos.GetShopStorage(stockMove.ShopId, stockMove.ToStorageId, "")
		stockMoveRowDetail.ToStorageName = toStorage.Name
		stockMoveRowDetailList = append(stockMoveRowDetailList, *stockMoveRowDetail)
	}
	if page > 0 && size > 0 {
		return *common.SuccessWithPage(stockMoveRowDetailList, resultPage.Count, page, size, resultPage.Total)
	} else {
		return *common.Success(stockMoveRowDetailList)
	}
}

func GetStockMoveRow(shopId, skuId, moveNo string) common.Result {
	stockMove, err := nacos.GetStockMove(0, shopId, skuId, moveNo)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if stockMove == nil {
		return *common.Error(-1, "查无数据")
	}
	stockMoveRowDetail := new(vo.StockMoveRowDetail)
	stockMoveRowDetail.Id = stockMove.Id
	stockMoveRowDetail.ShopId = stockMove.ShopId
	stockMoveRowDetail.MoveNo = stockMove.MoveNo
	stockMoveRowDetail.SkuId = stockMove.SkuId
	stockMoveRowDetail.Number = stockMove.Number
	stockMoveRowDetail.Unit = stockMove.Unit
	stockMoveRowDetail.FromStorageId = stockMove.FromStorageId
	stockMoveRowDetail.ToStorageId = stockMove.ToStorageId
	stockMoveRowDetail.MoveTime = utils.GormTimeFormat(stockMove.MoveTime)
	stockMoveRowDetail.Operator = stockMove.Operator
	stockMoveRowDetail.Remark = stockMove.Remark
	productSku, _ := nacos.GetProductSku(stockMove.ShopId, stockMove.SkuId, "")
	stockMoveRowDetail.SkuName = productSku.SkuName
	stockMoveRowDetail.Sku = productSku.Sku
	stockMoveRowDetail.ProductId = productSku.ProductId
	stockMoveRowDetail.PriceList = productSku.PriceList
	stockMoveRowDetail.SkuGuid = productSku.SkuGuid
	stockMoveRowDetail.BarCode = productSku.BarCode
	fromStorage, _ := nacos.GetShopStorage(stockMove.ShopId, stockMove.FromStorageId, "")
	stockMoveRowDetail.FromStorageName = fromStorage.Name
	toStorage, _ := nacos.GetShopStorage(stockMove.ShopId, stockMove.ToStorageId, "")
	stockMoveRowDetail.ToStorageName = toStorage.Name
	return *common.Success(stockMoveRowDetail)
}

func MoveSkuStock(shopId, skuId, moveNo, unit, fromStorageId, toStorageId, operator, remark string, number int) common.Result {
	if shopId == "" {
		return *common.Error(-1, "商户账号不可为空")
	}
	if skuId == "" {
		return *common.Error(-1, "要移库的货品编码不可为空")
	}
	if moveNo == "" {
		moveNo = "MOVE" + shopId + time.Now().Format("20060102150405")
	}
	if unit == "" {
		return *common.Error(-1, "商品单位不可为空")
	}
	if fromStorageId == "" {
		return *common.Error(-1, "原仓库编码不可为空")
	}
	if toStorageId == "" {
		return *common.Error(-1, "目标仓库编码不可为空")
	}
	if operator == "" {
		return *common.Error(-1, "移库操作员名称不可为空")
	}
	if number == 0 {
		return *common.Error(-1, "要移库的货品数量不可为0")
	}
	productSku, _ := nacos.GetProductSku(shopId, skuId, "")
	if productSku == nil {
		return *common.Error(-1, "货品编码错误或不存在")
	}
	if unit != productSku.BaseUnit {
		productUnit, _ := nacos.GetProductUnit(0, shopId, productSku.ProductId, unit, productSku.BaseUnit)
		if productUnit == nil {
			return *common.Error(-1, "要移动的货品单位"+unit+"与货品基本单位"+productSku.BaseUnit+"之间的换算规则不存在")
		}
	}
	fromStorage, _ := nacos.GetShopStorage(shopId, fromStorageId, "")
	if fromStorage == nil {
		return *common.Error(-1, "原仓库编码不存在")
	}
	toStorage, _ := nacos.GetShopStorage(shopId, toStorageId, "")
	if toStorage == nil {
		return *common.Error(-1, "目标仓库编码不存在")
	}
	skuStock, _ := nacos.GetSkuStock(shopId, skuId, fromStorageId)
	if skuStock == nil {
		return *common.Error(-1, "原仓库中货品"+productSku.SkuName+"无库存，无法完成移库")
	}
	//从原库中出库
	result := nacos.InOutSkuStock(shopId, skuId, unit, fromStorageId, -number, 0, 0)
	if result.Status != 1 {
		logs.Error("移库过程出库失败:{}", result)
		return result
	}
	//在目标库中以原价格入库
	result = nacos.InOutSkuStock(shopId, skuId, unit, toStorageId, number, skuStock.CostPrice, skuStock.LastPrice)
	if result.Status != 1 {
		logs.Error("移库过程入库失败:{}", result)
		return result
	}
	//保存移库单
	return nacos.SaveStockMove(shopId, skuId, moveNo, unit, fromStorageId, toStorageId, operator, remark, number)
}

func DeleteStockMove(id int) common.Result {
	stockMove, _ := nacos.GetStockMove(id, "", "", "")
	if stockMove == nil {
		return *common.Error(-1, "移库单id不存在或已删除")
	}
	skuStock, _ := nacos.GetSkuStock(stockMove.ShopId, stockMove.SkuId, stockMove.ToStorageId)
	if skuStock == nil {
		return *common.Error(-1, "目标仓库中此货品库存记录不存在，无法完成移库单删除")
	}
	//从目标库中出库
	result := nacos.InOutSkuStock(stockMove.ShopId, stockMove.SkuId, stockMove.Unit, stockMove.ToStorageId, -stockMove.Number, 0, 0)
	if result.Status != 1 {
		logs.Error("从目标库出库失败:{}", result)
		return result
	}
	//在原库库中以原价格入库
	result = nacos.InOutSkuStock(stockMove.ShopId, stockMove.SkuId, stockMove.Unit, stockMove.FromStorageId, stockMove.Number, skuStock.CostPrice, skuStock.LastPrice)
	if result.Status != 1 {
		logs.Error("在原库中入库失败:{}", result)
		return result
	}
	//保存移库单
	return nacos.DeleteStockMove(id)
}
