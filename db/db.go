package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/vlam1/acceptance_test_example/config"
)

// Client specifies all field to create connection with the database
type Client struct {
	Conn *sql.DB
}

// NewDBClient ...
func NewDBClient(c config.Configuration) *Client {
	dbConf := &mysql.Config{
		User:                 c.DBUser,
		Passwd:               c.DBPass,
		Addr:                 fmt.Sprintf("%s:%d", c.DBHost, c.DBPort),
		Net:                  "tcp",
		ReadTimeout:          c.DBReadTimeout,
		WriteTimeout:         c.DBWriteTimeout,
		AllowNativePasswords: true,
	}

	conn, err := sql.Open("mysql", dbConf.FormatDSN())
	if err != nil {
		// for this example, I'm just using log.Fatal but might
		// want to have more detail logging and handling
		// we don't want the service to stop just because the db is
		// temporary down or something
		log.Fatalf("failed to connect to DB: %s", err.Error())
	}

	return &Client{Conn: conn}
}

// WriteSomething to the DB
func (db *Client) WriteSomething() error {
	//simple insert statment
	// irl, you would want something like proxysql as an
	// abstraction layer to set up your rules and
	// use prepare statement and other things to prevent
	// sql injection
	stmt := "INSERT IGNORE INTO test.test_table (txt) VALUES ('something')"
	_, err := db.Conn.Exec(stmt)
	return err
}

// GetIDs from the DB based on that field
func (db *Client) GetIDs() (int, error) {
	stmt := "SELECT id FROM test.test_table WHERE txt='something'"
	res, err := db.Conn.Query(stmt)
	if err != nil {
		return -1, err
	}

	var id int
	for res.Next() {
		err = res.Scan(&id)
		if err != nil {
			return -1, err
		}
	}

	return id, nil
}
