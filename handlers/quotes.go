package handlers

import (
	"database/sql"
	"log"

	"github.com/PaulTheBlur/GERM/factories"
	"github.com/gorilla/mux"
)

func AddHandlers(l *log.Logger, sm *mux.Router, db *sql.DB) {
	qh := factories.NewQuoteHeaders(l, db)

	// grQH := sm.Methods(http.MethodGet).Subrouter()
	// grQH.HandleFunc("/quoteheader/{id:[0-9]+}", qh.GetQuoteHeaders)
	sm.HandleFunc("/quoteheader/{id:[0-9]+}", qh.GetQuoteHeaders).Methods("GET")
	l.Println("GET /quoteheader added")

	//	putRouter.Use(ph.MiddlewareProductsValidation)
}
