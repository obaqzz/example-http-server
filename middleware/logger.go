package middleware

import (
    "log"
    "net/http"
    "time"
)

func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("started %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		
        next.ServeHTTP(w, r)

        log.Printf("completed in %v from %s", time.Since(start), r.RemoteAddr)
    })
}