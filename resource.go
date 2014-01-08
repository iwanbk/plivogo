package plivogo

import (
	simplejson "github.com/bitly/go-simplejson"
)

type Resource struct {
	jsStr string
	jsObj *simplejson.Json
}

func NewResource(jsStr string) (*Resource, error) {
	r := new(Resource)
	r.jsStr = jsStr

	jsObj, err := simplejson.NewJson([]byte(jsStr))
	if err != nil {
		return nil, err
	}
	r.jsObj = jsObj
	return r, nil
}

func (r *Resource) GetString(key string) (string, error) {
	return r.jsObj.Get(key).MustString(), nil
}

func (r *Resource) GetBool(key string) (bool, error) {
	return r.jsObj.Get(key).MustBool(), nil
}
