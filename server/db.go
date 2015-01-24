package server

// vim:sw=4:ts=4

import (
	"errors"
	"time"

	"appengine"
	"appengine/datastore"
)

type Uid string

type User struct {
	Uid
}

type Entry struct {
	Time       time.Time
	WasIdle    bool
	Idle       time.Duration
	App, Title string
}

// func withdraw(c appengine.Context, id string, amount, other int) error {
// 	key := datastore.NewKey(c, "BankAccount", "", 1, nil)
// 	account := BankAccount{}
// 	err := datastore.Get(c, key, &account)
// 	if err != nil {
// 		return errors.New("No such account")
// 	}
// 	if account.Balance >= amount {
// 		account.Balance -= amount
// 	} else {
// 		return errors.New("insufficient funds")
// 	}
// 	return nil
// }

func AddEntry(c appengine.Context, u *User, e *Entry) (*datastore.Key, error) {
	user := datastore.NewKey(c, "User", string(u.Uid), 0, nil)
	entry := datastore.NewKey(c, "Entry", "", 0, user)
	var err error
	if entry, err = datastore.Put(c, entry, e); err != nil {
		return nil, errors.New("Could not put the entry")
	}
	return entry, nil
}
