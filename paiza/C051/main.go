package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func maxInt(a []int, arg int) int {
	sort.Sort(sort.IntSlice(a))
	return a[len(a)-arg]
}

// 1行・数値・空白区切り対応
func TestPaizaGetNum() {
	nums := PaizaGetNums()

	var first_int int
	// first, _ := strconv.Atoi(strconv.Itoa(maxInt(nums, 1)) + strconv.Itoa(maxInt(nums, 3)))
	first := strconv.Itoa(maxInt(nums, 1)) + strconv.Itoa(maxInt(nums, 3))
	first_int, _ = strconv.Atoi(first)
	// second, _ := strconv.Atoi(strconv.Itoa(maxInt(nums, 2)) + strconv.Itoa(maxInt(nums, 4)))

	// total := first + second

	fmt.Printf(first_int)
	// total2, _ := strconv.Atoi(total)

	// fmt.Printf(total2)

}

func main() {
	TestPaizaGetNum()
	// 2 9 3 8
}
