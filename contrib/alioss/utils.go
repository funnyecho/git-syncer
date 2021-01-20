package alioss

import (
	"fmt"
	"time"
)

// JSONTime jsonable of time.Time
type JSONTime time.Time

// MarshalJSON stringify from time.Time
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan 18 16:41:44 2021 +0800"))
	return []byte(stamp), nil
}
