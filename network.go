package main

import (
	"fmt"
	"net/http"
	"time"
	"context"
)

type Networker struct {
	requests []*http.Request
	REQUEST_TIMEOUT time.Duration
}


func (n *Networker) AddSimpleRequest(method string, url string) {
	ctx, _ := context.WithTimeout(context.Background(), n.REQUEST_TIMEOUT)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	n.requests = append(n.requests, req.WithContext(ctx))
}
func NewNetworker() Networker{
	return Networker{
		requests:[]*http.Request{},
		REQUEST_TIMEOUT: 2 * time.Second,
	}
}
func NewNetworkerWithTimeout(timeout time.Duration) Networker{
	return Networker{
		requests:[]*http.Request{},
		REQUEST_TIMEOUT: timeout,
	}
}

func (n *Networker) GetPool(max_threads int) WorkerPool{
	pool := WorkerPool{}
	chunkSize := (len(n.requests)) / max_threads
	fmt.Println("Chunck size : ", chunkSize)

	for i := 0; i < len(n.requests); i += chunkSize {
		end := i + chunkSize

		if end > len(n.requests) {
			end = len(n.requests)
		}
		fmt.Println(i, end)

		pool.workers = append(pool.workers, Worker{n.requests[i:end]})
	}
	return pool


}