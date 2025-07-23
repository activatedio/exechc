package exechc

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type server struct {
	svr *http.Server
}

func (s *server) Start() error {
	log.Printf("Listening on %s\n", s.svr.Addr)
	err := s.svr.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}
	log.Println("Shut down")
	return nil
}

func (s *server) Shutdown() error {
	return s.svr.Shutdown(context.Background())
}

// NewServer constructs a server
func NewServer(cfg *Runtime, chk Checker) Server {

	return &server{
		svr: &http.Server{
			Addr:              fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			ReadHeaderTimeout: 2 * time.Second,
			Handler: http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
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
				_, err = rw.Write([]byte("SERVING\n"))
				Must(err)
			}),
		},
	}
}
