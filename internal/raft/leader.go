package raft

type Leader struct {
	Raft *Raft
}

func (l *Leader) Active() {
	for range l.Raft.Timer.Bridge {
		l.heartBeat()
	}
}

func (l *Leader) heartBeat() {
	l.Raft.Net.Send(HeartBeat())
}
