package main

import (
	"net/http"

	"github.com/markbates/pkger"
)

func (s *server) routes() {
	staticHandler := http.FileServer(pkger.Dir("/dist"))
	// static assets
	s.router.PathPrefix("/static/").Handler(http.StripPrefix(s.config.PathPrefix, staticHandler))
	s.router.PathPrefix("/img/").Handler(http.StripPrefix(s.config.PathPrefix, staticHandler))
	s.router.PathPrefix("/lang/").Handler(http.StripPrefix(s.config.PathPrefix, staticHandler))
	// api (currently internal to UI application)
	s.router.HandleFunc("/api/hostnames", GetHostnames).Methods("GET")
	s.router.HandleFunc("/api/hostnames/{domain}", GetHostname).Methods("GET")
	s.router.HandleFunc("/api/hostnames", CreateHostname).Methods("POST")
	s.router.HandleFunc("/api/hostnames/{domain}", UpdateHostname).Methods("PUT")
	s.router.HandleFunc("/api/hostnames/{domain}", DeleteHostname).Methods("DELETE")
	s.router.HandleFunc("/api/triggerUpdate", TriggerUpdate).Methods("GET")
	// handle index.html
	s.router.PathPrefix("/").HandlerFunc(s.handleIndex())
}
