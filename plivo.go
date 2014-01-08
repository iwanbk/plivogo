package plivogo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	API_URL = "https://api.plivo.com/v1/Account/"
)

type Plivo struct {
	authId    string
	authToken string
	Account   *AccountClient
	Endpoint  *EndpointClient
}

func NewPlivo(authId, authToken string) *Plivo {
	p := new(Plivo)
	p.authId = authId
	p.authToken = authToken
	p.Account = NewAccountClient(authId, authToken)
	p.Endpoint = NewEndpointClient(authId, authToken)
	return p
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
