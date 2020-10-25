package agent

import (
	"lll/study/log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const ERROR_WAIT_STATUS  syscall.WaitStatus = 0xFFFF

// 执行 cmd , 并检查脚本是否执行完毕

// 在15s 内不停检查，脚本是否成功执行，如果成功执行则返回，如果没有检测到，休息会再执行
func waitTimeout(cmd *exec.Cmd,timeout time.Duration) (status syscall.WaitStatus) {

	deadline := time.Now().Add(timeout)
	for retry:=0;retry<10;retry++ {

		// 查看制定子进程的状态
		wpid , err := syscall.Wait4(cmd.Process.Pid,&status,syscall.WNOHANG|syscall.WUNTRACED,nil)
		if err!=nil || wpid <0{
			log.Logger.Errorf("failed to wait %s,%v,%s",cmd.Path,cmd.Args,err)
			status = ERROR_WAIT_STATUS
			break
		}
		if wpid > 0 {
			log.Logger.Infof("script execute success")
			break
		}

		if time.Now().Before(deadline) {
			time.Sleep(500 *time.Millisecond)
			continue
		}

		err = cmd.Process.Kill()
		if err!=nil{
			log.Logger.Errorf("failed to kill %s %v: %s", cmd.Path, cmd.Args, err)
			// failed to kill process
		}
		time.Sleep(500 *time.Millisecond)
	}
	return
}


func execute(cmd *exec.Cmd) syscall.WaitStatus {
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Start()
	if err!=nil{
		//log
		return 11111
	}
	return waitTimeout(cmd,15*time.Second)
}

func ExecuteCommand(args string) bool{
	cmdArray := strings.Split(args," ")
	cmd := exec.Command(cmdArray[0],cmdArray[1:]...)
	status :=execute(cmd)
	if status!=0{
		return false
	}
	return true
}
