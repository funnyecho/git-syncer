package alioss

import (
	"bytes"
	"encoding/json"
	"path/filepath"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

/*
Log item format:

SHA1: 664c7be795e0dce15586207234bdb2ab0d7da844
Executor: SamHwang1990 <samhwang1990@gmail.com>
Date:   Mon Jan 18 16:41:44 2021 +0800

 upload [path/to/uploaded/file] [file size] [file sha1]
 delete [path/to/deleted/file] [file size] [file sha1]

*/

// LogInfo log info
type LogInfo struct {
	SHA1     string         `json:"SHA1"`
	RefSHA1  string         `json:"RefSHA1"`
	Executor string         `json:"Executor"`
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
func (a *Alioss) pushLog(info LogInfo) error {
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

	logPath := filepath.Join(objectLogDir, info.SHA1)
	_, uploadErr := a.uploadObject(logPath, bytes.NewReader(jsonLog))
	if uploadErr != nil {
		return errors.NewError(
			errors.WithCode(exitcode.ContribSyncFailed),
			errors.WithErr(uploadErr),
			errors.WithMsgf("failed to upload log file to %s", logPath),
		)
	}

	headLogErr := a.putSymlink(logPath, objectHeadLinkFile)
	if headLogErr != nil {
		return errors.NewError(
			errors.WithCode(exitcode.ContribUnknown),
			errors.WithErr(headLogErr),
			errors.WithMsgf("failed to link head log file from %s to %s", logPath, objectHeadLinkFile),
		)
	}

	return nil
}

// PeekLog get head log
func (a *Alioss) peekLog() (*LogInfo, error) {
	headLogReader, headLogReaderErr := a.getObject(objectHeadLinkFile)
	if headLogReaderErr != nil {
		return nil, errors.NewError(
			errors.WithErr(headLogReaderErr),
			errors.WithMsg("failed to download head log file"),
			errors.WithCode(exitcode.ContribForbidden),
		)
	}

	infoBuf := new(bytes.Buffer)
	_, infoReadErr := infoBuf.ReadFrom(headLogReader)
	if infoReadErr != nil {
		return nil, errors.NewError(
			errors.WithErr(infoReadErr),
			errors.WithMsg("failed to read contrib head log file"),
			errors.WithCode(exitcode.ContribUnknown),
		)
	}

	var info LogInfo
	jsonErr := json.Unmarshal(infoBuf.Bytes(), &info)
	if jsonErr != nil {
		return nil, errors.NewError(
			errors.WithErr(jsonErr),
			errors.WithMsgf("failed to unmarshal contrib head log file: %s", infoBuf.String()),
			errors.WithCode(exitcode.Unknown),
		)
	}

	return &info, nil
}
