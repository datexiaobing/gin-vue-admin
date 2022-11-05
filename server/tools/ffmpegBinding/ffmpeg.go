package ffmpegBinding

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	log "github.com/flipped-aurora/gin-vue-admin/server/go-admin-core/logger"
	"github.com/flipped-aurora/gin-vue-admin/server/tools/ffmpegBinding/transcoder/utils"

	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type Information struct {
	Progress chan Progress
	Error    chan error
	Cmd      *exec.Cmd
}

type Progress struct {
	FramesProcessed string
	CurrentTime     string
	CurrentBitrate  string
	Progress        float64
	Speed           string
}

type FFmpeg struct {
	config      *Config
	input       string
	timeout     int
	output      string
	options     [][]string
	isShCms     bool
	metadata    *Metadata
	information *Information
	Progress    float64
	UserAgent   []string
}

func New(config *Config) *FFmpeg {
	FFmpeg := &FFmpeg{}
	FFmpeg.config = config
	return FFmpeg
}

func (e *FFmpeg) SetInput(input string) *FFmpeg {
	e.input = input
	return e
}

func (e *FFmpeg) SetUserAgent(userAgent string) *FFmpeg {
	list := make([]string, 0)
	list = append(list, "-user_agent")
	list = append(list, userAgent)
	e.UserAgent = list
	return e
}

func (e *FFmpeg) IsShCms() *FFmpeg {
	e.isShCms = true
	return e
}

func (e *FFmpeg) SetTimeout(time int) *FFmpeg {
	e.timeout = time
	return e
}

func (e *FFmpeg) SetOutput(output string) *FFmpeg {
	e.output = output
	return e
}

func (e *FFmpeg) WithOptions(opts Options) *FFmpeg {
	e.options = [][]string{opts.Parame}
	return e
}

func (t *FFmpeg) GetMetadata(information *Information) (Metadata, error) {

	var metadata Metadata

	if t.config.FfprobeBinPath == "" {
		t.config.FfprobeBinPath = "ffprobe"
	}
	var outb, errb bytes.Buffer

	input := t.input
	args := make([]string, 0)
	if len(t.UserAgent) > 0 {
		args = append(args, t.UserAgent[0], t.UserAgent[1])
	}
	if t.timeout > 0 {
		args = append(args, "-timeout", strconv.Itoa(t.timeout*1000000))
		args = append(args, "-rw_timeout", strconv.Itoa(t.timeout*1000000))

	}

	args = append(args, "-i", input, "-print_format", "json", "-show_format", "-show_streams", "-show_error")

	information.Cmd = exec.Command(t.config.FfprobeBinPath, args...)
	information.Cmd.Stdout = &outb
	information.Cmd.Stderr = &errb

	err := information.Cmd.Run()
	if err != nil {
		return metadata, fmt.Errorf("error executing (%s) with args (%s) | error: %s | message: %s %s", t.config.FfprobeBinPath, args, err, outb.String(), errb.String())
	}

	if err = json.Unmarshal([]byte(outb.String()), &metadata); err != nil {
		return metadata, err
	}

	t.metadata = &metadata
	return metadata, nil

}

func (t *FFmpeg) validate() error {

	if t.input == "" {
		return errors.New("missing input option")
	}
	//if t.output == "" {
	//	return errors.New("missing input option")
	//}

	return nil
}

func (t *FFmpeg) close() {
	defer close(t.information.Progress)
}

func (t *FFmpeg) SetMetadata() {

}

func (t *FFmpeg) Img(information *Information) {
	if information == nil {
		information = &Information{}
	}

	t.information = information
	information.Error = make(chan error)
	information.Progress = make(chan Progress)

	options := t.options
	err := t.validate()
	if err != nil {
		defer t.close()
		go func(err error, information *Information) {
			information.Error <- err
		}(err, information)
		return
	}

	// Append input file and standard options
	args := make([]string, 0)
	if len(t.UserAgent) > 0 {
		args = append(args, t.UserAgent[0], t.UserAgent[1])
	}
	if t.timeout > 0 {
		//args = append(args, "-timeout", strconv.Itoa(t.timeout*1000))
		args = append(args, "-rw_timeout", strconv.Itoa(t.timeout*1000000))
	}
	if len(t.input) > 0 {
		args = append(args, "-i", t.input)
	}
	for _, v := range options {
		args = append(args, v...)
	}
	if len(t.output) > 0 {
		args = append(args, t.output)
	}

	if t.config.FfmpegBinPath == "" {
		t.config.FfmpegBinPath = "ffmpeg"
	}

	var outb, errb bytes.Buffer
	information.Cmd = exec.Command(t.config.FfmpegBinPath, args...)
	information.Cmd.Stdout = &outb
	information.Cmd.Stderr = &errb

	defer t.close()
	err = information.Cmd.Run()
	go func(err error) {
		information.Error <- err
	}(err)
}

func (t *FFmpeg) Run(information *Information) {

	if information == nil {
		information = &Information{}
	}

	t.information = information
	information.Error = make(chan error)
	information.Progress = make(chan Progress)
	options := t.options

	var stderrIn io.ReadCloser

	err := t.validate()
	if err != nil {
		log.Error(err)
		defer t.close()
		go func(err error, information *Information) {
			information.Error <- err
		}(err, information)
		return
	}

	if t.metadata == nil {
		_, err = t.GetMetadata(information)
	}

	if err != nil {
		log.Error(err)
		defer t.close()
		go func(err error, information *Information) {
			information.Error <- err
		}(err, information)
		return
	}

	// Append input file and standard options
	args := make([]string, 0)
	if len(t.UserAgent) > 0 {
		args = append(args, t.UserAgent[0], t.UserAgent[1])
	}
	if t.timeout > 0 {
		//args = append(args, "-timeout", strconv.Itoa(t.timeout*1000))
		args = append(args, "-rw_timeout", strconv.Itoa(t.timeout*1000000))
	}

	if len(t.input) > 0 {
		args = append(args, "-i", t.input)
	}
	for _, v := range options {
		args = append(args, v...)
	}
	if len(t.output) > 0 {
		args = append(args, t.output)
	}

	if t.config.FfmpegBinPath == "" {
		t.config.FfmpegBinPath = "ffmpeg"
	}

	if t.isShCms {
		shcms := ""
		var shcmss []string
		for _, v := range options {
			shcmss = append(shcmss, v...)
		}
		for _, v := range shcmss {
			shcms += v + " "
		}

		information.Cmd = exec.Command("sh", "-c", t.config.FfmpegBinPath+" "+shcms)
	} else {
		information.Cmd = exec.Command(t.config.FfmpegBinPath, args...)
	}

	stdoutIn, err := information.Cmd.StdoutPipe()
	if err != nil {
		log.Error("args:", args, "error:", err)
		defer t.close()
		go func(err error, information *Information) {
			information.Error <- err
		}(err, information)
		return
	}

	stderrIn, err = information.Cmd.StderrPipe()
	if err != nil {
		log.Error("args:", args, "error:", err)
		defer t.close()
		go func(err error, information *Information) {
			information.Error <- err
		}(err, information)
		return
	}

	// Start process
	err = information.Cmd.Start()
	if err != nil {
		log.Error("args:", args, "error:", err)
		defer t.close()
		go func(err error, information *Information) {
			information.Error <- err
		}(err, information)
		return
	}

	if t.config.ProgressEnabled {
		go func() {
			t.progress(stderrIn, information.Progress)
		}()
	}

	go func() {
		defer t.close()
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			if t.config.Verbose {
				CopyAndCapture(os.Stderr, stderrIn)
			}
			wg.Done()
		}()
		go func() {
			if t.config.Verbose {
				CopyAndCapture(os.Stderr, stdoutIn)
			}
			wg.Done()
		}()
		wg.Wait()

		err = information.Cmd.Wait()
		go func(err error) {
			if err != nil {
				log.Error("args:", args, "error:", err)
			}
			information.Error <- err
		}(err)

	}()

}

// progress sends through given channel the transcoding status
func (t *FFmpeg) progress(stream io.ReadCloser, out chan Progress) {

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
		mProgress := new(Progress)
		line := scanner.Text()
		log.Debug(string(scanner.Bytes()))
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
			dursec, _ := strconv.ParseFloat(t.metadata.Format.Duration, 64)

			progress := (timesec * 100) / dursec
			mProgress.Progress = progress

			mProgress.CurrentBitrate = currentBitrate
			mProgress.FramesProcessed = framesProcessed
			mProgress.CurrentTime = currentTime
			mProgress.Speed = currentSpeed
			ProtectRun(func() {
				t.Progress = progress
				out <- *mProgress
			})
		}
	}
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
		default:
		}
	}()
	entry()
}

func CopyAndCapture(w io.Writer, r io.Reader) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			w.Write(d)
		}
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
	}
}
