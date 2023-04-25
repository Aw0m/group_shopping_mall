package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/rdm"
)

func GetCommodityDetail(ctx *gin.Context, commodityId int64) (commodityInfo rdm.CommodityInfo, sellerInfo rdm.SellerInfo, commentInfo rdm.CommentInfo, err error) {
	// 1. 拉取commodity并校验，保证commodity存在且未被删除
	commodityList, err := dao.BatchGetCommodityByIdList(ctx, client_db.GetDB(), []int64{commodityId})
	if err != nil {
		err = errors.Errorf("get commodity from db err! err:%s", err.Error())
		return
	}
	if len(commodityList) == 0 {
		err = errors.Errorf("commodity not exist! commodityId:%d", commodityId)
		return
	}
	commodity := commodityList[0]
	if commodity.IsDeleted {
		err = errors.Errorf("commodity is deleted! commodityId:%d", commodityId)
		return
	}
	commodityInfo = rdm.CommodityInfo{
		CommodityId:   commodityId,
		CommodityName: commodity.CommodityName,
		Price:         commodity.Price,
		ImageURL:      commodity.ImageURL,
	}

	// 2. 拉取commodity的seller
	sellerList, err := dao.BatchGetSellerByIdList(ctx, client_db.GetDB(), []int64{commodity.SellerId})
	if err != nil {
		err = errors.Errorf("get seller from db err! err:%s", err.Error())
		return
	}
	if len(sellerList) == 0 {
		err = errors.Errorf("seller not exist! sellerId:%d", commodity.SellerId)
		return
	}
	seller := sellerList[0]
	sellerInfo = rdm.SellerInfo{
		SellerId:   seller.SellerId,
		SellerName: seller.SellerName,
	}

	//TODO 3. 拉取commodity的comment
	commentInfo = rdm.CommentInfo{
		CommentId:          0,
		CommenterName:      "小星星",
		CommenterAvatarURL: "https://mall-1301454934.cos.ap-shanghai.myqcloud.com/xt%2Cjpg.jpg",
		CommentText:        "太好吃了！",
		ScoreRes:           5,
	}

	return
}
