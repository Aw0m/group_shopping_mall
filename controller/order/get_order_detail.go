package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetOrderDetail(ctx *gin.Context) {
	req, err := checkGetOrderDetailReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetOrderDetailRsp)
	rsp.Order, err = biz.GetOrderDetail(ctx, req.OrderId)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetOrderDetailReq(ctx *gin.Context) (req *rdm.GetOrderDetailReq, retErr error) {
	req = new(rdm.GetOrderDetailReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
