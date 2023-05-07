package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetOrderCount(ctx *gin.Context) {
	req, err := checkGetOrderCountReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetOrderCountRsp)
	rsp.OrderCount, err = biz.GetOrderCount(ctx, req.StatusList)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetOrderCountReq(ctx *gin.Context) (req *rdm.GetOrderCountReq, retErr error) {
	req = new(rdm.GetOrderCountReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
