package factories

import (
	"log"
	"net/http"
	"strconv"

	"database/sql"

	"github.com/PaulTheBlur/GERM/data"
	"github.com/gorilla/mux"
)

func NewQuoteHeaders(l *log.Logger, db *sql.DB) *QuoteHeadersIntf {
	return &QuoteHeadersIntf{l, db}
}

func (p *QuoteHeadersIntf) GetQuoteHeaders(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET QuoteHeaders")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert event module id", http.StatusBadRequest)
		return
	}

	lp, err := data.GetQuoteHeaders(p, id)
	if err != nil {
		http.Error(rw, "Unable to get quote headers", http.StatusInternalServerError)
		return
	}
	err = lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

/*
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusBadRequest)
		return
	}

}

type KeyProduct struct{}

func (p Products) MiddlewareProductsValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmargin json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)

	})
}
*/
