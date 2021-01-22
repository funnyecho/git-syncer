package alioss

import (
	"errors"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// AsOssError try to convert to oss service error
func AsOssError(err error) *oss.ServiceError {
	var ossErr oss.ServiceError

	if errors.As(err, &ossErr) {
		return &ossErr
	}

	return nil
}

// IsObjectNotFoundErr object not found err
func IsObjectNotFoundErr(err error) bool {
	if ossErr := AsOssError(err); ossErr != nil {
		return ossErr.StatusCode == 404
	}

	return false
}
