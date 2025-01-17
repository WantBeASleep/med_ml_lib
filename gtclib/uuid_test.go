package gtclib_test

import (
	"testing"

	"github.com/WantBeASleep/goooool/gtclib"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMustStringPointerToPointer_Success(t *testing.T) {
	t.Parallel()
	validUUIDString := "123e4567-e89b-12d3-a456-426614174000"

	tests := []struct {
		name     string
		value    *string
		expected *uuid.UUID
	}{
		{"Nil pointer", nil, nil},
		{"Valid UUID string", &validUUIDString, uuidPointer(uuid.MustParse(validUUIDString))},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := gtclib.Uuid.MustStringPointerToPointer(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTestMustStringPointerToPointer_Panic(t *testing.T) {
	t.Parallel()
	
	invalidUUIDString := "yare broke nigga!"

	assert.Panics(t, func() {gtclib.Uuid.MustStringPointerToPointer(&invalidUUIDString)})
}

func uuidPointer(uuid uuid.UUID) *uuid.UUID {
	return &uuid
}