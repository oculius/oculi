package enum

import (
	"database/sql/driver"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	Test BaseEnum
)

const TestEnumKey = "test_enum"

func (t Test) Name() string {
	return BaseEnum(t).AbstractName(TestEnumKey)
}

func (t Test) Code() string {
	return BaseEnum(t).AbstractCode(TestEnumKey)
}

func (t Test) Value() (driver.Value, error) {
	return BaseEnum(t).AbstractValue(t)
}

func (t *Test) Scan(val interface{}) error {
	idx, err := BaseEnum(*t).AbstractScan(val, TestEnumKey)
	if err != nil {
		return err
	}
	*t = Test(idx)
	return nil
}

func (t Test) MarshalJSON() ([]byte, error) {
	return BaseEnum(t).AbstractMarshalJSON(t)
}

func (t *Test) UnmarshalJSON(val []byte) error {
	idx, err := BaseEnum(*t).AbstractUnmarshalJSON(val, TestEnumKey)
	if err != nil {
		return err
	}
	*t = Test(idx)
	return nil
}

func TestEnum(t *testing.T) {

	t.Run("normal enum registration", func(t *testing.T) {
		enumCollection = map[string][]Single{}
		err := Create[Test, *Test](TestEnumKey, []Single{
			New("Start", "start"),
			New("Ongoing", "ongoing"),
			New("Stop", "stop"),
			New("Crashed", "crashed"),
		})
		assert.Nil(t, err)
		assert.Equal(t, "Start", Test(1).Name())
		assert.Equal(t, "Ongoing", Test(2).Name())
		assert.Equal(t, "stop", Test(3).Code())
		assert.Equal(t, "crashed", Test(4).Code())
	})

	t.Run("when register same key", func(t *testing.T) {
		enumCollection = map[string][]Single{}
		err := Create[Test, *Test](TestEnumKey, []Single{
			New("Start", "start"),
			New("Ongoing", "ongoing"),
			New("Stop", "stop"),
			New("Crashed", "crashed"),
		})
		assert.Nil(t, err)
		err = Create[Test, *Test](TestEnumKey, []Single{
			New("Red", "red"),
			New("Green", "green"),
			New("Blue", "blue"),
		})
		assert.Error(t, err)
		assert.True(t, ErrEnumKeyRegistered(nil, nil).Equal(err))
	})

	t.Run("scan and value", func(t *testing.T) {
		enumCollection = map[string][]Single{}
		err := Create[Test, *Test](TestEnumKey, []Single{
			New("Red", "red"),
			New("Green", "green"),
			New("Blue", "blue"),
		})
		assert.Nil(t, err)

		val, err := Test(1).Value()
		assert.Equal(t, "red", val)
		assert.Nil(t, err)

		val, err = Test(4).Value()
		assert.Equal(t, "", val)
		assert.Nil(t, err)

		enum := Test(0)
		err = enum.Scan("blue")
		assert.Equal(t, Test(3), enum)
		assert.Nil(t, err)

		enum = Test(0)
		err = enum.Scan("bluee")
		assert.Equal(t, Test(0), enum)
		assert.Error(t, err)
		assert.True(t, ErrEnumNotFound(nil, nil).Equal(err))
	})
}
