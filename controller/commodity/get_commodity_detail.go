package commodity

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetCommodityDetail(ctx *gin.Context) {
	req, err := checkGetCommodityDetailReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetCommodityDetailRsp)
	rsp.CommodityInfo, rsp.SellerInfo, rsp.CommentInfo, err = biz.GetCommodityDetail(ctx, req.CommodityId)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetCommodityDetailReq(ctx *gin.Context) (req *rdm.GetCommodityDetailReq, retErr error) {
	req = new(rdm.GetCommodityDetailReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
