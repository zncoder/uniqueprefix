package uniqueprefix

import (
	"reflect"
	"testing"
)

func TestPrefixes(t *testing.T) {
	testcases := []struct {
		input  []string
		output []string
		unique bool
	}{
		{[]string{"a"}, []string{"a"}, true},
		{[]string{"abc"}, []string{"a"}, true},
		{[]string{"a", "a"}, []string{"a", "a"}, false},
		{[]string{"ab", "b", "ab", "abc", "abde"}, []string{"ab", "b", "ab", "abc", "abd"}, false},
		{[]string{"a", "ab"}, []string{"a", "ab"}, true},
		{[]string{"a", "ab", "ac"}, []string{"a", "ab", "ac"}, true},
		{[]string{"a", "abcd", "ae", "abfg"}, []string{"a", "abc", "ae", "abf"}, true},
		{[]string{"a", "bc"}, []string{"a", "b"}, true},
	}

	for i, tc := range testcases {
		got, uniq := Prefixes(tc.input...)
		if !reflect.DeepEqual(got, tc.output) {
			t.Fatalf("case %d: want:%v got:%v", i, tc.output, got)
		}
		if uniq != tc.unique {
			t.Fatalf("case %d: unique want:%v got:%v", i, tc.unique, uniq)
		}
	}
}

func TestTrie(t *testing.T) {
	tr := NewTrie()
	input := []struct {
		s     string
		added bool
	}{
		{"a", true},
		{"ab", true},
		{"abc", true},
		{"ab", false},
		{"abc", false},
		{"ac", true},
		{"ba", true},
		{"c", true},
		{"da", true},
		{"abcdef", true},
	}
	for i, x := range input {
		ok := tr.Add(x.s)
		if ok != x.added {
			t.Fatalf("input %d add want:%v got:%v", i, x.added, ok)
		}
	}

	output := []struct {
		s      string
		prefix string
		ok     bool
	}{
		{"bc", "b", false},
		{"", "", true},
		{"a", "a", true},
		{"abc", "abc", true},
		{"ab", "ab", true},
		{"ac", "ac", true},
		{"acd", "ac", false},
		{"ba", "b", true},
		{"c", "c", true},
		{"abcdef", "abcd", true},
		{"da", "d", true},
		{"xyz", "", false},
	}
	for i, x := range output {
		got, ok := tr.Prefix(x.s)
		if got != x.prefix {
			t.Fatalf("output %d: want:%q got:%q", i, x.prefix, got)
		}
		if ok != x.ok {
			t.Fatalf("output %d: ok want:%v got:%v", i, x.ok, ok)
		}
	}
}
