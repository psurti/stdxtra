# stdxtra

A Go library providing extra standard utilities for Go applications.

## Features

Currently, the library includes:

- `domain/entity`: A generic Entity type with ID support
  - Provides a flexible Entity structure with different ID types (string, integer)
  - Includes a builder pattern for entity creation with options

- `domain/measurement`: Types for handling various measurements
  - `Temperature`: Handles temperature values with conversion between Celsius, Fahrenheit, and Kelvin
  - `Measurable`: Interface for all measurement types with string representation and equality comparison

- `domain/value`: A generic Value type implementation
  - Holds any type of value with Get/Set methods
  - Provides both object methods and package functions

## Requirements

- Go 1.24 or higher (for generics support)

## Installation

```bash
go get github.com/psurti/stdxtra
```

## Usage

### Value Package

```go
package main

import "github.com/psurti/stdxtra/domain/value"

func main() {
    // Create a new value
    v := value.New(42)

    // Get the value
    val := v.Get() // 42

    // Set a new value
    v.Set(100)
    val = v.Get() // 100

    // Package functions are also available
    var v2 value.Value[string]
    value.Set(&v2, "hello")
    str := value.Get(&v2) // "hello"
}
```

### Entity Package

```go
package main

import "github.com/psurti/stdxtra/domain/entity"

func main() {
    // Create an entity with a string ID
    e1 := entity.New[string](entity.WithStrID[string]("user-123"))
    id := e1.ID() // Returns the ID interface
    strID := e1.StrID() // Returns "user-123"

    // Create an entity with an integer ID
    e2 := entity.New[int](entity.WithIntID[int](42))

    // Use the entities...
    _ = id
    _ = strID
    _ = e2
}
```

### Measurement Package

```go
package main

import (
    "fmt"
    "github.com/psurti/stdxtra/domain/measurement"
)

func main() {
    // Create a temperature in Celsius
    temp := measurement.NewTemperature(25.0, measurement.Celsius)

    // Convert to other units
    fahrenheit := temp.Fahrenheit() // 77.0
    kelvin := temp.Kelvin() // 298.15

    // String representation
    fmt.Println(temp) // "25.00C"

    // Create a temperature in Fahrenheit
    tempF := measurement.NewTemperature(98.6, measurement.Fahrenheit)
    celsius := tempF.Celsius() // 37.0

    // Use the values...
    _ = fahrenheit
    _ = kelvin
    _ = celsius
}
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
