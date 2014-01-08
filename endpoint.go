package plivogo

type EndpointClient struct {
	authId    string
	authToken string
}

func NewEndpointClient(authId, authToken string) *EndpointClient {
	e := new(EndpointClient)
	e.authId = authId
	e.authToken = authToken

	return e
}
