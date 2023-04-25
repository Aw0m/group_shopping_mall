package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/utils/utils"
)

func AddToCart(ctx *gin.Context, commodityId int64, commodityNum int) (retErr error) {
	//1. 获取用户本的购物车信息
	user, err := GetUserinfo(ctx)
	if err != nil {
		err = errors.WithMessage(err, "get user info failed")
		return err
	}

	//2. 拉取商品基本信息
	commodityList, err := dao.BatchGetCommodityByIdList(ctx, client_db.GetDB(), []int64{commodityId})
	if err != nil {
		err = errors.WithMessage(err, "get commodity from db failed")
		return err
	}
	if len(commodityList) == 0 {
		err = errors.WithMessage(err, "commodity not exist")
		return err
	}
	commodity := commodityList[0]

	//3. 拉取用户购物车信息。由于后续需要对购物车进行更新，因此需要开启事务
	tx := client_db.GetDB().Begin()
	cartList, err := dao.GetCartListByUserId(ctx, tx, user.UserId, false)
	if err != nil {
		tx.Rollback()
		err = errors.WithMessage(err, "get cart from db failed")
		return err
	}

	//4. 如果该商品已经被添加过，则更新数量
	for _, cart := range cartList {
		if cart.CommodityId != commodityId {
			continue
		}

		err = dao.UpdateCart(ctx, tx, cart.CartId, map[string]any{"commodity_num": cart.CommodityNum + commodityNum})
		if err != nil {
			tx.Rollback()
			err = errors.WithMessage(err, "update cart from db failed")
			return err
		}
		return tx.Commit().Error
	}

	//5. 如果该商品未被添加过，则新增一条购物车记录
	cartNew := bdm.Cart{
		CartId:       utils.GetSnowflakeId(),
		CustomerId:   user.UserId,
		CommodityId:  commodityId,
		InitPrice:    commodity.Price,
		IsDeleted:    false,
		CommodityNum: commodityNum,
	}
	err = dao.CreateCart(ctx, tx, cartNew)
	if err != nil {
		tx.Rollback()
		err = errors.WithMessage(err, "create cart from db failed")
		return err
	}
	return tx.Commit().Error
}
