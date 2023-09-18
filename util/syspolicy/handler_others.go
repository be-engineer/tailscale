package syspolicy

import (
	"sync"
	"sync/atomic"
)

var (
	handlerOnce sync.Once
	handler     atomic.Value
)

func init() {
	handler.Store(defaultHandler{})
}

// RegisterHandler registers a handler for reading policies from OS-specific storage.
func RegisterHandler(h Handler) bool {
	handlerOnce.Do(func() {
		handler.Store(h)
	})
	return h == handler.Load()
}

func getHandler() Handler {
	return handler.Load().(Handler)
}
