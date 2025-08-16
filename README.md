# go-uddf

A Go library for parsing and validating Universal Dive Data Format (UDDF) files.

## Overview

UDDF is an XML-based format used to store dive computer data, dive logs, and related diving information. This library provides:

- XML parsing of UDDF files into Go structs
- Validation using struct tags
- Flexible data handling for non-standard implementations

## Pointer Usage

This library extensively uses pointers for both optional and required fields to maintain maximum flexibility with UDDF data. Many manufacturers do not strictly follow the UDDF standard, and pointers allow the library to distinguish between missing data and zero values. Validation is used to enforce standard compliance when needed, rather than strict struct requirements.

## Usage

```go
package main

import (
    "log"
    "github.com/Flipez/go-uddf"
)

func main() {
    // Parse from file
    data, err := uddf.ParseFile("dive.uddf")
    if err != nil {
        log.Fatal(err)
    }

    // Parse from bytes
    content := []byte(`<uddf version="3.0">...</uddf>`)
    data, err = uddf.Parse(content)
    if err != nil {
        log.Fatal(err)
    }

    // Validate structure
    if err := data.Validate(); err != nil {
        log.Printf("Validation failed: %v", err)
    }
}
```

## Data Types

### Custom Types

- `FlexibleFloat`: Handles numeric fields that may contain invalid data
- `Time`: Supports multiple datetime formats for broad compatibility

## Validation

The library uses `github.com/go-playground/validator/v10` for field validation. Validation tags enforce:

- Required fields (e.g., `validate:"required"`)
- Value ranges (e.g., `validate:"min=0,max=1"`)
- Enumerated values (e.g., `validate:"oneof=recreation training scientific"`)

## Testing

Run tests with:

```bash
go test
```

Test data is provided in the `testdata/` directory for various scenarios including invalid data handling.