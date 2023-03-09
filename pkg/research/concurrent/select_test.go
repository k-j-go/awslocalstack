package concurrent

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test_Select_One(t *testing.T) {
	ch := make(chan bool, 0)
	go func(ch chan bool) {
		time.Sleep(time.Second * 1)
		ch <- true
	}(ch)
	select {
	case ret, ok := <-ch:
		fmt.Println(ret, ok)
	case <-time.After(time.Second * 5):
		fmt.Println("timeout")
	}
}

func Test_loop(t *testing.T) {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				//Very important
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
