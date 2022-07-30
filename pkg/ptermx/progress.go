package ptermx

import "github.com/pterm/pterm"

type ProgressWriter struct {
	p *pterm.ProgressbarPrinter
}

func (pw ProgressWriter) Write(bytes []byte) (n int, err error) {
	n = len(bytes)
	pw.p.Add(n)
	return
}

func NewProgressWriter(p *pterm.ProgressbarPrinter) ProgressWriter {
	return ProgressWriter{p: p}
}
