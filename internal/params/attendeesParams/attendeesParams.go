package attendeesParams

type AttendeeAuthReq struct {
	Token string `validate:"required"`
}
