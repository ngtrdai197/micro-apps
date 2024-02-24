package user_trpt

type UserTranporter interface {
}

type userTransporter struct {
}

func New() UserTranporter {
	return &userTransporter{}
}
