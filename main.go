package atcoderLib

import (
	"bufio"
	"os"
	"strconv"
)

func BufScan(n int) []string {
	var in []string
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		sc.Scan()
		in = append(in, sc.Text())
	}

	return in
}

func IntBufScan(n int) []int {
	var inInts []int
	inStrings := BufScan(n)
	for _, i := range inStrings {
		tmp, _ := strconv.Atoi(i)
		inInts = append(inInts, tmp)
	}
	return inInts
}
