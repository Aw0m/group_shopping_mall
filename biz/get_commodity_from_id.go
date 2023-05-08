package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
)

func GetCommodityFromID(ctx *gin.Context, commodityId int64) (commodity bdm.Commodity, retErr error) {
	commodityList, err := dao.BatchGetCommodityByIdList(ctx, client_db.GetDB(), []int64{commodityId})
	if err != nil {
		retErr = errors.WithMessage(err, "dao.BatchGetCommodityByIdList failed")
		return
	}
	if len(commodityList) == 0 {
		retErr = errors.Errorf("commodity not exist! commodityId:%d", commodityId)
		return
	}

	commodity = commodityList[0]
	return commodity, retErr
}
