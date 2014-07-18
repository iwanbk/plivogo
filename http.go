package plivogo

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ApiError struct {
	StatusCode int
	RawError   string
}

type httpClient struct {
	authID    string
	authToken string
	client    *http.Client
}

func newHttpClient(authID, authToken string) *httpClient {
	h := &httpClient{
		authID:    authID,
		authToken: authToken,
		client:    http.DefaultClient,
	}
	return h
}

func (h *httpClient) Get(url string, v interface{}) (*Response, error) {
	log.Printf("doing GET request to :%s\n", url)
	//create request object
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//set auth
	req.SetBasicAuth(h.authID, h.authToken)

	//do request
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}

	//read response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	response.Raw = b

	//decode response
	err = json.Unmarshal(b, &v)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (h *httpClient) Post(url string, body io.Reader, v interface{}) (*Response, error) {
	//create request object
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	//set auth
	req.SetBasicAuth(h.authID, h.authToken)

	//do request
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}

	//read response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	response.Raw = b

	//decode response
	err = json.Unmarshal(b, &v)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (e ApiError) Error() string {
	return string(e.StatusCode) + ":" + e.RawError
}

func get(authId, authToken, path, data string) (int, string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API_URL+authId+path, nil)
	if err != nil {
		log.Println(err)
		return -1, "", err
	}
	req.SetBasicAuth(authId, authToken)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return -1, "", err
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return -1, "", err
	}
	return resp.StatusCode, string(contents), nil
}

func getExpectUnmarshal(authId, authToken, path, data string, v interface{}, expectedCode int) error {
	code, resp, err := get(authId, authToken, path, data)
	if err != nil {
		return err
	}
	if code != expectedCode {
		apiErr := ApiError{code, resp}
		return apiErr
	}
	return json.Unmarshal([]byte(resp), v)
}

func postdelete(method string, authId, authToken, path, data string) (int, string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, API_URL+authId+path, strings.NewReader(data))
	if err != nil {
		log.Println(err)
		return -1, "", err
	}

	req.Header.Add("content-type", "application/json")
	req.SetBasicAuth(authId, authToken)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return -1, "", err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return -1, "", err
	}
	log.Printf("contents = %s\n", string(contents))
	return resp.StatusCode, string(contents), nil
}

func postDeleteExpectUnmarshal(method string, authId, authToken, path, data string, v interface{}, expectedCode int) error {
	code, resp, err := postdelete(method, authId, authToken, path, data)
	if err != nil {
		return err
	}
	if code != expectedCode {
		apiErr := ApiError{code, resp}
		return apiErr
	}
	if v == nil {
		return nil
	}
	return json.Unmarshal([]byte(resp), v)
}

func postExpectUnmarshal(authId, authToken, path, data string, v interface{}, expectedCode int) error {
	return postDeleteExpectUnmarshal("POST", authId, authToken, path, data, v, expectedCode)
}

func deleteExpectUnmarshal(authId, authToken, path, data string, v interface{}, expectedCode int) error {
	return postDeleteExpectUnmarshal("DELETE", authId, authToken, path, data, v, expectedCode)
}
