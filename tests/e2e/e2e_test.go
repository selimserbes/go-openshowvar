package openshowvar

import (
	"strconv"
	"testing"
	"time"

	"github.com/selimserbes/go-openshowvar/openshowvar"
	"github.com/stretchr/testify/assert"
)

// WARNING: These tests use real connections to test the functionality of the library that implements the OpenShowVar protocol.
// Since they work with real data, running these tests may lead to unintended consequences in robot systems.
// Please run these tests only in development environments and with caution.
// Make sure to adjust your connection configurations and test data to fit your own environment.

// Tests writing and reading operations with a real robot.
// This test connects to an OpenShowVar server running on a real robot,
// writes a value to a specified variable, waits for a short duration,
// then reads the same variable to ensure the value matches the written one.
func TestWriteAndReadWithRealRobot(t *testing.T) {
	// Establishing a connection to OpenShowVar.
	osv := openshowvar.NewOpenShowVar("10.145.173.160", 7000)
	err := osv.Connect()
	assert.NoError(t, err, "Connection failed")

	// Defining the value to be written and the variable name.
	value := "2"
	variableName := "COUNT"

	// Performing the variable writing operation.
	_, err = osv.Write(variableName, value)
	assert.NoError(t, err, "Variable writing failed")

	// Waiting for the value to propagate.
	time.Sleep(1 * time.Second)

	// Performing the variable reading operation.
	readValue, err := osv.Read(variableName)
	assert.NoError(t, err, "Variable reading failed")
	assert.Equal(t, value, readValue, "Read value is different than expected")

	// Closing the OpenShowVar connection.
	osv.Disconnect()
}

// Tests error handling in OpenShowVar operations.
// This test checks error handling behavior for various scenarios such as
// connecting with an invalid IP address, reading from a non-existent variable,
// and disconnecting when there is no active connection.
func TestOpenShowVarErrorHandling(t *testing.T) {
	// Testing connection with an invalid IP address.
	osv := openshowvar.NewOpenShowVar("invalid_ip_address", 7000)
	err := osv.Connect()
	assert.Error(t, err, "Successful connection with an invalid IP address")

	// Testing reading a non-existent variable.
	osv = openshowvar.NewOpenShowVar("10.145.173.160", 7000)
	err = osv.Connect()
	assert.NoError(t, err)

	// Testing reading to a non-existent variable.
	_, err = osv.Read("non_existing_var")
	assert.Error(t, err, "Successful reading of a non-existent variable")

	// Testing writing to a non-existent variable.
	_, err = osv.Write("non_existing_var", "value")
	assert.Error(t, err, "Successful writing of a non-existent variable")
	// Testing disconnecting without an active connection.
	osv.Disconnect()
	assert.Nil(t, osv.Conn, "Successful disconnection without an active connection")
}

// Tests the performance of OpenShowVar operations.
// This test measures the performance of write and read operations by repeatedly writing
// a variable and immediately reading it back, ensuring that the written and read values match.
func TestOpenShowVarPerformance(t *testing.T) {
	// Establishing a connection to OpenShowVar.
	osv := openshowvar.NewOpenShowVar("192.168.1.10", 7000)
	err := osv.Connect()
	assert.NoError(t, err)

	// Writing and reading variable values in a loop.
	for i := 0; i < 100; i++ {
		variableName := "existing_var"
		_, err := osv.Write(variableName, strconv.Itoa(i))
		assert.NoError(t, err)

		readValue, err := osv.Read(variableName)
		assert.NoError(t, err)
		assert.Equal(t, strconv.Itoa(i), readValue, "Written and read values do not match")
	}

	// Closing the OpenShowVar connection.
	osv.Disconnect()
}
