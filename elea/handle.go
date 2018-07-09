package elea

import (
	"net/http"
	"sync"
)

var (
	Handles = make(map[string]http.HandlerFunc)
	mu      sync.RWMutex
)

type Handle struct{}

func (h *Handle) GetHandle() map[string]http.HandlerFunc {
	return Handles
}

func (h *Handle) Register(pattern string, handleFunc http.HandlerFunc) {
	defer mu.Unlock()
	mu.Lock()
	Handles[pattern] = handleFunc
}
