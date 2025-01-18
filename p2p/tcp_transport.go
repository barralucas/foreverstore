package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAdrresss string
	listener       net.Listener

	mu             sync.RWMutex
	peers          map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAdrresss: listenAddr,
		peers:          make(map[net.Addr]Peer),
	}
}