package internal

import (
	"log"
	"net/http"
	"time"
)

func Logging(req *http.Request) {
	start := time.Now()
	log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
}
