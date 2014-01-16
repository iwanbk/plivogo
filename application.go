package plivogo

var basePath = "/Application/"

type applicationClient struct {
	authId    string
	authToken string
}

//Application Client Constructor
func NewApplicationClient(authId, authToken string) *applicationClient {
	a := new(applicationClient)
	a.authId = authId
	a.authToken = authToken

	return a
}

//Application give details of Plivo Application
type Application struct {
	Id                string `json:"app_id"`
	Name              string `json:"app_name"`
	DefaultApp        bool   `json:"default_app"`
	HangupUrl         string `json:"hangup_url"`
	HangupMethod      string `json:"hangup_method"`
	AnswerUrl         string `json:"answer_url"`
	AnswerMethod      string `json:"answer_method"`
	MessageUrl        string `json:"message_url"`
	MessageMethod     string `json:"message_method"`
	FallbackAnswerUrl string `json:"fallback_answer_url"`
	FallbackMethod    string `json:"fallback_method"`
}

//Aplication list
type ApplicationList struct {
	Meta ResourceMeta  `json:"meta"`
	List []Application `json:"objects"`
}

type appCreateResponse struct {
	Message string `json:"message"`
	AppId   string `json:"app_id"`
}

//Get application
func (ac *applicationClient) Get(id string) (*Application, error) {
	app := Application{}
	path := basePath + id + "/"
	err := getExpectUnmarshal(ac.authId, ac.authToken, path, "{}", &app, 200)
	return &app, err
}

func (ac *applicationClient) GetList(params *Params) (*ApplicationList, error) {
	appList := ApplicationList{}
	err := getExpectUnmarshal(ac.authId, ac.authToken, basePath, params.Dumps(), &appList, 200)
	return &appList, err
}

//Create Plivo Application
func (ac *applicationClient) Create(params *Params) (string, error) {
	acr := appCreateResponse{}
	err := postExpectUnmarshal(ac.authId, ac.authToken, basePath, params.Dumps(), &acr, 201)
	if err != nil {
		return "", err
	}
	return acr.AppId, nil
}

//Delete plivo application
func (ac *applicationClient) Delete(id string) error {
	return deleteExpectUnmarshal(ac.authId, ac.authToken, basePath+id+"/", "{}", nil, 204)
}

//Modify Plivo Application
func (ac *applicationClient) Modify(id string, params *Params) error {
	return postExpectUnmarshal(ac.authId, ac.authToken, basePath+id+"/", params.Dumps(), nil, 202)
}
