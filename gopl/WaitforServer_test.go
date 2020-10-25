package gopl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Get(url string) error{
	resp , err := http.Get(url)
	if err!=nil{
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return fmt.Errorf("invalid status %s",resp.Status)
	}

	raw , err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return err
	}
	fmt.Println(string(raw))
	return nil
}


func WaitforServer(url string) error{
	const timeout = time.Minute
	deadline := time.Now().Add(timeout)

	for tries:=0;time.Now().Before(deadline);tries++{
		err := Get(url)
		if err!=nil{
			fmt.Printf("%s\n",err.Error())
		}else{
			return nil
		}
		time.Sleep(time.Second << uint(tries))
	}
	return nil
}

func TestWait(t *testing.T){

	err := WaitforServer("http://www.baidu.com")
	if err!=nil{
		t.Errorf(err.Error())
	}
}
