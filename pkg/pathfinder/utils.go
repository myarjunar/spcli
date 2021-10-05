package pathfinder

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type AdjacencyList map[string]map[string]int

func AdjacencyListFromCsv(csvPath string) AdjacencyList {
	adjacencyList := AdjacencyList{}

	csvFile, err := os.Open(csvPath)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for idx, line := range csvLines {
		if idx > 0 {
			start := line[0]
			end := line[1]
			cost := line[2]
			costInt, _ := strconv.Atoi(cost)

			if existingMap, ok := adjacencyList[start]; ok {
				existingMap[end] = costInt
			} else {
				adjacencyList[start] = map[string]int{
					end: costInt,
				}
			}
		}
	}

	return adjacencyList
}
