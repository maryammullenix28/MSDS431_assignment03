package main

import (
	"reflect"
	"testing"
)

func TestExtractColumn(t *testing.T) {
	// Test data
	data := [][]string{
		{"value", "income", "age"},
		{"1", "50000", "30"},
		{"2", "60000", "35"},
		{"3", "70000", "40"},
	}

	expected := []float64{1, 2, 3}

	columnName := "value"
	result := extractColumn(data, columnName)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ExtractColumn failed, expected %v but got %v", expected, result)
	}
}

func TestComputeMean(t *testing.T) {
	// Test data
	values := []float64{1, 2, 3, 4, 5}

	expected := 3.0

	result := computeMean(values)

	if result != expected {
		t.Errorf("ComputeMean failed, expected %f but got %f", expected, result)
	}
}
