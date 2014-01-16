package plivogo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ApiError struct {
	StatusCode int
	RawError   string
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
