package utils

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
)

var (
	cache = make(map[string]string)
	once  = make(map[string]*sync.Once)
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateID(length int) string {
	id := make([]byte, length)
	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
	}
	return string(id)
}

func GetEnv(key, fallback string) string {
	if once[key] == nil {
		once[key] = &sync.Once{}
	}

	once[key].Do(func() {
		val := os.Getenv(key)
		if val == "" {
			val = fallback
		}
		cache[key] = val

		// Log de depuración
		pc, file, line, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()
		fmt.Printf("[ENV] %s = %s (desde %s:%d → %s)\n", key, val, file, line, funcName)
	})

	return cache[key]
}
