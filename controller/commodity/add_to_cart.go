package commodity

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func AddToCart(ctx *gin.Context) {
	req, err := checkAddToCartReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.AddToCartRsp)
	err = biz.AddToCart(ctx, req.CommodityId, req.CommodityNum)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkAddToCartReq(ctx *gin.Context) (req *rdm.AddToCartReq, retErr error) {
	req = new(rdm.AddToCartReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
