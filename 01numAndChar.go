/*
@Time        :2021/07/15 17:01:05
@Author      :Reid
@Version     :1.0
@Desc        :交替打印数字和字母
*/

package interview

import (
	"fmt"
	"sync"
)

func PrintNumAndChar() {

	var wg sync.WaitGroup
	num, char := make(chan struct{}), make(chan struct{})
	exit := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 1
		// for select 用法
		for {
			select {
			case <-num:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				char <- struct{}{}
			case <-exit:
				return
			}
		}

		// for range 用法
		// for _ = range num {
		// 	fmt.Print(i)
		// 	i++
		// 	fmt.Print(i)
		// 	i++
		// 	char <- struct{}{}
		// }
	}()

	wg.Add(1)
	go func() {
		j := 'A'
		for {
			select {
			case <-char:
				if j < 'Z' {
					fmt.Print(string(j))
					j++
					fmt.Print(string(j))
					j++
					num <- struct{}{}
				} else {
					wg.Done()
					exit <- struct{}{}
					return
				}
			}
		}
	}()
	num <- struct{}{}
	wg.Wait()
}
