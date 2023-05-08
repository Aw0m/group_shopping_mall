package biz

import (
	"github.com/gin-gonic/gin"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
)

func GetCommodityCount(ctx *gin.Context) (commodityCount int64, retErr error) {
	commodityCount, retErr = dao.GetCommodityCount(ctx, client_db.GetDB())
	return
}
