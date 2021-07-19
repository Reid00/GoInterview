/*
@Time        :2021/07/19 16:49:07
@Author      :Reid
@Version     :1.0
@Desc        :多协程查询切片问题
*/

package interview

import (
	"context"
	"fmt"
	"time"
)

func SearchAction() {
	timer := time.NewTimer(5 * time.Second)
	data := []int{1, 2, 3, 10, 999, 8, 345, 7, 98, 33, 66, 77, 88, 68, 96}
	dataLen := len(data)
	size := 3
	target := 345

	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan bool)

	for i := 0; i < dataLen; i += size {
		end := i + size

		if end >= dataLen {
			end = dataLen - 1
		}
		go SearchTarget(ctx, data[i:end], target, resultChan)
	}

	select {
	case <-timer.C:
		fmt.Println("Timeout! Not Found! !")
		cancel()
	case <-resultChan:
		fmt.Println("Found it!")
		cancel()
	}

	time.Sleep(2 * time.Second)
}

func SearchTarget(ctx context.Context, data []int, target int, resultChan chan<- bool) {
	for _, v := range data {
		select {
		case <-ctx.Done():
			fmt.Println("Task Cancelded")
			return
		default:
		}

		// 模拟一个耗时查找，这里只是比对值，真实开发中可以是其他操作
		fmt.Println("looked num is: ", v)
		time.Sleep(1500 * time.Millisecond)
		if target == v {
			resultChan <- true
			return
		}
	}
}
