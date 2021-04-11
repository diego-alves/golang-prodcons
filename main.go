package main

import (
	"fmt"
	"sync"
	"time"
)

func say(queue chan string, w *sync.WaitGroup, name string) {
	i := 0
	for v := range queue {
		i = i + 1
		fmt.Printf("%s %2d - %s\n", name, i, v)
		time.Sleep(50000)
		w.Done()
	}
}

func main() {
	queue := make(chan string, 100)
	var w sync.WaitGroup

	go say(queue, &w, "A")
	go say(queue, &w, "B")
	go say(queue, &w, "C")

	for _, v := range GetExternalIdentities("").Edges {
		w.Add(1)
		queue <- v.Node.User.Login
	}

	w.Wait()
	close(queue)
	fmt.Println("End of Job")
}
