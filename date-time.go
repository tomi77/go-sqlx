package sqlx

import (
	"database/sql/driver"
	"time"
)

// NullDateTime represents an time.Time that may be null. NullDateTime implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullDateTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (n *NullDateTime) Scan(value interface{}) error {
	n.Time, n.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullDateTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}
