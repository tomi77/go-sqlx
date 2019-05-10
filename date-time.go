package sqlx

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// NullDateTime represents an time.Time that may be null. NullDateTime implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullDateTime struct {
	DateTime time.Time
	Valid    bool // Valid is true if DateTime is not NULL
}

// Scan implements the Scanner interface.
func (n *NullDateTime) Scan(value interface{}) error {
	if value == nil {
		n.DateTime = time.Time{}
		n.Valid = false
	} else {
		date := fmt.Sprintf("%s", value)
		var err error
		n.DateTime, err = time.Parse("2006-01-02 15:04:05 -0700 MST", date)
		n.Valid = err == nil
	}
	return nil
}

// Value implements the driver Valuer interface.
func (n NullDateTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.DateTime, nil
}
