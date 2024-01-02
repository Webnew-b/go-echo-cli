package eventParams

type EventFetchRequest struct {
	OrgSlug string `validate:"required"`
	EvSlug  string `validate:"required"`
}

type EventFetchDao struct {
	*EventFetchRequest
	OrgId int64
}

type FetchEventRegParam struct {
	RegId    int64
	TicketId int64
}
