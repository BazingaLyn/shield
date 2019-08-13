package core

import "sync"

type StringMap struct {
	mu       sync.RWMutex
	innerMap map[string]string
}
