package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func WXLogin(ctx *gin.Context) {
	req, err := checkWXLogin(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.WXLoginResponse)
	rsp.User, rsp.Token, err = biz.WXLogin(ctx, req.Code, req.Username, req.AvatarURL, req.Gender)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)

	return
}

func checkWXLogin(ctx *gin.Context) (req *rdm.WXLoginRequest, retErr error) {
	req = &rdm.WXLoginRequest{}
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	if req.Code == "" {
		return req, errors.New("wrong param! code is nil")
	}

	gender := req.Gender
	if gender != "m" && gender != "f" && gender != "u" {
		return req, errors.Errorf("wrong param! gender is %s, not m/f/u", gender)
	}

	return req, nil
}
