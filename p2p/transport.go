package p2p

// Peer is an interface that represents the remote node.
type Peer interface {
	ID() string
	Address() string
}

// Transport is anything that handles the communication
// between the nodes in the network. This can be of the
// form of a TCP connection, a websocket, or any other
// form of communication.
type Transport interface {
	ListenAndServe() error
	HandlePeer(Peer) error
}
