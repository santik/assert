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

func TestStringNotEmpty_withEmptyStringAndCustomMessage_ShouldReturnCustomError(t *testing.T) {
	s := ""
	assert := Assert{}
	customMessage := "custom message"
	err := assert.stringNotEmpty(s, customMessage)

	if err.Error() != customMessage {
		t.Error("Expected custom message")
	}
}