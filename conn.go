package sparrow

import "errors"

type Conn struct {
	state          uint8
	SequenceNumber uint64
	startResultCh  chan error
	readCh         chan []byte
}

func (c *Conn) Write(b []byte) (n int, err error) {
	return len(b), s.send(c, Segment{
		// TODO, header
		Payload: b,
	})
}

func (c *Conn) Read(b []byte) (n int, err error) {
	// TODO
	b = <-c.readCh
	return len(b), nil
}

func (c *Conn) Close() error {
	return errors.New("unimplemented conn.close")
}
