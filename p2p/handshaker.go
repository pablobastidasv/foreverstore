package p2p

// HandshakeFunc... ?
type HandshakeFunc func(Peer) error

// NOPHandshakeFunc... ?
func NOPHandshakeFunc(Peer) error {return nil}
