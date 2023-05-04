package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func OrderReceipt(ctx *gin.Context) {
	req, err := checkOrderReceiptReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.OrderReceiptRsp)
	err = biz.OrderReceipt(ctx, req.OrderId)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkOrderReceiptReq(ctx *gin.Context) (req *rdm.OrderReceiptReq, retErr error) {
	req = new(rdm.OrderReceiptReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
