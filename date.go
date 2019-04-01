package sqlx

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// NullDate represents an time.Time that may be null. NullDate implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullDate struct {
	Date  time.Time
	Valid bool // Valid is true if Date is not NULL
}

// Scan implements the Scanner interface.
func (n *NullDate) Scan(value interface{}) error {
	if value == nil {
		n.Date = time.Time{}
		n.Valid = false
	} else {
		date := fmt.Sprintf("%s", value)
		parts := strings.Split(date, " ")
		var err error
		n.Date, err = time.Parse("2006-01-02", parts[0])
		n.Valid = err == nil
	}
	return nil
}

// Value implements the driver Valuer interface.
func (n NullDate) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Date, nil
}
