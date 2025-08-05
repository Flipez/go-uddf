package uddf

import (
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("nil UDDF should return error", func(t *testing.T) {
		var u *UDDF
		err := u.Validate()
		if err == nil {
			t.Error("expected error for nil UDDF, got nil")
		}
	})

	t.Run("valid UDDF file should pass validation", func(t *testing.T) {
		uddf, err := ParseFile("testdata/valid.uddf")
		if err != nil {
			t.Fatalf("failed to parse valid UDDF file: %v", err)
		}

		err = uddf.Validate()
		if err != nil {
			t.Errorf("expected no validation errors for valid UDDF, got: %v", err)
		}
	})

	t.Run("invalid gas fraction should fail validation", func(t *testing.T) {
		uddf, err := ParseFile("testdata/invalid_gas_fraction.uddf")
		if err != nil {
			t.Fatalf("failed to parse UDDF file: %v", err)
		}

		err = uddf.Validate()
		if err == nil {
			t.Error("expected validation error for invalid gas fraction > 1.0, got nil")
		}
	})

	t.Run("missing required Mix name should fail validation", func(t *testing.T) {
		uddf, err := ParseFile("testdata/missing_mix_name.uddf")
		if err != nil {
			t.Fatalf("failed to parse UDDF file: %v", err)
		}

		err = uddf.Validate()
		if err == nil {
			t.Error("expected validation error for missing required Mix name, got nil")
		}
	})

	t.Run("invalid Problems keyword should fail validation", func(t *testing.T) {
		uddf, err := ParseFile("testdata/invalid_problems.uddf")
		if err != nil {
			t.Fatalf("failed to parse UDDF file: %v", err)
		}

		err = uddf.Validate()
		if err == nil {
			t.Error("expected validation error for invalid Problem keyword, got nil")
		}
	})

	t.Run("invalid Program keyword should fail validation", func(t *testing.T) {
		uddf, err := ParseFile("testdata/invalid_program.uddf")
		if err != nil {
			t.Fatalf("failed to parse UDDF file: %v", err)
		}

		err = uddf.Validate()
		if err == nil {
			t.Error("expected validation error for invalid Program keyword, got nil")
		}
	})
}

func TestParseFile(t *testing.T) {
	t.Run("should parse valid UDDF file", func(t *testing.T) {
		uddf, err := ParseFile("testdata/valid.uddf")
		if err != nil {
			t.Fatalf("failed to parse valid UDDF file: %v", err)
		}

		if uddf.Version != "3.2.3" {
			t.Errorf("expected version '3.2.3', got '%s'", uddf.Version)
		}

		if uddf.GasDefinitions == nil || len(uddf.GasDefinitions.Mixes) == 0 {
			t.Error("expected gas definitions with mixes")
		}

		if len(uddf.ProfileData.RepetitionGroup) == 0 {
			t.Error("expected at least one repetition group")
		}
	})

	t.Run("should return error for non-existent file", func(t *testing.T) {
		_, err := ParseFile("testdata/non_existent.uddf")
		if err == nil {
			t.Error("expected error for non-existent file, got nil")
		}
	})
}