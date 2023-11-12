package maputils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToArray(t *testing.T) {
	t.Run("should return empty array if value is nil", func(tt *testing.T) {
		var val map[string]string
		got := ToArray(val)
		assert.Equal(tt, 0, len(got))
	})
	t.Run("should return value if value is not nil", func(tt *testing.T) {
		tt.Run("empty", func(t *testing.T) {
			val := map[string]string{}
			got := ToArray(val)
			assert.Equal(t, 0, len(got))
		})
		tt.Run("non-empty", func(t *testing.T) {
			val := map[string]string{"key": "value"}
			got := ToArray(val)
			assert.Equal(t, 1, len(got))
			assert.Equal(t, "value", got[0])
		})
		tt.Run("non-empty multi value", func(t *testing.T) {
			val := map[string]string{"key": "value", "foo": "bar"}
			got := ToArray(val)
			assert.Equal(t, 2, len(got))
			assert.Equal(t, "value", got[0])
			assert.Equal(t, "bar", got[1])
		})
		tt.Run("non-empty int", func(t *testing.T) {
			val := map[string]int{"key": 123, "foo": 456, "bar": 789}
			got := ToArray(val)
			assert.Equal(t, 3, len(got))
			assert.Equal(t, []int{123, 456, 789}, got)
		})
	})
}
