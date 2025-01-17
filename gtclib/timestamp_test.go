package gtclib_test

import (
	"testing"
	"time"

	"github.com/WantBeASleep/goooool/gtclib"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTimePointerToPointer(t *testing.T) {
	t.Parallel()

	now := time.Now()

	tests := []struct {
		name     string
		value    *time.Time
		expected *timestamppb.Timestamp
	}{
		{"Nil pointer", nil, nil},
		{"Valid pointer", timePointer(now), timestamppb.New(now)},
		{"Zero value pointer", timePointer(time.Time{}), timestamppb.New(time.Time{})},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := gtclib.Timestamp.TimePointerToPointer(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPointerToTimePointer(t *testing.T) {
	t.Parallel()

	now := time.Now()
	now = now.UTC()

	tests := []struct {
		name     string
		value    *timestamppb.Timestamp
		expected *time.Time
	}{
		{"Nil timestamp", nil, nil},
		{"Valid timestamp", timestamppb.New(now), timePointer(now)},
		{"Zero timestamp", timestamppb.New(time.Time{}), timePointer(time.Time{})},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := gtclib.Timestamp.PointerToTimePointer(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}
