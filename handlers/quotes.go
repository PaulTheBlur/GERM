package handlers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PaulTheBlur/Projects/GERM/handlers"
	"github.com/PaulTheBlur/Projects/GERM/data"
	"github.com/gorilla/mux"
)

func AddHandlers(l log.logger, sm *mux.Router) {
	qh := handlers.QuoteHeaders(l)
	
	grQH := sm.Methods(http.MethodGet).Subrouter()
	grQH.HandleFunc("/quoteheader/{id:[0-9]+}", qh.GetQuoteHeaders)

//	putRouter := sm.Methods(http.MethodPut).Subrouter()
//	putRouter.HandleFunc("/products/{id:[0-9]+}", qh.UpdateProduct)
//	putRouter.Use(ph.MiddlewareProductsValidation)

//	postRouter := sm.Methods(http.MethodPost).Subrouter()
//	postRouter.HandleFunc("/products", qh.AddProduct)
//	postRouter.Use(ph.MiddlewareProductsValidation)
	
}

