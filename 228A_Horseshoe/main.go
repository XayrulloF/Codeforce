package main
import (
	"fmt"
)

func inputFunc(sl map[int]int) {
	for i:=0; i<4; i++ {
		var num int
		fmt.Scan(&num)
		sl[num] = 1
	}
}

func findFunc(sl map[int]int, summ *int) {
	var ss int = len(sl)
	*summ = 4 - ss
}

func main() {
	var summ int = 0
	sl := make(map[int]int)
	inputFunc(sl)
	findFunc(sl, &summ)
	fmt.Println(summ)
}