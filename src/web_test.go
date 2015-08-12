package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"io/ioutil"
	"contact"
)

func TestAddContacts(t *testing.T) {
	// t.Error("This is test error")
	ts := httptest.NewServer(http.HandlerFunc(addressBookHandler))
	defer ts.Close()

	resp, _ := http.Post(ts.URL+"/contacts", "application/json", strings.NewReader(`{		
		    "name" : "Kotchaphan",
		    "tel" : "000",
		    "email" : "kotchaphan.m@kkumail.com"
		}`))
	if resp.StatusCode != http.StatusOK {
		t.Error("Expect status code is 200 OK")
	}
	if len(contacts) != 1 {
		t.Error("Expect contacts has 1 contact")
	}
	c := contacts[0]
	if c.Name != "Kotchaphan" {
		t.Error("Expect name is Kotchaphan but got", c.Name)
	}
	if c.Tel != "000" {
		t.Error("Expect tel is 000011 but got", c.Tel)
	}
	if c.Email != "kotchaphan.m@kkumail.com" {
		t.Error("Expect emial is kotchaphan.m@kkumail.com but got", c.Email)
	}
}

func TestGetContact(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(addressBookHandler))
	defer ts.Close()

	contacts = []contact.Contact{
		contact.Contact{
			Name : "Kotchaphan",
			Tel : "000",
			Email : "kotchaphan.m@kkumail.com",
		},
	}
	resp ,_ := http.Get(ts.URL + "/contacts?name=Kotchaphan")
	if resp.StatusCode != http.StatusOK {
		t.Error("Expect status is 200")
	}

	b , _ := ioutil.ReadAll(resp.Body)
	if string(b) != `{"name":"Kotchaphan","tel":"000","email":"kotchaphan.m@kkumail.com"}` + "\n" {
		t.Error(`Expect json {"name" : "Kotchaphan" , "tel" : "000" , "email" : "kotchaphan.m@kkumail.com"} but got ` , string(b))
	}

	resp , _ =  http.Get(ts.URL + "/contacts?name=haha")
	if resp.StatusCode != http.StatusNotFound{
		t.Error("Expect status code is 404")
	}
}
