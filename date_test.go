package sqlx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNullDate(t *testing.T) {
	_, err := Exec(`INSERT INTO test_nulldate(id, field) VALUES($1, $2)`, 1, NullDate{Date: time.Time{}, Valid: false})
	assert.Empty(t, err)
	_, err = Exec(`INSERT INTO test_nulldate(id, field) VALUES($1, $2)`, 2, NullDate{Date: time.Date(2018, 10, 5, 0, 0, 0, 0, time.UTC), Valid: true})
	assert.Empty(t, err)

	var r NullDate

	err = QueryRow(`SELECT field FROM test_nulldate WHERE id = $1`, 1).Scan(&r)
	assert.Empty(t, err)
	assert.False(t, r.Valid)
	assert.Empty(t, r.Date)

	err = QueryRow(`SELECT field FROM test_nulldate WHERE id = $1`, 2).Scan(&r)
	assert.Empty(t, err)
	assert.True(t, r.Valid)
	assert.Equal(t, time.Date(2018, 10, 5, 0, 0, 0, 0, time.UTC), r.Date)
}
