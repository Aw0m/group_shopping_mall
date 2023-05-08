package commodity

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func UpdateCommodity(ctx *gin.Context) {
	req, err := checkUpdateCommodityReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.UpdateCommodityRsp)
	err = biz.UpdateCommodity(ctx, req)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkUpdateCommodityReq(ctx *gin.Context) (req *rdm.UpdateCommodityReq, retErr error) {
	req = new(rdm.UpdateCommodityReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
