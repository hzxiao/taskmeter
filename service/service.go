package service

import "github.com/hzxiao/taskmeter/pkg/errno"

func newArgInvalidError(format string, message ...interface{}) error {
	err := errno.New(errno.ErrServiceArgInvalid, nil)
	err.Addf(format, message...)
	return err
}
