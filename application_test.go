package plivogo

import (
	"testing"
)

func TestGetApplication(t *testing.T) {
	p := NewPlivo(authId, authToken)
	app, err := p.Application.Get("13066057776063802")
	if err != nil {
		t.Fatal("Get Application failed:", err.Error())
	}
	if app.Name != "Demo Speak" {
		t.Fatal("Invalid app name :" + app.Name)
	}
}

func TestGetApplicationList(t *testing.T) {
	p := NewPlivo(authId, authToken)
	params, _ := NewParams()
	appList, err := p.Application.GetList(params)
	if err != nil {
		t.Error("Get Application list failed:", err.Error())
	}
	t.Logf("Number of application = %d\n", len(appList.List))
	for _, app := range appList.List {
		t.Log("app name = " + app.Name)
	}
}

func TestApplicationCrud(t *testing.T) {
	plivo := NewPlivo(authId, authToken)

	//create test
	params, _ := NewParams()
	params.Set("answer_url", "http://google.com")
	params.Set("app_name", "monster")

	appId, err := plivo.Application.Create(params)
	if err != nil {
		t.Fatal("Application create failed :" + err.Error())
	}

	//edit
	newAnswerUrl := "http://answer.com"
	params, _ = NewParams()
	params.Set("answer_url", newAnswerUrl)
	err = plivo.Application.Modify(appId, params)
	if err != nil {
		t.Error("Application modification failed :" + err.Error())
	}

	//verify our change
	app, err := plivo.Application.Get(appId)
	if err != nil {
		t.Error("Application get failed :" + err.Error())
	}
	if app.AnswerUrl != newAnswerUrl {
		t.Error("Application edit failed. variable not changed")
	}
	//delete test
	err = plivo.Application.Delete(appId)
	if err != nil {
		t.Error("Application delete failed :" + err.Error())
	}
}
