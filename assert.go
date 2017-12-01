package assert

import "errors"

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

func (Assert)returnError(message string, defaultMessage string) error  {
	if message == "" {
		message = defaultMessage
	}

	return errors.New(message)
}

