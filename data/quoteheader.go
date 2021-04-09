package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PaulTheBlur/GERM/types"
)

type QuoteHeadersIntf types.QuoteHeadersIntf

type QuoteHeader struct {
	EventID            int       `json:"eventid"`
	EventDescription   string    `json:"eventdescription"`
	ArrivalDate        time.Time `json:"arrivaldate"`
	StartDate          time.Time `json:"startdate"`
	EndDate            time.Time `json:"enddate"`
	QuoteID            int       `json:"quoteid"`
	EventModuleID      int       `json:"eventmoduleid"`
	Description        string    `json:"description"`
	QuoteDate          time.Time `json:"quotedate"`
	FinalQuote         bool      `json:"finalquote"`
	ForecastFactor     float32   `json:"forecastfactor"`
	TotalQuote         float32   `json:"totalquote"`
	TotalNet           float32   `json:"totalnet"`
	TotalGross         float32   `json:"totalgross"`
	UnknownChargeTypes bool      `json:"unknownchargetypes"`
	NegativeElements   bool      `json:"negativeelements"`
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

var ErrPQuoteHeaderNotFound = fmt.Errorf("quote header not found")

func GetQuoteHeaders(p *types.QuoteHeadersIntf, id int) (QuoteHeaders, error) {
	p.L.Println("Get QuoteHeader ", id)
	// Add in get here
	ctx := context.Background()

	// Check if database is alive.

	err := p.DB.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := fmt.Sprintf("exec dbo.PulseGetEventModuleQuotes @EventModuleID")

	// Execute query
	rows, err := p.DB.QueryContext(ctx, tsql, sql.Named("EventModuleID", id))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// // Iterate through the result set.

	bks := make([]*QuoteHeader, 0)
	for rows.Next() {
		item := new(QuoteHeader)
		err := rows.Scan(&item.EventID, &item.EventDescription, &item.ArrivalDate, &item.StartDate, &item.EndDate, &item.QuoteID, &item.EventModuleID, &item.Description, &item.QuoteDate, &item.FinalQuote, &item.ForecastFactor, &item.TotalQuote, &item.TotalNet, &item.TotalGross, &item.UnknownChargeTypes, &item.NegativeElements)
		if err != nil {
			return nil, fmt.Errorf(http.StatusText(500), 500)
		}
		bks = append(bks, item)
	}

	return bks, nil
}
