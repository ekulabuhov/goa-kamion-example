package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("users", func() {
	Title("User Service")
	Description("Service for managing users")
	Server("users", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("users", func() {
	Method("invite", func() {
		Payload(func() {
			Attribute("user", LoginInviteBody)
		})

		Result(AuthUserResponse)

		HTTP(func() {
			POST("/login/invite")
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

var LoginInviteBody = Type("LoginInviteBody", func() {
	Description("Type sent to login/invite method")

	Attribute("email", String, func() {
		Description("New user email")
		Example("test@test.com")
	})
	Attribute("password", String, func() {
		Description("New user password")
		Example("Password!1")
	})
	Attribute("organisationID", String, func() {
		Description("ID of the organisation he was invited to")
		Example("0763c56a-673b-4d30-9777-43cb1720ddfb")
	})

	Required("email", "password", "organisationID")
})

var AuthUserResponse = Type("AuthUserResponse", func() {
	Description("Response sent by login/invite method")

	Attribute("status", String, func() {
		Description("Status of the performed request")
		Enum("success", "failure")
	})
})
