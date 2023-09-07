package webserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
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

	fmt.Println("Server starting at addr:", s.Address)
	err := s.srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server Shutdown")
	} else if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) Shutdown(goCtx context.Context) {
	if s.srv == nil {
		fmt.Println("Server not yet started.")
		return
	}

	if err := s.srv.Shutdown(goCtx); err != nil {
		fmt.Printf("Server failed to shutdown: %v", err.Error())
	}
}
