package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "a[b+c=\\"
	fmt.Println(str)
	fmt.Println(regexp.QuoteMeta(str))

	query := "UPDATE accounts SET password = $1, email = $2, updated_at = now() WHERE id = $3"
	reQuery := regexp.QuoteMeta(query)
	fmt.Println(query)
	fmt.Println(reQuery)

	rep := regexp.MustCompile(`[\s]{2,}`)
	fmt.Println(rep)
}
