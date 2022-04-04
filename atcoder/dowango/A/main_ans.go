package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

type Music struct {
	Name string
	Time int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	playlist := make([]*Music, n) // [<nil> <nil> <nil>]
	for i := 0; i < n; i++ {
		var name string
		var time int
		fmt.Scanf("%s %d", &name, &time)
		playlist[i] = &Music{
			Name: name,
			Time: time,
		}
	}

	var x string
	fmt.Scanf("%s", &x)

	total := 0
	start := false
	for _, music := range playlist {
		if start {
			total += music.Time
		}
		if music.Name == x {
			start = true
		}
	}

	fmt.Printf("%d", total)
}
