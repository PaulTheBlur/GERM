package types

import (
	"database/sql"
	"log"
)

type QuoteHeadersIntf struct {
	l  *log.Logger
	db *sql.DB
}
