package plivogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiURL = "https://api.plivo.com/v1/Account"
)

var (
	API_URL = "https://api.plivo.com/v1/Account/"
)

type Plivo struct {
	authId      string
	authToken   string
	Application *applicationClient
	Endpoint    *endpointClient
}

func NewPlivo(authId, authToken string) *Plivo {
	p := new(Plivo)
	p.authId = authId
	p.authToken = authToken
	p.Application = NewApplicationClient(authId, authToken)
	p.Endpoint = NewEndpointClient(authId, authToken)
	return p
}

//Client represents a Plivo API client
type Client struct {
	authID    string
	authToken string
	hc        *httpClient
	baseURL   string

	Account *AccountService
}

//Response represent a Plivo API HTTP response
type Response struct {
	*http.Response
	Raw []byte
}

//ResourceMeta represent meta info of plivo API response
type ResourceMeta struct {
	Previous   int `json:"previous"`
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Next       int `json:"next"`
}

//GenericResponse represents plivo generic response
//Usually returned when we do PUT/POST
type GenericResponse struct {
	Message string `json:"message"`
	ApiID   string `json:"api_id"`
}

//NewClient returns a new Plivo API client.
func NewClient(authID, authToken string) *Client {
	c := &Client{
		authID:    authID,
		authToken: authToken,
		baseURL:   apiURL,
	}

	hc := newHttpClient(authID, authToken)

	c.hc = hc
	c.Account = &AccountService{client: c}

	return c
}

func (c *Client) Get(path string, v interface{}) (*Response, error) {
	return c.hc.Get(c.baseURL+"/"+c.authID+"/"+path, v)
}

func (c *Client) Post(path string, body, v interface{}) (*Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("post body=%s\n", string(b))
	r := bytes.NewReader(b)

	return c.hc.Post(c.baseURL+"/"+c.authID+"/"+path, r, v)
}

func newResponse(r *http.Response) *Response {
	resp := &Response{Response: r}
	return resp
}
