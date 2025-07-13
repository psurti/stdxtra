# stdxtra

A Go library providing extra standard utilities for Go applications.

## Features

Currently, the library includes:

- `value` package: A generic Value type implementation that can hold any type of value with Get/Set methods

## Requirements

- Go 1.24 or higher (for generics support)

## Installation

```bash
go get github.com/yourusername/stdxtra
```

## Usage

```go
import "github.com/yourusername/stdxtra/domain/value"

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
```

## License

[Add your license here]

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.