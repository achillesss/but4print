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

	NewButer(
		nil,
		format0,
		arg0,
	).
		Color(COLOR_CYAN, false).
		Show(SET_BOLD).
		Print()

	NewButer(nil, format1, arg1()).
		Color(COLOR_RED, true).Color(COLOR_BLACK, false).
		Show(SET_UNDERLINE).
		Print()
}
