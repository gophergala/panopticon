package server

// vim:sw=4:ts=4

import (
	"testing"

	"appengine"
	"appengine/aetest"
)

func TestMyFunction(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req1, err := inst.NewRequest("POST", "/title", nil)
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	c1 := appengine.NewContext(req1)

	// What now?
	_ = c1
}
