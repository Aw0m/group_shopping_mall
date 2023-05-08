package commodity

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetCommodityFromCategory(ctx *gin.Context) {
	req, err := checkGetCommodityFromCategoryReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetCommodityFromCategoryRsp)
	rsp.CommodityList, err = biz.GetCommodityFromCategory(ctx, req.CategoryId, req.PageNum, req.PageSize, req.WithDeleted)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetCommodityFromCategoryReq(ctx *gin.Context) (req *rdm.GetCommodityFromCategoryReq, retErr error) {
	req = new(rdm.GetCommodityFromCategoryReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
