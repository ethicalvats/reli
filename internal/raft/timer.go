package raft

import (
	"log"
	"time"
)

type VTimer struct {
	Done    bool
	Current time.Time
	Bridge  chan time.Time
	dc      chan bool
	Timeout int
}

func (t *VTimer) Run() {
	log.Println("[VERBOSE] timer run")

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	t.dc = make(chan bool)
	defer close(t.dc)
	t.Reset()

	for {
		select {
		case d := <-t.dc:
			if d {
				t.Done = true
			}
		case tc := <-ticker.C:
			t.Bridge <- tc
		}
	}
}

func (t *VTimer) Reset() {
	go func() {
		log.Println(" [VERBOSE] timer reset ")
		t.Done = false
		time.Sleep(time.Duration(t.Timeout) * time.Millisecond)
		t.dc <- true
	}()
}
