package gopl

import (
	"fmt"
	"net/http"
	"testing"
)

// 要实现某个接口，就要实现其所有的方法，这个在实际生产过程中非常的麻烦
// 详见 http  HandlerFunc



type db map[string]int
func (d db)hello(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"hello %s","world")
}

func TestHandle(t *testing.T){
	d := db{"shoes":20,"sock":10}
	http.HandleFunc("/hello",d.hello)
	http.ListenAndServe("127.0.0.1:8080",nil)
}
