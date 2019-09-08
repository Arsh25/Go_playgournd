package main

import (
	"reflect"
	"testing"
)

func checkCorrect(expected, result []int) (passed bool) {
	passed = reflect.DeepEqual(result, expected)
	return
}

func TestFibo0(t *testing.T) {
	result := fiboSeries(0)
	expected := []int{1}
	if !checkCorrect(expected, result) {
		t.Errorf("fiboSeries 0 returned %v. Expected %v", result, expected)
	}
}

func TestFibo1(t *testing.T) {
	result := fiboSeries(1)
	expected := []int{1, 1}
	if !checkCorrect(expected, result) {
		t.Errorf("fibo 1 returned %v. Expected %v", result, expected)
	}
}

func TestFiboGreaterThan1(t *testing.T) {
	num := 2
	result := fiboSeries(num)
	expected := []int{1, 1, 2}
	if !checkCorrect(expected, result) {
		t.Errorf("fiboSeries %v returned %v. Expected %v", num, result, expected)
	}
	num = 3
	result = fiboSeries(num)
	expected = append(expected, 3)
	if !checkCorrect(expected, result) {
		t.Errorf("fiboSeries %v returned %v. Expected %v", num, result, expected)
	}
	num = 4
	result = fiboSeries(num)
	expected = append(expected, 5)
	if !checkCorrect(expected, result) {
		t.Errorf("fiboSeries %v returned %v. Expected %v", num, result, expected)
	}
	num = 5
	result = fiboSeries(num)
	expected = append(expected, 8)
	if !checkCorrect(expected, result) {
		t.Errorf("fiboSeries %v returned %v. Expected %v", num, result, expected)
	}
}
