package gitrepo_test

import (
	"fmt"
	"testing"

	"github.com/funnyecho/git-syncer/repository/gitrepo"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter/gittertest"
	"github.com/stretchr/testify/assert"
)

func TestHeadReader(t *testing.T) {
	t.Run("TestGetHead", func(t *testing.T) {
		tcs := []struct {
			name            string
			getHead         func() (string, error)
			getHeadSHA1     func() (string, error)
			getSymbolicHead func() (string, error)
			expectHeadErred bool
			expectHead      string
		}{
			{
				name: "return non-empty symbolic head",
				getSymbolicHead: func() (string, error) {
					return "symbolic/head", nil
				},
				getHead: func() (string, error) {
					assert.Fail(t, "shall not step into `getHead`")
					return "", nil
				},
				getHeadSHA1: func() (string, error) {
					assert.Fail(t, "shall not step into `getHeadSHA1`")
					return "", nil
				},
				expectHeadErred: false,
				expectHead:      "symbolic/head",
			},
			{
				name: "return non-empty head",
				getSymbolicHead: func() (string, error) {
					return "", nil
				},
				getHead: func() (string, error) {
					return "head", nil
				},
				getHeadSHA1: func() (string, error) {
					assert.Fail(t, "shall not step into `getHeadSHA1`")
					return "", nil
				},
				expectHeadErred: false,
				expectHead:      "head",
			},
			{
				name: "return non-empty head sha1",
				getSymbolicHead: func() (string, error) {
					return "", nil
				},
				getHead: func() (string, error) {
					return "", nil
				},
				getHeadSHA1: func() (string, error) {
					return "abcd1234", nil
				},
				expectHeadErred: false,
				expectHead:      "abcd1234",
			},
			{
				name: "failed with empty symbolic head, head and head sha1",
				getSymbolicHead: func() (string, error) {
					return "", nil
				},
				getHead: func() (string, error) {
					return "", nil
				},
				getHeadSHA1: func() (string, error) {
					return "", nil
				},
				expectHeadErred: true,
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				repo, _ := gitrepo.New(
					gitrepo.WithGitter(&headTestGitter{
						gittertest.New(),
						nil,
						tc.getHead,
						tc.getHeadSHA1,
						tc.getSymbolicHead,
					}),
				)

				head, err := repo.GetHead()
				if tc.expectHeadErred {
					assert.Error(t, err)
					return
				}

				assert.Equal(t, tc.expectHead, head)
			})
		}
	})

	t.Run("TestGetHeadSHA1", func(t *testing.T) {
		tcs := []struct {
			name            string
			getHeadSHA1     func() (string, error)
			expectHeadErred bool
			expectHead      string
		}{
			{
				name: "return non-empty head sha1",
				getHeadSHA1: func() (string, error) {
					return "abcd1234", nil
				},
				expectHeadErred: false,
				expectHead:      "abcd1234",
			},
			{
				name: "failed to get sha1 from gitter",
				getHeadSHA1: func() (string, error) {
					return "", fmt.Errorf("failed to get sha1 from gitter")
				},
				expectHeadErred: true,
			},
			{
				name: "failed with empty head sha1",
				getHeadSHA1: func() (string, error) {
					return "", nil
				},
				expectHeadErred: true,
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				repo, _ := gitrepo.New(
					gitrepo.WithGitter(&headTestGitter{
						gittertest.New(),
						nil,
						nil,
						tc.getHeadSHA1,
						nil,
					}),
				)

				head, err := repo.GetHeadSHA1()
				if tc.expectHeadErred {
					assert.Error(t, err)
					return
				}

				assert.Equal(t, tc.expectHead, head)
			})
		}
	})
}

func TestHeadWriter(t *testing.T) {

}

type headTestGitter struct {
	gitter.Gitter
	checkout        func(head string) error
	getHead         func() (string, error)
	getHeadSHA1     func() (string, error)
	getSymbolicHead func() (string, error)
}

func (g *headTestGitter) Checkout(head string) error {
	if g.checkout != nil {
		return g.checkout(head)
	}
	return nil
}

func (g *headTestGitter) GetHead() (string, error) {
	if g.getHead != nil {
		return g.getHead()
	}
	return "", nil
}
func (g *headTestGitter) GetHeadSHA1() (string, error) {
	if g.getHeadSHA1 != nil {
		return g.getHeadSHA1()
	}
	return "", nil
}

func (g *headTestGitter) GetSymbolicHead() (string, error) {
	if g.getSymbolicHead != nil {
		return g.getSymbolicHead()
	}
	return "", nil
}
