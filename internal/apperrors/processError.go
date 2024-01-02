package apperrors

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/tools/logUtil"
)

var valid = validator.New()

func ValidateStruct(i interface{}, errMsg string) error {
	err := valid.Struct(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errMsg)
	}
	return nil
}

func HandleError(err error) error {
	logUtil.Log.Println(err.Error())
	// todo 确认错误处理
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
