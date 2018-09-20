package sqlx

import "database/sql/driver"

// NullInt32 represents an int32 that may be null. NullInt32 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullInt32 struct {
	Int32 int32
	Valid bool // Valid is true if Int32 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullInt32) Scan(value interface{}) error {
	if value == nil {
		n.Int32, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Int32, value)
}

// Value implements the driver Valuer interface.
func (n NullInt32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int32, nil
}

// NullInt16 represents an int32 that may be null. NullInt16 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
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

// NullInt8 represents an int32 that may be null. NullInt8 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
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

// NullInt represents an int32 that may be null. NullInt implements the Scanner interface so it can be used as a scan destination, similar to NullString.
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

// NullUInt64 represents an int32 that may be null. NullUInt64 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullUInt64 struct {
	UInt64 uint64
	Valid  bool // Valid is true if UInt64 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUInt64) Scan(value interface{}) error {
	if value == nil {
		n.UInt64, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.UInt64, value)
}

// Value implements the driver Valuer interface.
func (n NullUInt64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.UInt64, nil
}

// NullUInt32 represents an int32 that may be null. NullUInt32 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullUInt32 struct {
	UInt32 uint32
	Valid  bool // Valid is true if UInt32 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUInt32) Scan(value interface{}) error {
	if value == nil {
		n.UInt32, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.UInt32, value)
}

// Value implements the driver Valuer interface.
func (n NullUInt32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.UInt32, nil
}

// NullUInt16 represents an int32 that may be null. NullUInt16 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullUInt16 struct {
	UInt16 uint16
	Valid  bool // Valid is true if UInt16 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUInt16) Scan(value interface{}) error {
	if value == nil {
		n.UInt16, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.UInt16, value)
}

// Value implements the driver Valuer interface.
func (n NullUInt16) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.UInt16, nil
}

// NullUInt8 represents an int32 that may be null. NullUInt8 implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullUInt8 struct {
	UInt8 uint8
	Valid bool // Valid is true if UInt8 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUInt8) Scan(value interface{}) error {
	if value == nil {
		n.UInt8, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.UInt8, value)
}

// Value implements the driver Valuer interface.
func (n NullUInt8) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.UInt8, nil
}

// NullUInt represents an int32 that may be null. NullUInt implements the Scanner interface so it can be used as a scan destination, similar to NullString.
type NullUInt struct {
	UInt  uint
	Valid bool // Valid is true if UInt is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUInt) Scan(value interface{}) error {
	if value == nil {
		n.UInt, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.UInt, value)
}

// Value implements the driver Valuer interface.
func (n NullUInt) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.UInt, nil
}

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

// NullByte represents an int32 that may be null. NullByte implements the Scanner interface so it can be used as a scan destination, similar to NullString.
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

// NullRune represents an int32 that may be null. NullRune implements the Scanner interface so it can be used as a scan destination, similar to NullString.
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
