package exechc

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type server struct {
	svr *http.Server
}

func (s *server) Start() error {
	log.Printf("Listening on %s\n", s.svr.Addr)
	s.svr.ListenAndServe()
	log.Println("Shut down")
	return nil
}

func (s *server) Shutdown() error {
	return s.svr.Shutdown(context.Background())
}

func NewServer(cfg *Runtime, chk Checker) Server {

	return &server{
		svr: &http.Server{
			Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
				ok, err := chk.Check()

				if err != nil {
					http.Error(rw, err.Error(), http.StatusInternalServerError)
					return
				}

				if !ok {
					http.Error(rw, "check failed", http.StatusInternalServerError)
					return
				}

				rw.Header().Set("Content-Type", "text/plain")
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte("SERVING\n"))
			}),
		},
	}
}
