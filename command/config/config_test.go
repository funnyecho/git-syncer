package config_test

import (
	"testing"

	"github.com/funnyecho/git-syncer/command/config"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	tcs := []struct {
		name          string
		key           string
		getConfig     func(key string) (string, error)
		expectErred   bool
		expectErrCode int
		expectValue   string
	}{
		{
			"failed when key is empty",
			"",
			nil,
			true,
			exitcode.MissingArguments,
			"",
		},
		{
			"config reader failed",
			"abc",
			func(key string) (string, error) {
				assert.Equal(t, "abc", key)
				return "", errors.NewError(errors.WithCode(exitcode.RepoUnknown))
			},
			true,
			exitcode.RepoUnknown,
			"",
		},
		{
			"failed when returning empty config",
			"abc",
			func(key string) (string, error) {
				assert.Equal(t, "abc", key)
				return "", nil
			},
			true,
			exitcode.RepoConfigNotFound,
			"",
		},
		{
			"return non-empty config",
			"abc",
			func(key string) (string, error) {
				assert.Equal(t, "abc", key)
				return "foobar", nil
			},
			false,
			exitcode.Nil,
			"foobar",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			v, err := config.GetConfig(tc.key, &mockConfigReader{
				tc.getConfig,
			})

			if tc.expectErred {
				assert.Error(t, err)
				assert.Equal(t, tc.expectErrCode, errors.GetErrorCode(err))
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectValue, v)
			}
		})
	}
}

func TestUpdateConfig(t *testing.T) {
	tcs := []struct {
		name          string
		key           string
		value         string
		updateErr     error
		expectErred   bool
		expectErrCode int
	}{
		{
			"failed when key is empty",
			"",
			"d",
			nil,
			true,
			exitcode.MissingArguments,
		},
		{
			"failed when update config failed",
			"foo",
			"bar",
			errors.NewError(errors.WithCode(exitcode.RepoUnknown)),
			true,
			exitcode.RepoUnknown,
		},
		{
			"complete when update config succesfully",
			"foo",
			"bar",
			nil,
			false,
			exitcode.Nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := config.UpdateConfig(tc.key, tc.value, &mockConfigWriter{
				func(key, value string) error {
					assert.Equal(t, tc.key, key)
					return tc.updateErr
				},
			})

			if tc.expectErred {
				assert.Error(t, err)
				assert.Equal(t, tc.expectErrCode, errors.GetErrorCode(err))
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

type mockConfigReader struct {
	getConfig func(key string) (string, error)
}

func (c *mockConfigReader) GetConfig(key string) (string, error) {
	if c.getConfig != nil {
		return c.getConfig(key)
	}

	return "", nil
}

type mockConfigWriter struct {
	setConfig func(key, value string) error
}

func (m *mockConfigWriter) SetConfig(key, value string) error {
	if m.setConfig != nil {
		return m.setConfig(key, value)
	}

	return nil
}
