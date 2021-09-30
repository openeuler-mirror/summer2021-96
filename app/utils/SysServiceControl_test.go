package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestSysServiceControl(t *testing.T) {
	SysServiceControl("redis-server","status")
}

func TestSysServiceControl2(t *testing.T) {
	serverName := "mysql"
	operation := "restart"
	cmd := exec.Command("service",serverName,operation)
	_ , err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			log.Printf("systemctl finished with non-zero: %v\n", exitErr)
			log.Println(exitErr)
		} else {
			log.Printf("failed to run systemctl: %v", err)
			os.Exit(1)
			log.Println(err)
		}
	}
}

func TestName(t *testing.T) {
	log.Println(getCPUmodel())
	serviceRestart()
}

func getCPUmodel() string {
	cmd := "cat /proc/cpuinfo | egrep '^model name' | uniq | awk '{print substr($0, index($0,$4))}'"
	out, err := exec.Command("bash","-c",cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	return string(out)
}
func serviceRestart() string {
	reCmd := "echo 'youzc' | sudo -S service mysql restart"
	cmd := "cat /proc/cpuinfo | egrep '^model name' | uniq | awk '{print substr($0, index($0,$4))}'"
	out, err := exec.Command("bash","-c",cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}

	log.Println(out)

	out1, err := exec.Command("bash","-c",reCmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", reCmd)
	}

	return string(out1)
}