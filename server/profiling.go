package server

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// Profiling server accessible @ http://127.0.0.1:6060/debug/pprof/
//
// server will seperate go routine, so that profiling server will not blocked by proxy server
func InitProfilingServer() {
	go newProfilingServer()
}

func newProfilingServer() {
	log.Println("Starting Profiling Server @http://127.0.0.1:6060/debug/pprof/")

	err := http.ListenAndServe("localhost:6060", nil)
	if err != nil {
		log.Printf("Failed to start profiling server: %v", err)
	}
}
