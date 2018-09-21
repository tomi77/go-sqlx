package sqlx

import "database/sql/driver"

// NullUint8 represents an uint8 that may be null. NullUint8 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullUint8 struct {
	Uint8 uint8
	Valid bool // Valid is true if Uint8 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUint8) Scan(value interface{}) error {
	if value == nil {
		n.Uint8, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Uint8, value)
}

// Value implements the driver Valuer interface.
func (n NullUint8) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Uint8, nil
}
