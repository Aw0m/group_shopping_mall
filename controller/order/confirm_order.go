package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func ConfirmOrder(ctx *gin.Context) {
	req, err := checkConfirmOrderReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.ConfirmOrderRsp)
	err = biz.ConfirmOrder(ctx, req.OrderId)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkConfirmOrderReq(ctx *gin.Context) (req *rdm.ConfirmOrderReq, retErr error) {
	req = new(rdm.ConfirmOrderReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
