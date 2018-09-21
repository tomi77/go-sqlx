package sqlx

import "database/sql/driver"

// NullInt16 represents an int16 that may be null. NullInt16 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullInt16 struct {
	Int16 int16
	Valid bool // Valid is true if Int16 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullInt16) Scan(value interface{}) error {
	if value == nil {
		n.Int16, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Int16, value)
}

// Value implements the driver Valuer interface.
func (n NullInt16) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int16, nil
}
