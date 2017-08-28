package but

import (
	"fmt"
	"strings"
)

// but is short for beautiful

func combineColor(color colorName, isBackgroundColor bool) outPutSet {
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

func (x *printer) Color(color colorName, isBackgroundColor bool) Buter {
	c := combineColor(color, isBackgroundColor)
	return x.updateSets(c)
}

func (x *printer) Show(sets ...outPutSet) Buter {
	for _, set := range sets {
		x.updateSets(set)
	}
	return x
}

func (x *printer) Print() {
	f := x.format
	var returns string

	for strings.HasSuffix(f, "\n") {
		f = strings.TrimSuffix(f, "\n")
		returns += "\n"
	}

	f = x.prefix + f + x.sufix + returns

	x.p(f, x.args...)
}

func NewButer(format string, args ...interface{}) Buter {
	return &printer{p: fmt.Printf, format: format, args: args}
}
