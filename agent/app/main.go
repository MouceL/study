package main

import (
	"context"
	"flag"
	"fmt"
	"lll/study/log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
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

	// 初始化
}


func hook(cancel context.CancelFunc){
	sig := make(chan os.Signal)
	signal.Notify(sig,syscall.SIGKILL,syscall.SIGTERM)
	ch := <- sig
	log.Logger.Infof("get signal %v",ch)
	cancel()
}