package but

import (
	"testing"
	"time"
)

func TestButer(t *testing.T) {
	format0 := "%s\t\t"
	format1 := "Now:\t%s\n\n\n"
	arg0 := "Hello, world!"

	arg1 := func() string {
		return time.Now().String()
	}

	NewButer(format0, arg0).
		Color(COLOR_CYAN, false).
		Show(SET_BOLD).
		Print()

	NewButer(format1, arg1()).
		Color(COLOR_RED, true).
		Show(SET_REVERSAL, SET_UNDERLINE).
		Print()
}
