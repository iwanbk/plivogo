package plivogo

import (
	"testing"
)

func TestGetEndpoint(t *testing.T) {
	p := NewPlivo(authId, authToken)

	ep, err := p.Endpoint.Get("85785376675499")
	if err != nil {
		t.Fatal("Get endpoint failed:", err.Error())
	}
    if ep.Alias != "appsub" {
        t.Fatal("invalid alias=", ep.Alias)
    }
}
