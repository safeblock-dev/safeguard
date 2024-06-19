# safeguard

Package `safeguard` provides utilities for error handling and reporting, with support for capturing and processing errors and panics. It offers functions to handle single errors, multiple errors, and errors returned from functions, while allowing customization for error handling scenarios.

## Installation

```bash
go get github.com/safeblock-dev/safeguard
```

## Features

- **Catch**: Captures and handles errors and panics.
- **CatchFn**: Captures and handles errors and panics from a function call.
- **CollectErrors**: Collects provided errors and converts panics to errors.
- **SkipErr**: Creates a specific error type (`skipErr`) for errors to be skipped during processing.
- **Customizable Error Handling**: Supports handlers for single errors, multiple errors, and slices of errors.

## Usage

### Catching and Handling Errors

```go
package main

import (
	"fmt"
	"github.com/safeblock-dev/safeguard"
)

func main() {
	// Example of Catch function
	safeguard.Catch(someFunction())
	
	// Example of CatchFn function
	safeguard.CatchFn(func() error {
		// Function logic that may return an error
		return someFunction()
	})
}

func someFunction() error {
	// Example function that may return an error
	return fmt.Errorf("example error")
}
```

### Custom Error Handling

```go
package main

import (
	"fmt"
	"github.com/safeblock-dev/safeguard"
)

func main() {
	// Example of handling specific errors
	err := someFunction()
	safeguard.Catch(err, handleSpecificError)
}

func handleSpecificError(err error) {
	// Example handler for a specific error type
	fmt.Println("Handling error:", err)
}
```

### Skipping Specific Errors

```go
package main

import (
	"fmt"
	"github.com/safeblock-dev/safeguard"
)

func main() {
	// Example of skipping specific errors
	err := fmt.Errorf("error to skip")
	safeguard.Catch(err, safeguard.SkipErr(err))
}
```
