package pathfinder

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathfinder(t *testing.T) {
	spTests := map[string]struct {
		sequence     string
		expectedDist int
		expectedErr  error
	}{
		"FindShortestPath Success":            {sequence: "BE", expectedDist: 18, expectedErr: nil},
		"FindShortestPath Fail Not Exist":     {sequence: "AZ", expectedDist: 0, expectedErr: errors.New("end node is not in the graph")},
		"FindShortestPath Fail Not Connected": {sequence: "CA", expectedDist: 9999, expectedErr: errors.New("nodes are not connected")},
	}

	plTests := map[string]struct {
		sequence     string
		expectedDist int
		expectedErr  error
	}{
		"FindPathLength Success": {sequence: "AEBCD", expectedDist: 62, expectedErr: nil},
		"FindPathLength Fail":    {sequence: "AZ", expectedDist: 0, expectedErr: errors.New("path not found")},
	}

	for name, test := range spTests {
		t.Run(name, func(t *testing.T) {
			filepath := "./_testdata/graph.csv"
			pf := NewPathfinder(filepath, test.sequence)
			shortestDistance, err := pf.FindShortestPath()
			assert.Equal(t, test.expectedDist, shortestDistance)
			assert.Equal(t, test.expectedErr, err)
		})
	}

	for name, test := range plTests {
		t.Run(name, func(t *testing.T) {
			filepath := "./_testdata/graph.csv"
			pf := NewPathfinder(filepath, test.sequence)
			totalDistance, err := pf.FindPathLength()
			assert.Equal(t, test.expectedDist, totalDistance)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
