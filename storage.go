package main

import "sync"

var (
	urlMap = make(map[string]string)
	mutex  sync.RWMutex
)

func saveURL(id, url string) {
	mutex.Lock()
	defer mutex.Unlock()
	urlMap[id] = url
}

func getURL(id string) string {
	mutex.RLock()
	defer mutex.RUnlock()
	return urlMap[id]
}
