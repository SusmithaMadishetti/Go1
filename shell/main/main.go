package main

import (
	"os"
	"time"
)

func main() {
	var scheduler Scheduler = &SchedulerImpl{}

	os.Setenv("SEC", "tick")
	os.Setenv("MIN", "tock")
	os.Setenv("HR", "bong")

	clocks := []Clock{
		{scheduler.oneSec, time.Second},
		{scheduler.oneMin, time.Minute},
		{scheduler.oneHour, time.Hour},
	}

	done := make(chan bool)
	ts := make([]*time.Ticker, len(clocks))
	for i, p := range clocks {
		ts[i] = scheduler.schedule(p.f, p.interval, done)
	}

	time.Sleep(10800 * time.Second)
	close(done)

	for _, t := range ts {
		t.Stop()
	}
}
