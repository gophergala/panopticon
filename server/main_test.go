package server

// vim:sw=4:ts=4

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"appengine"
	"appengine/aetest"
	"appengine/datastore"
	"appengine/user"
)

var now = time.Now().Round(time.Millisecond)
var sampleUser = "lmc"
var idle = 10 * time.Second
var sampleEntry = Entry{
	Time:    now,
	WasIdle: true,
	Idle:    idle,
	App:     "Chrome",
	Title:   "localhost"}

var inst aetest.Instance

func init() {
	var err error
	log.Printf("init()")
	inst, err = aetest.NewInstance(nil)
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}
}

func TestAddEntry(t *testing.T) {
	log.Printf("TestAddEntry()")
	req, err := inst.NewRequest("PUT", "http://i-dont-remember.appspot.com/api/v1/add_entry", nil)
	if err != nil {
		t.Fatalf("Couldn't create a GET request: %v", err)
	}

	c := appengine.NewContext(req)
	entry, err := AddEntry(c, sampleUser, &sampleEntry)
	if err != nil {
		t.Fatal(err)
	}

	var e Entry
	if err := datastore.Get(c, entry, &e); err != nil {
		t.Fatal(err)
	}
	if e.Time != now || e.WasIdle != true || e.Idle != idle || e.App != "Chrome" || e.Title != "localhost" {
		t.Fatal(errors.New(fmt.Sprintf("Wrong entry returned: stored this: \n%v\ngot this: \n%v", sampleEntry, e)))
	}
}

func TestApiAddEntry(t *testing.T) {
	jsonSampleEntry, err := json.Marshal(&sampleEntry)
	if err != nil {
		t.Fatalf("Couldn't marshal sampleEntry: %v", err)
	}
	req, err := inst.NewRequest("PUT",
		"http://i-dont-remember.appspot.com/api/v1/add_entry",
		bytes.NewBuffer(jsonSampleEntry))
	if err != nil {
		t.Fatalf("Failed to create req: %v", err)
	}
	aetest.Login(&user.User{Email: "larry@theclapp.org"}, req)

	w := httptest.NewRecorder()
	apiAddEntry(w, req)
	// fmt.Printf("TestApiAddEntry: %d - %s - %v\n", w.Code, w.Body.String(), w.Header())
	if w.Code != 200 {
		t.Fatalf("w.Code expected to be 200, not %v", w.Code)
	}
	if w.Header().Get("Location") != "agtkZXZ-dGVzdGFwcHIuCxIEVXNlciISbGFycnlAdGhlY2xhcHAub3JnDAsSBUVudHJ5GICAgICAgIAJDA" {
		t.Fatalf("Location header wrong: %v", w.Header().Get("Location"))
	}
}

// func TestGetEntry(t *testing.T) {
// 	c, err := aetest.NewContext(nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer c.Close()
// }

func TestShutdown(t *testing.T) {
	inst.Close()
}
