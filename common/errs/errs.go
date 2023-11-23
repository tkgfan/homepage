// author gmfan
// date 2023/7/8
package errs

import "fmt"

type CodeErr struct {
	Code uint32
	Msg  string
}

func NewCodeErr(code uint32) *CodeErr {
	return &CodeErr{
		Code: code,
		Msg:  codeMsg[code],
	}
}

func NewCodeErrMgs(code uint32, msg string) *CodeErr {
	return &CodeErr{
		Code: code,
		Msg:  msg,
	}
}

func (t *CodeErr) Error() string {
	return fmt.Sprintf("Code: %d,Msg: %s", t.Code, t.Msg)
}

func IsCodeErr(e any) (err *CodeErr, ok bool) {
	err, ok = e.(*CodeErr)
	return
}
