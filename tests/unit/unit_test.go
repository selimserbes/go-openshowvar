package test

import (
	"net"
	"testing"

	"github.com/selimserbes/go-openshowvar/pkg/openshowvar"
	"github.com/stretchr/testify/assert"
)

// Helper function to start a mock server.
//
// This function contains the necessary code to start a TCP server for use in tests.
// It returns a TCP listener and a channel to signal when to stop the server.
// Instead of a real server, a simulated server is created for use in tests.
// Details of the server implementation are not covered here.
func startMockServer() (net.Listener, chan bool) {
	// Bind a TCP listener to a random port on localhost.
	listener, _ := net.Listen("tcp", "127.0.0.1:0")
	// Create a channel to signal when to stop the server.
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				conn, err := listener.Accept()
				if err != nil {
					continue
				}
				go func(conn net.Conn) {
					defer conn.Close()
					buf := make([]byte, 1024)
					n, _ := conn.Read(buf)
					// Process the request and send back the appropriate response.
					response := processRequest(buf[:n])
					conn.Write(response)
				}(conn)
			}
		}
	}()

	return listener, stop
}

// Processes a mock request and generates the appropriate response.
//
// This function is responsible for processing a mock request and generating
// the appropriate response. In this implementation, the request is simply echoed back.
func processRequest(request []byte) []byte {
	// Here we should process the request and generate the appropriate response.
	// Check the length of the incoming data.
	if len(request) < 7 {
		// Return an error if the length is not sufficient.
		return []byte("invalid data length")
	}

	// For simplicity, we will just echo back the request.
	response := request

	// Modify the response to match the expected behavior in tests.
	if request[4] == 1 {
		// For write requests, echo only the value part with the necessary header.
		varNameLen := int(request[5])<<8 | int(request[6])
		valLen := int(request[7+varNameLen])<<8 | int(request[7+varNameLen+1])
		response = append(
			[]byte{0, 0, 0, byte(3 + valLen), 1, byte(valLen >> 8), byte(valLen & 0xFF)},
			request[7+varNameLen+2:]...)
	}

	return response
}

// Tests the `Connect` method of the `OpenShowVar` struct for successful connection.
func TestConnect(t *testing.T) {
	// Start a mock server
	listener, stop := startMockServer()
	defer close(stop)
	addr := listener.Addr().(*net.TCPAddr)

	// Create an `OpenShowVar` instance and connect to the mock server.
	osv := openshowvar.NewOpenShowVar(addr.IP.String(), addr.Port)
	err := osv.Connect()
	assert.NoError(t, err)
	// Check if connection is established.
	assert.NotNil(t, osv.Conn)
}

// Tests the `Send` method of the `OpenShowVar` struct.
func TestSend(t *testing.T) {
	// Start a mock server.
	listener, stop := startMockServer()
	defer close(stop)
	addr := listener.Addr().(*net.TCPAddr)

	// Create an `OpenShowVar` instance and connect to the mock server.
	osv := openshowvar.NewOpenShowVar(addr.IP.String(), addr.Port)
	osv.Connect()

	// Test sending data to the mock server.
	response, err := osv.Send("existing_var", "new_value")
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

// Tests the `Read` method of the `OpenShowVar` struct.
func TestRead(t *testing.T) {
	// Start a mock server.
	listener, stop := startMockServer()
	defer close(stop)
	addr := listener.Addr().(*net.TCPAddr)

	// Create an `OpenShowVar` instance and connect to the mock server.
	osv := openshowvar.NewOpenShowVar(addr.IP.String(), addr.Port)
	osv.Connect()

	// Test reading data from the mock server.
	response, err := osv.Read("existing_var")
	assert.NoError(t, err)
	assert.Equal(t, "existing_var", response)
}

// Tests the `Write` method of the `OpenShowVar` struct.
func TestWrite(t *testing.T) {
	// Start a mock server.
	listener, stop := startMockServer()
	defer close(stop)
	addr := listener.Addr().(*net.TCPAddr)

	// Create an `OpenShowVar` instance and connect to the mock server.
	osv := openshowvar.NewOpenShowVar(addr.IP.String(), addr.Port)
	osv.Connect()

	// Test writing data to the mock server.
	response, err := osv.Write("existing_var", "new_value")
	assert.NoError(t, err)
	assert.Equal(t, "new_value", response)
}

// Tests the `Disconnect` method of the `OpenShowVar` struct.
func TestDisconnect(t *testing.T) {
	// Start a mock server
	listener, stop := startMockServer()
	defer close(stop)
	addr := listener.Addr().(*net.TCPAddr)

	// Create an `OpenShowVar` instance and connect to the mock server.
	osv := openshowvar.NewOpenShowVar(addr.IP.String(), addr.Port)
	osv.Connect()
	// Disconnect from the mock server.
	osv.Disconnect()

	// Check if connection is closed.
	assert.Nil(t, osv.Conn)
}
