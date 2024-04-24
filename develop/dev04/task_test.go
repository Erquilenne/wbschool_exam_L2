package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramGroups(t *testing.T) {
	words := []string{"пятак", "листок", "кулон", "столик", "пятка", "тяпка", "слиток", "клоун", "уклон"}
	expected := map[string][]string{
		"пятак":  []string{"пятак", "пятка", "тяпка"},
		"листок": []string{"листок", "столик", "слиток"},
		"кулон":  []string{"кулон", "клоун", "уклон"},
	}

	result := findAnagramGroups(&words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected result. Got: %v, Want: %v", result, expected)
	}
}

func TestFindAnagramGroupsUpperCase(t *testing.T) {
	words := []string{"ПЯТАК", "листок", "КУлон", "Столик", "ПЯТКа", "тяпка", "СЛиТОК", "клоун", "уклон"}
	expected := map[string][]string{
		"пятак":  []string{"пятак", "пятка", "тяпка"},
		"листок": []string{"листок", "столик", "слиток"},
		"кулон":  []string{"кулон", "клоун", "уклон"},
	}

	result := findAnagramGroups(&words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected result. Got: %v, Want: %v", result, expected)
	}
}
func TestFindAnagramGroupsOnceElement(t *testing.T) {
	words := []string{"цветок", "пятак", "листок", "кулон", "столик", "пятка", "тяпка", "слиток", "клоун", "уклон"}
	expected := map[string][]string{
		"пятак":  []string{"пятак", "пятка", "тяпка"},
		"листок": []string{"листок", "столик", "слиток"},
		"кулон":  []string{"кулон", "клоун", "уклон"},
	}

	result := findAnagramGroups(&words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected result. Got: %v, Want: %v", result, expected)
	}
}
