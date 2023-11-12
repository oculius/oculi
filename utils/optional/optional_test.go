package optional

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	t.Run("should return nil if value is nil", func(tt *testing.T) {
		tt.Run("int", func(t *testing.T) {
			var val *int
			got := Zero(val)
			assert.Nil(t, got)
		})
		tt.Run("uint", func(t *testing.T) {
			var val *uint
			got := Zero(val)
			assert.Nil(t, got)
		})
		tt.Run("int8", func(t *testing.T) {
			var val *int8
			got := Zero(val)
			assert.Nil(t, got)
		})
		tt.Run("duration", func(t *testing.T) {
			var val *time.Duration
			got := Zero(val)
			assert.Nil(t, got)
		})
	})
	t.Run("should return nil if value is zero", func(tt *testing.T) {
		tt.Run("int", func(t *testing.T) {
			val := 0
			got := Zero(&val)
			assert.Nil(t, got)
		})
		tt.Run("uint16", func(t *testing.T) {
			val := uint16(0)
			got := Zero(&val)
			assert.Nil(t, got)
		})
		tt.Run("int16", func(t *testing.T) {
			val := int16(0)
			got := Zero(&val)
			assert.Nil(t, got)
		})
		tt.Run("duration", func(t *testing.T) {
			val := 0 * time.Millisecond
			got := Zero(&val)
			assert.Nil(t, got)
		})
	})

	t.Run("should return value if value is not nil", func(tt *testing.T) {
		tt.Run("int", func(t *testing.T) {
			val := -23
			got := Zero(&val)
			assert.Equal(t, &val, got)
		})
		tt.Run("uint32", func(t *testing.T) {
			val := uint32(24567)
			got := Zero(&val)
			assert.Equal(t, &val, got)
		})
		tt.Run("int64", func(t *testing.T) {
			val := int64(124)
			got := Zero(&val)
			assert.Equal(t, &val, got)
		})
		tt.Run("duration", func(t *testing.T) {
			val := 10 * time.Second
			got := Zero(&val)
			assert.Equal(t, &val, got)
		})
	})
}

func TestEmpty(t *testing.T) {
	t.Run("should return nil if value is nil", func(tt *testing.T) {
		tt.Run("string", func(t *testing.T) {
			var val *string
			got := Empty(val)
			assert.Nil(t, got)
		})
		tt.Run("time", func(t *testing.T) {
			var val *time.Time
			got := Empty(val)
			assert.Nil(t, got)
		})
	})

	t.Run("should return value if value is not nil", func(tt *testing.T) {
		tt.Run("string", func(t *testing.T) {
			val := "hello"
			got := Empty(&val)
			assert.Equal(t, &val, got)
		})
		tt.Run("time", func(t *testing.T) {
			val := time.Now()
			got := Empty(&val)
			assert.Equal(t, &val, got)
		})
	})

	t.Run("should return nil if value is empty", func(tt *testing.T) {
		tt.Run("string", func(t *testing.T) {
			val := ""
			got := Empty(&val)
			assert.Nil(t, got)
		})
		tt.Run("time", func(t *testing.T) {
			val := time.Time{}
			got := Empty(&val)
			assert.Nil(t, got)
		})
	})
}

func TestTrim(t *testing.T) {
	t.Run("should return nil if value is nil", func(tt *testing.T) {
		var val *string
		got := Trim(val)
		assert.Nil(t, got)
	})
	t.Run("should return value if value is not nil", func(tt *testing.T) {
		val := "hello"
		got := Trim(&val)
		assert.Equal(t, &val, got)
	})
	t.Run("should return nil if value is empty", func(tt *testing.T) {
		val := ""
		got := Trim(&val)
		assert.Nil(t, got)
	})
}
