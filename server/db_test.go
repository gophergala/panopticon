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

// func TestWithdrawLowBal(t *testing.T) {
// 	c, err := aetest.NewContext(nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer c.Close()
// 	key := datastore.NewKey(c, "BankAccount", "", 1, nil)
// 	if _, err := datastore.Put(c, key, &BankAccount{100}); err != nil {
// 		t.Fatal(err)
// 	}
//
// 	err = withdraw(c, "myid", 128, 0)
// 	if err == nil || err.Error() != "insufficient funds" {
// 		t.Errorf("Error: %v; want insufficient funds error", err)
// 	}
//
// 	b := BankAccount{}
// 	if err := datastore.Get(c, key, &b); err != nil {
// 		t.Fatal(err)
// 	}
// 	if bal, want := b.Balance, 100; bal != want {
// 		t.Errorf("Balance %d, want %d", bal, want)
// 	}
// }

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
