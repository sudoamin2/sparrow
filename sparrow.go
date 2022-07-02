package sparrow

import "net"

func Dial(addr net.Addr) (*Conn, error) {
	c := &client{
		server: addr,
		addr:   findClientAddr(),
	}
	return c.connect()
}
