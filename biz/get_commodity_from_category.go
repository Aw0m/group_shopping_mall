package biz

import (
	"github.com/gin-gonic/gin"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
)

func GetCommodityFromCategory(ctx *gin.Context, categoryId int64, pageNum, pageSize int, WithDeleted bool) (commodityList []bdm.Commodity, retErr error) {
	if pageNum <= 0 || pageSize <= 0 {
		return nil, nil
	}

	offset := pageSize * (pageNum - 1)
	limit := pageSize
	commodityList, retErr = dao.GetCommodityFromCategory(ctx, client_db.GetDB(), categoryId, limit, offset, WithDeleted)
	if retErr != nil {
		return nil, retErr
	}

	return commodityList, nil
}
