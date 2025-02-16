package gtc_test

import (
	"database/sql"
	"testing"

	"github.com/WantBeASleep/med_ml_lib/gtc"
	"github.com/stretchr/testify/assert"
)

func TestPointerToSql_Float64(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		value    *float64
		expected sql.NullFloat64
	}{
		{"Nil pointer_Float64", nil, sql.NullFloat64{}},
		{"Valid pointer_Float64", floatPointer(10.5), sql.NullFloat64{Valid: true, Float64: 10.5}},
		{"Zero value pointer_Float64", floatPointer(0), sql.NullFloat64{Valid: true, Float64: 0}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := gtc.Float64.PointerToSql(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSqlToPointer_Float64(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		value    sql.NullFloat64
		expected *float64
	}{
		{"Invalid sql.NullFloat64_Float64", sql.NullFloat64{}, nil},
		{"Valid sql.NullFloat64 (10.5)_Float64", sql.NullFloat64{Valid: true, Float64: 10.5}, floatPointer(10.5)},
		{"Valid sql.NullFloat64 (0)_Float64", sql.NullFloat64{Valid: true, Float64: 0}, floatPointer(0)},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := gtc.Float64.SqlToPointer(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func floatPointer(f float64) *float64 {
	return &f
}
