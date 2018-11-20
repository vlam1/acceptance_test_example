package myapp

import (
	"log"
	"net/http"

	"github.com/vlam1/acceptance_test_example/config"
	"github.com/vlam1/acceptance_test_example/db"
)

type Server struct {
	config config.Configuration
	db     *db.Client
}

// NewServer ...
func NewServer(c config.Configuration) *Server {
	return &Server{
		config: c,
		db:     db.NewDBClient(c),
	}
}

// Start the service
func (srv *Server) Start() {
	defer srv.disconnectDB()
	log.Println("started service: myapp")

	http.ListenAndServe(":8081", nil)
}

func (srv *Server) disconnectDB() {
	srv.db.Conn.Close()
}
