package utils

import (
	"log"
	"os"
	"os/exec"
)

// SysServiceControl 服务控制工具，用于控制系统服务
// @Param serverName string 要操作的服务名称
// @Param operation string 要操作的类型(start,stop...)
// @return bool 命令是否执行成功
// @return error 异常
func SysServiceControl(serverName string,operation string) (bool,error){
	cmd := exec.Command("systemctl", operation, serverName)
	_ , err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			log.Printf("systemctl finished with non-zero: %v\n", exitErr)
			return false,exitErr
		} else {
			log.Printf("failed to run systemctl: %v", err)
			os.Exit(1)
			return false,err
		}
	}
	return true,nil
}

