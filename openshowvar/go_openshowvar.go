package openshowvar

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strconv"
)

// OpenShowVar struct is used to connect to a robot control system and read/write variable values over a TCP connection.
type OpenShowVar struct {
	TCP_IP   string
	TCP_PORT int
	Conn     net.Conn
}

// NewOpenShowVar creates a new instance of OpenShowVar.
//
// Parameters:
// - TCP_IP: IP address of the TCP server to connect to.
// - TCP_PORT: Port number of the TCP server to connect to.
//
// Returns: A new instance of OpenShowVar.
func NewOpenShowVar(TCP_IP string, TCP_PORT int) *OpenShowVar {
	return &OpenShowVar{
		TCP_IP:   TCP_IP,
		TCP_PORT: TCP_PORT,
	}
}

// Connect establishes a TCP connection to the server.
//
// Returns: nil if the connection is successful, otherwise an error.
func (osv *OpenShowVar) Connect() error {
	// Establish a TCP connection
	conn, err := net.Dial("tcp", osv.TCP_IP+":"+strconv.Itoa(osv.TCP_PORT))
	if err != nil {
		return fmt.Errorf("connection error: %v", err)
	}
	// Save the connection
	osv.Conn = conn
	return nil
}

// Send sends a request to read/write a variable value.
//
// Parameters:
// - varname: The name of the variable.
// - val: The value to write (leave empty to read).
//
// Returns: The response from the server or an error.
func (osv *OpenShowVar) Send(varname string, val string) ([]byte, error) {
	var msg []byte
	temp := make([]byte, 0)

	// If value is provided, prepare the message for writing
	if val != "" {
		valLen := len(val)
		msg = append(msg, byte((valLen&0xff00)>>8)) // MSB of value length
		msg = append(msg, byte(valLen&0x00ff))      // LSB of value length
		msg = append(msg, []byte(val)...)           // Value in ASCII
	}

	// Prepare the message with the variable name
	if len(varname) > 0 {
		varNameLen := len(varname)
		temp = append(temp, byte(varNameLen>>8))     // MSB of variable name length
		temp = append(temp, byte(varNameLen&0x00ff)) // LSB of variable name length
		temp = append(temp, []byte(varname)...)      // Variable name in ASCII
	}

	// Determine if the operation is a read or write
	if val != "" {
		temp = append([]byte{1}, temp...) // Write operation
	} else {
		temp = append([]byte{0}, temp...) // Read operation
	}

	// Combine variable name and value parts into final message
	msg = append(temp, msg...)

	// Create the message header
	var msgLen uint16 = uint16(len(msg))
	header := make([]byte, 4)

	// Message ID is set to 0, followed by message length
	binary.BigEndian.PutUint16(header[0:2], 0)
	binary.BigEndian.PutUint16(header[2:4], msgLen)

	// Complete request to be sent
	request := append(header, msg...)

	fmt.Printf("Sent request: %x\n", request)

	// Ensure the connection is established
	if osv.Conn == nil {
		return nil, errors.New("not connected to server")
	}

	// Send the request
	_, err := osv.Conn.Write(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	// Read the response
	response := make([]byte, 1024)
	n, err := osv.Conn.Read(response)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Trim the response to the actual data size
	response = response[:n]
	fmt.Printf("Received response: %x\n", response)

	// Filter visible characters from the response
	visibleChars := make([]byte, 0)
	for _, b := range response {
		if b >= 32 && b <= 126 {
			visibleChars = append(visibleChars, b)
		}
	}
	responseStr := string(visibleChars)
	if responseStr == "" || response[len(response)-1] == 0 {
		return nil, errors.New("variable not found in response")
	}

	return response, nil
}

// Read reads the value of a specified variable.
//
// Parameters:
// - varname: The name of the variable to read.
//
// Returns: The value of the variable as a string or an error.
func (osv *OpenShowVar) Read(varname string) (string, error) {
	// Check if the variable name is provided
	if varname == "" {
		return "", errors.New("empty variable name")
	}

	// Send a request to read the variable
	response, err := osv.Send(varname, "")
	if err != nil {
		return "", err
	}

	// Ensure the response has a valid length
	if len(response) < 7 {
		return "", errors.New("invalid response length")
	}

	// Extract the length of the variable value
	valLen := binary.BigEndian.Uint16(response[5:7])
	if len(response) < int(7+valLen) {
		return "", errors.New("response length does not match value length")
	}

	// Extract and return the variable value
	varValue := string(response[7 : 7+valLen])
	return varValue, nil
}

// Write writes a value to a specified variable.
//
// Parameters:
// - varname: The name of the variable to write.
// - val: The value to write.
//
// Returns: The written value as a string or an error.
func (osv *OpenShowVar) Write(varname string, val string) (string, error) {
	// Check if the variable name and value are provided
	if varname == "" {
		return "", errors.New("empty variable name")
	}
	if val == "" {
		return "", errors.New("empty value")
	}

	// Send a request to write the variable
	response, err := osv.Send(varname, val)
	if err != nil {
		return "", err
	}

	// Ensure the response has a valid length
	if len(response) < 7 {
		return "", errors.New("invalid response length")
	}

	// Extract the length of the variable value
	valLen := binary.BigEndian.Uint16(response[5:7])
	if len(response) < int(7+valLen) {
		return "", errors.New("response length does not match value length")
	}

	// Extract and return the written variable value
	varValue := string(response[7 : 7+valLen])
	return varValue, nil
}

// Disconnect terminates the TCP connection.
func (osv *OpenShowVar) Disconnect() {
	// Close the connection if it exists
	if osv.Conn != nil {
		osv.Conn.Close()
		osv.Conn = nil
	}
}
