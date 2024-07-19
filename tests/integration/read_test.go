package test

import (
	"testing"

	"github.com/selimserbes/go-openshowvar/openshowvar"
	"github.com/stretchr/testify/assert"
)

// WARNING: These tests use real connections to test the functionality of the library that implements the OpenShowVar protocol.
// Since they work with real data, running these tests may lead to unintended consequences in robot systems.
// Please run these tests only in development environments and with caution.
// Make sure to adjust your connection configurations and test data to fit your own environment.

// Test reading a variable successfully.
func TestReadSuccess(t *testing.T) {
	// Creating an instance of OpenShowVar with a valid IP address and port
	osv := openshowvar.NewOpenShowVar("10.145.173.160", 7000)

	// Attempting to connect to OpenShowVar server
	err := osv.Connect()
	assert.NoError(t, err)

	// Successfully read a variable
	result, err := osv.Read("existing_var")
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

// Test attempting to read a variable with an empty name.
func TestReadEmptyVariableName(t *testing.T) {
	// Creating an instance of OpenShowVar with a valid IP address and port
	osv := openshowvar.NewOpenShowVar("10.145.173.160", 7000)

	// Attempting to connect to OpenShowVar server
	err := osv.Connect()
	assert.NoError(t, err)

	// An error should be returned when attempting to read a variable with an empty name
	result, err := osv.Read("")
	assert.Error(t, err)
	assert.Empty(t, result)
}

// Test attempting to read a non-existent variable.
func TestReadVariableNotFound(t *testing.T) {
	// Creating an instance of OpenShowVar with a valid IP address and port
	osv := openshowvar.NewOpenShowVar("10.145.173.160", 7000)

	// Attempting to connect to OpenShowVar server
	err := osv.Connect()
	assert.NoError(t, err)

	// An error should be returned when attempting to read a variable that does not exist
	result, err := osv.Read("non_existing_var")
	assert.Error(t, err)
	assert.Empty(t, result)
}
