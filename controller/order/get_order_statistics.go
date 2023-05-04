package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetOrderStatistics(ctx *gin.Context) {
	_, err := checkGetOrderStatisticsReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetOrderStatisticsRsp)
	rsp.PendingNum, rsp.UnshippedNum, rsp.ShippedNum, rsp.CompletedNum, err = biz.GetOrderStatistics(ctx)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetOrderStatisticsReq(ctx *gin.Context) (req *rdm.GetOrderStatisticsReq, retErr error) {
	req = new(rdm.GetOrderStatisticsReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
