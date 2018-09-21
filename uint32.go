package sqlx

import "database/sql/driver"

// NullUint32 represents an uint32 that may be null. NullUint32 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullUint32 struct {
	Uint32 uint32
	Valid  bool // Valid is true if Uint32 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUint32) Scan(value interface{}) error {
	if value == nil {
		n.Uint32, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Uint32, value)
}

// Value implements the driver Valuer interface.
func (n NullUint32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Uint32, nil
}
