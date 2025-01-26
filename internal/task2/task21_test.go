package task2

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestTask## - запускает тест функции Run##
// Тест требует доработки в зависимости от условия задачи
func TestTask21(t *testing.T) {
	// Указать диапазон тестов
	for i := 8; i <= 41; i++ {
		file, err := os.Open(fmt.Sprintf("tests/%d", i))
		if err != nil {
			continue
		}
		defer file.Close()

		t.Run(fmt.Sprintf("Test:%d", i), func(t *testing.T) {
			in := bufio.NewReader(file)

			expected, err := os.ReadFile(fmt.Sprintf("tests/%d.a", i))
			require.Nil(t, err)

			var buffer bytes.Buffer
			out := bufio.NewWriter(&buffer)

			// Указать тестируемую функцию задачи # (вариант #)
			Run21(in, out)

			out.Flush()

			// Код, требующий доработки в зависимости от условия задачи
			result, err := io.ReadAll(bufio.NewReader(&buffer))
			require.Nil(t, err)

			require.Equal(t, string(expected), string(result))
		})
	}
}

// BenchmarkTask## - бенчмарк для задачи # (вариант #)
func BenchmarkTask21(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Task21()
	}
}
