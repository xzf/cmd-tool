package cmdtool

import (
	"bytes"
	"os/exec"
)

type CMDInterface interface {
	Get() []CMD
}

type CMD struct {
	Name     string
	Logic    func(arg ...string)
	Describe string
}

//func AddCmd() {
//
//}

func AddCmd() {

}

func RunCmd(binName string, args ...string) (string, error) {
	outPutBuff := bytes.NewBuffer(nil)
	cmd := exec.Command(binName, args...)
	cmd.Stdout = outPutBuff
	err := cmd.Run()
	return outPutBuff.String(), err
}
