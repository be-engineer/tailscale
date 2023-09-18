package syspolicy

import (
	"errors"
)

// Handler reads system policies from OS-specific storage.
type Handler interface {
	// ReadBool returns a boolean whether the given policy key exists or not on device settings.
	ReadBool(key string) (bool, error)
	// ReadString reads the policy settings value string given the key.
	ReadString(key string) (string, error)
}

// ErrNoSuchKey is returned when the specified key does not have a value set.
var ErrNoSuchKey = errors.New("no such key")

type defaultHandler struct{}

func (defaultHandler) ReadBool(_ string) (bool, error) {
	return false, ErrNoSuchKey
}

func (defaultHandler) ReadString(_ string) (string, error) {
	return "", ErrNoSuchKey
}
