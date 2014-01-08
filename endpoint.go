package plivogo

type EndpointClient struct {
	authId    string
	authToken string
}

func NewEndpointClient(authId, authToken string) *EndpointClient {
	e := new(EndpointClient)
	e.authId = authId
	e.authToken = authToken

	return e
}

type Endpoint struct {
	resourceStr    string
	resource       *Resource
	Alias          string
	Id             string
	ApplicationUri string
	SipRegistered  string
	Password       string
	SipUri         string
	SubAccount     string
	Username       string
	ResourceUri    string
}

func NewEndpoint(jsonStr string) (*Endpoint, error) {
    r, err := NewResource(jsonStr)
    if err != nil {
        return nil, err
    }
    e := new(Endpoint)
    e.resourceStr = jsonStr
    e.resource = r

    if e.Alias, err = r.GetString("alias"); err != nil {
        return nil, err
    }
    if e.Id, err = r.GetString("endpoint_id"); err != nil {
        return nil, err
    }
    if e.ApplicationUri, err = r.GetString("application"); err != nil {
        return nil, err
    }
    if e.SipRegistered, err = r.GetString("sip_registered"); err != nil {
        return nil, err
    }
    if e.Password, err = r.GetString("password"); err != nil {
        return nil, err
    }
    if e.SipUri, err = r.GetString("sip_uri"); err != nil {
        return nil, err
    }
    if e.SubAccount, err = r.GetString("sub_account"); err != nil {
        return nil, err
    }
    if e.Username, err = r.GetString("username"); err != nil {
        return nil, err
    }
    if e.ResourceUri, err = r.GetString("resource_uri"); err != nil {
        return nil, err
    }
    return e, nil
}
func (ec *EndpointClient) Get(id string) (*Endpoint, error) {
	jsonStr, err := Get(ec.authId, ec.authToken, "/Endpoint/"+id+"/", "{}")
	if err != nil {
        return nil, err
    }
    e, err := NewEndpoint(jsonStr)
    if err != nil {
        return nil, err
    }
    return e, nil
}
