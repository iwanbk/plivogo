package plivogo

import (
	"testing"
)

const (
	authId    = "MAMJFLMZJKMZE0OTZHNT"
	authToken = "YmE1N2NiMDhiNTZlMWE1YjU3NzAwYmYyYTVmYjg3"
)

func TestGet(t *testing.T) {
	p := NewPlivo(authId, authToken)
	account, err := p.Account.Get()
	if err != nil {
		t.Fatal("Get account failed:", err.Error())
	}
	if account.Name != "Iwan Budi Kusnanto" {
		t.Fatal("Account name invalid:", account.Name)
	}
	if account.AuthId != "MAMJFLMZJKMZE0OTZHNT" {
		t.Fatal("Invalid auth id:", account.AuthId)
	}
}
