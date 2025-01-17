package gtc_test

import (
	"database/sql"
	"testing"

	"github.com/WantBeASleep/med_ml_lib/gtc"
	"github.com/stretchr/testify/assert"
)

func TestPointerToSql_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		value    *string
		expected sql.NullString
	}{
		{"Nil pointer_String", nil, sql.NullString{}},
		{"Valid pointer_String", stringPointer("hello"), sql.NullString{String: "hello", Valid: true}},
		{"Empty string pointer_String", stringPointer(""), sql.NullString{String: "", Valid: true}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := gtc.String.PointerToSql(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSqlToPointer_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		value    sql.NullString
		expected *string
	}{
		{"Invalid sql.NullString_String", sql.NullString{}, nil},
		{"Valid sql.NullString (hello)_String", sql.NullString{String: "hello", Valid: true}, stringPointer("hello")},
		{"Valid sql.NullString (empty)_String", sql.NullString{String: "", Valid: true}, stringPointer("")},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := gtc.String.SqlToPointer(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func stringPointer(s string) *string {
	return &s
}
