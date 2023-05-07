package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func UpdateOrderInfo(ctx *gin.Context) {
	req, err := checkUpdateOrderInfoReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.UpdateOrderInfoRsp)
	err = biz.UpdateOrderInfo(ctx,
		req.OrderId, req.SellerName, req.CommodityName,
		req.DeliverymanName, req.DeliverymanPhoneNum, req.PhoneNum,
		req.ConsigneeName, req.AddressDetailInfo, req.Price, req.Quantity,
		req.Status, req.IsDeleted,
	)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkUpdateOrderInfoReq(ctx *gin.Context) (req *rdm.UpdateOrderInfoReq, retErr error) {
	req = new(rdm.UpdateOrderInfoReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
