package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func CancelOrder(ctx *gin.Context) {
	req, err := checkCancelOrderReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.CancelOrderRsp)
	err = biz.CancelOrder(ctx, req.OrderId)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkCancelOrderReq(ctx *gin.Context) (req *rdm.CancelOrderReq, retErr error) {
	req = new(rdm.CancelOrderReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
