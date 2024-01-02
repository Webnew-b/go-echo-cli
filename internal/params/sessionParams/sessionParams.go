package sessionParams

import "gorm.io/gorm"

type SessionsRegCreate struct {
	Transaction    *gorm.DB
	SessionId      int64
	RegistrationId int64
}
