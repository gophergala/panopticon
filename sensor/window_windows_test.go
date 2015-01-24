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
		Convey("Titles should match", func() {
			So(title, ShouldEqual, expected)
		})
		Convey("Lengths should match", func() {
			So(len(title), ShouldEqual, len(expected))
		})
	})
}
