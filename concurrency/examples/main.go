package main

import (
	"fmt"

	"../../concurrency"
)

func main() {
	// pub-sub example
	fmt.Println("---pub-sub example---")
	ps := concurrency.CreatePubSub()
	ch := ps.Subscribe("go")
	ps.Publish("go", "pub/sub pattern")
	fmt.Println(<-ch)
	ps.Close()
	err := ps.Publish("go", "test")
	fmt.Println(err)

	// workers pool example
	fmt.Println("---workers pool example---")
	tasks := []*concurrency.Task{
		concurrency.NewTask(func() error {
			fmt.Println("task 1")
			return nil
		}),
		concurrency.NewTask(func() error {
			fmt.Println("task 2")
			return nil
		}),
		concurrency.NewTask(func() error {
			fmt.Println("task 3")
			return nil
		}),
	}

	p := concurrency.NewPool(tasks, 10)
	p.Run()

	var numErrors int
	for _, task := range p.Tasks {
		if task.Err != nil {
			fmt.Println(task.Err)
			numErrors++
		}
		if numErrors >= 1 {
			fmt.Println("Has errors")
			break
		}
	}

	// generator
	fmt.Println("---generator example---")
	for i := range concurrency.Count(1, 3) {
		fmt.Println(i)
	}
}
