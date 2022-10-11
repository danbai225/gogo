package csgo

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Inject(dllName string) string {
	cmd := runCmd(fmt.Sprintf("danbai.exe %s", dllName))
	if cmd != "" {
		fmt.Println(cmd)
	}
	return ""
}
func Inject2(pid int32, dllName string) string {
	cmd := runCmd(fmt.Sprintf("danbai2.exe %d %s", pid, dllName))
	if cmd != "" {
		fmt.Println(cmd)
	}
	return ""
}
func runCmd(cmdStr string) string {
	list := strings.Split(cmdStr, " ")
	list = append([]string{"/C"}, list...)
	cmd := exec.Command("cmd", list...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String()
	} else {
		return out.String()
	}
}
