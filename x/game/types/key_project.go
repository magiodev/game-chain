package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProjectKeyPrefix is the prefix to retrieve all Project
	ProjectKeyPrefix = "Project/value/"
)

// ProjectKey returns the store key to retrieve a Project from the index fields
func ProjectKey(
	symbol string,
) []byte {
	var key []byte

	symbolBytes := []byte(symbol)
	key = append(key, symbolBytes...)
	key = append(key, []byte("/")...)

	return key
}
