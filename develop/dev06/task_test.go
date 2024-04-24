package main

import (
	"reflect"
	"testing"
)

func TestGetSelectedColumnsWithFields(t *testing.T) {
	lines := []string{
		"1\t2\t3\t4",
		"5\t6\t7\t8",
	}
	fields := "1,2"
	delimiter := "\t"
	separated := false

	expectedOutput := []string{
		"1\t2",
		"5\t6",
	}

	output := getSelectedColumns(lines, fields, delimiter, separated)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Output was incorrect, got: %v, expected: %v", output, expectedOutput)
	}
}

func TestGetSelectedColumnsWithDelimiter(t *testing.T) {
	lines := []string{
		"test1:test2:test3:test4",
		"5:6:7:8",
		"test9:test10:test11:test12",
	}
	fields := "1,3"
	delimiter := ":"
	separated := false

	expectedOutput := []string{
		"test1:test3",
		"5:7",
		"test9:test11",
	}

	output := getSelectedColumns(lines, fields, delimiter, separated)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Output was incorrect, got: %v, expected: %v", output, expectedOutput)
	}
}

func TestGetSelectedColumnsWithDelimiterWithNotEqualStrings(t *testing.T) {
	lines := []string{
		"test1:test2:test3:test4",
		"5:6:7:8",
		"test9test12",
	}
	fields := "1,3"
	delimiter := ":"
	separated := false

	expectedOutput := []string{
		"test1:test3",
		"5:7",
		"test9test12",
	}

	output := getSelectedColumns(lines, fields, delimiter, separated)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Output was incorrect, got: %v, expected: %v", output, expectedOutput)
	}
}

func TestGetSelectedColumnsWithSeparated(t *testing.T) {
	lines := []string{
		"test1:test2:test3:test4",
		"5:6:7:8",
		"test9test12",
	}
	fields := "1,3"
	delimiter := ":"
	separated := true

	expectedOutput := []string{
		"test1:test3",
		"5:7",
	}

	output := getSelectedColumns(lines, fields, delimiter, separated)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Output was incorrect, got: %v, expected: %v", output, expectedOutput)
	}
}
