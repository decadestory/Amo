package ammodel

import (
	"time"
)

type Time time.Time

func (t Time) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil
}

// UnmarshalJSON is joson.Unmarshaler
func (ct *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t, _ := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*ct = Time(t)
	return err
}
