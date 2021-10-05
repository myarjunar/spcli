package pathfinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {
	t.Run("AdjacencyListFromCsv", func(t *testing.T) {
		filepath := "./_testdata/graph.csv"

		expectedAdjacencyList := AdjacencyList{
			"A": {"B": 15, "D": 50, "E": 37},
			"B": {"C": 4},
			"C": {"D": 8, "E": 22},
			"D": {"C": 10, "E": 6},
			"E": {"B": 13},
		}

		adjacencyList := AdjacencyListFromCsv(filepath)

		assert.Equal(t, expectedAdjacencyList, adjacencyList)
	})
}
