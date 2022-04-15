package perftimer

var globalTimer *Timer

func SetPoint(msg string) {
	if globalTimer == nil {
		globalTimer = New()
	}
	globalTimer.SetPoint(msg)
}

func Report() {
	if globalTimer != nil {
		globalTimer.Report()
	}
}

func Reset() {
	if globalTimer == nil {
		globalTimer = New()
		return
	}
	globalTimer.Reset()
}
