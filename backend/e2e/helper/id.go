package helper

import (
	"fmt"
	"testing"
)

func GenerateIDs(t *testing.T, count int) []string {
	t.Helper()

	ids := make([]string, count)

	for i := 0; i < count; i++ {
		ids[i] = fmt.Sprintf("00000000-0000-0000-0000-000000000%03d", i)
	}

	return ids
}
