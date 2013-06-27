package plivogo

import (
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	API_URL = "https://api.plivo.com/v1/Account/"
)

type Params struct {
	js *simplejson.Json
}

func NewParams() (*Params, error) {
	p := new(Params)
	js, err := simplejson.NewJson([]byte("{}"))
	if err != nil {
		return nil, err
	}
	p.js = js
	return p, nil
}

func (p *Params) Set(key, value string) {
	p.js.Set(key, value)
}

func (p *Params) Dumps() string {
	b, err := p.js.MarshalJSON()
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(b)
}

type Account struct {
	authId    string
	authToken string
}

func (a *Account) Get() (string, error) {
	return doGet(a.authId, a.authToken, "/", "{}")
}

func NewAccount(authId, authToken string) *Account {
	a := new(Account)
	a.authId = authId
	a.authToken = authToken
	return a
}

type Client struct {
	authId    string
	authToken string
	Account   *Account
	Call      *Call
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

func (c *Call) Call(params string) (string, error) {
	return doPost(c.authId, c.authToken, c.path, params)
}

func NewClient(authId, authToken string) (*Client, error) {
	c := new(Client)
	c.authId = authId
	c.authToken = authToken
	c.Account = NewAccount(authId, authToken)
	c.Call = NewCall(authId, authToken)
	return c, nil
}

func doGet(authId, authToken, path, data string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API_URL+authId+path, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	req.SetBasicAuth(authId, authToken)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Printf("%s\n", string(contents))
	return string(contents), nil
}

func doPost(authId, authToken, path, data string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", API_URL+authId+path, strings.NewReader(data))
	if err != nil {
		log.Println(err)
		return "", err
	}

	req.Header.Add("content-type", "application/json")
	req.SetBasicAuth(authId, authToken)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Printf("%s\n", string(contents))
	return string(contents), nil
}
