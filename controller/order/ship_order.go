package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func ShipOrder(ctx *gin.Context) {
	req, err := checkShipOrderReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.ShipOrderRsp)
	err = biz.ShipOrder(ctx, req.OrderId, req.DeliverymanName, req.DeliverymanPhoneNum)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkShipOrderReq(ctx *gin.Context) (req *rdm.ShipOrderReq, retErr error) {
	req = new(rdm.ShipOrderReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
