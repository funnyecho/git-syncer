package alioss

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJSONTime(t *testing.T) {
	ut, _ := time.Parse(JSONTimeLayout, "2006-01-02T15:04:05+08:00")

	jt := JSONTime{ut}

	marshalJT, marshalErr := jt.MarshalJSON()
	assert.Nil(t, marshalErr)
	assert.Equal(t, `"2006-01-02T15:04:05+08:00"`, string(marshalJT))

	ut = ut.Add(time.Hour * 24)
	unmarshalErr := jt.UnmarshalJSON([]byte(ut.Format(JSONTimeLayout)))

	assert.Nil(t, unmarshalErr)
	assert.Equal(t, 3, jt.Day())
}
