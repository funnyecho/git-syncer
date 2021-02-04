package alioss

import (
	"bytes"
	"encoding/json"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

// LogInfo log info
type LogInfo struct {
	SHA1     string         `json:"SHA1"`
	RefSHA1  string         `json:"RefSHA1"`
	Date     JSONTime       `json:"Date"`
	Uploaded []UploadedFile `json:"Uploaded"`
	Deleted  []string       `json:"Deleted"`
}

// UploadedFile uploaded file to log
type UploadedFile struct {
	Path string `json:"Path"`
	Size int64  `json:"Size"`
	SHA1 string `json:"SHA1"`
}

// PushLog log to contrib
func (a *Alioss) PushLog(info LogInfo) error {
	if info.SHA1 == "" {
		return errors.NewError(
			errors.WithCode(exitcode.MissingArguments),
			errors.WithMsg("sha1 of log requried"),
		)
	}

	jsonLog, jsonErr := json.Marshal(info)
	if jsonErr != nil {
		return errors.NewError(
			errors.WithCode(exitcode.Unknown),
			errors.WithErr(jsonErr),
			errors.WithMsgf("failed to marshal log info: %v", info),
		)
	}

	logPath := filepath.Join(ObjectLogDir, info.SHA1)
	_, uploadErr := a.uploadObject(logPath, bytes.NewReader(jsonLog), oss.ObjectACL(oss.ACLPrivate))
	if uploadErr != nil {
		return errors.NewError(
			errors.WithCode(exitcode.ContribSyncFailed),
			errors.WithErr(uploadErr),
			errors.WithMsgf("failed to upload log file to %s", logPath),
		)
	}

	headLogErr := a.putSymlink(logPath, ObjectHeadLinkFile, oss.ObjectACL(oss.ACLPrivate))
	if headLogErr != nil {
		return errors.NewError(
			errors.WithCode(exitcode.ContribForbidden),
			errors.WithErr(headLogErr),
			errors.WithMsgf("failed to link head log file from %s to %s", logPath, ObjectHeadLinkFile),
		)
	}

	return nil
}

// PeekLog get head log
func (a *Alioss) PeekLog() (*LogInfo, error) {
	headLogReader, headLogReaderErr := a.getObject(ObjectHeadLinkFile)
	if headLogReaderErr != nil {
		if IsObjectNotFoundErr(headLogReaderErr) {
			return &LogInfo{}, nil
		}

		return nil, errors.NewError(
			errors.WithErr(headLogReaderErr),
			errors.WithMsg("failed to download head log file"),
		)
	}

	infoBuf := new(bytes.Buffer)
	_, infoReadErr := infoBuf.ReadFrom(headLogReader)
	if infoReadErr != nil {
		return nil, errors.NewError(
			errors.WithErr(infoReadErr),
			errors.WithMsg("failed to read contrib head log file"),
			errors.WithCode(exitcode.ContribForbidden),
		)
	}

	var info LogInfo
	jsonErr := json.Unmarshal(infoBuf.Bytes(), &info)
	if jsonErr != nil {
		return nil, errors.NewError(
			errors.WithErr(jsonErr),
			errors.WithMsgf("failed to unmarshal contrib head log file: %s", infoBuf.String()),
			errors.WithCode(exitcode.ContribInvalidLog),
		)
	}

	return &info, nil
}
