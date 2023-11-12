package arrayutils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMap(t *testing.T) {
	t.Run("should return empty map if value is nil", func(tt *testing.T) {
		var val []string
		got := ToMap(val, func(v string) int {
			return len(v)
		})
		assert.Equal(tt, 0, len(got))
	})
	t.Run("should return value if value is not nil", func(tt *testing.T) {
		tt.Run("empty", func(t *testing.T) {
			val := []string{}
			got := ToMap(val, func(v string) int {
				return len(v)
			})
			assert.Equal(t, 0, len(got))
		})

		tt.Run("non-empty", func(t *testing.T) {
			val := []string{"key"}
			got := ToMap(val, func(v string) int {
				return len(v)
			})
			assert.Equal(t, 1, len(got))
			assert.Equal(t, "key", got[3][0])
		})

		tt.Run("non-empty multi value", func(t *testing.T) {
			val := []string{"key", "foo"}
			got := ToMap(val, func(v string) int {
				return len(v)
			})

			assert.Equal(t, 1, len(got))
			assert.Equal(t, 2, len(got[3]))
			assert.Equal(t, "key", got[3][0])
			assert.Equal(t, "foo", got[3][1])
		})

		tt.Run("non-empty multi value (2)", func(t *testing.T) {
			val := []string{"key", "foo", "qwerty"}
			got := ToMap(val, func(v string) int {
				return len(v)
			})
			assert.Equal(t, 2, len(got))
			assert.Equal(t, 2, len(got[3]))
			assert.Equal(t, 1, len(got[6]))
			assert.Equal(t, "key", got[3][0])
			assert.Equal(t, "foo", got[3][1])
			assert.Equal(t, "qwerty", got[6][0])
		})
	})
}

func TestToMapUnique(t *testing.T) {
	t.Run("should return empty map if value is nil", func(tt *testing.T) {
		var val []string
		got, omitted := ToMapUnique(val, func(v string) int {
			return len(v)
		})
		assert.Equal(tt, 0, len(got))
		assert.Equal(tt, 0, omitted)
	})

	t.Run("should return value if value is not nil", func(tt *testing.T) {
		tt.Run("empty", func(t *testing.T) {
			val := []string{}
			got, omitted := ToMapUnique(val, func(v string) int {
				return len(v)
			})
			assert.Equal(t, 0, len(got))
			assert.Equal(t, 0, omitted)
		})

		tt.Run("non-empty", func(t *testing.T) {
			val := []string{"key"}
			got, omitted := ToMapUnique(val, func(v string) int {
				return len(v)
			})
			assert.Equal(t, 1, len(got))
			assert.Equal(t, 0, omitted)
			assert.Equal(t, "key", got[3])
		})

		tt.Run("non-empty multi value", func(t *testing.T) {
			val := []string{"key", "foo"}
			got, omitted := ToMapUnique(val, func(v string) int {
				return len(v)
			})
			assert.Equal(t, 1, len(got))
			assert.Equal(t, 1, omitted)
			assert.Equal(t, "key", got[3])
		})

		tt.Run("non-empty multi value (2)", func(t *testing.T) {
			val := []string{"key", "foo", "qwerty"}
			got, omitted := ToMapUnique(val, func(v string) int {
				return len(v)
			})
			assert.Equal(t, 2, len(got))
			assert.Equal(t, 1, omitted)
			assert.Equal(t, "key", got[3])
			assert.Equal(t, "qwerty", got[6])
		})
	})
}

func TestUnique(t *testing.T) {
	t.Run("should return empty array if value is nil", func(tt *testing.T) {
		var val []string
		got, omitted := Unique(val, func(v string) int {
			return len(v)
		})
		assert.Equal(tt, 0, len(got))
		assert.Equal(tt, 0, omitted)
	})
	t.Run("should return value if value is not nil", func(tt *testing.T) {
		tt.Run("empty", func(t *testing.T) {
			val := []string{}
			got, omitted := Unique(val, func(v string) int {
				return len(v)
			})
			assert.Equal(t, 0, len(got))
			assert.Equal(t, 0, omitted)
		})
		tt.Run("non-empty", func(t *testing.T) {
			val := []string{"key"}
			got, omitted := Unique(val, Identity[string])
			assert.Equal(t, 1, len(got))
			assert.Equal(t, 0, omitted)
			assert.Equal(t, "key", got[0])
		})
		tt.Run("non-empty multi value", func(t *testing.T) {
			val := []string{"key", "foo"}
			got, omitted := Unique(val, Identity[string])
			sort.Strings(got)
			assert.Equal(t, 2, len(got))
			assert.Equal(t, 0, omitted)
			assert.Equal(t, "key", got[1])
			assert.Equal(t, "foo", got[0])
		})
		tt.Run("non-empty multi value (2)", func(t *testing.T) {
			val := []string{"key", "foo", "qwerty", "foo"}
			got, omitted := Unique(val, Identity[string])
			sort.Strings(got)
			assert.Equal(t, 3, len(got))
			assert.Equal(t, 1, omitted)
			assert.Equal(t, "key", got[1])
			assert.Equal(t, "foo", got[0])
			assert.Equal(t, "qwerty", got[2])
		})
	})
}
