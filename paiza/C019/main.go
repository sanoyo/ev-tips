package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 空白や空文字だけの値を除去したSplit()
func SplitWithoutEmpty(stringTargeted string, delim string) (stringReturned []string) {
	stringSplited := strings.Split(stringTargeted, delim)

	for _, str := range stringSplited {
		if str != "" {
			stringReturned = append(stringReturned, str)
		}
	}

	return
}

// 空白や空文字だけの値を除去したSplitN()
func SplitWithoutEmptyN(stringTargeted string, delim string, n int) (stringReturned []string) {
	stringSplited := strings.SplitN(stringTargeted, delim, n)

	for _, str := range stringSplited {
		if str != "" {
			stringReturned = append(stringReturned, str)
		}
	}

	return
}

// デリミタで分割して整数値スライスを取得
func SplitAndConvertToInt(stringTargeted, delim string) (intSlices []int, err error) {
	// 分割
	stringSplited := SplitWithoutEmpty(stringTargeted, delim)

	// 整数スライスに保存
	for i := range stringSplited {
		var iparam int
		iparam, err = strconv.Atoi(stringSplited[i])
		if err != nil {
			return
		}
		intSlices = append(intSlices, iparam)
	}
	return
}

/**
 * paiza専用: 行指定なしデータ列読み込み
 *
 */
func PaizaSequenceGetsAsInt() (line [][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		data, err := SplitAndConvertToInt(str, " ")
		if err != nil || data == nil {
			break
		}
		line = append(line, data)
	}

	return
}

/**
 * paiza専用: 行指定なし文字列読み込み
 *
 */
func PaizaSequenceGets() (lines [][]string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		data := SplitWithoutEmpty(str, " ")
		if data == nil {
			break
		}
		lines = append(lines, data)
	}

	return
}

/**
 * paiza専用: 「文字 数値 数値 数値...」の取得
 *
 */
func PaizaStrFirstIntAfter(data string, delim string) (head string, body []int) {
	list := SplitWithoutEmptyN(data, delim, 2)
	head = list[0]
	body, _ = SplitAndConvertToInt(list[1], delim)
	return
}

/**
 * paiza専用: 文字列一括取得・最初X,Y、以下Y行文字列
 *
 */
func PaizaGetsXY() (x, y int, line []string) {
	scanner := bufio.NewScanner(os.Stdin)

	first := true
	for scanner.Scan() {
		str := scanner.Text()
		if first {
			xy, _ := SplitAndConvertToInt(str, " ")
			x = xy[0]
			y = xy[1]
			first = false
		} else {
			line = append(line, strings.TrimSpace(str))
		}
	}

	return
}

/**
 * paiza専用: 文字列一括取得・最初行数、以下文字列
 *
 */
func PaizaGets() (count int, line []string) {
	scanner := bufio.NewScanner(os.Stdin)

	n := -1
	for scanner.Scan() {
		str := scanner.Text()
		if n == -1 {
			count, _ = strconv.Atoi(strings.TrimSpace(str))
		} else {
			line = append(line, strings.TrimSpace(str))
		}
		n += 1

		if n >= count {
			break
		}
	}

	return
}

/**
 * paiza専用: 数字列1行取得・1整数以上対応
 *
 */
func PaizaGetNums() (intReturned []int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()
	intReturned, _ = SplitAndConvertToInt(str, " ")

	return
}

// 標準入力受け取り定義
var sc = bufio.NewScanner(os.Stdin)

func inputText() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := inputText()
	input, _ := strconv.Atoi(s)

	intSlice := 0
	for i := 1; i < input; i++ {
		if input%i == 0 {
			intSlice += i
		}
	}

	if intSlice == input {
		fmt.Println("perfect")
	} else if (input - intSlice) == 1 {
		fmt.Println("nearly")
	} else {
		fmt.Println("neither")
	}
}
