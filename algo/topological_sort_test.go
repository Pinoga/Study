package algo

import (
	"testing"
	. "turing/ds"

	"github.com/stretchr/testify/assert"
)

func TestTopologicalSort(t *testing.T) {
	testCases := []struct {
		value         Graph
		expectedIndex int
		expectedOrder []NodeValue
		desc          string
	}{
		{
			value:         DAG(),
			expectedIndex: -1,
			expectedOrder: []NodeValue{0, 2, 3, 9, 4, 1, 7, 8, 5, 6},
			desc:          "DAG",
		},
		{
			value:         DirectedCyclicGraph(),
			expectedIndex: 1,
			expectedOrder: []NodeValue{0, 2, 3, 9, 4, 1, 7, 8, 5, 6},
			desc:          "One Cycle",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var actual []NodeValue = make([]NodeValue, tC.value.Size())
			lastIndex := TopologicalSortRecursive(tC.value, 0, actual, tC.value.Size()-1)
			assert.Equal(t, tC.expectedIndex, lastIndex)
			assert.Equal(t, tC.expectedOrder, actual)
		})
	}
}
