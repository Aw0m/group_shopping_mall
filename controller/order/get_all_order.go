package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetAllOrder(ctx *gin.Context) {
	req, err := checkGetAllOrderReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetAllOrderRsp)
	rsp.OrderList, err = biz.GetAllOrder(ctx, req.StatusList, req.PageSize, req.PageNum)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetAllOrderReq(ctx *gin.Context) (req *rdm.GetAllOrderReq, retErr error) {
	req = new(rdm.GetAllOrderReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
