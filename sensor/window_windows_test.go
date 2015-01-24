package sensor

// vim:ts=4

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWindowTitle(t *testing.T) {
	// This is the empirically determined handle of a mintty in $HOME.  This
	// would have to be updated every time said mintty is restarted.
	testHandle = 3278308
	expected := "~"

	Convey("Get the current window title", t, func() {
		title := WindowTitle()
		// expected := "lmcdell || ~/src/goget/src/github.com/gophergala/panopticon/sensor"
		// expected := "GoConvey - Google Chrome"

		// Running using goconvey, with the editor focused.
		// expected := "window_windows_test.go (~\\src\\goget\\src\\github.com\\gophergala\\panopticon\\sensor) ((1) of 2) - GVIM"
		Convey("Titles should match", func() {
			So(title, ShouldEqual, expected)
		})
		Convey("Lengths should match", func() {
			So(len(title), ShouldEqual, len(expected))
		})
	})
}
