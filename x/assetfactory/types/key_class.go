package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ClassKeyPrefix is the prefix to retrieve all Class
	ClassKeyPrefix = "Class/value/"
)

// ClassKey returns the store key to retrieve a Class from the index fields
func ClassKey(
	symbol string,
) []byte {
	var key []byte

	symbolBytes := []byte(symbol)
	key = append(key, symbolBytes...)
	key = append(key, []byte("/")...)

	return key
}
