package sparrow

type Flag uint8

const (
	_ = iota

	// Client requests
	FlagStart = Flag(iota)
	FlagData  = Flag(iota)
	FlagEnd   = Flag(iota)

	// Server responses
	FlagStartOK  = Flag(iota)
	FlagStartErr = Flag(iota)
	FlagEndOK    = Flag(iota)
	FlagEndErr   = Flag(iota)
	// TODO
	FlagRetransmit = Flag(iota)
)
