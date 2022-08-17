package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ethicalvats/reli/internal/logutil"
	"github.com/ethicalvats/reli/internal/raft"
)

func main() {
	filter := &logutil.LevelFilter{
		Levels:   []logutil.LogLevel{"VERBOSE", "INFO", "FATAL"},
		MinLevel: logutil.LogLevel("INFO"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	ts := []int{500, 600, 700, 800, 900, 1000}

	go func() {
		r1 := raft.Raft{Label: "R1", Port: "9000", Timeout: ts[rand.Intn(len(ts))], Peers: []string{"9001", "9002"}}
		r1.Config()
	}()

	go func() {
		r2 := raft.Raft{Label: "R2", Port: "9001", Timeout: ts[rand.Intn(len(ts))], Peers: []string{"9000", "9002"}}
		r2.Config()
	}()

	go func() {
		r3 := raft.Raft{Label: "R3", Port: "9002", Timeout: ts[rand.Intn(len(ts))], Peers: []string{"9000", "9001"}}
		r3.Config()
	}()

	for {
		time.Sleep(10 * time.Second)
		log.Println(" [INFO] exiting program ")
		return
	}
}
