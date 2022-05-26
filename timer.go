package perftimer

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

const StartMsg = "start"

type Point struct {
	msg string
	t   time.Time
	dur time.Duration
}

type Timer struct {
	points    *sync.Map
	lastPoint string
	maxDur    time.Duration
}

func New() *Timer {
	t := new(Timer)
	t.Reset()
	return t
}

func (t *Timer) SetPoint(msg string) {
	now := time.Now()
	var dur time.Duration
	last, loaded := t.points.Load(t.lastPoint)
	if !loaded {
		dur = 0
	} else {
		lastPoint := last.(*Point)
		dur = now.Sub(lastPoint.t)
	}
	t.points.Store(msg, &Point{
		msg: msg,
		t:   now,
		dur: dur,
	})
	if dur > t.maxDur {
		t.maxDur = dur
	}
	t.lastPoint = msg
}

func (t *Timer) SetPointFrom(from, msg string) {
	now := time.Now()
	var dur time.Duration
	last, loaded := t.points.Load(from)
	if !loaded {
		last, _ = t.points.Load(StartMsg)
	}
	lastPoint := last.(*Point)
	dur = now.Sub(lastPoint.t)
	t.points.Store(msg, &Point{
		msg: msg,
		t:   now,
		dur: dur,
	})
	if dur > t.maxDur {
		t.maxDur = dur
	}
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
		"time",
		"duration",
		// "total",
	})

	points := make(PointList, 0)
	t.points.Range(func(key, value any) bool {
		points = append(points, value.(*Point))
		return true
	})

	sort.Sort(points)

	total := time.Duration(0)
	for i, p := range points {
		total += p.dur
		tab.AppendRow(table.Row{
			i,
			p.msg,
			p.t.Format(TimeFormat),
			p.dur.String(),
			// total.String(),
			arrow(t.maxDur, p.dur),
		})
	}

	fmt.Println(tab.Render())
}

func (t *Timer) Reset() {
	t.points = new(sync.Map)
	t.maxDur = 0
	t.lastPoint = StartMsg
	t.points.Store(StartMsg, &Point{
		msg: StartMsg,
		t:   time.Now(),
		dur: 0,
	})
}

type PointList []*Point

func (p PointList) Len() int           { return len(p) }
func (p PointList) Less(i, j int) bool { return p[i].t.Before(p[j].t) }
func (p PointList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
