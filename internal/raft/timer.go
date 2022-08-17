package raft

import (
	"log"
	"time"
)

type VTimer struct {
	Done      bool
	Current   time.Time
	Bridge    chan time.Time
	Timeout   int
	RaftLabel string
	nextT     time.Time
}

func (t *VTimer) Run() {
	log.Println("[VERBOSE] timer run", t.RaftLabel)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	t.Reset()

	for {
		select {
		case tc := <-ticker.C:
			t.Bridge <- tc
			u := time.Until(t.nextT)
			if u < time.Duration(time.Millisecond) {
				t.Done = true
			} else {
				t.Done = false
			}
		}
	}
}

func (t *VTimer) Reset() {
	tn := time.Now()
	t.nextT = tn.Add(time.Duration(t.Timeout) * time.Millisecond)
}

func (t *VTimer) IsDone() bool {
	return t.Done
}
