// Code generated by goa v3.7.3, DO NOT EDIT.
//
// users HTTP client encoders and decoders
//
// Command:
// $ goa gen goa-kamion-example/design

package client

import (
	"bytes"
	"context"
	users "goa-kamion-example/gen/users"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildInviteRequest instantiates a HTTP request object with method and path
// set to call the "users" service "invite" endpoint
func (c *Client) BuildInviteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: InviteUsersPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("users", "invite", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeInviteRequest returns an encoder for requests sent to the users invite
// server.
func EncodeInviteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*users.InvitePayload)
		if !ok {
			return goahttp.ErrInvalidType("users", "invite", "*users.InvitePayload", v)
		}
		body := NewInviteRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("users", "invite", err)
		}
		return nil
	}
}

// DecodeInviteResponse returns a decoder for responses returned by the users
// invite endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeInviteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body InviteResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("users", "invite", err)
			}
			err = ValidateInviteResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("users", "invite", err)
			}
			res := NewInviteAuthUserResponseOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("users", "invite", resp.StatusCode, string(body))
		}
	}
}

// marshalUsersLoginInviteBodyToLoginInviteBodyRequestBody builds a value of
// type *LoginInviteBodyRequestBody from a value of type *users.LoginInviteBody.
func marshalUsersLoginInviteBodyToLoginInviteBodyRequestBody(v *users.LoginInviteBody) *LoginInviteBodyRequestBody {
	if v == nil {
		return nil
	}
	res := &LoginInviteBodyRequestBody{
		Email:          v.Email,
		Password:       v.Password,
		OrganisationID: v.OrganisationID,
	}

	return res
}

// marshalLoginInviteBodyRequestBodyToUsersLoginInviteBody builds a value of
// type *users.LoginInviteBody from a value of type *LoginInviteBodyRequestBody.
func marshalLoginInviteBodyRequestBodyToUsersLoginInviteBody(v *LoginInviteBodyRequestBody) *users.LoginInviteBody {
	if v == nil {
		return nil
	}
	res := &users.LoginInviteBody{
		Email:          v.Email,
		Password:       v.Password,
		OrganisationID: v.OrganisationID,
	}

	return res
}
