package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_MaxRoute(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		tree := [][]int{
			{59},
			{73, 41},
			{52, 40, 53},
			{26, 53, 6, 34},
		}

		result := MaxRoute(tree)

		assert.Equal(t, 237, result)
	})

	t.Run("load input from file", func(t *testing.T) {
		data, err := os.ReadFile("files/hard.json")
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		var tree [][]int
		err = json.Unmarshal(data, &tree)
		if err != nil {
			t.Fatalf("failed to unmarshal json: %v", err)
		}
		expectedValue := 7273

		result := MaxRoute(tree)

		assert.Equal(t, expectedValue, result)
	})

}
