package raft

import (
	"log"
	"net"
)

type Network struct {
	Port     string
	listener net.Listener
	Peers    []string
}

func (n *Network) Connect(bc chan []byte) {
	listener, err := net.Listen("tcp", "localhost:"+n.Port)
	if err != nil {
		log.Println("[FATAL] TCP listen err ", err)
		return
	}
	log.Println(" [INFO] listening on port ", n.Port)
	defer listener.Close()

	n.listener = listener
	n.listen(bc)
}

func (n *Network) Send(p []byte) {
	for _, peer := range n.Peers {
		conn, err := net.Dial("tcp", "localhost:"+peer)
		if err != nil {
			log.Println(" [FATAL] tcp dial err ", err)
			return
		}
		log.Print("[VERBOSE] sending data to peer ", string(p), " ", peer)
		conn.Write(p)
		conn.Close()
	}
}

func (n *Network) listen(bc chan []byte) {

	for {
		conn, err := n.listener.Accept()
		if err != nil {
			log.Println("[INFO] conn err ", err)
		}
		go n.handleIncomingRequest(conn, bc)
	}
}

func (n *Network) handleIncomingRequest(c net.Conn, bc chan []byte) {
	b := make([]byte, 1024)
	_, err := c.Read(b)
	if err != nil {
		log.Println("[INFO] err reading the body ", err)
		return
	}
	log.Println(" [VERBOSE] received bytes by port ", string(b), " ", n.Port)
	bc <- b
	c.Write([]byte("Success"))
	c.Close()
}
