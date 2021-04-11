package main

import (
	"fmt"
	"sync"
	"time"
)

func consumer(queue chan []string, w *sync.WaitGroup, name string) {
	i := 0
	for v := range queue {
		i = i + 1
		if v[0] != "" {
			fmt.Printf("%s %2d - %s\n", name, i, v)
		}
		time.Sleep(50000)
		w.Done()
	}
}

func main() {
	queue := make(chan []string, 100)
	var w sync.WaitGroup

	go consumer(queue, &w, "A")
	go consumer(queue, &w, "B")
	go consumer(queue, &w, "C")

	var lastCursor string
	for {
		ids := GetExternalIdentities("", lastCursor)
		for _, v := range ids.Edges {
			w.Add(1)
			queue <- []string{v.Node.SamlIdentity.NameId, v.Node.User.Login}
		}
		lastCursor = ids.PageInfo.EndCursor
		if !ids.PageInfo.HasNextPage {
			break
		}
	}

	w.Wait()
	close(queue)
	fmt.Println("End of Job")
}
