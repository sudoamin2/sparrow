package sparrow

import (
	"net"
)

func Dial(addr net.Addr) (*Conn, error) {
	// TODO,
	// arp
	c := &Conn{
		state: 1, // waiting for start result
		// client/server ip:port, mac

		startResultCh: make(chan error),
		readCh: make(chan []byte),
	}

	c.SequenceNumber += 1
	if err := s.send(c, Segment{
		Flag:           FlagStart,
		SequenceNumber: c.SequenceNumber,
		ReqID:          0,
		Lenght:         15,
	}); err != nil {
		return nil, err
	}

	return c, <-c.startResultCh
}
