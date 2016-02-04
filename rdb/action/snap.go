package action

import (
	"sync"
	"time"
)

//A ActionSnap is telemetry
type actionSnap interface {
	Started() time.Time
	Finished() time.Time
	Done(err error)

	Start()
	Error() error
}

//A actionSnapDefault - telemetry perform actions for each database
type actionSnapDefault struct {
	mu       *sync.Mutex
	started  time.Time
	finished time.Time
	done     bool
	doneErr  error
}

//newSnap returns new ActionSnap
func newSnap() actionSnap {
	return &actionSnapDefault{
		mu: &sync.Mutex{},
	}
}

//Started returns a time when the launch of the action
func (a *actionSnapDefault) Started() time.Time {
	return a.started
}

//Finished returns the time when the action is complete
func (a *actionSnapDefault) Finished() time.Time {
	return a.finished
}

//Start it marks the start of the action
func (a *actionSnapDefault) Start() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.started = time.Now()
}

//Done it marks the completion of the action
func (a *actionSnapDefault) Done(err error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.done {
		return
	}
	a.done = true
	a.doneErr = err
	a.finished = time.Now()
}

//Error returns an error at the end or nil
func (a *actionSnapDefault) Error() error {
	return a.doneErr
}
