package uddf

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)

func Parse(data []byte) (*UDDF, error) {
	var uddf UDDF
	decoder := xml.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&uddf); err != nil {
		return nil, fmt.Errorf("failed to decode UDDF file: %w", err)
	}

	return &uddf, nil
}

func ParseFile(filename string) (*UDDF, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return Parse(data)
}

//func ParseReader(r io.Reader) (*UDDF, error)

// Validate validates the UDDF structure using the validation tags defined in the structs
func (u *UDDF) Validate() error {
	if u == nil {
		return fmt.Errorf("UDDF object is nil")
	}

	validate := validator.New()
	return validate.Struct(u)
}
