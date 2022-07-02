package sparrow

import (
	"errors"
	"sync"
)

var s sparrow

func init() {
	s = sparrow{
		conns: NewConnList(),
		rwm:   new(sync.RWMutex),

		pendingConns: make(chan Conn, 10),
	}
}

type sparrow struct {
	conns *ConnList
	rwm   *sync.RWMutex

	pendingConns chan Conn
}

func (s *sparrow) receive(segment []byte) {
	// find the conn by the headers
	// remove L2, L3 headers
	// TODO, convert
	sg := Segment{}
	c := Conn{}

	switch sg.Flag {
	case FlagStart:
		// TODO,
		s.conns.PushFront(c)
	case FlagData:
		select {
		case c.readCh <- sg.Payload:
		default:
		}
	case FlagEnd:

	// and the state of the connection was 1
	case FlagStartOK:
		c.startResultCh <- nil
		close(c.startResultCh)
	case FlagStartErr:
		// TODO, remove the conn
		// TODO, parse error and more details
		c.startResultCh <- errors.New("failed to connect")
		close(c.startResultCh)

	case FlagEndOK:
	case FlagEndErr:

	case FlagRetransmit:
	}
}

// TODO,
func (s *sparrow) send(conn *Conn, segment Segment) error {
	// add L2, L3 headers
	// call xdp.enqueue
	return errors.New("unimplemented send")
}
