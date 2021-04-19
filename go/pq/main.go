// ref: http://y-ogura.hatenablog.com/entry/2018/04/18/105758
package main

import (
	"fmt"
	"time"

	"github.com/lib/pq"
)

// Sample sample struct
type Sample struct {
	ID       int
	JoinDate pq.NullTime
}

func main() {
	// sample := &Sample{}
	value := map[string]interface{}{"id": 1, "join_date": pq.NullTime{Time: time.Now(), Valid: false}}
	fmt.Println(value)
	// r.Update(sample, value)
	// UPDATE "samples" SET "id" = '1', "join_date" = NULL WHERE "samples","id" = '1'
}
