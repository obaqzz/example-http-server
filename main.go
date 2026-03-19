package main

import (
    "log"
    "net/http"
    "web-server/middleware"
)

func main() {
    mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./static"))
    mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

    mux.HandleFunc("/", HomeHandler)

    loggedMux := middleware.Logger(mux) // wrap the mux with the logger middleware

    log.Println("starting server on :8080")
    err := http.ListenAndServe(":8080", loggedMux) // port:handler
	// err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", loggedMux)

    if err != nil {
        log.Fatal(err)
    }
}