package plivogo

import (
	"log"
)

type AccountClient struct {
	authId    string
	authToken string
}

func NewAccountClient(authId, authToken string) *AccountClient {
	a := new(AccountClient)
	a.authId = authId
	a.authToken = authToken
	return a
}

type Account struct {
	resourceStr  string
	resource     *Resource
	AuthId       string
	AccountType  string
	Address      string
	AutoRecharge bool
	BillingMode  string
	CashCredit   string
	City         string
	Name         string
	State        string
	Timezone     string
}

func NewAccount(jsonStr string) (*Account, error) {
	a := new(Account)

	r, err := NewResource(jsonStr)
	if err != nil {
		log.Println("Failed to parse account resource string:", err.Error())
		return nil, err
	}
	a.resourceStr = jsonStr
	a.resource = r

	if a.AccountType, err = r.GetString("account_type"); err != nil {
		return nil, err
	}
	if a.Address, err = r.GetString("address"); err != nil {
		return nil, err
	}
	if a.AuthId, err = r.GetString("auth_id"); err != nil {
		return nil, err
	}
	if a.AutoRecharge, err = r.GetBool("auto_recharge"); err != nil {
		return nil, err
	}
	if a.BillingMode, err = r.GetString("billing_mode"); err != nil {
		return nil, err
	}
	if a.CashCredit, err = r.GetString("cash_credit"); err != nil {
		return nil, err
	}
	if a.City, err = r.GetString("city"); err != nil {
		return nil, err
	}
	if a.Name, err = r.GetString("name"); err != nil {
		return nil, err
	}
	if a.State, err = r.GetString("state"); err != nil {
		return nil, err
	}
	if a.Timezone, err = r.GetString("timezone"); err != nil {
		return nil, err
	}
	return a, nil
}

//Get account resource
func (ac *AccountClient) Get() (*Account, error) {
	jsonStr, err := Get(ac.authId, ac.authToken, "/", "{}")
	if err != nil {
		return nil, err
	}

	a, err := NewAccount(jsonStr)
	if err != nil {
		return nil, err
	}
	return a, nil
}
