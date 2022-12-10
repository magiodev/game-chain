package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AdministratorKeyPrefix is the prefix to retrieve all Administrator
	AdministratorKeyPrefix = "Administrator/value/"
)

// AdministratorKey returns the store key to retrieve a Administrator from the index fields
func AdministratorKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
