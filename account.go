package plivogo

type accountClient struct {
	authId    string
	authToken string
	basePath  string
}

func NewAccountClient(authId, authToken string) *accountClient {
	a := new(accountClient)
	a.authId = authId
	a.authToken = authToken
	a.basePath = "/"
	return a
}

type Account struct {
	AuthId       string `json:"auth_id"`
	Type         string `json:"account_type"`
	Address      string `json:"address"`
	AutoRecharge bool   `json:"auto_recharge"`
	BillingMode  string `json:"billing_mode"`
	CashCredit   string `json:"cash_credit"`
	City         string `json:"city"`
	Name         string `json:"name"`
	State        string `json:"state"`
	Timezone     string `json:"timezone"`
}

//Get account resource
func (ac *accountClient) Get() (*Account, error) {
	a := Account{}
	err := getExpectUnmarshal(ac.authId, ac.authToken, ac.basePath, "{}", &a, 200)
	return &a, err
}

//Modify account
func (ac *accountClient) Modify(params *Params) error {
	return postExpectUnmarshal(ac.authId, ac.authToken, ac.basePath, params.Dumps(), nil, 202)
}
