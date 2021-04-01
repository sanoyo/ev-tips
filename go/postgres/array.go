// ref: https://github.com/lib/pq/blob/master/array.go

// pq.Array([]int{})
// golangの構造体に変換する
db.Query(`SELECT * FROM t WHERE id = ANY($1)`, pq.Array([]int{235, 401}))
