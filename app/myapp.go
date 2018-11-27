package myapp

import (
	"fmt"
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

func (srv *Server) handleWrite(w http.ResponseWriter, r *http.Request) {
	err := srv.db.WriteSomething()
	if err != nil {
		fmt.Fprintf(w, "unable to write to db: %s\n", err.Error())
		return
	}

	// irl, you would probably want to return something in json format
	fmt.Fprintln(w, "success")
}
func (srv *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	id, err := srv.db.GetIDs()
	if err != nil {
		fmt.Fprintf(w, "unable to read from db: %s\n", err.Error())
		return
	}
	fmt.Fprintf(w, "success: id=%d\n", id)
}

// Start the service
func (srv *Server) Start() {
	defer srv.disconnectDB()
	log.Println("started service: myapp")

	http.HandleFunc("/write", srv.handleWrite)
	http.HandleFunc("/get", srv.handleGet)
	http.ListenAndServe(":8081", nil)
}

func (srv *Server) disconnectDB() {
	srv.db.Conn.Close()
}
