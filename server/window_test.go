package main

// vim:ts=4

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWindowTitle(t *testing.T) {
	Convey("Get the current window title", t, func() {
		title := WindowTitle()
		expected := "lmcdell || ~/src/goget/src/github.com/gophergala/panopticon/server"
		Convey("titles should match", func() {
			So(title, ShouldEqual, expected)
		})
		Convey("lengths should match", func() {
			So(len(title), ShouldEqual, len(expected))
		})
	})
}
