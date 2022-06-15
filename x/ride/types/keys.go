package types

import "time"

const (
	// ModuleName defines the module name
	ModuleName = "ride"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ride"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	NextRideKey = "NextRide-value-"
)

const (
	RequestRideEventKey       = "NewRideRequest"
	RequestRideEventPassenger = "Passenger"
	RequestRideEventIndex     = "Index"
	RequestRideStartLocation  = "StartLocation"
)

const (
	AcceptRideEventKey     = "RideAccepted"
	AcceptRideEventDriver  = "Driver"
	AcceptRideEventIdValue = "IdValue"
)

const (
	FinishRideEventKey = "RideFinished"
)

const (
	TimeFormat = "2006-01-02 15:04:05.999999999 +0000 UTC"
	// TODO: Make this hardcoded value configurable.
	DeadlinePeriod = 5 * time.Minute
)

const (
	NoFifoIdKey = "-1"
)
