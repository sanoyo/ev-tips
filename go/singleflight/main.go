package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

func doSomething(name string) {
	v, err, shared := group.Do(name, func() (interface{}, error) {
		time.Sleep(5 * time.Millisecond)
		return time.Now(), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result: %s, shared: %t\n", v, shared)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doSomething("hogehoge")
		}()
		time.Sleep(time.Millisecond)
	}
	wg.Wait()
}
