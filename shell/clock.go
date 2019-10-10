// package main

// import (
// 	"log"
// 	"time"
// 	"os"
// )

// type Scheduler interface {
// 	schedule(f func(), interval time.Duration, done <-chan bool) *time.Ticker
// 	oneSec()
// 	oneMin()
// 	oneHour()
// }
// type SchedulerImpl struct {
// }

// type Clock struct {
// 	f        func()
// 	interval time.Duration
// }

// func main() {
// 	var scheduler Scheduler = &SchedulerImpl{}

// 	os.Setenv("SEC", "tick")
// 	os.Setenv("MIN", "tock")
// 	os.Setenv("HR", "bong")

// 	clocks := []Clock{
// 		{scheduler.oneSec, time.Second},
// 		{scheduler.oneMin, time.Minute},
// 		{scheduler.oneHour, time.Hour},
// 	}

// 	done := make(chan bool)
// 	ts := make([]*time.Ticker, len(clocks))
// 	for i, p := range clocks {
// 		ts[i] = scheduler.schedule(p.f, p.interval, done)
// 	}

// 	time.Sleep(10800 * time.Second)
// 	close(done)

// 	for _, t := range ts {
// 		t.Stop()
// 	}
// }

// func (t *SchedulerImpl)schedule(f func(), interval time.Duration,done <-chan bool) *time.Ticker {
// 	ticker := time.NewTicker(interval)
// 	counter := 1
// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				if(counter<=59) {
// 					f()
// 					counter++;
// 				}else{
// 					counter=1;
// 				}
// 			case <-done:
// 				return
// 			}
// 		}
// 	}()
// 	return ticker
// }

// func (t *SchedulerImpl)oneSec() {
// 	//user can to alter any of the printed values while the program is running by setting the environment variable
// 	log.Println(os.Getenv("SEC"))
// }

// func (t *SchedulerImpl)oneMin() {
// 	//user can to alter any of the printed values while the program is running by setting the environment variable
// 	log.Println(os.Getenv("MIN"))
// }
// func (t *SchedulerImpl) oneHour() {
// 	//user can to alter any of the printed values while the program is running by setting the environment variable
// 	log.Println(os.Getenv("HR"))
// }
