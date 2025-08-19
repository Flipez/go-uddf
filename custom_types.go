package uddf

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type FlexibleFloat struct {
	Value *float64
}

func (f *FlexibleFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	content = strings.TrimSpace(content)
	if content == "" {
		return nil // Empty content, leave as nil
	}

	// Try to parse as float
	if val, err := strconv.ParseFloat(content, 64); err == nil {
		f.Value = &val
		return nil
	}

	// If parsing fails, just ignore and leave as nil
	log.Printf("Skipping invalid float: %s", content)
	return nil
}

type Time time.Time

func (t *Time) parseTimeString(dateStr string) error {
	// Trim any whitespace
	dateStr = strings.TrimSpace(dateStr)

	// Try to parse the date in multiple formats
	formats := []string{
		time.RFC3339,           // 2006-01-02T15:04:05Z07:00
		"2006-01-02",           // YYYY-MM-DD
		"2006-01-02T15:04:05Z", // ISO 8601 UTC
		"2006-01-02T15:04:05",  // ISO 8601 UTC without timezone
		"2006-01-02T15:04",     // YYYY-MM-DDTHH:MM (without seconds)
		"2006",                 // YYYY (year only)
	}

	var parsedTime time.Time
	var err error

	for _, format := range formats {
		parsedTime, err = time.Parse(format, dateStr)
		if err == nil {
			break
		}
	}

	if err != nil {
		return fmt.Errorf("unable to parse datetime '%s': %w", dateStr, err)
	}

	*t = Time(parsedTime)
	return nil
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var dateStr string
	if err := d.DecodeElement(&dateStr, &start); err != nil {
		return err
	}
	return t.parseTimeString(dateStr)
}

func (t *Time) UnmarshalXMLAttr(attr xml.Attr) error {
	return t.parseTimeString(attr.Value)
}
