package gtc_test

import (
	"testing"

	"github.com/WantBeASleep/med_ml_lib/gtc"
	// Используем правильный импорт
	"github.com/stretchr/testify/assert" // Импортируем библиотеку testify
)

func TestValueToPointer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		value    int
		expected *int
	}{
		{"Positive integer", 1, intPointer(1)},
		{"Zero", 0, intPointer(0)},
		{"Negative integer", -1, intPointer(-1)},
	}

	for _, tt := range tests {
		tt := tt // создаем новую переменную для захвата цикла
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // Запуск тестов параллельно
			result := gtc.ValueToPointer(tt.value)
			assert.Equal(t, tt.expected, result) // Используем assert для утверждения
		})
	}
}

func TestValueToPointerZeroValue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		value    string
		expected *string
	}{
		{"Empty string", "", nil}, // пустая строка - нулевое значение
		{"Non-empty string", "test", stringPointer("test")},
	}

	for _, tt := range tests {
		tt := tt // создаем новую переменную для захвата цикла
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // Запуск тестов параллельно
			result := gtc.ValueToPointerZeroValue(tt.value)
			assert.Equal(t, tt.expected, result) // Используем assert для утверждения
		})
	}
}

func intPointer(i int) *int {
	return &i
}
