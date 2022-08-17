package raft

import "log"

type Leader struct {
	Raft *Raft
}

func (l *Leader) Active() {
	log.Println(" [INFO] leader is ", l.Raft.Label)
	for range l.Raft.Timer.Bridge {
		l.heartBeat()
	}
}

func (l *Leader) heartBeat() {
	l.Raft.Net.Send(HeartBeat())
}
