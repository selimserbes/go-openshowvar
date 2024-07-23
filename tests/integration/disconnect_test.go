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

// Tests the disconnection functionality by connecting to an OpenShowVar server,
// then disconnecting and verifying that the connection has been closed.
func TestDisconnect(t *testing.T) {
	// Establishing a connection to OpenShowVar.
	osv := openshowvar.NewOpenShowVar("192.168.1.10", 7000)
	err := osv.Connect()
	assert.NoError(t, err)

	// Disconnecting from OpenShowVar.
	osv.Disconnect()

	// Verifying that the connection has been successfully closed and is now set to nil.
	assert.Nil(t, osv.Conn)
}
