package perftimer

import (
	"testing"
	"time"
)

func Test_Timer(t *testing.T) {
	timer := New()

	time.Sleep(time.Second)
	timer.SetPoint("step 1")
	time.Sleep(2 * time.Second)
	timer.SetPoint("step 2")
	for i := 0; i < 1024*1024; i++ {
		_ = make([]byte, 1024)
	}
	timer.SetPoint("new 1M 1k buffer")

	timer.Report()
}
