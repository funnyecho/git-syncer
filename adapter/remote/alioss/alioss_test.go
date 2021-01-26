package alioss_test

// func TestGetHeadSHA1(t *testing.T) {
// 	tcs := []struct {
// 		name          string
// 		bucketMocking buckettest.Mocking
// 		expectErred   bool
// 		expectErrCode int
// 		expectSHA1    string
// 	}{
// 		{
// 			"failed when checking whether contrib locked failed",
// 			buckettest.Mocking{
// 				IsObjectExist: func(key string, options ...oss.Option) (bool, error) {
// 					assert.Equal(t, alioss.ObjectLockFile, key)
// 					return false, errors.NewError(errors.WithCode(exitcode.RemoteForbidden))
// 				},
// 			},
// 			true,
// 			exitcode.RemoteForbidden,
// 			"",
// 		},
// 		{
// 			"failed if contrib locked",
// 			buckettest.Mocking{
// 				IsObjectExist: func(key string, options ...oss.Option) (bool, error) {
// 					assert.Equal(t, alioss.ObjectLockFile, key)
// 					return true, nil
// 				},
// 			},
// 			true,
// 			exitcode.RemoteLocked,
// 			"",
// 		},
// 		{
// 			"failed when fetching head log",
// 			buckettest.Mocking{
// 				GetObject: func(key string, options ...oss.Option) (io.ReadCloser, error) {
// 					if key == alioss.ObjectHeadLinkFile {
// 						return nil, errors.NewError(errors.WithCode(exitcode.RemoteForbidden))
// 					}

// 					return ioutil.NopCloser(strings.NewReader("")), nil
// 				},
// 				PutObject: func(key string, reader io.Reader, options ...oss.Option) error {
// 					if key == alioss.ObjectLockFile {
// 						return nil
// 					}

// 					return errors.NewError(errors.WithCode(exitcode.RemoteSyncFailed))
// 				},
// 				IsObjectExist: func(key string, options ...oss.Option) (bool, error) {
// 					return false, nil
// 				},
// 			},
// 			true,
// 			exitcode.RemoteForbidden,
// 			"",
// 		},
// 		{
// 			"failed when head log is not valid json format",
// 			buckettest.Mocking{
// 				GetObject: func(key string, options ...oss.Option) (io.ReadCloser, error) {
// 					if key == alioss.ObjectHeadLinkFile {
// 						return ioutil.NopCloser(strings.NewReader("not a valid json file")), nil
// 					}

// 					return ioutil.NopCloser(strings.NewReader("")), nil
// 				},
// 				PutObject: func(key string, reader io.Reader, options ...oss.Option) error {
// 					if key == alioss.ObjectLockFile {
// 						return nil
// 					}

// 					return errors.NewError(errors.WithCode(exitcode.RemoteSyncFailed))
// 				},
// 			},
// 			true,
// 			exitcode.RemoteInvalidLog,
// 			"",
// 		},
// 		{
// 			"return sha1 in head log",
// 			buckettest.Mocking{
// 				GetObject: func(key string, options ...oss.Option) (io.ReadCloser, error) {
// 					if key == alioss.ObjectHeadLinkFile {
// 						logInfo := alioss.LogInfo{
// 							SHA1: "foobar",
// 							Date: alioss.JSONTime{time.Now()},
// 						}
// 						jsonInfo, _ := json.Marshal(logInfo)
// 						return ioutil.NopCloser(bytes.NewBuffer(jsonInfo)), nil
// 					}

// 					return ioutil.NopCloser(strings.NewReader("")), nil
// 				},
// 				PutObject: func(key string, reader io.Reader, options ...oss.Option) error {
// 					if key == alioss.ObjectLockFile {
// 						return nil
// 					}

// 					return errors.NewError(errors.WithCode(exitcode.RemoteSyncFailed))
// 				},
// 			},
// 			false,
// 			exitcode.Nil,
// 			"foobar",
// 		},
// 	}

// 	for _, tc := range tcs {
// 		t.Run(tc.name, func(t *testing.T) {
// 			opt := alioss.NewOptions(repotest.NewConfigReadWriter())
// 			bkt, _ := buckettest.New(tc.bucketMocking)

// 			c, _ := alioss.NewContrib(opt, bkt)

// 			sha1, err := c.GetHeadSHA1()

// 			if tc.expectErred {
// 				assert.Error(t, err)
// 				assert.Equal(t, tc.expectErrCode, errors.GetErrorCode(err))
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tc.expectSHA1, sha1)
// 			}
// 		})
// 	}
// }

// func TestSync(t *testing.T) {
// 	t.Run("failed test cases", func(t *testing.T) {
// 		tcs := []struct {
// 			name    string
// 			syncReq contrib.SyncReq

// 			bucketLockedUnknown      bool
// 			bucketLocked             bool
// 			bucketLockFailed         bool
// 			bucketUploadFailedCursor int
// 			bucketDeleteFailedCursor int
// 			bucketLogUploadFailed    bool
// 			bucketHeadLogLinkFailed  bool
// 			bucketUnlockFailed       bool

// 			expectErrCode int
// 			expectRes     contrib.SyncRes
// 		}{
// 			{
// 				name:                "failed when checking whether contrib locked",
// 				bucketLockedUnknown: true,
// 				expectErrCode:       exitcode.RemoteForbidden,
// 			},
// 			{
// 				name:          "failed when contrib locked",
// 				bucketLocked:  true,
// 				expectErrCode: exitcode.RemoteLocked,
// 			},
// 			{
// 				name:             "failed to upload lock to contrib",
// 				bucketLockFailed: true,
// 				expectErrCode:    exitcode.RemoteSyncFailed,
// 			},
// 			{
// 				name: "partial uploading failed",
// 				syncReq: contrib.SyncReq{
// 					SHA1:    "foobar",
// 					Uploads: []string{"foo", "bar", "zoo"},
// 					Deletes: []string{"foo", "bar", "zoo"},
// 				},
// 				bucketUploadFailedCursor: 2,
// 				expectErrCode:            exitcode.RemoteSyncFailed,
// 				expectRes: contrib.SyncRes{
// 					SHA1:     "",
// 					Uploaded: []string{"foo"},
// 					Deleted:  nil,
// 				},
// 			},
// 			{
// 				name: "partial deletion failed",
// 				syncReq: contrib.SyncReq{
// 					SHA1:    "foobar",
// 					Uploads: []string{"foo", "bar", "zoo"},
// 					Deletes: []string{"foo", "bar", "zoo"},
// 				},
// 				bucketDeleteFailedCursor: 2,
// 				expectErrCode:            exitcode.RemoteSyncFailed,
// 				expectRes: contrib.SyncRes{
// 					SHA1:     "",
// 					Uploaded: []string{"foo", "bar", "zoo"},
// 					Deleted:  []string{"foo"},
// 				},
// 			},
// 			{
// 				name: "upload contrib log failed",
// 				syncReq: contrib.SyncReq{
// 					SHA1:    "foobar",
// 					Uploads: []string{"foo", "bar", "zoo"},
// 					Deletes: []string{"foo", "bar", "zoo"},
// 				},
// 				bucketLogUploadFailed: true,
// 				expectErrCode:         exitcode.RemoteSyncFailed,
// 				expectRes: contrib.SyncRes{
// 					SHA1:     "",
// 					Uploaded: []string{"foo", "bar", "zoo"},
// 					Deleted:  []string{"foo", "bar", "zoo"},
// 				},
// 			},
// 			{
// 				name: "link head contrib log failed",
// 				syncReq: contrib.SyncReq{
// 					SHA1:    "foobar",
// 					Uploads: []string{"foo", "bar", "zoo"},
// 					Deletes: []string{"foo", "bar", "zoo"},
// 				},
// 				bucketHeadLogLinkFailed: true,
// 				expectErrCode:           exitcode.RemoteSyncFailed,
// 				expectRes: contrib.SyncRes{
// 					SHA1:     "",
// 					Uploaded: []string{"foo", "bar", "zoo"},
// 					Deleted:  []string{"foo", "bar", "zoo"},
// 				},
// 			},
// 			// {
// 			// 	name: "failed to unlock contrib",
// 			// 	syncReq: contrib.SyncReq{
// 			// 		SHA1:    "foobar",
// 			// 		Uploads: []string{"foo", "bar", "zoo"},
// 			// 		Deletes: []string{"foo", "bar", "zoo"},
// 			// 	},
// 			// 	bucketUnlockFailed: true,
// 			// 	expectErrCode:      exitcode.RemoteUnlock,
// 			// 	expectRes: contrib.SyncRes{
// 			// 		SHA1:     "foobar",
// 			// 		Uploaded: []string{"foo", "bar", "zoo"},
// 			// 		Deleted:  []string{"foo", "bar", "zoo"},
// 			// 	},
// 			// },
// 		}

// 		for _, tc := range tcs {
// 			t.Run(tc.name, func(t *testing.T) {
// 				uploadCursor := 0
// 				deleteCursor := 0

// 				opt := alioss.NewOptions(repotest.NewConfigReadWriter())
// 				bkt, _ := buckettest.New(buckettest.Mocking{
// 					GetObject: func(key string, options ...oss.Option) (io.ReadCloser, error) {
// 						return nil, nil
// 					},
// 					PutObject: func(key string, reader io.Reader, options ...oss.Option) error {
// 						if key == alioss.ObjectLockFile {
// 							if tc.bucketLockFailed {
// 								return errors.NewError(errors.WithCode(exitcode.RemoteForbidden))
// 							}
// 							return nil
// 						}

// 						if strings.Contains(key, alioss.ObjectLogDir) {
// 							if tc.bucketLogUploadFailed {
// 								return errors.NewError(errors.WithCode(exitcode.RemoteSyncFailed))
// 							}
// 							return nil
// 						}

// 						uploadCursor++
// 						if uploadCursor == tc.bucketUploadFailedCursor {
// 							return errors.NewError(errors.WithCode(exitcode.RemoteSyncFailed))
// 						}

// 						return nil
// 					},
// 					DeleteObject: func(key string, options ...oss.Option) error {
// 						if key == alioss.ObjectLockFile {
// 							if tc.bucketUnlockFailed {
// 								return errors.NewError(errors.WithCode(exitcode.RemoteForbidden))
// 							}
// 							return nil
// 						}

// 						deleteCursor++
// 						if deleteCursor == tc.bucketDeleteFailedCursor {
// 							return errors.NewError(errors.WithCode(exitcode.RemoteSyncFailed))
// 						}

// 						return nil
// 					},
// 					IsObjectExist: func(key string, options ...oss.Option) (bool, error) {
// 						if key == alioss.ObjectLockFile {
// 							if tc.bucketLockedUnknown {
// 								return false, errors.NewError(errors.WithCode(exitcode.RemoteForbidden))
// 							}

// 							if tc.bucketLocked {
// 								return true, nil
// 							}
// 						}

// 						return false, nil
// 					},
// 					PutSymlink: func(symObjectKey, targetObjectKey string, options ...oss.Option) error {
// 						if symObjectKey == alioss.ObjectHeadLinkFile {
// 							if tc.bucketHeadLogLinkFailed {
// 								return errors.NewError(errors.WithCode(exitcode.RemoteForbidden))
// 							}
// 						}
// 						return nil
// 					},
// 				})

// 				c, _ := alioss.NewContrib(opt, bkt)

// 				res, err := c.Sync(&tc.syncReq)

// 				assert.Error(t, err)
// 				// assert.Equal(t, tc.expectErrCode, errors.GetErrorCode(err))
// 				assert.Equal(t, contrib.SyncRes{}, res)
// 			})
// 		}
// 	})
// }
