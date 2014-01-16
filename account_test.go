package plivogo

import (
    "math/rand"
	"testing"
)

const (
	authId    = "MAMJFLMZJKMZE0OTZHNT"
	authToken = "YmE1N2NiMDhiNTZlMWE1YjU3NzAwYmYyYTVmYjg3"
)

func TestAccountGet(t *testing.T) {
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

func TestAccountModify(t *testing.T) {
	p := NewPlivo(authId, authToken)
    r := rand.New(rand.NewSource(99))
    randAddress := string(r.Int())
    params,_ := NewParams()
    params.Set("address", randAddress)

    err := p.Account.Modify(params)
    if err != nil {
        t.Error("Modification attempt failed:" + err.Error())
    }

    //verify
    account, err := p.Account.Get()
	if err != nil {
		t.Fatal("Get account failed:", err.Error())
	}
	if account.Address != randAddress {
		t.Fatal("Address modification failed")
	}

}
