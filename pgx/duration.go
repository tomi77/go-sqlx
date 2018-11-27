package pgx

import (
	"database/sql/driver"
	"time"

	timex "github.com/tomi77/go-timex"
)

// NullDuration represents an time.Duration that may be null. NullDuration implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullDuration struct {
	Duration time.Duration
	Valid    bool // Valid is true if Duration is not NULL
}

// Scan implements the Scanner interface.
func (n *NullDuration) Scan(value interface{}) error {
	n.Duration = 0
	n.Valid = false
	if value != nil {
		bDur, valid := value.([]byte)
		if valid {
			duration, err := timex.ParseDuration(string(bDur))
			n.Duration = duration
			n.Valid = err == nil
		}
	}
	return nil
}

// Value implements the driver Valuer interface.
func (n NullDuration) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return timex.FormatDuration(n.Duration), nil
}

// Duration represents an time.Duration. Duration implements the Scanner interface so it can be used as a scan destination.
type Duration struct {
	Duration time.Duration
}

// Scan implements the Scanner interface.
func (d *Duration) Scan(value interface{}) error {
	d.Duration = 0
	if value != nil {
		bDur, valid := value.([]byte)
		if valid {
			duration, _ := timex.ParseDuration(string(bDur))
			d.Duration = duration
		}
	}
	return nil
}

// Value implements the driver Valuer interface.
func (d Duration) Value() (driver.Value, error) {
	return timex.FormatDuration(d.Duration), nil
}
