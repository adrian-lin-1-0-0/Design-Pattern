package exporter

import "io"

type Base struct {
	output io.Writer
}

func (b *Base) SetOutput(output io.Writer) {
	b.output = output
}

func (b *Base) GetOutput() io.Writer {
	return b.output
}

func (b *Base) Export(message string) {
	b.output.Write([]byte(message))
}
