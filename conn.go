package sparrow

import "errors"

type Conn struct {
	// TODO, client or server
	xdp xdpConn
}

func (s *Conn) Write(b []byte) (n int, err error) {
	// client.write
	return 0, errors.New("unimplemented conn.Write")
}

func (s *Conn) Read(b []byte) (n int, err error) {
	// client.read
	return 0, errors.New("unimplemented conn.read")
}

func (s *Conn) Close() error {
	return errors.New("unimplemented conn.close")
}
