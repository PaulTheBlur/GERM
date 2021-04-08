package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type QuoteHeader struct {
	ID int `json:"id"`
}

type QuoteHeaders []*QuoteHeader

func (p *QuoteHeaders) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *QuoteHeaders) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func AddProduct(p *QuoteHeader) {
	// Add in here add
}

func UpdateProduct(id int, p *QuoteHeader) error {
	// Add in here update
	return nil
}

var ErrPQuoteHeaderNotFound = fmt.Errorf("Quote Header not found")

func GetQuoteHeaders(id int) QuoteHeaders {
	// Add in get here
	return nil
}
