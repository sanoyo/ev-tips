// https://christina04.hatenablog.com/entry/go-singleflight
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

func download(path string) ([]byte, error) {
	log.Println("download")
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func suppressDupCall(path string) ([]byte, error) {
	v, err, _ := group.Do(path, func() (interface{}, error) {
		return download(path)
	})
	if err != nil {
		return nil, err
	}
	return v.([]byte), nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	path := "https://www.gstatic.com/webp/gallery3/1.png"
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// img, err := download(path)
			img, err := suppressDupCall(path)
			fmt.Printf("result: %d, err: %v\n", len(img), err)
		}()
	}
	wg.Wait()
}
