package assert

import (
	"errors"
	"fmt"
)

type Assert struct {
}

func (assert Assert)stringNotEmpty(value string, message string) error  {
	if value != "" {
		return nil
	}

	return assert.returnError(message, "Expected not empty")
}

func (assert Assert)stringEmpty(value string, message string) error  {
	if value == "" {
		return nil
	}

	return assert.returnError(message, "Expected empty")
}

func (assert Assert)true(value bool, message string) error  {
	if value == true {
		return nil
	}

	return assert.returnError(message, "Expected true")
}

func (assert Assert)false(value bool, message string) error  {
	if value == false {
		return nil
	}

	return assert.returnError(message, "Expected false")
}

func (assert Assert)intEquals(value1 int, value2 int, message string) error  {
	if value1 == value2 {
		return nil
	}

	return assert.returnError(message, "Expected equal values")
}

func (assert Assert)intNotEquals(value1 int, value2 int, message string) error  {
	if value1 != value2 {
		return nil
	}

	return assert.returnError(message, "Expected not equal values")
}

func (assert Assert)intGreaterThan(value1 int, value2 int, message string) error  {
	if value1 > value2 {
		return nil
	}

	return assert.returnError(message, fmt.Sprintf("Expected %d greater than %d", value1, value2))
}

func (Assert)returnError(message string, defaultMessage string) error  {
	if message == "" {
		message = defaultMessage
	}

	return errors.New(message)
}

