package plivogo

var (
	API_URL = "https://api.plivo.com/v1/Account/"
)

type Plivo struct {
	authId      string
	authToken   string
	Account     *accountClient
	Application *applicationClient
	Endpoint    *endpointClient
}

func NewPlivo(authId, authToken string) *Plivo {
	p := new(Plivo)
	p.authId = authId
	p.authToken = authToken
	p.Account = NewAccountClient(authId, authToken)
	p.Application = NewApplicationClient(authId, authToken)
	p.Endpoint = NewEndpointClient(authId, authToken)
	return p
}

type ResourceMeta struct {
	Previous   int `json:"previous"`
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Next       int `json:"next"`
}

type Call struct {
	authId    string
	authToken string
	path      string
}

func NewCall(authId, authToken string) *Call {
	c := new(Call)
	c.authId = authId
	c.authToken = authToken
	c.path = "/Call/"
	return c
}

//func (c *Call) Call(params string) (string, error) {
//	return doPost(c.authId, c.authToken, c.path, params)
//}
