package assert

import (
	"testing"
)

func TestStringNotEmpty_withEmptyString_ShouldReturnError(t *testing.T) {
	s := ""
	assert := Assert{}
	err := assert.stringNotEmpty(s, "")

	if err == nil {
		t.Error("Expected error")
	}
}

func TestStringNotEmpty_withNotEmptyString_ShouldNotReturnError(t *testing.T) {
	s := "some"
	assert := Assert{}
	err := assert.stringNotEmpty(s, "")

	if err != nil {
		t.Error("Expected no error")
	}
}

func TestStringEmpty_withNotEmptyString_ShouldReturnError(t *testing.T) {
	s := "some"
	assert := Assert{}
	err := assert.stringEmpty(s, "")

	if err == nil {
		t.Error("Expected error")
	}
}

func TestStringEmpty_withEmptyString_ShouldNotReturnError(t *testing.T) {
	assert := Assert{}
	s := ""
	err := assert.stringEmpty(s, "")

	if err != nil {
		t.Error("Expected not error")
	}
}

func TestTrue_withNotTrue_ShouldReturnError(t *testing.T)  {
	assert := Assert{}
	value := false

	if assert.true(value, "") == nil {
		t.Error("Expected error")
	}
}

func TestTrue_withTrue_ShouldNotReturnError(t *testing.T)  {
	assert := Assert{}
	value := true

	if assert.true(value, "") != nil {
		t.Error("Expected no error")
	}
}

func TestFalse_withNotFalse_ShouldReturnError(t *testing.T)  {
	assert := Assert{}
	value := true

	if assert.false(value, "") == nil {
		t.Error("Expected error")
	}
}

func TestFalse_withFalse_ShouldNotReturnError(t *testing.T)  {
	assert := Assert{}
	value := false

	if assert.false(value, "") != nil {
		t.Error("Expected no error")
	}
}

func TestReturnError_withEmptyMessage_ShouldReturnErrorWithDefaultMessage(t *testing.T) {
	message := ""
	assert := Assert{}
	customMessage := "custom message"
	err := assert.returnError(message, customMessage)

	if err.Error() != customMessage {
		t.Error("Expected custom message")
	}
}

func TestIntEquals_WithEqualValues_ShouldNotReturnError(t *testing.T )  {
	assert := Assert{}
	value1 := 123
	value2 := 123

	if assert.intEquals(value1, value2, "") != nil {
		t.Error("Expected no error")
	}
}

func TestIntEquals_WithNotEqualValues_ShouldReturnError(t *testing.T )  {
	assert := Assert{}
	value1 := 123
	value2 := 1234

	if assert.intEquals(value1, value2, "") == nil {
		t.Error("Expected error")
	}
}

func TestIntNotEquals_WithEqualValues_ShouldReturnError(t *testing.T )  {
	assert := Assert{}
	value1 := 123
	value2 := 123

	if assert.intNotEquals(value1, value2, "") == nil {
		t.Error("Expected error")
	}
}

func TestIntNotEquals_WithNotEqualValues_ShouldNotReturnError(t *testing.T )  {
	assert := Assert{}
	value1 := 123
	value2 := 1234

	if assert.intNotEquals(value1, value2, "") != nil {
		t.Error("Expected no error")
	}
}

func TestIntGreaterThan_WithGreaterThanValues_ShouldNotReturnError(t *testing.T )  {
	assert := Assert{}
	value1 := 124
	value2 := 123

	if assert.intGreaterThan(value1, value2, "") != nil {
		t.Error("Expected no error")
	}
}

func TestIntGreaterThan_WithNotGreaterThanValues_ShouldNotReturnError(t *testing.T )  {
	assert := Assert{}
	value1 := 123
	value2 := 124

	if assert.intGreaterThan(value1, value2, "") == nil {
		t.Error("Expected error")
	}
}

func TestIntGreaterThan_WithEqualValues_ShouldNotReturnError(t *testing.T )  {
	assert := Assert{}
	value1 := 123
	value2 := 123

	if assert.intGreaterThan(value1, value2, "") == nil {
		t.Error("Expected error")
	}
}

func TestReturnError_withMessage_ShouldReturnErrorWithMessage(t *testing.T) {
	message := "message"
	assert := Assert{}
	customMessage := "custom message"
	err := assert.returnError(message, customMessage)

	if err.Error() != message {
		t.Error("Expected message")
	}
}