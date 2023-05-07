package category

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func DeleteCategory(ctx *gin.Context) {
	req, err := checkDeleteCategoryReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.DeleteCategoryRsp)
	err = biz.DeleteCategory(ctx, req.CategoryId)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkDeleteCategoryReq(ctx *gin.Context) (req *rdm.DeleteCategoryReq, retErr error) {
	req = new(rdm.DeleteCategoryReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
