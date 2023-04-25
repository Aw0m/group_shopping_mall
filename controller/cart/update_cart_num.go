package cart

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func UpdateCartNum(ctx *gin.Context) {
	req, err := checkUpdateCartNumReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.UpdateCartNumRsp)
	err = biz.UpdateCartNum(ctx, req.CartId, req.Quantity)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkUpdateCartNumReq(ctx *gin.Context) (req *rdm.UpdateCartNumReq, retErr error) {
	req = new(rdm.UpdateCartNumReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
