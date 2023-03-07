package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed dist
var f embed.FS

func (s *server) getFileSystem(useOS bool) (http.FileSystem, fs.FS) {
	if useOS {
		log.Println("using live mode")
		return http.FS(os.DirFS("dist")), fs.FS(os.DirFS("dist"))
	}

	fsys, err := fs.Sub(f, "dist")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys), fs.FS(fsys)
}

func (s *server) routes() {
	HFS, FSS := s.getFileSystem(embedUseOS)
	staticHandler := http.FileServer(HFS)
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
	s.router.PathPrefix("/").HandlerFunc(s.handleIndex(FSS))
}
