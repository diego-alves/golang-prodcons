package main

import "sync"

type ChannelWaitGroup struct {
	channel chan [2]string
	wg      *sync.WaitGroup
}

func (cwg ChannelWaitGroup) Add(first string, second string) {
	cwg.wg.Add(1)
	cwg.channel <- [2]string{first, second}
}

func (cwg ChannelWaitGroup) Wait() {
	cwg.wg.Wait()
	close(cwg.channel)
}

func (cwg ChannelWaitGroup) Done() {
	cwg.wg.Done()
}
