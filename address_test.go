package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Go is Great",
			expected: []string{"go", "is", "great"},
		},
		{
			input:    "   Multiple   Spaces   ",
			expected: []string{"multiple", "spaces"},
		},
		{
			input: "PikachuIchooseYou",
			expected: []string{"pikachuichooseyou"},
		},
	}

	for _, c := range cases {
	actual := cleanInput(c.input)

	if len(actual) != len(c.expected) {
		t.Errorf("Length of expected result doesn't match with input.\n Expected Length:%d\n Input Length:%d\n", len(actual), len(c.expected))
		t.Fail()
		continue
	}

	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
	
		if word != expectedWord {
			t.Errorf("Output doesn't match with expected result.\n Output: %s\nExpected:%s\n",word,expectedWord)
			t.Fail()
		}
	}
}
}
