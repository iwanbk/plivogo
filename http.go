package plivogo

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Get(authId, authToken, path, data string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", API_URL+authId+path, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	req.SetBasicAuth(authId, authToken)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(contents), nil
}

func Post(authId, authToken, path, data string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", API_URL+authId+path, strings.NewReader(data))
	if err != nil {
		log.Println(err)
		return "", err
	}

	req.Header.Add("content-type", "application/json")
	req.SetBasicAuth(authId, authToken)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(contents), nil
}
