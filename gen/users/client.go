// Code generated by goa v3.7.3, DO NOT EDIT.
//
// users client
//
// Command:
// $ goa gen goa-kamion-example/design

package users

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "users" service client.
type Client struct {
	InviteEndpoint goa.Endpoint
}

// NewClient initializes a "users" service client given the endpoints.
func NewClient(invite goa.Endpoint) *Client {
	return &Client{
		InviteEndpoint: invite,
	}
}

// Invite calls the "invite" endpoint of the "users" service.
func (c *Client) Invite(ctx context.Context, p *InvitePayload) (res *AuthUserResponse, err error) {
	var ires interface{}
	ires, err = c.InviteEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*AuthUserResponse), nil
}
