package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/constant"
	"group_shopping_mall/utils/utils"
)

func CreateOrder(ctx *gin.Context, orderCommodityInfoList []rdm.OrderCommodityInfo) (orderList []bdm.Order, retErr error) {
	tx := client_db.GetDB().Begin()
	defer func() {
		if retErr != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// 1. 根据用户id获取用户信息
	user, err := GetUserinfo(ctx)
	if err != nil {
		retErr = errors.WithMessage(err, "get user info fail")
		return nil, retErr
	}

	// 2. 获取购买的商品信息
	commodityIdList := make([]int64, 0, len(orderCommodityInfoList))
	for _, info := range orderCommodityInfoList {
		commodityIdList = append(commodityIdList, info.CommodityId)
	}

	commodityList, err := dao.BatchGetCommodityByIdList(ctx, tx, commodityIdList)
	if err != nil {
		retErr = errors.WithMessage(err, "BatchGetCommodityByIdList from db err!")
		return nil, retErr
	}

	commodityMap := make(map[int64]bdm.Commodity)
	for _, commodity := range commodityList {
		commodityMap[commodity.CommodityId] = commodity
	}

	// 3. 批量拉取得到的商品的卖家信息
	sellerIdList := make([]int64, 0)
	for _, commodity := range commodityList {
		sellerIdList = append(sellerIdList, commodity.SellerId)
	}

	sellerList, err := dao.BatchGetSellerByIdList(ctx, tx, sellerIdList)
	if err != nil {
		retErr = errors.WithMessage(err, "BatchGetSellerByIdList from db err!")
		return nil, retErr
	}
	sellerMap := make(map[int64]bdm.Seller)
	for _, seller := range sellerList {
		sellerMap[seller.SellerId] = seller
	}

	// 4. 获取address信息
	addressList, err := GetAddressList(ctx)
	if err != nil {
		retErr = errors.WithMessage(err, "get addressList from redis err!")
		return nil, retErr
	}
	addressMap := make(map[int64]bdm.ShoppingAddress)
	for _, address := range addressList {
		addressMap[address.AddressId] = address
	}

	// 5. 为每个商品生成一个订单
	orderList = make([]bdm.Order, 0, len(orderCommodityInfoList))
	for _, info := range orderCommodityInfoList {
		commodity, ok := commodityMap[info.CommodityId]
		if !ok {
			continue
		}

		seller, ok := sellerMap[commodity.SellerId]
		if !ok {
			continue
		}

		address, ok := addressMap[user.AddressId]
		if !ok {
			continue
		}

		order := packOrder(&user, &seller, &address, &commodity, info.Quantity)
		orderList = append(orderList, order)
	}

	// 6. 将订单信息插入到数据库中
	err = dao.InsertOrders(ctx, tx, orderList)
	if err != nil {
		retErr = errors.WithMessage(err, "insert orderList to db err!")
		return nil, retErr
	}
	return orderList, nil
}

func packOrder(user *bdm.User, seller *bdm.Seller, address *bdm.ShoppingAddress, commodity *bdm.Commodity, quantity int) (retOrder bdm.Order) {
	return bdm.Order{
		OrderId:             utils.GetSnowflakeId(),
		SellerId:            seller.SellerId,
		SellerName:          seller.SellerName,
		CommodityId:         commodity.CommodityId,
		CommodityName:       commodity.CommodityName,
		CustomerId:          user.UserId,
		DeliverymanName:     "",
		DeliverymanPhoneNum: "",
		PhoneNum:            user.PhoneNum,
		ConsigneeName:       user.Username,
		AddressId:           address.AddressId,
		AddressDetailInfo:   address.DetailInfo,
		Price:               commodity.Price,
		Quantity:            quantity,
		Status:              constant.OrderStatus_Unconfirmed,
	}
}
