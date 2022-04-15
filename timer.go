package perftimer

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Point struct {
	msg  string
	t    time.Time
	dur  time.Duration
	next time.Time
}

type Timer struct {
	points []Point
	maxDur time.Duration
}

func New() *Timer {
	t := new(Timer)
	t.Reset()
	return t
	return t
}

func (t *Timer) SetPoint(msg string) {
	now := time.Now()
	var dur time.Duration
	if len(t.points) == 0 {
		dur = 0
	} else {
		dur = now.Sub(t.points[len(t.points)-1].t)
	}
	t.points = append(t.points, Point{
		msg: msg,
		t:   now,
		dur: dur,
	})
	if dur > t.maxDur {
		t.maxDur = dur
	}
	t.points[len(t.points)-1].next = time.Now()
}

// func (t *Timer) Report() {
// 	msgWidth := t.msgWidth()
// 	msgFormat := "%-" + strconv.Itoa(msgWidth) + "s"
// 	total := time.Duration(0)

// 	for _, p := range t.points {
// 		fmt.Printf(msgFormat, p.msg)
// 		fmt.Printf("%-32s", p.dur.String())
// 		total += p.dur
// 		fmt.Printf("%-32s", total.String())
// 		fmt.Print("\n")
// 	}
// }
func (t *Timer) Report() {

	tab := table.NewWriter()
	tab.AppendHeader(table.Row{
		"#",
		"step",
		// "time",
		"duration",
		"total",
	})

	total := time.Duration(0)
	for i, p := range t.points[1:] {
		total += p.dur
		tab.AppendRow(table.Row{
			i,
			p.msg,
			// p.t.Format(TimeFormat),
			p.dur.String(),
			total.String(),
			arrow(t.maxDur, p.dur),
		})
	}

	fmt.Println(tab.Render())
}

func (t *Timer) Reset() {
	t.points = make([]Point, 0)
	t.maxDur = 0
	t.SetPoint("start")
}
