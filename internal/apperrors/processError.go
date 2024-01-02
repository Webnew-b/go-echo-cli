package apperrors

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/internal/apperrors/eventError"
	"wscmakebygo.com/internal/apperrors/organizerError"
	"wscmakebygo.com/internal/apperrors/registrationsError"
	"wscmakebygo.com/internal/apperrors/ticketsError"
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
	switch {
	case errors.Is(err, &eventError.EventSlugNotFoundError{}),
		errors.Is(err, &organizerError.OrganizerSlugNotFoundError{}),
		errors.Is(err, &attendeesError.NotLogin{}),
		errors.Is(err, &ticketsError.NotAvailable{}),
		errors.Is(err, &registrationsError.AlreadyRegistrar{}):
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
