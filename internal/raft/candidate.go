package raft

import (
	"log"
	"time"
)

type Candidate struct {
	Current  time.Time
	Raft     *Raft
	isLeader bool
}

func (c *Candidate) Wait() {
	for range c.Raft.Timer.Bridge {
		if c.Raft.Timer.IsDone() && !c.isLeader {
			log.Println("[INFO] timer expired!", c.Raft.Label)
			c.vote()
		}

		if c.isLeader {
			l := &Leader{Raft: c.Raft}
			l.Active()
		}
	}
}

func (c *Candidate) vote() {
	log.Println(" [INFO] voting ", c.Raft.Label)
	c.isLeader = true
	c.Raft.Timer.Reset()
}
