package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func CreateOrder(ctx *gin.Context) {
	req, err := checkCreateOrderReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.CreateOrderRsp)
	rsp.OrderList, err = biz.CreateOrder(ctx, req.OrderCommodityInfoList)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkCreateOrderReq(ctx *gin.Context) (req *rdm.CreateOrderReq, retErr error) {
	req = new(rdm.CreateOrderReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
