package perftimer

import (
	"strings"
	"time"
)

var (
	Arrow       = ">"
	ArrowLength = 16
	TimeFormat  = "04:05.999999"
)

func (t *Timer) msgWidth() int {
	w := 0
	for _, p := range t.points {
		l := len(p.msg)
		if l > w {
			w = l
		}
	}
	w = (w + 3) / 4

	return w
}

func arrow(max, cur time.Duration) string {
	n := (cur * 16) / max
	return strings.Repeat(Arrow, int(n))
}
