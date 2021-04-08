package handlers

import (
	"log"

	"github.com/PaulTheBlur/GERM/factories"
	"github.com/gorilla/mux"
)

func AddHandlers(l *log.Logger, sm *mux.Router) {
	qh := factories.NewQuoteHeaders(l)

	// grQH := sm.Methods(http.MethodGet).Subrouter()
	// grQH.HandleFunc("/quoteheader/{id:[0-9]+}", qh.GetQuoteHeaders)
	sm.HandleFunc("/quoteheader/{id:[0-9]+}", qh.GetQuoteHeaders).Methods("GET")
	l.Println("GET /quoteheader added")

	//	putRouter.Use(ph.MiddlewareProductsValidation)
}
