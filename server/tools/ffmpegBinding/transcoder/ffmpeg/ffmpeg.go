package ffmpeg

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/tools/ffmpegBinding/transcoder"
	"github.com/flipped-aurora/gin-vue-admin/server/tools/ffmpegBinding/transcoder/utils"
)

// Transcoder ...
type Transcoder struct {
	config           *Config
	input            string
	timeout          int
	output           []string
	options          [][]string
	arguments        []string
	metadata         transcoder.Metadata
	inputPipeReader  *io.ReadCloser
	outputPipeReader *io.ReadCloser
	inputPipeWriter  *io.WriteCloser
	outputPipeWriter *io.WriteCloser
}

// New ...
func New(cfg *Config) transcoder.Transcoder {
	return &Transcoder{config: cfg}
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			//fmt.Println("runtime error:", err)
		default: // 非运行时错误
			//fmt.Println("error:", err)
		}
	}()
	entry()
}

func (t *Transcoder) Run(information *transcoder.Information) {

	information.Error = make(chan error)
	information.Progress = make(chan transcoder.Progress)
	options := t.options
	defer t.closePipes()

	var stderrIn io.ReadCloser

	err := t.validate()
	if err != nil {
		go func(err error, information *transcoder.Information) {
			information.Error <- err
		}(err, information)
		return
	}

	_, err = t.GetMetadata(information)
	if err != nil {
		go func(err error, information *transcoder.Information) {
			information.Error <- err
		}(err, information)
		return
	}

	// Append input file and standard options
	args := make([]string, 0)

	if t.timeout > 0 {
		args = append(args, "-timeout", strconv.Itoa(t.timeout))
	}
	args = append(args, "-i", t.input)
	for _, v := range options {
		args = append(args, v...)
	}

	args = append(args, t.output...)

	fmt.Println("t.output:", t.output)

	fmt.Println("args:", args)

	information.Cmd = exec.Command(t.config.FfmpegBinPath, args...)

	if t.config.ProgressEnabled && !t.config.Verbose {
		stderrIn, err = information.Cmd.StderrPipe()
		if err != nil {
			go func(err error, information *transcoder.Information) {
				information.Error <- err
			}(err, information)
			return
		}
	}

	if t.config.Verbose {
		information.Cmd.Stderr = os.Stdout
	}

	// Start process
	err = information.Cmd.Start()
	if err != nil {
		go func(err error, information *transcoder.Information) {
			information.Error <- err
		}(err, information)
		return
	}

	if t.config.ProgressEnabled && !t.config.Verbose {
		go func() {
			t.progress(stderrIn, information.Progress)
		}()

		go func() {
			defer close(information.Progress)
			err = information.Cmd.Wait()
			if err != nil {
				go func(err error) {
					information.Error <- err
				}(err)
			} else {
				close(information.Error)
			}

		}()
	} else {
		err = information.Cmd.Wait()
		if err != nil {
			go func(err error) {
				information.Error <- err
			}(err)
		}
	}

}

//func (t *Transcoder) Start(opts transcoder.Options,cmd *exec.Cmd) (<-chan transcoder.Progress, error ,<-chan error) {
//
//	done := make(chan error)
//	var stderrIn io.ReadCloser
//
//	out := make(chan transcoder.Progress)
//
//	defer t.closePipes()
//
//	// Validates config
//	err := t.validate()
//	if err != nil {
//		return nil, err ,done
//	}
//
//	var mcmd exec.Cmd
//	// Get file metadata
//	_, err = t.GetMetadata(&mcmd)
//	if err != nil {
//		return nil, err ,done
//	}
//
//	// Append input file and standard options
//	args := make([]string,0)
//
//	if t.timeout >0 {
//		args = append(args, "-timeout", strconv.Itoa(t.timeout))
//	}
//	args = append(args,"-i", t.input)
//	args = append(args,opts.GetStrArguments()... )
//	outputLength := len(t.output)
//	optionsLength := len(t.options)
//
//	if outputLength == 1 && optionsLength == 0 {
//		// Just append the 1 output file we've got
//		args = append(args, t.output[0])
//	} else {
//		for index, out := range t.output {
//			// Get executable flags
//			// If we are at the last output file but still have several options, append them all at once
//			if index == outputLength-1 && outputLength < optionsLength {
//				for i := index; i < len(t.options); i++ {
//					args = append(args, t.options[i]...)
//				}
//				// Otherwise just append the current options
//			} else {
//				args = append(args, t.options[index]...)
//			}
//
//			// Append output flag
//			args = append(args, out)
//		}
//	}
//
//	// Initialize command
//	cmd = exec.Command(t.config.FfmpegBinPath, args...)
//
//	// If progresss enabled, get stderr pipe and start progress process
//	if t.config.ProgressEnabled && !t.config.Verbose {
//		stderrIn, err = cmd.StderrPipe()
//		if err != nil {
//			return nil, fmt.Errorf("Failed getting transcoding progress (%s) with args (%s) with error %s", t.config.FfmpegBinPath, args, err),done
//		}
//	}
//
//	if t.config.Verbose {
//		cmd.Stderr = os.Stdout
//	}
//
//	// Start process
//	err = cmd.Start()
//	if err != nil {
//		return nil, fmt.Errorf("Failed starting transcoding (%s) with args (%s) with error %s", t.config.FfmpegBinPath, args, err),done
//	}
//
//	if t.config.ProgressEnabled && !t.config.Verbose {
//		go func() {
//			t.progress(stderrIn, out)
//		}()
//
//		go func() {
//			defer close(out)
//			err = cmd.Wait()
//			if err!=nil {
//				go func(err error) {
//					done <- err
//				}(err)
//			}else{
//				close(done)
//			}
//
//		}()
//	} else {
//		err = cmd.Wait()
//	}
//
//	return out,err,done
//}

func (t *Transcoder) TimeOut(arg int) transcoder.Transcoder {
	t.timeout = arg
	return t
}

// Input ...
func (t *Transcoder) Input(arg string) transcoder.Transcoder {
	t.input = arg
	return t
}

// Output ...
func (t *Transcoder) Output(arg string) transcoder.Transcoder {
	t.output = append(t.output, arg)
	return t
}

// InputPipe ...
func (t *Transcoder) InputPipe(w *io.WriteCloser, r *io.ReadCloser) transcoder.Transcoder {
	if &t.input == nil {
		t.inputPipeWriter = w
		t.inputPipeReader = r
	}
	return t
}

// OutputPipe ...
func (t *Transcoder) OutputPipe(w *io.WriteCloser, r *io.ReadCloser) transcoder.Transcoder {
	if &t.output == nil {
		t.outputPipeWriter = w
		t.outputPipeReader = r
	}
	return t
}

// WithOptions Sets the options object
func (t *Transcoder) WithOptions(opts transcoder.Options) transcoder.Transcoder {
	t.arguments = opts.GetStrArguments()
	t.options = [][]string{t.arguments}
	return t
}

// validate ...
func (t *Transcoder) validate() error {
	if t.config.FfmpegBinPath == "" {
		return errors.New("ffmpeg binary path not found")
	}

	if t.input == "" {
		return errors.New("missing input option")
	}

	outputLength := len(t.output)

	if outputLength == 0 {
		return errors.New("missing output option")
	}

	// length of output files being greater than length of options would produce an invalid ffmpegManage command
	// unless there is only 1 output file, which obviously wouldn't be a problem
	if outputLength > len(t.options) && outputLength != 1 {
		return errors.New("number of options and output files does not match")
	}

	for index, output := range t.output {
		if output == "" {
			return fmt.Errorf("output at index %d is an empty string", index)
		}
	}

	return nil
}

// GetMetadata Returns metadata for the specified input file
func (t *Transcoder) GetMetadata(information *transcoder.Information) (transcoder.Metadata, error) {

	if t.config.FfprobeBinPath != "" {
		var outb, errb bytes.Buffer

		input := t.input

		if t.inputPipeReader != nil {
			input = "pipe:"
		}
		args := make([]string, 0)
		if t.timeout > 0 {
			args = append(args, []string{"-timeout", strconv.Itoa(t.timeout)}...)
		}

		args = append(args, "-i", input, "-print_format", "json", "-show_format", "-show_streams", "-show_error")

		information.Cmd = exec.Command(t.config.FfprobeBinPath, args...)
		information.Cmd.Stdout = &outb
		information.Cmd.Stderr = &errb

		err := information.Cmd.Run()
		if err != nil {
			return nil, fmt.Errorf("error executing (%s) with args (%s) | error: %s | message: %s %s", t.config.FfprobeBinPath, args, err, outb.String(), errb.String())
		}

		var metadata Metadata

		if err = json.Unmarshal([]byte(outb.String()), &metadata); err != nil {
			return nil, err
		}

		t.metadata = metadata

		return metadata, nil
	}

	return nil, errors.New("ffprobe binary not found")
}

// progress sends through given channel the transcoding status
func (t *Transcoder) progress(stream io.ReadCloser, out chan transcoder.Progress) {

	defer stream.Close()

	split := func(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, data[0:i], nil
		}
		if i := bytes.IndexByte(data, '\r'); i >= 0 {
			// We have a cr terminated line
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	}

	scanner := bufio.NewScanner(stream)
	scanner.Split(split)

	buf := make([]byte, 2)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan() {
		Progress := new(Progress)
		line := scanner.Text()

		if strings.Contains(line, "frame=") && strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") {
			var re = regexp.MustCompile(`=\s+`)
			st := re.ReplaceAllString(line, `=`)

			f := strings.Fields(st)

			var framesProcessed string
			var currentTime string
			var currentBitrate string
			var currentSpeed string

			for j := 0; j < len(f); j++ {
				field := f[j]
				fieldSplit := strings.Split(field, "=")

				if len(fieldSplit) > 1 {
					fieldname := strings.Split(field, "=")[0]
					fieldvalue := strings.Split(field, "=")[1]

					if fieldname == "frame" {
						framesProcessed = fieldvalue
					}

					if fieldname == "time" {
						currentTime = fieldvalue
					}

					if fieldname == "bitrate" {
						currentBitrate = fieldvalue
					}
					if fieldname == "speed" {
						currentSpeed = fieldvalue
					}
				}
			}

			timesec := utils.DurToSec(currentTime)
			dursec, _ := strconv.ParseFloat(t.metadata.GetFormat().GetDuration(), 64)

			progress := (timesec * 100) / dursec
			Progress.Progress = progress

			Progress.CurrentBitrate = currentBitrate
			Progress.FramesProcessed = framesProcessed
			Progress.CurrentTime = currentTime
			Progress.Speed = currentSpeed
			ProtectRun(func() {
				out <- *Progress
			})
		}
	}
}

// closePipes Closes pipes if opened
func (t *Transcoder) closePipes() {
	if t.inputPipeReader != nil {
		ipr := *t.inputPipeReader
		ipr.Close()
	}

	if t.outputPipeWriter != nil {
		opr := *t.outputPipeWriter
		opr.Close()
	}
}
