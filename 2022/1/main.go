package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(arr *[]int64) (sum int64) {
	for _, val := range *arr {
		sum += val
	}
	return
}

func insert(arr *[]int64, num int64) {
	for i, val := range *arr {
		if num > val {
			temp := (*arr)[i]
			(*arr)[i] = num
			num = temp
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var mostCals []int64 = []int64{0, 0, 0}

	var currCals int64 = 0

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			//fmt.Println("---")
			insert(&mostCals, currCals)
			currCals = 0
			continue
		}
		currCals += num
		//fmt.Println(num)
	}
	fmt.Println(sum(&mostCals))
}
