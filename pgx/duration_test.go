package pgx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNullDuration(t *testing.T) {
	d, _ := time.ParseDuration("1h")

	_, err := Exec(`INSERT INTO test_nullduration(id, field) VALUES($1, $2)`, 1, NullDuration{Duration: 0, Valid: false})
	assert.Empty(t, err)
	_, err = Exec(`INSERT INTO test_nullduration(id, field) VALUES($1, $2)`, 2, NullDuration{Duration: d, Valid: true})
	assert.Empty(t, err)

	var r NullDuration

	err = QueryRow(`SELECT field FROM test_nullduration WHERE id = $1`, 1).Scan(&r)
	assert.Empty(t, err)
	assert.False(t, r.Valid)
	assert.Empty(t, r.Duration)

	err = QueryRow(`SELECT field FROM test_nullduration WHERE id = $1`, 2).Scan(&r)
	assert.Empty(t, err)
	assert.True(t, r.Valid)
	assert.Equal(t, d, r.Duration)
}

func TestDuration(t *testing.T) {
	d, _ := time.ParseDuration("1h")

	_, err := Exec(`INSERT INTO test_duration(id, field) VALUES($1, $2)`, 1, Duration{Duration: d})
	assert.Empty(t, err)

	var r Duration

	err = QueryRow(`SELECT field FROM test_duration WHERE id = 1`).Scan(&r)
	assert.Empty(t, err)
	assert.Equal(t, d, r.Duration)
}
