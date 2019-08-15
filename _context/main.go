package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel() // 防止任务比超时时间短导致资源未释放
	// 启动协程
	go task(ctx)
	// 主协程需要等待，否则直接退出
	time.Sleep(time.Second * 4)
}

func task(ctx context.Context) {

	ch := make(chan struct{}, 0)

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("finish")
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout")

	}
}
