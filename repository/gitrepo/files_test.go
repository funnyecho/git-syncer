package gitrepo_test

import (
	"fmt"
	"testing"

	"github.com/funnyecho/git-syncer/repository/gitrepo"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter/gitter_test"
	"github.com/stretchr/testify/assert"
)

func TestListAllFiles(t *testing.T) {
	tcs := []struct {
		name                  string
		syncRoot              string
		getHeadSHA1           func() (string, error)
		getUnoPorcelainStatus func() (status []string, err error)
		getVersion            func() (majorVersion, minorVersion int, err error)
		listFiles             func(path string) ([]string, error)
		expectErred           bool
		expectHead            string
		expectFiles           []string
	}{
		{
			name: "failed to get git version",
			getVersion: func() (majorVersion int, minorVersion int, err error) {
				return 0, 0, fmt.Errorf("failed to get git version")
			},
			expectErred: true,
		},
		{
			name: "failed when git version < 1.7",
			getVersion: func() (majorVersion int, minorVersion int, err error) {
				return 0, 10, nil
			},
			expectErred: true,
		},
		{
			name: "failed when git version < 1.7",
			getVersion: func() (majorVersion int, minorVersion int, err error) {
				return 1, 6, nil
			},
			expectErred: true,
		},
		{
			name: "failed to check whether repo is dirty",
			getUnoPorcelainStatus: func() (status []string, err error) {
				return nil, fmt.Errorf("failed to check whether repo is dirty")
			},
			expectErred: true,
		},
		{
			name: "failed when repository is dirty",
			getUnoPorcelainStatus: func() (status []string, err error) {
				return []string{"file_a", "file_b"}, nil
			},
			expectErred: true,
		},
		{
			name: "failed to get head sha1",
			getHeadSHA1: func() (string, error) {
				return "", fmt.Errorf("failed to get head sha1")
			},
			expectErred: true,
		},
		{
			name:     "failed to list files",
			syncRoot: "./foo/bar/foo",
			listFiles: func(path string) ([]string, error) {
				assert.Equal(t, "./foo/bar/foo", path)
				return nil, fmt.Errorf("failed to list files with path: %s", path)
			},
			expectErred: true,
		},
		{
			name: "success to list files",
			getHeadSHA1: func() (string, error) {
				return "abcd1234", nil
			},
			listFiles: func(path string) ([]string, error) {
				assert.Equal(t, gitrepo.DefaultSyncRoot, path)
				return []string{"foo", "bar", "zoo"}, nil
			},
			expectErred: false,
			expectHead:  "abcd1234",
			expectFiles: []string{"foo", "bar", "zoo"},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			repo, _ := gitrepo.New(
				gitrepo.WithGitter(&filesTestGitter{
					gitter_test.New(),
					tc.syncRoot,
					tc.getHeadSHA1,
					tc.getUnoPorcelainStatus,
					tc.getVersion,
					tc.listFiles,
				}),
			)

			head, files, err := repo.ListAllFiles()
			if tc.expectErred {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tc.expectHead, head)
			assert.Equal(t, tc.expectFiles, files)
		})
	}
}

func TestListChangedFiles(t *testing.T) {

}

type filesTestGitter struct {
	gitter.Gitter
	syncRoot              string
	getHeadSHA1           func() (string, error)
	getUnoPorcelainStatus func() (status []string, err error)
	getVersion            func() (majorVersion, minorVersion int, err error)
	listFiles             func(path string) ([]string, error)
}

func (g *filesTestGitter) GetUnoPorcelainStatus() (status []string, err error) {
	if g.getUnoPorcelainStatus != nil {
		return g.getUnoPorcelainStatus()
	}

	return nil, nil
}

func (g *filesTestGitter) GetVersion() (majorVersion, minorVersion int, err error) {
	if g.getVersion != nil {
		return g.getVersion()
	}

	return 1, 8, nil
}

func (g *filesTestGitter) GetHeadSHA1() (string, error) {
	if g.getHeadSHA1 != nil {
		return g.getHeadSHA1()
	}

	return "", nil
}

func (g *filesTestGitter) ListFiles(path string) ([]string, error) {
	if g.listFiles != nil {
		return g.listFiles(path)
	}

	return nil, nil
}

func (g *filesTestGitter) ConfigGet(key string, _ gitter.ConfigGetOptions) (string, error) {
	if key == gitrepo.ConfigSyncRoot && g.syncRoot != "" {
		return g.syncRoot, nil
	}

	return "", nil
}
