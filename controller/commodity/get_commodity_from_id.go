package commodity

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetCommodityFromID(ctx *gin.Context) {
	req, err := checkGetCommodityFromIDReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetCommodityFromIDRsp)
	rsp.Commodity, err = biz.GetCommodityFromID(ctx, req.CommodityId)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetCommodityFromIDReq(ctx *gin.Context) (req *rdm.GetCommodityFromIDReq, retErr error) {
	req = new(rdm.GetCommodityFromIDReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
