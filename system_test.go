/**
 * @file system_test.go
 * @author Mikhail Klementyev jollheef<AT>riseup.net
 * @license GNU GPLv3
 * @date July, 2016
 */

package system

import (
	"fmt"
	"testing"
)

func systemPanic(stdout, stderr string, ret int, err error) {
	panic(fmt.Sprintf("stdout: '%v', stderr: '%v', ret: '%v', err: '%v'",
		stdout, stderr, ret, err))
}

func TestSystem1(*testing.T) {
	stdout, stderr, ret, err := System("/bin/true")
	if ret != 0 || err != nil || stdout != "" || stderr != "" {
		systemPanic(stdout, stderr, ret, err)
	}
}

func TestSystem2(*testing.T) {
	stdout, stderr, ret, err := System("/bin/echo", "hello")
	if ret != 0 || err != nil || stdout != "hello\n" || stderr != "" {
		systemPanic(stdout, stderr, ret, err)
	}
}

func TestSystem3(*testing.T) {
	stdout, stderr, ret, err := System("/bin/false")
	if ret != 1 || err == nil || stdout != "" || stderr != "" {
		systemPanic(stdout, stderr, ret, err)
	}
}

func TestSystem4(*testing.T) {
	stdout, stderr, ret, err := System("/invalid/path/to/bin")
	if ret != 0 || err == nil || stdout != "" || stderr != "" {
		systemPanic(stdout, stderr, ret, err)
	}
}

func TestSystem5(*testing.T) {
	stdout, stderr, ret, err := System("/invalid/path/to/")
	if ret != 0 || err == nil || stdout != "" || stderr != "" {
		systemPanic(stdout, stderr, ret, err)
	}
}

func TestSystem6(*testing.T) {
	stdout, stderr, ret, err := System("grep")
	if ret != 2 || err == nil || stdout != "" || stderr == "" {
		systemPanic(stdout, stderr, ret, err)
	}
}
