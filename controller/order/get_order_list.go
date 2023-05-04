package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetOrderList(ctx *gin.Context) {
	req, err := checkGetOrderListReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetOrderListRsp)
	rsp.OrderList, err = biz.GetOrderList(ctx, req.StatusList)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetOrderListReq(ctx *gin.Context) (req *rdm.GetOrderListReq, retErr error) {
	req = new(rdm.GetOrderListReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
