package raft

func Begin(r *Raft) {

	go r.Timer.Run()
	c := &Candidate{Raft: r, isLeader: false}
	c.Wait()
}
