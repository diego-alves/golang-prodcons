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

func consumer(queue *ChannelWaitGroup) {
	for v := range queue.channel {
		if v[0] != "" {
			fmt.Print(".")
			// fmt.Printf("%s %3d - %s\n", name, i, v)
		}
		time.Sleep(time.Millisecond * 50)
		queue.Done()
	}
}

func producer(todos *ChannelWaitGroup, queue *ChannelWaitGroup) {
	for v := range todos.channel {
		fmt.Printf("\nGet More " + v[0] + "\n")
		ids := GetExternalIdentities(v[0], v[1])
		for _, v := range ids.Edges {
			queue.Add(v.Node.SamlIdentity.NameId, v.Node.User.Login)
		}
		if ids.PageInfo.HasNextPage {
			todos.Add(v[0], ids.PageInfo.EndCursor)
		}
		todos.Done()
	}
}

func main() {
	orgs := ChannelWaitGroup{
		channel: make(chan [2]string, 10),
		wg:      new(sync.WaitGroup),
	}
	queue := ChannelWaitGroup{
		channel: make(chan [2]string, 1000),
		wg:      new(sync.WaitGroup),
	}

	//setup consumers
	go consumer(&queue)
	go consumer(&queue)
	go consumer(&queue)

	//setup producers
	go producer(&orgs, &queue)

	// starts producer
	for _, org := range organizations {
		orgs.Add(org, "")
	}

	orgs.Wait()
	fmt.Println("\nProducers is Done")
	queue.Wait()
	fmt.Println("\nEnd of Job")
}
