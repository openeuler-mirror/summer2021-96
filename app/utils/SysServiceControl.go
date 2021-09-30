package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// SysServiceControlN 服务控制工具，用于控制系统服务
// @Param serverName string 要操作的服务名称
// @Param operation string 要操作的类型(start,stop...)
// @return bool 命令是否执行成功
// @return error 异常
func SysServiceControl(serverName string,operation string) (bool,string){
	//cmd := exec.Command("service",serverName,operation
	//_ , err := cmd.CombinedOutput())
	cmd := "echo 'youzc' | sudo -S service " + serverName + " " + operation // 使用管道输入密码
	out, err := exec.Command("bash","-c",cmd).Output()


	if err != nil {
		return false,fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	return true,string(out);

	//if err != nil {
	//	if exitErr, ok := err.(*exec.ExitError); ok {
	//		log.Printf("systemctl finished with non-zero: %v\n", exitErr)
	//		return false,exitErr
	//	} else {
	//		log.Printf("failed to run systemctl: %v", err)
	//		os.Exit(1)
	//		return false,err
	//	}
	//}
	//return true,nil
}

func SysServiceControlOld(serverName string,operation string) (bool,error){
	cmd := exec.Command("service",serverName,operation)
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