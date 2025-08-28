package jsontypes

import (
	"database/sql/driver"
	"time"
)

// Time is a custom time type that serializes to JSON with millisecond precision (3 decimal places)
type Time time.Time

// MarshalJSON formatiert mit nur 3 Nachkommastellen
func (t Time) MarshalJSON() ([]byte, error) {
	truncated := time.Time(t).Truncate(time.Millisecond)
	return []byte(`"` + truncated.Format(time.RFC3339) + `"`), nil
}

// UnmarshalJSON parses a JSON string into Time
func (t *Time) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return nil
	}
	
	parsed, err := time.Parse(time.RFC3339, string(data[1:len(data)-1]))
	if err != nil {
		// Try with nanoseconds if RFC3339 fails
		parsed, err = time.Parse(time.RFC3339Nano, string(data[1:len(data)-1]))
		if err != nil {
			return err
		}
	}
	
	*t = Time(parsed)
	return nil
}

// GORM Scan interface - reads from database
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		*t = Time{}
		return nil
	}
	if v, ok := value.(time.Time); ok {
		*t = Time(v)
		return nil
	}
	return nil
}

// GORM Value interface - writes to database
func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// String returns the time formatted as RFC3339 with millisecond precision
func (t Time) String() string {
	return time.Time(t).Truncate(time.Millisecond).Format(time.RFC3339)
}

// Time returns the underlying time.Time
func (t Time) Time() time.Time {
	return time.Time(t)
}

// Now returns the current time as Time type
func Now() Time {
	return Time(time.Now())
}

// Parse parses a time string and returns a Time
func Parse(layout, value string) (Time, error) {
	t, err := time.Parse(layout, value)
	return Time(t), err
}