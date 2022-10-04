package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func g1() error {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		if i == 5 {
			return fmt.Errorf("g1 error")
		}
		fmt.Printf("g1: %d秒\n", i)
	}
	return nil
}

func g2() error {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("g2: %d秒\n", i)
	}
	return nil
}

func funcWithRecover(f func() error) func() error {
	return func() error {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("my recover:", err)
			}
		}()
		err := f()
		if err != nil {
			return err
		}
		return nil
	}
}

func main() {
	errgroup := &errgroup.Group{}
	funArrary := []func() error{g1, g2}
	for j := 0; j < len(funArrary); j++ {
		//	这里如果传j是传地址进去，会出现问题，因为j是循环变量，最后的值是2，所以最后的结果都是2
		start := j
		errgroup.Go(funcWithRecover(func() error {
			fmt.Printf("start address: %p\n", &start)
			err := funArrary[start]()
			if err != nil {
				return err
			}
			return nil
		}))
	}
	err := errgroup.Wait()
	if err != nil {
		fmt.Println(err)
	}
	return
}
