// Ref:
// https://github.com/gorilla/csrf
// http://matope.hatenablog.com/entry/2019/06/05/144435
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

func main() {
	// Hash keys should be at least 32 bytes long
	var hashKey = []byte("very-secret")
	fmt.Println(hashKey)
	// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
	// Shorter keys may weaken the encryption used.
	var blockKey = []byte("a-lot-secret")
	fmt.Println(blockKey)
	var s = securecookie.New(hashKey, blockKey)
	fmt.Println(s)

	value := map[string]string{
		"foo": "bar",
	}
	if encoded, err := s.Encode("cookie-name", value); err == nil {
		cookie := &http.Cookie{
			Name:     "cookie-name",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}
		fmt.Println("cookie: ", cookie)
	}
}
