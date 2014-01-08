package plivogo

import (
	simplejson "github.com/bitly/go-simplejson"
	"log"
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
