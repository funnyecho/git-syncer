package syncer

import (
	"github.com/funnyecho/git-syncer/pkg/gitter"
)

var wipLocalSHA1 = ""

func SetupLocalSHA1() (err error) {
	sha1, err := getLocalSHA1()
	if err != nil {
		return err
	}

	wipLocalSHA1 = sha1
	return nil
}

func getLocalSHA1() (string, error) {
	return gitter.GetHeadSHA1()
}