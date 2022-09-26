package main

import "testing"

type PalindromeTest struct {
	value  string
	result bool
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}
func TestIsPalindrome(t *testing.T) {
	testData := []PalindromeTest{
		{value: "ana", result: true},
		{value: "bob", result: true},
		{value: "car", result: false},
		{value: "poop", result: true},
		{value: "poop", result: true},
	}
	for _, obj := range testData {
		if result := IsPalindrome(obj.value); result != obj.result {
			t.Errorf("IsPalindrome(%v) = %v, want %v", obj.value, result, obj.result)
		}
	}
}
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
