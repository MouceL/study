package main

import (
	"context"
	"flag"
	"fmt"
	"lll/study/agent"
	"lll/study/log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)
const (
	EXIT_NORMAL     = 0
	EXIT_INIT_ERROR = 1
)
var v bool

func init(){
	flag.BoolVar(&v,"v",false,"version")
	flag.Parse()
}

func main(){

	if v {
		log.Logger.Infof("agent version is v1.0")
		fmt.Fprintf(os.Stdout,"agent version is v1.0")
		os.Exit(0)
	}
	// 最多用1c
	runtime.GOMAXPROCS(1)
	_ ,cancel:=context.WithCancel(context.Background())

	go hook(cancel)

	status := loop()
	log.Logger.Info("exist soon")
	os.Exit(status)
}

// 初始化 + 巡检
func loop()int{

	manager := agent.NewManage()
	if err := manager.Start();err!=nil{
		log.Logger.Errorf("start manager failed, %s",err.Error())
		return EXIT_INIT_ERROR
	}

	for{
		if isDelete() {
			log.Logger.Error("agent is deleted")
			break
		}
		if isHealthy(){
			log.Logger.Info("agent is healthy")
			time.Sleep(5*time.Second)
		}else {
			log.Logger.Error("agent is not healthy")
			time.Sleep(time.Second)
			break
		}
	}

	return EXIT_NORMAL
}



func isDelete() bool{
	return false
}

func isHealthy() bool{
	return true
}


func hook(cancel context.CancelFunc){
	sig := make(chan os.Signal)
	signal.Notify(sig,syscall.SIGKILL,syscall.SIGTERM)
	log.Logger.Infof("get signal %s , exit",<-sig)
	cancel()
}