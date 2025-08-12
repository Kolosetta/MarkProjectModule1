package events

import (
	"log"
	"net/http"
)

func startPprof() {
	go func() {
		log.Println("pprof запущен на http://localhost:6060/debug/pprof/")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Println("pprof error:", err)
		}
	}()
}
