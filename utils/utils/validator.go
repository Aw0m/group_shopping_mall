package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

var validate *validator.Validate

func ValidateInit() {
	validate = validator.New()
}

func ReqBindAndCheck(ctx *gin.Context, req interface{}) error {
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		return errors.Errorf("wrong param! bind err: %s", err.Error())
	}

	err = validate.Struct(req)
	if err != nil {
		return errors.Errorf("wrong param! validate err: %s", err.Error())
	}

	return nil
}
