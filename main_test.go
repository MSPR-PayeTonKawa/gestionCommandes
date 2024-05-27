package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleFunctionToTest(t *testing.T) {
	// Arrange
	expected := "Hello World"

	// Act
	result := simpleFunctionToTest("World")

	// Assert
	assert.Equal(t, expected, result)
}
