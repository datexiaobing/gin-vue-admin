package transcoder

import (
	"io"
	"os/exec"
)

type  Information struct {
	Progress         chan Progress
	Error            chan error
	Cmd              *exec.Cmd
}

// Transcoder ...
type Transcoder interface {
	Run(information *Information)
	Input(i string) Transcoder
	TimeOut(i int) Transcoder
	InputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	Output(o string) Transcoder
	OutputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	WithOptions(opts Options) Transcoder
	GetMetadata(information *Information) (Metadata, error)
}

