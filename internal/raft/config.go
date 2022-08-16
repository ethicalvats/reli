package raft

import (
	"log"
	"time"
)

type Raft struct {
	Label   string
	Port    string
	Net     *Network
	Timer   *VTimer
	Timeout int
	Peers   []string
}

func (r *Raft) Config() {
	log.Println("[INFO] Raft Config ", r)
	b := make(chan time.Time)
	vt := &VTimer{Bridge: b, Timeout: r.Timeout}
	r.Timer = vt
	bytec := make(chan []byte)
	n := &Network{Port: r.Port, Peers: r.Peers}
	go n.Connect(bytec)
	r.Net = n
	go Begin(r)
	for bc := range bytec {
		log.Println(" [VERBOSE] raft accepted bytes ", string(bc))
		r.Timer.Reset()
	}
}
