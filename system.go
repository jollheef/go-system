/**
 * @file system.go
 * @author Mikhail Klementyev jollheef<AT>riseup.net
 * @license GNU GPLv3
 * @date July, 2016
 * @brief Provide ~execl/~system -like helper
 */

package system

import (
	"fmt"
	"io"
	"os/exec"
)

func readBytesUntilEOF(pipe io.ReadCloser) (buf []byte, err error) {

	bufSize := 1024

	for err != io.EOF {
		stdout := make([]byte, bufSize)
		var n int

		n, err = pipe.Read(stdout)
		if err != nil && err != io.EOF {
			return
		}

		buf = append(buf, stdout[:n]...)
	}

	if err == io.EOF {
		err = nil
	}

	return
}

func readUntilEOF(pipe io.ReadCloser) (str string, err error) {
	buf, err := readBytesUntilEOF(pipe)
	str = string(buf)
	return
}

func System(name string, arg ...string) (stdout string, stderr string,
	ret int, err error) {

	cmd := exec.Command(name, arg...)

	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return
	}

	cmd.Start()

	stdout, err = readUntilEOF(outPipe)
	if err != nil {
		return
	}

	stderr, err = readUntilEOF(errPipe)
	if err != nil {
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Sscanf(err.Error(), "exit status %d", &ret)
		return
	}

	return
}
