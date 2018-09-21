package sqlx

import "database/sql/driver"

// NullRune represents an rune that may be null. NullRune implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullRune struct {
	Rune  rune
	Valid bool // Valid is true if Rune is not NULL
}

// Scan implements the Scanner interface.
func (n *NullRune) Scan(value interface{}) error {
	if value == nil {
		n.Rune, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Rune, value)
}

// Value implements the driver Valuer interface.
func (n NullRune) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Rune, nil
}
