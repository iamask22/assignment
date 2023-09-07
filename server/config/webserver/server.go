package webserver

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Server represents a server mux
type Server struct {
	*mux.Router
	Address string
	srv     *http.Server
}

func New(applicationInitializer func(router *mux.Router)) *Server {

	port := "8080"

	r := mux.NewRouter()

	applicationInitializer(r)

	return &Server{
		Router:  r,
		Address: ":" + port,
	}
}

func (s *Server) ServeHTTP() {
	s.srv = &http.Server{
		Handler:      s.Router,
		Addr:         s.Address,
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
	}

	log.Println("Server starting at addr:", s.Address)
	err := s.srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server Shutdown")
	} else if err != nil {
		log.Println(err)
	}
}

func (s *Server) Shutdown(goCtx context.Context) {
	if s.srv == nil {
		log.Println("Server not yet started.")
		return
	}

	if err := s.srv.Shutdown(goCtx); err != nil {
		log.Printf("Server failed to shutdown: %v", err.Error())
	}
}
