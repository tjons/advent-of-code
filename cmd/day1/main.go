package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	nums := []int{}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num := scanner.Text()

		convNum, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, convNum)
	}

	if scanErr := scanner.Err(); scanErr != nil {
		log.Fatal(err)
	}

	var a, b, c int

	for _, num := range nums {
		for _, bNum := range nums {
			for _, cNum := range nums {
				if num != bNum && bNum != cNum && num + bNum + cNum == 2020 {
					a = num
					b = bNum
					c = cNum
				}
			}
		}
	}

	log.Println(a * b * c)
	return
}