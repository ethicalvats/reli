package raft

import (
	"time"
)

func HeartBeat() []byte {
	return []byte("[PULSE] " + time.Now().String())
}
