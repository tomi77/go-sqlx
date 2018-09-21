package sqlx

import "database/sql/driver"

// NullUint16 represents an uint16 that may be null. NullUint16 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullUint16 struct {
	Uint16 uint16
	Valid  bool // Valid is true if Uint16 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUint16) Scan(value interface{}) error {
	if value == nil {
		n.Uint16, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Uint16, value)
}

// Value implements the driver Valuer interface.
func (n NullUint16) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Uint16, nil
}
