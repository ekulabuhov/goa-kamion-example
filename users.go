package usersapi

import (
	"context"
	users "goa-kamion-example/gen/users"
	"log"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	logger *log.Logger
}

// NewUsers returns the users service implementation.
func NewUsers(logger *log.Logger) users.Service {
	return &userssrvc{logger}
}

// Invite implements invite.
func (s *userssrvc) Invite(ctx context.Context, p *users.InvitePayload) (res *users.AuthUserResponse, err error) {
	res = &users.AuthUserResponse{}
	s.logger.Print("users.invite")
	return
}
