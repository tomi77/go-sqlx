package sqlx

import "database/sql/driver"

// NullByte represents an byte that may be null. NullByte implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullByte struct {
	Byte  byte
	Valid bool // Valid is true if Byte is not NULL
}

// Scan implements the Scanner interface.
func (n *NullByte) Scan(value interface{}) error {
	if value == nil {
		n.Byte, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Byte, value)
}

// Value implements the driver Valuer interface.
func (n NullByte) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Byte, nil
}
