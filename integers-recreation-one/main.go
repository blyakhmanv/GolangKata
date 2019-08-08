package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(ListSquared(42, 42))

}

//ListSquared - kata function
func ListSquared(m, n int) [][]int {
	result := make([][]int, 0, 0)
	var tempint int
	var tempstr string

	for i := m; i <= n; i++ {
		tempint = 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				tempint += j * j
			}
		}
		tempstr = strconv.FormatInt(int64(tempint), 16)
		if strings.HasSuffix(tempstr, "0") || strings.HasSuffix(tempstr, "1") || strings.HasSuffix(tempstr, "4") || strings.HasSuffix(tempstr, "9") {
			if tempint == int(math.Sqrt(float64(tempint)))*int(math.Sqrt(float64(tempint))) {
				result = append(result, []int{i, tempint})
			}
		}
	}

	return result
}
