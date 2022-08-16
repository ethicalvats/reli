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
		if c.Raft.Timer.Done && !c.isLeader {
			log.Println("[INFO] timer expired!")
			c.vote()
		}

		if c.isLeader {
			l := &Leader{Raft: c.Raft}
			l.Active()
		}
	}
}

func (c *Candidate) vote() {
	log.Println(" [INFO] voting ")
	c.isLeader = true
	c.Raft.Timer.Reset()
}
