package log

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func MyCallerFuncHandler(h Handler) Handler {
	return FuncHandler(func(r *Record) error {
		r.Ctx = append(r.Ctx, "fn", fmt.Sprintf("%s:%n:%d", r.Call, r.Call, r.Call))
		return h.Log(r)
	})
}
func MyStreamHandler2(wr io.Writer) Handler {
	fmtr := TerminalFormat(true)
	h := FuncHandler(func(r *Record) error {
		_, err := wr.Write(fmtr.Format(r))
		return err
	})
	return LazyHandler(SyncHandler(MyCallerFuncHandler(h)))
}
func TestRoot(t *testing.T) {
	Root().SetHandler(LvlFilterHandler(LvlTrace, MyStreamHandler2(os.Stderr)))
	Trace("hello,world")
}
