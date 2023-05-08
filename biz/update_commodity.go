package biz

import (
	"github.com/gin-gonic/gin"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/rdm"
)

func UpdateCommodity(ctx *gin.Context, req *rdm.UpdateCommodityReq) (retErr error) {
	commodityID := req.CommodityId

	updateMap := make(map[string]any)
	if req.CommodityName != nil {
		updateMap["commodity_name"] = *req.CommodityName
	}

	if req.Price != nil {
		updateMap["price"] = *req.Price
	}

	if req.CategoryID != nil {
		updateMap["category_id"] = *req.CategoryID
	}

	if req.CommodityStatus != nil {
		updateMap["commodity_status"] = *req.CommodityStatus
	}

	if req.IsDeleted != nil {
		updateMap["is_deleted"] = *req.IsDeleted
	}

	retErr = dao.UpdateCommodity(ctx, client_db.GetDB(), commodityID, updateMap)
	return
}
