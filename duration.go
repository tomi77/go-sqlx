package sqlx

import (
	"database/sql/driver"
	"time"
)

// NullDuration represents an time.Duration that may be null. NullDuration implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullDuration struct {
	Duration time.Duration
	Valid    bool // Valid is true if Duration is not NULL
}

// Scan implements the Scanner interface.
func (n *NullDuration) Scan(value interface{}) error {
	n.Duration, n.Valid = value.(time.Duration)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullDuration) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Duration.Seconds(), nil
}

// Duration represents an time.Duration that not may be null. Duration implements the Scanner interface so it can be used as a scan destination.
type Duration struct {
	Duration time.Duration
}

// Scan implements the Scanner interface.
func (d *Duration) Scan(value interface{}) error {
	d.Duration, _ = value.(time.Duration)
	return nil
}

// Value implements the driver Valuer interface.
func (d Duration) Value() (driver.Value, error) {
	return d.Duration.Seconds(), nil
}
