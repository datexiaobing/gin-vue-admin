package tools

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

// 执行Shell 命令
func ExecShell(command string) (out string, errMsg string, err error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		command = "/C  " + command
		args := strings.Split(command, " ")
		fmt.Println("args:", args)
		cmd = exec.Command("cmd", args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("err", err)
		}
		s := ConvertByte2String(output, GB18030)
		fmt.Println(string(output))
		fmt.Println(s)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}
	// output, err := cmd.CombinedOutput()
	// fmt.Println(string(output), err)
	// 输出转码命令
	fmt.Println("执行命令：", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	out = stdout.String()

	errMsg = stderr.String()
	err = cmd.Run()
	return out, errMsg, err
}
