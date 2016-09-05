package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 1)
	fmt.Println("done")

	done <- true
}

type job struct {
	Name string
	Result string
}

func worker2(j chan job) {
	for {
		myJob := <- j
		if myJob.Name == "close" {
			time.Sleep(time.Second * 1)
			myJob.Result = "종료"
		} else {
			fmt.Println("작업수행 : ", myJob.Name)
			myJob.Result = "완료"
		}
		j <- myJob
	}
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done

	done2 := make(chan job, 1)
	go worker2(done2)

	for _, i := range []string{"job 1", "job 2"} {
		done2 <- job{Name:i}
		result := <- done2
		fmt.Println("Job: ", result.Name, result.Result)
	}
	done2 <- job{Name:"close"}
	result := <- done2
	fmt.Println("결과: ", result)
}
