package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PaulTheBlur/GERM/handlers"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

var db *sql.DB

var server = "services"
var port = 1433
var user = "go_access"
var password = "Go4GoodAgain"
var database = "Core_Kx2021_0"

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		l.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		l.Fatal(err.Error())
	}
	fmt.Printf("Connected to %v/%v!\n", server, database)

	sm := mux.NewRouter()

	handlers.AddHandlers(l, sm, db)

	http.Handle("/", sm)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	s.Shutdown(tc)
}
