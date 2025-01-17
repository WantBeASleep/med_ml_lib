package gtclib_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/WantBeASleep/goooool/gtclib"
	"github.com/stretchr/testify/assert"
)

func TestPointerToSql_Time(t *testing.T) {
	t.Parallel()

	now := time.Now()

	tests := []struct {
		name     string
		value    *time.Time
		expected sql.NullTime
	}{
		{"Nil pointer_Time", nil, sql.NullTime{}},
		{"Valid pointer_Time", timePointer(now), sql.NullTime{Time: now, Valid: true}},
		{"Zero value pointer_Time", timePointer(time.Time{}), sql.NullTime{Time: time.Time{}, Valid: true}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := gtclib.Time.PointerToSql(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSqlToPointer_Time(t *testing.T) {
	t.Parallel()

	now := time.Now()
	
	tests := []struct {
		name     string
		value    sql.NullTime
		expected *time.Time
	}{
		{"Invalid sql.NullTime_Time", sql.NullTime{}, nil},
		{"Valid sql.NullTime (now)_Time", sql.NullTime{Time: now, Valid: true}, timePointer(now)},
		{"Valid sql.NullTime (zero)_Time", sql.NullTime{Time: time.Time{}, Valid: true}, timePointer(time.Time{})},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := gtclib.Time.SqlToPointer(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func timePointer(t time.Time) *time.Time {
	return &t
}
