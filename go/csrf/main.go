// Ref:
// https://github.com/gorilla/csrf
// http://matope.hatenablog.com/entry/2019/06/05/144435
package main

import (
	"net/http"

	"github.com/gorilla/csrf"
)

func main() {
	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"))
	http.ListenAndServe(":8000", CSRF(r))
}
