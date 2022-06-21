package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RatingStructKeyPrefix is the prefix to retrieve all RatingStruct
	RatingStructKeyPrefix = "RatingStruct/value/"
)

// RatingStructKey returns the store key to retrieve a RatingStruct from the index fields
func RatingStructKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
