package main

import (
	"github.com/voltgizerz/go-concurrent/config"
)

var (
	log = config.SetupLog()
)

func main() {
	// Concurrent
	// StartConcurrent()

	// Generics
	log.Println(toString(42))          // "42"
	log.Println(toString(int32(42)))   // "42"
	log.Println(toString(int64(42)))   // "42"
	log.Println(isEqual("asd", "Asd")) // false
	log.Println(isEqual(1, 1))         // true
}
