package gopl

import (
	"errors"
	"fmt"
	"lll/study/log"
	"testing"
)


func test() (err error) {

	defer func() {
		p := recover()
		if p != nil{
			err = errors.New(fmt.Sprintf("%s",p))
		}
	}()
	n:=3
	for i:=3;i>=-1;i--{
		fmt.Printf("%d\n",n/i)
	}
	return
}


func TestPanic(t *testing.T){
	log.Logger.Info("it is a test")
	err := test()
	if err!=nil{
		fmt.Println(err.Error())
		log.Logger.Debug(err.Error())
	}
}
