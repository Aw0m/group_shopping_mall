package biz

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/middleware"
)

func GetCartList(ctx *gin.Context) (retItemList []rdm.CartInfoItem, retErr error) {
	// 1. 拉取指定用户的cart列表
	_, userid, err := middleware.GetID(ctx)
	if err != nil {
		err = errors.WithMessage(err, "GetID from token err!")
	}

	cartList, err := dao.GetCartListByUserId(ctx, client_db.GetDB(), userid, false)
	if err != nil {
		err = errors.WithMessage(err, "GetCartListByUserId from db err!")
		return nil, err
	}

	commodityMap := make(map[int64]bdm.Commodity)
	sellerMap := make(map[int64]bdm.Seller)
	// 2. 遍历cart列表，获取对应的commodity信息
	commodityIdList := make([]int64, 0, len(cartList))
	for _, cart := range cartList {
		commodityIdList = append(commodityIdList, cart.CommodityId)
	}
	commodityList, err := batchGetObjByIds(ctx, client_db.GetDB(), commodityIdList,
		dao.BatchGetCommodityByIdList,
		func(commodity *bdm.Commodity) int64 {
			return commodity.CommodityId
		})
	if err != nil {
		err = errors.WithMessage(err, "BatchGetCommodityByIdList from db err!")
		return nil, err
	}
	for _, commodity := range commodityList {
		commodityMap[commodity.CommodityId] = commodity
	}

	// 3. 获得所有的Seller信息
	sellerIdList := make([]int64, 0, len(commodityList))
	for _, commodity := range commodityList {
		sellerIdList = append(sellerIdList, commodity.SellerId)
	}
	sellerList, err := batchGetObjByIds(ctx, client_db.GetDB(), sellerIdList,
		dao.BatchGetSellerByIdList,
		func(seller *bdm.Seller) int64 {
			return seller.SellerId
		})
	if err != nil {
		err = errors.WithMessage(err, "BatchGetSellerByIdList from db err!")
		return nil, err
	}
	for _, seller := range sellerList {
		sellerMap[seller.SellerId] = seller
	}

	// 4.组装成CartInfoItem
	cartInfoItemMap := make(map[int64]rdm.CartInfoItem) // key: sellerId, value: cartInfoItem
	for _, cart := range cartList {
		// 不存在或者已经下架的commodity应该被过滤掉
		commodity, ok := commodityMap[cart.CommodityId]
		if !ok || commodity.IsDeleted {
			continue
		}

		// 不存在或者已经下架的seller应该被过滤掉
		seller, ok := sellerMap[commodity.SellerId]
		if !ok || seller.IsDeleted {
			continue
		}

		// item以seller维度进行分组。需要先判断该seller是否存在，不存在则先初始化一个item
		item, ok := cartInfoItemMap[commodity.SellerId]
		if !ok {
			item = rdm.CartInfoItem{
				SellerID:      seller.SellerId,
				SellerName:    seller.SellerName,
				CommodityList: make([]rdm.CartCommodityItem, 0, 1), // 此时至少存在1个item，因此初始化cap为1
			}
		}
		item.CommodityList = append(item.CommodityList, rdm.CartCommodityItem{
			Commodity: commodity,
			Cart:      cart,
		})
		cartInfoItemMap[commodity.SellerId] = item
	}

	for _, item := range cartInfoItemMap {
		retItemList = append(retItemList, item)
	}
	return retItemList, nil
}

func batchGetObjByIds[T any](ctx *gin.Context, db *gorm.DB, idList []int64, daoGetObjByIdList func(context.Context, *gorm.DB, []int64) ([]T, error), getObjId func(*T) int64) (objMap map[int64]T, err error) {
	objMap = make(map[int64]T)
	objList, err := daoGetObjByIdList(ctx, db, idList)
	if err != nil {
		return nil, errors.WithMessage(err, "find obj by ids fail")
	}
	for _, obj := range objList {
		objMap[getObjId(&obj)] = obj
	}
	return objMap, nil
}
