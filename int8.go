package sqlx

import "database/sql/driver"

// NullInt8 represents an int8 that may be null. NullInt8 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullInt8 struct {
	Int8  int8
	Valid bool // Valid is true if Int8 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullInt8) Scan(value interface{}) error {
	if value == nil {
		n.Int8, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Int8, value)
}

// Value implements the driver Valuer interface.
func (n NullInt8) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int8, nil
}
