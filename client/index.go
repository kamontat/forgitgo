package client

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"

	util "github.com/kamontat/my_settings/settings/utils"
)

func ExecShellScript(filepath string) error {
	return rawCommandWithDefaultSTD("bash", filepath)
}

func rawCommandWithDefaultSTD(name string, arg ...string) (err error) {
	out, stderr, err := rawCommandWithReturn(name, arg...)
	if err == nil {
		util.GetLogger().Debug("execute", fmt.Sprintf("%s %s", name, arg))
		if len(out) > 0 {
			util.GetLogger().Debug("out", out)
		}
	} else {
		util.GetLogger().WithError(err).WithField(logrus.Fields{
			"command": name,
			"args":    arg,
		}).Error("Execute", stderr)
	}
	return
}

func rawCommandWithReturn(name string, arg ...string) (strout string, strerr string, err error) {
	var stdout, stderr bytes.Buffer
	// fmt.Println(name, strings.Join(arg, " "))
	cmd := exec.Command(name, arg...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	strout = stdout.String()
	strerr = stderr.String()
	return
}

func rawCommandWithCustomSTD(name string, in *os.File, out *os.File, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = in
	cmd.Stdout = out

	err = cmd.Run()
	return
}
