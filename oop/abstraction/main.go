package main

import (
	"context"
	"fmt"
	"time"
)

type Stream interface {
	Buf([]byte) error
	BufCtx(ctx *context.Context, body []byte) error
}

type abstractStreamer struct  {

}

func (as *abstractStreamer) BufCtx(ctx *context.Context, body []byte) error {
	panic("implement me")
	return nil
}

type Streamer struct {
	 abstractStreamer
}

func (s *Streamer) Buf(body []byte) error {
	fmt.Println("buffered : ", body)
	return nil
}


//func (s *Streamer) BufCtx(ctx *context.Context, body []byte) error {
//	fmt.Println("buffered with context : ", body)
//	return nil
//}
//

func main() {
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	stream := &Streamer{}
	data := []byte("hello")
	stream.Buf(data)
	stream.BufCtx(&ctx,data)

}