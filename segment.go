package sparrow

type Segment struct {
	Flag           Flag
	SequenceNumber uint64 
	ReqID          uint32
	Lenght         uint16
	Payload        []byte
}
