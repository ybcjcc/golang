package xerror_test

import (
	"git.medlinker.com/golang/xerror"
	"git.medlinker.com/golang/xerror/ecode"
	"git.medlinker.com/golang/xerror/xcode"
	"testing"
)

func TestNew(t *testing.T) {
	t.Log(xcode.New(500))
	t.Log(xerror.New("this is new error"))

	err := xerror.WithXCode(xcode.InternalServerError)
	t.Log(err, " / ", err.ToMessage(&ecode.Config{}))
}

func TestNewWithCode(t *testing.T) {
	err := xerror.WithCode(100401, "木有登陆")
	t.Log(err)
	t.Log(err.ToMessage(nil))
}

func TestNewWithXCode(t *testing.T) {
	err := xerror.WithXCode(xcode.InternalServerError)
	t.Log(err)
	t.Log(err.ToMessage(nil))
}

func TestWithXCodeMessage(t *testing.T) {
	err := xerror.WithXCodeMessage(xcode.InternalServerError, "变更消息")
	t.Log(err)
	t.Log(err.ToMessage(nil))
}
