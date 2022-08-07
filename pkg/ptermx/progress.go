package ptermx

import (
	"github.com/fzdwx/infinite/components"
)

type ProgressWriter struct {
	p *components.Progress
}

func (pw ProgressWriter) Write(bytes []byte) (n int, err error) {
	n = len(bytes)
	pw.p.Incr(int64(n))
	return
}

func NewProgressWriter(p *components.Progress) ProgressWriter {
	return ProgressWriter{p: p}
}
