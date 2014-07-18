package plivogo

import (
	"fmt"
	"net/http"
)

const (
	accountBasePath = ""
)

//AccountService handles communication with Plivo Account API
//
//See : https://www.plivo.com/docs/api/account/
type AccountService struct {
	client *Client
}

type Account struct {
	AuthId       string `json:"auth_id,omitempty"`
	Type         string `json:"account_type,omitempty"`
	Address      string `json:"address,omitempty"`
	AutoRecharge bool   `json:"auto_recharge,omitempty"`
	BillingMode  string `json:"billing_mode,omitempty"`
	CashCredit   string `json:"cash_credits,omitempty"`
	City         string `json:"city,omitempty"`
	Name         string `json:"name,omitempty"`
	State        string `json:"state,omitempty"`
	Timezone     string `json:"timezone,omitempty"`

	//other fields
	ApiID       string `json:"api_id"`
	ResourceURI string `json:"resource_uri"`
}

//Get account resource
func (a *AccountService) Get() (*Account, *Response, error) {
	var ac Account

	resp, err := a.client.Get(accountBasePath, &ac)

	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, resp, fmt.Errorf("invalid status code:%d", resp.StatusCode)
	}
	return &ac, resp, nil
}

func (as *AccountService) Modify(params map[string]interface{}) (*GenericResponse, *Response, error) {
	var gr GenericResponse

	//do POST
	resp, err := as.client.Post(accountBasePath, params, &gr)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusAccepted {
		return nil, resp, fmt.Errorf("invalid status code : %d", resp.StatusCode)
	}

	return &gr, resp, nil
}
