package main

import (
	"bufio"
	"fmt"
	"os"
	"flag"
	"strconv"
)

func main() {

	var filePath string
	var nLinesTotal int64

	flag.StringVar(&filePath, "file_path", "data/stack-distance-list.csv", "Input path")
	flag.Int64Var(&nLinesTotal, "n_lines", 0, "Number of lines in the input file")
	flag.Parse()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// The token number used to represent INF on the input file
	INF := int64(9223372036854775807)

	nLinesBefore := int64(0)
	previousStackDist := int64(-1)
	previousHR := float64(-1)

	header := "stackDist,hitRatio"
	fmt.Printf("%s\n", header)

	for scanner.Scan() {
		line := scanner.Text()
		stackDist, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Printf("Error converting line to int64: %v\n", err)
			continue
		}

		if stackDist == INF {
			fmt.Printf("-1,0.00\n")
			nLinesBefore += 1
			continue
		}

		// The capacity need for items with a Stack Distance X is X+1
		stackDist += 1

		// Assuming a the list of Stack Distances is descending from INF to 1
		hitRatio := float64(nLinesTotal-nLinesBefore) / float64(nLinesTotal)

		if previousStackDist == stackDist {
			// The same Stack Distance should have the same Hit Ratio
			fmt.Printf("%v,%.2f\n", stackDist, previousHR)
		} else {
			fmt.Printf("%v,%.2f\n", stackDist, hitRatio)
			previousStackDist = stackDist
			previousHR = hitRatio
		}

		nLinesBefore += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
}
