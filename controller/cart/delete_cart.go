package cart

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func DeleteCart(ctx *gin.Context) {
	req, err := checkDeleteCartReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.DeleteCartRsp)
	err = biz.DeleteCart(ctx, req.CartId)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkDeleteCartReq(ctx *gin.Context) (req *rdm.DeleteCartReq, retErr error) {
	req = new(rdm.DeleteCartReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
