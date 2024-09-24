# go_openshowvar

`go_openshowvar` is a Go library for interacting with Kuka robots over TCP/IP using the OpenShowVar protocol.

## Overview

`go-openshowvar` is a Go library that facilitates connecting to robot systems via TCP/IP to perform read and write operations. Targeting the KukaVarProxy server, it allows access to Kuka robots using the OpenShowVar protocol. This library is designed for use in robot control and monitoring applications, providing reliable communication over TCP/IP and extensive functionality for managing robot variables.

## Related Repositories

This project is inspired by the following repositories:

- [KUKAVARPROXY](https://github.com/ImtsSrl/KUKAVARPROXY)
- [JOpenShowVar](https://github.com/aauc-mechlab/JOpenShowVar)
- [kukavarproxy-msg-format](https://github.com/ahmad-saeed/kukavarproxy-msg-format)

## KukaVarProxy Message Format

The communication with KukaVarProxy follows this message format:

- msg ID in HEX (2 bytes)
- msg length in HEX (2 bytes)
- read (0) or write (1) indicator (1 byte)
- variable name length in HEX (2 bytes)
- variable name in ASCII (# bytes)
- variable value length in HEX (2 bytes)
- variable value in ASCII (# bytes)

## Installation

To install the `go-openshowvar` library, use the following command:

- **On Linux/macOS/Windows**:
  Open your terminal and execute:
  ```terminal
  go get github.com/selimserbes/go-openshowvar@latest
  ```

## Usage

### KukaVarProxy

This library is designed to connect to Kuka robots via the KukaVarProxy server. KukaVarProxy is server software used to access Kuka robot system variables over TCP/IP.

1. To install and configure KukaVarProxy, please follow the instructions provided in the [KukaVarProxy GitHub repository](https://github.com/ImtsSrl/KUKAVARPROXY).

2. **Ensure that `go_openshowvar` is configured to use port `7000` to connect to the KukaVarProxy server.** KukaVarProxy listens on this port to communicate with Kuka robots.

3. Create your own program using the `go_openshowvar` library to connect to your robot. Below is an example:

```go
package main

import (
	"fmt"
	"log"

	"github.com/selimserbes/go-openshowvar/pkg/openshowvar"
)

func main() {
	// Create a new OpenShowVar instance with the IP address and port of the TCP server.
	osv := openshowvar.NewOpenShowVar("192.168.1.10", 7000)

	// Connect to the TCP server.
	err := osv.Connect()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer osv.Disconnect()

	// Defining the value to be written and the variable name.
	varName := "existing_var"
	newValue := "new_value"

	// Read a variable value.
	initialValue, err := osv.Read(varName)
	if err != nil {
		log.Fatalf("Failed to read variable: %v", err)
	}
	fmt.Printf("Initial value of %s: %s\n", varName, initialValue)

	// Write a new value to the variable.
	_, err = osv.Write(varName, newValue)
	if err != nil {
		log.Fatalf("Failed to write variable: %v", err)
	}
	fmt.Printf("Written new value to %s: %s\n", varName, newValue)

	// Read the variable value again to verify the change.
	updatedValue, err := osv.Read(varName)
	if err != nil {
		log.Fatalf("Failed to read variable: %v", err)
	}
	fmt.Printf("Updated value of %s: %s\n", varName, updatedValue)
}
```

## Contributing

Contributions are welcome! If you'd like to contribute to `go_openshowvar`, please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
