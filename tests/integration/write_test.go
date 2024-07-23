package test

import (
	"testing"

	"github.com/selimserbes/go-openshowvar/pkg/openshowvar"
	"github.com/stretchr/testify/assert"
)

// WARNING: These tests use real connections to test the functionality of the library that implements the OpenShowVar protocol.
// Since they work with real data, running these tests may lead to unintended consequences in robot systems.
// Please run these tests only in development environments and with caution.
// Make sure to adjust your connection configurations and test data to fit your own environment.

// Test writing a value to a variable successfully.
func TestWriteSuccess(t *testing.T) {
	// Establishing a connection to OpenShowVar.
	osv := openshowvar.NewOpenShowVar("192.168.1.10", 7000)
	err := osv.Connect()
	assert.NoError(t, err)

	// Write a new value
	_, err = osv.Write("existing_var", "new_value")
	assert.NoError(t, err)
}

// Test attempting to write to a variable with an empty name.
func TestWriteEmptyVariableName(t *testing.T) {
	// Establishing a connection to OpenShowVar.
	osv := openshowvar.NewOpenShowVar("192.168.1.10", 7000)
	err := osv.Connect()
	assert.NoError(t, err)

	// An error should be returned when attempting to write to a variable with an empty name.
	_, err = osv.Write("", "new_value")
	assert.Error(t, err)
}

// Test attempting to write an empty value to a variable.
func TestWriteEmptyValue(t *testing.T) {
	// Establishing a connection to OpenShowVar.
	osv := openshowvar.NewOpenShowVar("192.168.1.10", 7000)
	err := osv.Connect()
	assert.NoError(t, err)

	// An error should be returned when attempting to write an empty value.
	_, err = osv.Write("existing_var", "")
	assert.Error(t, err)
}

// Test attempting to write a value to a non-existent variable.
func TestWriteVariableNotFound(t *testing.T) {
	// Establishing a connection to OpenShowVar.
	osv := openshowvar.NewOpenShowVar("192.168.1.10", 7000)
	err := osv.Connect()
	assert.NoError(t, err)

	// An error should be returned when attempting to write a value to a variable that does not exist.
	_, err = osv.Write("non_existing_var", "new_value")
	assert.Error(t, err)
}
