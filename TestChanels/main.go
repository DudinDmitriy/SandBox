package main

import (
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	var wg1 sync.WaitGroup
	//ss := []int{1, 2, 3, 4, 5, 6, 7, 13, 11, 8, 9}
	ss := []int{11, 8, 9}
	in1 := make(chan interface{},2)
	out1 := make(chan interface{},2)
	out2 := make(chan interface{},2)

	go func() {
		for i := range ss {
			time.Sleep(time.Millisecond * 10)
			in1 <- ss[i]
		}
		close(in1)
	}()

	wg.Add(1)
	go readchanel(in1, out1, &wg)
	wg.Add(1)
	go readchanel2(out1, out2, &wg)

	//wg.Add(1)
	go func() {
		//defer wg.Done()
		for i2 := range out2 {
			println("Read chanel2: ", i2.(int))
		}
	}()

	//wg.Wait()

	chekgor := func(ch1 chan interface{}, wg *sync.WaitGroup) {
		defer wg.Done()
		end := false
		readch := false
		for !end {
			select {
			case val,more:=<-ch1:
				if more{
				readch = true
				ch1<-val
				time.Sleep(time.Millisecond * 10)
				}else{
					end = true
				}
			default:
				if readch {
					end = true
				}
			}
		}
	}

	wg1.Add(3)
	go chekgor(in1, &wg1)
	go chekgor(out1, &wg1)
	go chekgor(out2, &wg1)
	wg1.Wait()

}

func readchanel(in, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in {
		time.Sleep(time.Millisecond * 1000)
		out <- val
		println("Get chanel1: ", val.(int))
	}
	close(out)
	println("Close chanel1")
	time.Sleep(time.Millisecond * 1000)
}

func readchanel2(in, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in {
		time.Sleep(time.Millisecond * 10)
		out <- val
		println("Get chanel2 : ", val.(int))
	}
	close(out)
	println("Close chanel2")
}
