package main

import (
	"reflect"
	"testing"
)

func TestSortLines(t *testing.T) {
	lines := []string{"3", "1", "2"}
	sorted := sortLines(lines, 0, true, false, false)

	expected := []string{"1", "2", "3"}
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, sorted)
	}
}
func TestSortLinesNumericReverseUnique(t *testing.T) {
	lines := []string{"3", "3", "1", "2", "2"}
	sorted := sortLines(lines, 0, true, true, true)

	expected := []string{"3", "2", "1"}
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, sorted)
	}
}

func TestSortLinesNonNumericReverseUnique(t *testing.T) {
	lines := []string{"c", "a", "b"}
	sorted := sortLines(lines, 0, false, true, true)

	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, sorted)
	}
}

func TestSortLinesWithKeyAndNum(t *testing.T) {
	lines := []string{"5 vyacheslav", "1 ivan", "2 vasiliy", "15 maxim", "3 alex", "4 dmitry"}
	sorted := sortLines(lines, 1, true, false, false)

	expected := []string{"1 ivan", "2 vasiliy", "3 alex", "4 dmitry", "5 vyacheslav", "15 maxim"}
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, sorted)
	}
}

func TestSortLinesWithKeyAndNumAndReverse(t *testing.T) {
	lines := []string{"5 vyacheslav", "1 ivan", "2 vasiliy", "2 vasiliy", "15 maxim", "3 alex", "4 dmitry"}
	sorted := sortLines(lines, 1, true, true, true)

	expected := []string{"15 maxim", "5 vyacheslav", "4 dmitry", "3 alex", "2 vasiliy", "1 ivan"}
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, sorted)
	}
}
