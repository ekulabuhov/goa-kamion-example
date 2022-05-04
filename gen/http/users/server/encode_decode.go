// Code generated by goa v3.7.3, DO NOT EDIT.
//
// users HTTP server encoders and decoders
//
// Command:
// $ goa gen goa-kamion-example/design

package server

import (
	"context"
	users "goa-kamion-example/gen/users"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeInviteResponse returns an encoder for responses returned by the users
// invite endpoint.
func EncodeInviteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*users.AuthUserResponse)
		enc := encoder(ctx, w)
		body := NewInviteResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeInviteRequest returns a decoder for requests sent to the users invite
// endpoint.
func DecodeInviteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body InviteRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateInviteRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewInvitePayload(&body)

		return payload, nil
	}
}

// unmarshalLoginInviteBodyRequestBodyToUsersLoginInviteBody builds a value of
// type *users.LoginInviteBody from a value of type *LoginInviteBodyRequestBody.
func unmarshalLoginInviteBodyRequestBodyToUsersLoginInviteBody(v *LoginInviteBodyRequestBody) *users.LoginInviteBody {
	if v == nil {
		return nil
	}
	res := &users.LoginInviteBody{
		Email:          *v.Email,
		Password:       *v.Password,
		OrganisationID: *v.OrganisationID,
	}

	return res
}
