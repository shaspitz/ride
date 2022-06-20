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

	StoredRideMutualStake = "MutualStake"
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
	DeadlinePeriod = 2 * time.Minute
)

const (
	NoFifoIdKey = "-1"
)

// Event keys representing the expiration events of stored rides in three different cases.
// 1. A requested but unaccepted ride has auto-expired, with no driver assigned yet.
// 2. An active ride has auto-expired, with an assigned driver and stake already put up.
// 3. A finished ride has auto-expired, where applicable payouts were made to the driver and/or passenger.
const (
	RideRequestExpiredEventKey  = "RideRequestExpired"
	ActiveRideExpiredEventKey   = "ActiveRideExpired"
	FinishedRideExpiredEventKey = "FinishedRideExpired"
)
