package trisoban

import (
	"time"
)

var sw *Stopwatch

// Stopwatch to keep track of your best times
type Stopwatch struct {
	starttime time.Time
	stoptime  time.Time
}

// InitTimer Initialize the timer and start it.
func InitTimer() *Stopwatch {
	sw = new(Stopwatch)
	return sw.Start()
}

// Start start the timer, set starttime value and return it.
func (sw *Stopwatch) Start() *Stopwatch {
	sw.starttime = time.Now()
	return sw
}

// Stop set the value of the stoptime to pinpoint when the timer stopped.
func (sw *Stopwatch) Stop() {
	sw.stoptime = time.Now()
}

// CheckpointTime subtract the start time of the current time and return it as a string.
func (sw *Stopwatch) CheckpointTime() string {
	return time.Now().Sub(sw.starttime).String()
}

// FinishTime subtract the start time of the stoptime to get the total time.
func (sw *Stopwatch) FinishTime() string {
	return sw.stoptime.Sub(sw.starttime).String()
}
