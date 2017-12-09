package but

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// but is short for beautiful

func combineColor(color ColorName, isBackgroundColor bool) outPutSet {
	if color >= COLOR_BLACK && color <= COLOR_WHITE {
		if isBackgroundColor {
			return COLOR_BACKGROUND + outPutSet(color)
		}
		return COLOR_FOREGROUND + outPutSet(color)
	}
	return 0
}

func (x *printer) updateSets(b outPutSet) *printer {

	pre := "\033["
	su := "m"
	if strings.HasPrefix(x.prefix, pre) && strings.HasSuffix(x.prefix, su) {
		s := strings.TrimPrefix(x.prefix, pre)
		s = strings.TrimSuffix(s, su)
		body := strings.Split(s, ";")
		body = append(body, fmt.Sprintf("%d", b))
		x.prefix = strings.Replace(PRINTER_FORMAT, "{{params}}", strings.Join(body, ";"), -1)
	} else {
		x.prefix = strings.Replace(PRINTER_FORMAT, "{{params}}", fmt.Sprintf("%d", b), -1)
	}
	x.updateSufix()

	return x
}

func (x *printer) updateSufix() {
	if x.sufix != PRINTER_DEFAULT {
		x.sufix = PRINTER_DEFAULT
	}
}

func (x *printer) Color(color ColorName, isBackgroundColor bool) Buter {
	c := combineColor(color, isBackgroundColor)
	if c > 0 {
		return x.updateSets(c)
	}
	return x
}

func (x *printer) Show(sets ...outPutSet) Buter {
	for _, set := range sets {
		x.updateSets(set)
	}
	return x
}

func (x *printer) Print() {
	f, args := x.formating()
	x.p(x.w, f, args...)
}

func (x *printer) formating() (formation string, args []interface{}) {
	f := x.format
	var returns string

	for strings.HasSuffix(f, "\n") {
		f = strings.TrimSuffix(f, "\n")
		returns += "\n"
	}

	f = x.prefix + f + x.sufix + returns

	return f, x.args
}

func (x *printer) String() string {
	f, args := x.formating()
	return fmt.Sprintf(f, args...)
}

func NewButer(w io.Writer, format string, args ...interface{}) Buter {
	if w == nil {
		w = os.Stdout
	}
	return &printer{w: w, p: func(w io.Writer, format string, args ...interface{}) { fmt.Fprintf(w, format, args...) }, format: format, args: args}
}
