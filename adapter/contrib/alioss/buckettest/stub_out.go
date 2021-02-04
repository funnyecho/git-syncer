package buckettest

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// StubOut mock calling output
type StubOut struct {
	GetObjectCallTimes   int
	GetObjectCallKey     []string
	GetObjectCallOptions [][]oss.Option

	PutObjectCallTimes   int
	PutObjectCallKey     []string
	PutObjectCallReader  []io.Reader
	PutObjectCallOptions [][]oss.Option

	DeleteObjectCallTimes   int
	DeleteObjectCallKey     []string
	DeleteObjectCallOptions [][]oss.Option

	IsObjectExistCallTimes   int
	IsObjectExistCallKey     []string
	IsObjectExistCallOptions [][]oss.Option

	PutSymlinkCallTimes     int
	PutSymlinkCallSymKey    []string
	PutSymlinkCallTargetKey []string
	PutSymlinkCallOptions   [][]oss.Option
}
