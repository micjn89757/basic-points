package concurrency

import (
	"fmt"
)

/**
实现一个交替打印
*/

func PrintNum(num int) {
	g1Chan := make(chan struct{}, 1)
	g2Chan := make(chan struct{}, 1)

	nums := make(chan int, 100)
	wg.Add(2)
	go g1(nums, g1Chan, g2Chan)
	go g2(nums, g1Chan, g2Chan)
	g1Chan <- struct{}{}
	for i := 1; i <= num; i++ {
		nums <- i
	}
	close(nums) // 注意要关闭通道

	wg.Wait()
}

func g1(numsChan chan int, g1Chan, g2Chan chan struct{}) {
	defer wg.Done()
	for {
		<- g1Chan
		v, ok := <- numsChan
		if ok {
			fmt.Println("g1", v)
			g2Chan <- struct{}{}
		} else {
			break
		}
	}
}

func g2(numsChan chan int, g1Chan, g2Chan chan struct{}) {
	defer wg.Done()

	for {
		<- g2Chan
		v, ok := <- numsChan
		if ok {
			fmt.Println("g2", v)
			g1Chan <- struct{}{}
		} else {
			break
		}
	}
}
