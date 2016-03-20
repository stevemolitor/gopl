package anagram

import "testing"

func TestIsAnagram(t *testing.T) {
    var tests = []struct {
        s1   string
        s2   string
        want bool
    }{
        {"aba", "aba", true},
        {"abc", "abcd", false},
        {"heart", "earth", true},
        {"", "a", false},
        {"a", "", false},
        {"", "", true},
    }

    for _, test := range tests {
        if got := IsAnagram(test.s1, test.s2); got != test.want {
            t.Errorf("IsAnagram(%s, %s): %t, expected %t",
                test.s1, test.s1, got, test.want)
        }
    }
}
