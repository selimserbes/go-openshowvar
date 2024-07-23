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

// Connects to the OpenShowVar server with a valid IP address and port,
// then checks if the connection is successful.
func TestConnectSuccess(t *testing.T) {
	// Creating an instance of OpenShowVar with a valid IP address and port.
	osv := openshowvar.NewOpenShowVar("192.168.1.10", 7000)

	// The connection attempt should be successful.
	err := osv.Connect()
	assert.NoError(t, err)

	// The `conn` field should be non-nil after a successful connection.
	assert.NotNil(t, osv.Conn)
}

// Tries to connect to the OpenShowVar server with an invalid IP address,
// then verifies that the connection attempt fails.
func TestConnectFailure(t *testing.T) {
	// Creating an instance of OpenShowVar with an invalid IP address and port.
	osv := openshowvar.NewOpenShowVar("invalid_ip", 7000)

	// The connection attempt should fail.
	err := osv.Connect()
	assert.Error(t, err)

	// The `conn` field should be nil after a failed connection attempt.
	assert.Nil(t, osv.Conn)
}
