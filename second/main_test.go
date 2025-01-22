package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_Decode(t *testing.T) {

	t.Run("input = LLRR= output = 210122", func(t *testing.T) {
		msg := "LLRR="
		expected := "210122"

		result := Decode(msg)

		assert.Equal(t, expected, result)
	})

	t.Run("input = ==RLL output = 000210", func(t *testing.T) {
		msg := "==RLL"
		expected := "000210"

		result := Decode(msg)

		assert.Equal(t, expected, result)
	})

	t.Run("input = =LLRR output = 221012", func(t *testing.T) {
		msg := "=LLRR"
		expected := "221012"

		result := Decode(msg)

		assert.Equal(t, expected, result)
	})

	t.Run("input = RRL=R output = 012001", func(t *testing.T) {
		msg := "RRL=R"
		expected := "012001"

		result := Decode(msg)

		assert.Equal(t, expected, result)
	})
}
