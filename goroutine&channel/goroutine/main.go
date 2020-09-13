package main

import (
	"fmt"
	"time"
)

type twoWay struct {
	in int
	out int
}

func hello(pipe chan *twoWay) {
	for {
		time.Sleep(time.Millisecond * 200)
		select {
		case v := <- pipe:
			if v.in != 1 {
				pipe <- &twoWay{
					in:  0,
					out: 2,
				}
				break
			}
			fmt.Println(v.in)
			pipe <- &twoWay{
				in:  0,
				out: 2,
			}
		}
	}
}

func main() {
	pipe := make(chan *twoWay,1)
	go hello(pipe)
	pipe <- &twoWay{
		in:  0,
		out: 0,
	}
	for {
		time.Sleep(time.Millisecond * 200)
		select {
		case v := <- pipe :
			if v.out != 2 {
				pipe <- &twoWay{
					in:  1,
					out: 0,
				}
				break
			}
			fmt.Println(v.out)
			pipe <- &twoWay{
				in:  1,
				out: 0,
			}
		}
	}

}
