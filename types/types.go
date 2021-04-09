package types

import (
	"database/sql"
	"log"
)

type QuoteHeadersIntf struct {
	L  *log.Logger
	DB *sql.DB
}
