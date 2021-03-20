package main

import (
	"fmt"
)

type Fuzzer struct {
	MAX_THREADS int
	wordlist []string
	target string

}

func (fz *Fuzzer) AddWordlist(wordlist []string) {
	fz.wordlist = append(fz.wordlist, wordlist...)
}

func (fz *Fuzzer) SetTarget(target string) {
	fz.target = target
}

func (fz *Fuzzer) Run() {


	// TODO: Generate requests by using the wordlist

	networker := NewNetworkerWithTimeout(2 * time.Second)
	networker.AddSimpleRequest("GET","https://google.com/1")
	networker.AddSimpleRequest("GET","https://google.com/2")
	networker.AddSimpleRequest("GET","https://google.com/3")
	networker.AddSimpleRequest("GET","https://google.com/4")
	networker.AddSimpleRequest("GET","https://google.com/5")
	networker.AddSimpleRequest("GET","https://google.com/6")
	pool := networker.GetPool(2) // 2 is max concurrent workers, so all requests will be splitted to , in this case, 2 workers
	//Pool is a " worker-executer", the pool will exexute workers asynchronously
	var results = pool.Run()
	//Then, pool collect all workers results and put them in results var
	for _, r := range results {
		fmt.Println("Founded ", r)
	}
}

func NewFuzzer(target string) Fuzzer {
	return Fuzzer {
		MAX_THREADS: 5,
		wordlist: []string{},
		target: target,
	}
}
