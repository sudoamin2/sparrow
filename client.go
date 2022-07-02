package sparrow

import (
	"errors"
	"net"
)

type client struct {
	server net.Addr
	addr   net.Addr
}

// try to init xdp or use an intialized global xdp
// find the mac of server
// send start packet

func (c *client) connect() (*Conn, error) {
	// conn.Client = c
	return nil, errors.New("unimplemented client.connect")
}

func findClientAddr() net.Addr {
	return nil
}
