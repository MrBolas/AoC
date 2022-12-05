package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var highestCalories, elfcount int
	var highestElves []int

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic("read error")
		}
		line = strings.TrimSuffix(line, "\r\n")
		fmt.Printf("Evaluated line: %s\n", line)

		if line == "" {

			// add top3
			if elfcount > highestCalories {
				highestCalories = elfcount
				fmt.Printf("new high score: %d \n", highestCalories)
			}

			highestElves = append(highestElves, elfcount)

			elfcount = 0
			continue
		}

		newCallories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		elfcount += newCallories
	}

	sort.Ints(highestElves)

	fmt.Println(highestCalories)

	fmt.Printf("%d\n", highestElves[len(highestElves)-3:])

	fmt.Printf("%d\n", highestElves[len(highestElves)-1]+
		highestElves[len(highestElves)-2]+
		highestElves[len(highestElves)-3])

}
