package alioss

import (
	"fmt"
	"strings"
	"time"
)

// JSONTime jsonable of time.Time
type JSONTime struct {
	time.Time
}

// JSONTimeLayout time format
const JSONTimeLayout = time.RFC3339

// MarshalJSON stringify from time.Time
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t.Time).Format(JSONTimeLayout))
	return []byte(stamp), nil
}

// UnmarshalJSON parse to time.Time
func (t JSONTime) UnmarshalJSON(bs []byte) error {
	tm, tErr := time.Parse(JSONTimeLayout, strings.Trim(string(bs), "\""))
	if tErr != nil {
		return tErr
	}

	t.Time = tm
	return nil
}
