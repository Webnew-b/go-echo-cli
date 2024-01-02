package middleware

import (
	"github.com/labstack/echo/v4"
	"wscmakebygo.com/internal/apperrors"
	"wscmakebygo.com/internal/apperrors/attendeesError"
	"wscmakebygo.com/internal/dao/attendeesDao"
	"wscmakebygo.com/internal/params/attendeesParams"
)

func UserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := attendeesParams.AttendeeAuthReq{
			Token: c.QueryParam("token"),
		}
		err := apperrors.ValidateStruct(token, attendeesError.ErrInvalidTokenMessage)
		if err != nil {
			return err
		}
		id, err := attendeesDao.FetchAttendeeIdByCache(token.Token)
		if err != nil {
			return apperrors.HandleError(err)
		}
		attendee, err := attendeesDao.FetchAttendeeById(id)
		if err != nil {
			return err
		}
		c.Set("attendee", attendee)
		return next(c)
	}
}
