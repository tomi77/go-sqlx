package sqlx

import "database/sql/driver"

// NullUint64 represents an uint64 that may be null. NullUint64 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullUint64 struct {
	Uint64 uint64
	Valid  bool // Valid is true if Uint64 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUint64) Scan(value interface{}) error {
	if value == nil {
		n.Uint64, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Uint64, value)
}

// Value implements the driver Valuer interface.
func (n NullUint64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Uint64, nil
}
