package biz

import (
	"github.com/gin-gonic/gin"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
)

func GetUserCount(ctx *gin.Context) (userCount int64, retErr error) {
	return dao.GetUserCount(ctx, client_db.GetDB())
}
