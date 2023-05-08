package commodity

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetCommodityCount(ctx *gin.Context) {
	_, err := checkGetCommodityCountReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetCommodityCountRsp)
	rsp.CommodityCount, err = biz.GetCommodityCount(ctx)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetCommodityCountReq(ctx *gin.Context) (req *rdm.GetCommodityCountReq, retErr error) {
	req = new(rdm.GetCommodityCountReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
