package main 

import (
	"fmt"
	"time"
)

func main() {
	fuzzer := NewFuzzer("https://google.com/$MTR")
	fuzzer.AddWordlist([]string{"admin","wp-admin"}) //Useless for now, need to code it !
	fuzzer.Run()
}



