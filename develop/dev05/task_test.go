package main

import (
	"reflect"
	"testing"
)

func TestFindMatches(t *testing.T) {
	lines := []string{
		"Line 1",
		"Line 2",
		"Matching line",
		"Line 4",
		"Line 5",
	}
	pattern := "Matching"
	options := map[string]interface{}{
		"ignoreCase": false,
		"fixed":      false,
		"invert":     false,
		"before":     0,
		"after":      0,
		"context":    0,
	}

	matchingLines, addedLines := findMatches(lines, pattern, options)

	expectedMatchingLines := map[int]string{
		2: "Matching line",
	}

	expectedAddedLines := []int{2}

	if !reflect.DeepEqual(matchingLines, expectedMatchingLines) {
		t.Errorf("Expected matching lines: %v, got: %v", expectedMatchingLines, matchingLines)
	}

	if !reflect.DeepEqual(addedLines, expectedAddedLines) {
		t.Errorf("Expected added lines: %v, got: %v", expectedAddedLines, addedLines)
	}
}

func TestFindMatchesWithIgnoreCase(t *testing.T) {
	lines := []string{
		"Line 1",
		"Line 2",
		"Matching line",
		"Line 4",
		"Line 5",
	}
	pattern := "matching"
	options := map[string]interface{}{
		"ignoreCase": true,
		"fixed":      false,
		"invert":     false,
		"before":     0,
		"after":      0,
		"context":    0,
	}

	matchingLines, addedLines := findMatches(lines, pattern, options)

	expectedMatchingLines := map[int]string{
		2: "Matching line",
	}

	expectedAddedLines := []int{2}

	if !reflect.DeepEqual(matchingLines, expectedMatchingLines) {
		t.Errorf("Expected matching lines: %v, got: %v", expectedMatchingLines, matchingLines)
	}

	if !reflect.DeepEqual(addedLines, expectedAddedLines) {
		t.Errorf("Expected added lines: %v, got: %v", expectedAddedLines, addedLines)
	}
}

func TestFindMatchesWithFixed(t *testing.T) {
	lines := []string{
		"Line 1",
		"line",
		"Matching line",
		"Line 4",
		"Line 5",
	}
	pattern := "line"
	options := map[string]interface{}{
		"ignoreCase": false,
		"fixed":      true,
		"invert":     false,
		"before":     0,
		"after":      0,
		"context":    0,
	}

	matchingLines, addedLines := findMatches(lines, pattern, options)

	expectedMatchingLines := map[int]string{
		1: "line",
	}

	expectedAddedLines := []int{1}

	if !reflect.DeepEqual(matchingLines, expectedMatchingLines) {
		t.Errorf("Expected matching lines: %v, got: %v", expectedMatchingLines, matchingLines)
	}

	if !reflect.DeepEqual(addedLines, expectedAddedLines) {
		t.Errorf("Expected added lines: %v, got: %v", expectedAddedLines, addedLines)
	}
}

func TestFindMatchesWithBefore(t *testing.T) {
	lines := []string{
		"Line 1",
		"line",
		"Matching line",
		"Line 4",
		"Line 5",
	}
	pattern := "Matching"
	options := map[string]interface{}{
		"ignoreCase": false,
		"fixed":      false,
		"invert":     false,
		"before":     2,
		"after":      0,
		"context":    0,
	}

	matchingLines, addedLines := findMatches(lines, pattern, options)

	expectedMatchingLines := map[int]string{
		0: "Line 1",
		1: "line",
		2: "Matching line",
	}

	expectedAddedLines := []int{0, 1, 2}

	if !reflect.DeepEqual(matchingLines, expectedMatchingLines) {
		t.Errorf("Expected matching lines: %v, got: %v", expectedMatchingLines, matchingLines)
	}

	if !reflect.DeepEqual(addedLines, expectedAddedLines) {
		t.Errorf("Expected added lines: %v, got: %v", expectedAddedLines, addedLines)
	}
}

func TestFindMatchesWithAfter(t *testing.T) {
	lines := []string{
		"Line 1",
		"line",
		"Matching line",
		"Line 4",
		"Line 5",
	}
	pattern := "Matching"
	options := map[string]interface{}{
		"ignoreCase": false,
		"fixed":      false,
		"invert":     false,
		"before":     0,
		"after":      2,
		"context":    0,
	}

	matchingLines, addedLines := findMatches(lines, pattern, options)

	expectedMatchingLines := map[int]string{
		2: "Matching line",
		3: "Line 4",
		4: "Line 5",
	}

	expectedAddedLines := []int{2, 3, 4}

	if !reflect.DeepEqual(matchingLines, expectedMatchingLines) {
		t.Errorf("Expected matching lines: %v, got: %v", expectedMatchingLines, matchingLines)
	}

	if !reflect.DeepEqual(addedLines, expectedAddedLines) {
		t.Errorf("Expected added lines: %v, got: %v", expectedAddedLines, addedLines)
	}
}

func TestFindMatchesWithContext(t *testing.T) {
	lines := []string{
		"Line 1",
		"line",
		"Matching line",
		"Line 4",
		"Line 5",
	}
	pattern := "Matching"
	options := map[string]interface{}{
		"ignoreCase": false,
		"fixed":      false,
		"invert":     false,
		"before":     0,
		"after":      0,
		"context":    2,
	}

	matchingLines, addedLines := findMatches(lines, pattern, options)

	expectedMatchingLines := map[int]string{
		0: "Line 1",
		1: "line",
		2: "Matching line",
		3: "Line 4",
		4: "Line 5",
	}

	expectedAddedLines := []int{0, 1, 2, 3, 4}

	if !reflect.DeepEqual(matchingLines, expectedMatchingLines) {
		t.Errorf("Expected matching lines: %v, got: %v", expectedMatchingLines, matchingLines)
	}

	if !reflect.DeepEqual(addedLines, expectedAddedLines) {
		t.Errorf("Expected added lines: %v, got: %v", expectedAddedLines, addedLines)
	}
}

func TestFindMatchesWithContextAndInvert(t *testing.T) {
	lines := []string{
		"Line 1",
		"line",
		"Matching line",
		"Line 4",
		"Line 5",
	}
	pattern := "Matching"
	options := map[string]interface{}{
		"ignoreCase": false,
		"fixed":      false,
		"invert":     true,
		"before":     0,
		"after":      0,
		"context":    1,
	}

	matchingLines, addedLines := findMatches(lines, pattern, options)

	expectedMatchingLines := map[int]string{
		0: "Line 1",
		1: "line",
		2: "Matching line",
		3: "Line 4",
		4: "Line 5",
	}

	expectedAddedLines := []int{0, 1, 2, 3, 4}

	if !reflect.DeepEqual(matchingLines, expectedMatchingLines) {
		t.Errorf("Expected matching lines: %v, got: %v", expectedMatchingLines, matchingLines)
	}

	if !reflect.DeepEqual(addedLines, expectedAddedLines) {
		t.Errorf("Expected added lines: %v, got: %v", expectedAddedLines, addedLines)
	}
}

func TestFindMatchesWithContextAndBefore(t *testing.T) {
	lines := []string{
		"Line 1",
		"line",
		"Matching line",
		"Line 4",
		"Line 5",
	}
	pattern := "Matching"
	options := map[string]interface{}{
		"ignoreCase": false,
		"fixed":      false,
		"invert":     false,
		"before":     2,
		"after":      0,
		"context":    1,
	}

	matchingLines, addedLines := findMatches(lines, pattern, options)

	expectedMatchingLines := map[int]string{
		0: "Line 1",
		1: "line",
		2: "Matching line",
		3: "Line 4",
	}

	expectedAddedLines := []int{0, 1, 2, 3}

	if !reflect.DeepEqual(matchingLines, expectedMatchingLines) {
		t.Errorf("Expected matching lines: %v, got: %v", expectedMatchingLines, matchingLines)
	}

	if !reflect.DeepEqual(addedLines, expectedAddedLines) {
		t.Errorf("Expected added lines: %v, got: %v", expectedAddedLines, addedLines)
	}
}
