package server

// vim:sw=4:ts=4

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"appengine/aetest"
	"appengine/datastore"
)

func TestAddEntry(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	now := time.Now()
	idle, _ := time.ParseDuration("10s")
	entry, err := AddEntry(c, &User{"lmc"}, &Entry{Time: now,
		WasIdle: true, Idle: idle, App: "Chrome", Title: "localhost"})
	if err != nil {
		t.Fatal(err)
	}

	var e Entry
	if err := datastore.Get(c, entry, &e); err != nil {
		t.Fatal(err)
	}
	if e.Time != now || e.WasIdle != true || e.Idle != idle || e.App != "Chrome" || e.Title != "localhost" {
		t.Fatal(errors.New(fmt.Sprintf("Wrong entry returned: %v", e)))
	}
}
