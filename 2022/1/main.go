package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var mostCals int64 = 0
	mostElf := 1

	var currCals int64 = 0
	currElf := 1

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			//fmt.Println("---")
			if currCals > mostCals {
				mostCals = currCals
				mostElf = currElf
			}
			currElf++
			currCals = 0
			continue
		}
		currCals += num
		//fmt.Println(num)
	}
	fmt.Println(mostCals, " : ", mostElf)
}
