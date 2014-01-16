package plivogo

//endpoint client
type endpointClient struct {
	authId    string
	authToken string
	basePath  string
}

//Endpoint Client Constructor
func NewEndpointClient(authId, authToken string) *endpointClient {
	e := new(endpointClient)
	e.authId = authId
	e.authToken = authToken
	e.basePath = "/Endpoint/"

	return e
}

//Endpoint hold details about Plivo Endpoint
type Endpoint struct {
	Alias         string `json:"alias"`
	Id            string `json:"endpoint_id"`
	Application   string `json:"application"`
	SipRegistered string `json:"sip_registered"`
	Password      string `json:"password"`
	SipUri        string `json:"sip_uri"`
	SubAccount    string `json:"sub_account"`
	Username      string `json:"username"`
	ResourceUri   string `json:"resource_uri"`
	resourceStr   string
	resource      *Resource
}

//EndpointList is a list of Plivo Endpoint based on some search criteria
type EndpointList struct {
	Meta ResourceMeta `json:"meta"`
	List []Endpoint   `json:"objects"`
}

//EndpointCreateResponse hold details about information returned when creating endpoint
type EndpointCreateResponse struct {
	Username string `json:"username"`
	Alias    string `json:"alias"`
	Id       string `json:"endpoint_id"`
	message  string `json:"message"`
	apiId    string `json:"api_id"`
}

//Get an endpoint
func (ec *endpointClient) Get(id string) (*Endpoint, error) {
	ep := Endpoint{}
	err := getExpectUnmarshal(ec.authId, ec.authToken, ec.basePath+id+"/", "{}", &ep, 200)
	return &ep, err
}

//Get endpoint list
func (ec *endpointClient) GetList(params *Params) (*EndpointList, error) {
	el := EndpointList{}
	err := getExpectUnmarshal(ec.authId, ec.authToken, ec.basePath, params.Dumps(), &el, 200)
	return &el, err
}

//Create an endpoint
func (ec *endpointClient) Create(params *Params) (*EndpointCreateResponse, error) {
	ecr := EndpointCreateResponse{}
	err := postExpectUnmarshal(ec.authId, ec.authToken, ec.basePath, params.Dumps(), &ecr, 201)
	return &ecr, err
}

//Delete an endpoint
func (ec *endpointClient) Delete(id string) error {
	return deleteExpectUnmarshal(ec.authId, ec.authToken, ec.basePath+id+"/", "{}", nil, 204)
}

//Modify Plivo Endpoint
func (ec *endpointClient) Modify(id string, params *Params) error {
	return postExpectUnmarshal(ec.authId, ec.authToken, ec.basePath+id+"/", params.Dumps(), nil, 202)
}
