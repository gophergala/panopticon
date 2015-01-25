package sensor

// vim:ts=4

import (
	"log"
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

// Not sure how to test this.  I guess, just make sure it's non-zero?
func TestGetLastInputInfo(t *testing.T) {
	Convey("Get tick count of the most recent keyboard activity", t, func() {
		if ticks, err := GetLastInputInfo(); err != nil {
			So(err, ShouldBeNil) // obviously an automatic failure
		} else {
			Convey("tick count is non-zero", func() {
				log.Printf("GetLastInputInfo is %v", ticks)
				So(ticks, ShouldBeGreaterThan, 0)
			})
		}
	})
}

// Not sure how to test this, either.  I guess, just make sure it's non-zero,
// too?  Which of course will fail if you run the test with the mouse in the
// upper left corner of the screen.
func TestGetMouseMovePointsEx(t *testing.T) {
	Convey("Get the mouse position", t, func() {
		if pos, err := GetCursorPos(); err != nil {
			So(err, ShouldBeNil) // obviously an automatic failure
		} else {
			// log.Printf("Mouse pos is %v", pos)
			Convey("Mouse pos x is non-zero", func() {
				So(pos.X, ShouldBeGreaterThan, 0)
			})
			Convey("Mouse pos y is non-zero", func() {
				So(pos.Y, ShouldBeGreaterThan, 0)
			})
		}
	})
}
