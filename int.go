package sqlx

import "database/sql/driver"

// NullInt represents an int that may be null. NullInt implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullInt struct {
	Int   int
	Valid bool // Valid is true if Int is not NULL
}

// Scan implements the Scanner interface.
func (n *NullInt) Scan(value interface{}) error {
	if value == nil {
		n.Int, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Int, value)
}

// Value implements the driver Valuer interface.
func (n NullInt) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int, nil
}
