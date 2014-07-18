package plivogo

import (
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"testing"
)

const (
	authId    = "MAMJFLMZJKMZE0OTZHNT"
	authToken = "YmE1N2NiMDhiNTZlMWE1YjU3NzAwYmYyYTVmYjg3"
)

func TestAccountGetReal(t *testing.T) {
	c := NewClient(authId, authToken)

	_, resp, err := c.Account.Get()
	if err != nil {
		t.Fatalf("raw response=%s", string(resp.Raw))
		t.Fatalf("Account.Get returened error:%v", err)
	}
}

func TestAccountGet(t *testing.T) {
	testSetup()
	defer testTeardown()

	mux.HandleFunc("/"+client.authID+"/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		resp :=
			`
			{
				"account_type": "standard",
				"address": "address",
				"api_id": "40cf0424-d887-11e3-8af9-22000ac988ec",
				"auth_id": "MANJQZNJLJZWY3ZJK5ZW",
				"auto_recharge": false,
				"billing_mode": "prepaid",
				"cash_credits": "186.29384",
				"city": "San Francisco",
				"name": "John Doe",
				"resource_uri": "/v1/Account/MANJQZNJLJZWY3ZJK5ZW/",
				"state": "",
				"timezone": "America/Los_Angeles"
			}
		`
		fmt.Fprint(w, resp)
	})

	account, resp, err := client.Account.Get()

	if err != nil {
		t.Errorf("Account. Get returned error: %v\nraw=%s", err, string(resp.Raw))
	}

	want := &Account{
		Type:         "standard",
		Address:      "address",
		ApiID:        "40cf0424-d887-11e3-8af9-22000ac988ec",
		AuthId:       "MANJQZNJLJZWY3ZJK5ZW",
		AutoRecharge: false,
		BillingMode:  "prepaid",
		CashCredit:   "186.29384",
		City:         "San Francisco",
		Name:         "John Doe",
		ResourceURI:  "/v1/Account/MANJQZNJLJZWY3ZJK5ZW/",
		State:        "",
		Timezone:     "America/Los_Angeles",
	}

	if !reflect.DeepEqual(account, want) {
		t.Errorf("Account.Get returned %+v, want %+v", account, want)
	}

}

func TestAccountModifyReal(t *testing.T) {
	c := NewClient(authId, authToken)

	r := rand.New(rand.NewSource(99))
	randAddress := string(r.Int())

	//modify
	params := map[string]interface{}{
		"address": randAddress,
	}

	_, resp, err := c.Account.Modify(params)
	if err != nil {
		t.Fatalf("raw response=%s", string(resp.Raw))
		t.Fatalf("Account.Modify returened error:%v", err)
	}

	//verify
	account, _, err := c.Account.Get()

	if err != nil {
		t.Errorf("Account.Get failed:%v", err)
	}

	if account.Address != randAddress {
		t.Fatal("Address modification failed")
	}
}
