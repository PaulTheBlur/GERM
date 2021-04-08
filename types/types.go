package types

import (
	"database/sql"
	"log"
)

type QuoteHeadersIntf struct {
	l  *log.Logger
	db *sql.DB
}

func NewQuoteHeadersIntf(value QuoteHeadersIntf) QuoteHeadersIntf {
	return QuoteHeadersIntf(value)
}
