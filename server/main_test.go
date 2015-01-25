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

func TestRoot(t *testing.T) {
	jsonSampleEntry, err := json.Marshal(&sampleEntry)
	if err != nil {
		t.Fatalf("Couldn't marshal sampleEntry: %v", err)
	}
	req1, err := inst.NewRequest("PUT", "/add_entry", bytes.NewBuffer(jsonSampleEntry))
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	// c1, _ = aetest.NewContext(req1)
	// _ = c1

	w := httptest.NewRecorder()
	Root(w, req1)
	fmt.Printf("%d - %s", w.Code, w.Body.String())
}

func TestAddEntry(t *testing.T) {
	log.Printf("TestAddEntry()")
	req, err := inst.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatalf("Couldn't create a GET request: %v", err)
	}

	c := appengine.NewContext(req)
	entry, err := AddEntry(c, sampleUser, &sampleEntry)
	if err != nil {
		t.Fatal(err)
	}

	// c := appengine.NewContext(req)
	var e Entry
	if err := datastore.Get(c, entry, &e); err != nil {
		t.Fatal(err)
	}
	if e.Time != now || e.WasIdle != true || e.Idle != idle || e.App != "Chrome" || e.Title != "localhost" {
		t.Fatal(errors.New(fmt.Sprintf("Wrong entry returned: stored this: \n%v\ngot this: \n%v", sampleEntry, e)))
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
