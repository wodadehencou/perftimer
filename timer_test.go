package perftimer_test

import (
	"testing"
	"time"

	"github.com/wodadehencou/perftimer"
)

func Test_Timer(t *testing.T) {
	timer := perftimer.New()

	time.Sleep(time.Second)
	timer.SetPoint("step 1")
	time.Sleep(2 * time.Second)
	timer.SetPoint("step 2")
	for i := 0; i < 1024*1024; i++ {
		_ = make([]byte, 1024)
	}
	timer.SetPoint("new 1M 1k buffer")

	timer.Report()

	timer.Reset()
	timer.SetPoint("again start")
	timer.Report()
}

func Test_Global(t *testing.T) {
	perftimer.Reset()
	time.Sleep(time.Second)
	perftimer.SetPoint("step 1")
	time.Sleep(2 * time.Second)
	perftimer.SetPoint("step 2")
	for i := 0; i < 1024*1024; i++ {
		_ = make([]byte, 1024)
	}
	perftimer.SetPoint("new 1M 1k buffer")

	perftimer.Report()

	perftimer.Reset()
	perftimer.SetPoint("again start")
	perftimer.Report()
}

func Test_TimerFrom(t *testing.T) {
	timer := perftimer.New()

	time.Sleep(time.Second)
	timer.SetPoint("step 1")
	time.Sleep(2 * time.Second)
	timer.SetPoint("step 2")
	for i := 0; i < 1024*1024; i++ {
		_ = make([]byte, 1024)
	}
	timer.SetPoint("new 1M 1k buffer")
	timer.SetPointFrom("step 1", "from step 1")
	timer.SetPointFrom(perftimer.StartMsg, "total")

	timer.Report()

	timer.Reset()
	timer.SetPoint("again start")
	timer.Report()
}
