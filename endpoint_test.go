package plivogo

import (
	"testing"
)

func TestGetEndpoint(t *testing.T) {
	p := NewPlivo(authId, authToken)

	ep, err := p.Endpoint.Get("85785376675499")
	if err != nil {
		t.Error("Get endpoint failed:" + err.Error())
	}
	if ep.Alias != "appsub" {
		t.Error("invalid alias=" + ep.Alias)
	}
}

func TestGetEndpointList(t *testing.T) {
	p := NewPlivo(authId, authToken)
	params, _ := NewParams()

	epl, err := p.Endpoint.GetList(params)
	if err != nil {
		t.Error("Get endpoint list failed:" + err.Error())
	}
	t.Logf("Number of endpoint = %d\n", len(epl.List))
	for _, ep := range epl.List {
		t.Log("endpoint username = " + ep.Username)
	}
}

func TestEndpointCrud(t *testing.T) {
	p := NewPlivo(authId, authToken)

	//create test
	params, _ := NewParams()
	params.Set("username", "gotest")
	params.Set("alias", "gotest")
	params.Set("password", "gotest")

	ecr, err := p.Endpoint.Create(params)
	if err != nil {
		t.Fatal("Endpoint create failed:" + err.Error())
	}

	endpointId := ecr.Id

	//edit it
	params, _ = NewParams()
	params.Set("alias", "newalias")
	err = p.Endpoint.Modify(endpointId, params)
	if err != nil {
		t.Error("Endpoint " + endpointId + " edit failed :" + err.Error())
	}

	//verify our changes
	ep, err := p.Endpoint.Get(endpointId)
	if err != nil {
		t.Error("Get endpoint " + endpointId + " failed :" + err.Error())
	}
	if ep.Alias != "newalias" {
		t.Error("Endpoint " + endpointId + " edit failed : alias is not changed")
	}

	//delete
	err = p.Endpoint.Delete(endpointId)
	if err != nil {
		t.Log("Endpoint " + endpointId + " delete failed :" + err.Error())
	}
}
