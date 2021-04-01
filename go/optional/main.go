// ref: 
// https://stackoverflow.com/questions/30716354/how-do-i-do-a-literal-int64-in-go/30716481#30716481
// https://stackoverflow.com/questions/30731687/how-do-i-represent-an-optional-string-in-go

package main

import (
	"fmt"
)

type myStr struct {
    name string
	url *string
}

func main() {
    f := func(s string) *string {
        return &s
    }
    test := myStr{
		name: "sano",
        url: f("12345"),
    }
	fmt.Println(test)
}