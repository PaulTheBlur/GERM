package data

import (
	"context"
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

func GetQuoteHeaders(p *QuoteHeadersIntf, id int) (QuoteHeaders, error) {
	// Add in get here

	ctx := context.Background()

	// Check if database is alive.
	err := p.db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var name, location string
		var id int

		// Get values from row.
		err := rows.Scan(&id, &name, &location)
		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
		count++
	}

	return count, nil
}
