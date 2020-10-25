package gopl

import (
	"fmt"
	"testing"
	"time"
)



func trance()func(){
	start := time.Now()
	fmt.Println(start)
	return func(){
		fmt.Println(time.Since(start))
	}
}

func slowoperation(){
	defer trance()()

	time.Sleep(2*time.Second)

}


func TestDefer(t *testing.T){
	slowoperation()
}
