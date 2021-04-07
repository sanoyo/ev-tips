// ref:
// https://qiita.com/nirasan/items/2160be0a1d1c7ccb5e65
// https://qiita.com/immrshc/items/1d1c64d05f7e72e31a98

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http.Handlerを使うことができる
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 使い方は下記のような感じ
// HandlerFunc(func(ResponseWriter, *http.Request))
