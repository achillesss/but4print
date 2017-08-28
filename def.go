package but

type Buter interface {
	Color(color colorName, isBackgroundColor bool) Buter
	Show(...outPutSet) Buter
	Print()
}

type colorName int
type outPutSet int

const (
	// 0(黑)、1(红)、2(绿)、 3(黄)、4(蓝)、5(洋红)、6(青)、7(白)

	COLOR_BLACK colorName = iota
	COLOR_RED
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_BLUE
	COLOR_MAGENTA
	COLOR_CYAN
	COLOR_WHITE

	// 显示：0(默认)、1(粗体/高亮)、22(非粗体)、4(单条下划线)、24(无下划线)、5(闪烁)、25(无闪烁)、7(反显、翻转前景色和背景色)、27(无反显)

	SET_DEFAULT      outPutSet = 0
	SET_BOLD         outPutSet = 1
	SET_UNDERLINE    outPutSet = 4
	SET_BLINK        outPutSet = 5
	SET_REVERSAL     outPutSet = 7
	SET_UNBOLD       outPutSet = 22
	SET_NO_UNDERLINE outPutSet = 24
	SET_NO_BLINK     outPutSet = 25
	SET_UNREVERSAL   outPutSet = 27
	COLOR_FOREGROUND outPutSet = 30
	COLOR_BACKGROUND outPutSet = 40

	PRINTER_FORMAT  = "\033[{{params}}m"
	PRINTER_DEFAULT = "\033[0m"
)

type printer struct {
	p      func(format string, args ...interface{}) (n int, err error)
	format string
	args   []interface{}
	prefix string
	sufix  string
}
