package alioss

import (
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/repository"
)

func NewContribFactory() contrib.Factory {
	return func(options interface{ repository.ConfigReader }) contrib.Contrib {
		c := &alioss{
			ConfigReader: options,
		}

		return c
	}
}

type alioss struct {
	repository.ConfigReader
}

func (a *alioss) GetHeadSHA1() (string, error) {
	panic("implement me")
}

func (a *alioss) Sync(reqx *contrib.SyncReq) (contrib.SyncRes, error) {
	panic("implement me")
}
