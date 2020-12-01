package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var (
		ret []int
		str string
		i   int
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str = scanner.Text()
		i, err = strconv.Atoi(str)
		if err != nil {
			return ret, err
		}
		ret = append(ret, i)
	}

	return ret, scanner.Err()
}

func main() {
	data, err := readLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var (
		first, second int
	)

outer:
	for idx := range data {
		first = data[idx]

		for j := idx + 1; j < len(data); j++ {
			second = data[j]
			if first+second == 2020 {
				break outer
			}
		}
	}

	log.Printf("first %d; second %d; mul %d", first, second, first*second)
}
