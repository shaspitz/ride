package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredRideKeyPrefix is the prefix to retrieve all StoredRide
	StoredRideKeyPrefix = "StoredRide/value/"
)

// StoredRideKey returns the store key to retrieve a StoredRide from the index fields
func StoredRideKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
