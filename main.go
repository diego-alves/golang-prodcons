package main

import (
	"fmt"
	"sync"
	"time"
)

var organizations = [...]string{
	"",
	"",
}

func consumer(queue chan []string, wg *sync.WaitGroup) {
	for v := range queue {
		if v[0] != "" {
			fmt.Print(".")
			// fmt.Printf("%s %3d - %s\n", name, i, v)
		}
		time.Sleep(time.Millisecond * 2)
		wg.Done()
	}
}

func producer(todos chan []string, queue chan []string, wg *sync.WaitGroup) {
	for v := range todos {
		fmt.Printf("\nGet More " + v[0] + "\n")
		ids := GetExternalIdentities(v[0], v[1])
		for _, v := range ids.Edges {
			wg.Add(1)
			queue <- []string{v.Node.SamlIdentity.NameId, v.Node.User.Login}
		}
		if ids.PageInfo.HasNextPage {
			wg.Add(1)
			todos <- []string{v[0], ids.PageInfo.EndCursor}
		}
		wg.Done()
	}
}

func main() {
	queue := make(chan []string, 1000)
	todos := make(chan []string, 10)
	var w sync.WaitGroup

	//setup consumers
	go consumer(queue, &w)

	//setup producers
	go producer(todos, queue, &w)

	// starts producer by org
	for _, org := range organizations {
		w.Add(1)
		todos <- []string{org, ""}
	}

	w.Wait()
	close(todos)
	close(queue)
	fmt.Println("End of Job")
}
