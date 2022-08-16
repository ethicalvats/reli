package main

import (
	"log"
	"os"
	"time"

	"github.com/ethicalvats/reli/internal/logutil"
	"github.com/ethicalvats/reli/internal/raft"
)

func main() {
	filter := &logutil.LevelFilter{
		Levels:   []logutil.LogLevel{"VERBOSE", "INFO", "FATAL"},
		MinLevel: logutil.LogLevel("VERBOSE"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	go func() {
		r1 := raft.Raft{Label: "R1", Port: "9000", Timeout: 1000, Peers: []string{"9001"}}
		r1.Config()
	}()

	go func() {
		r2 := raft.Raft{Label: "R2", Port: "9001", Timeout: 2000, Peers: []string{"9000"}}
		r2.Config()
	}()

	for {
		time.Sleep(10 * time.Second)
		log.Println(" [INFO] exiting program ")
		return
	}
}
