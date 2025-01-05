package main

import "testing"

func TestUnpackValue(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"aaa0b", "aab"},
		{"adas4lopa5", "adasssslopaaaaa"},
		{"d\n5abc", "d\n\n\n\n\nabc"},
	}
	for _, test := range tests {
		result := unpackValue(test.input)
		if result != test.expected {
			t.Errorf("unpackValue(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestPackValue(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"aaaabccddddde", "a4bc2d5e"},
		{"abcd", "abcd"},
		{"aaaabbbcce", "a4b3c2e"},
		{"aabkkllgggahjjhdi", "a2bk2l2g3ahj2hdi"},
	}
	for _, test := range tests {
		result := packValue(test.input)
		if result != test.expected {
			t.Errorf("unpackValue(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
