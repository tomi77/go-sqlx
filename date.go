package sqlx

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// NullDate represents an time.Time that may be null. NullDate implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullDate struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (n *NullDate) Scan(value interface{}) error {
	date := fmt.Sprintf("%s", value)
	var err error
	n.Time, err = time.Parse("2006-01-02", date)
	n.Valid = err == nil
	return nil
}

// Value implements the driver Valuer interface.
func (n NullDate) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}
