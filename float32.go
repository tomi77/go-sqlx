package sqlx

import "database/sql/driver"

// NullFloat32 represents an int32 that may be null. NullFloat32 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullFloat32 struct {
	Float32 float32
	Valid   bool // Valid is true if Float32 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullFloat32) Scan(value interface{}) error {
	if value == nil {
		n.Float32, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Float32, value)
}

// Value implements the driver Valuer interface.
func (n NullFloat32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Float32, nil
}
