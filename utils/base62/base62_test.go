package base62

import (
	"testing"
)

var testcases = []struct {
	num      int64
	expected string
}{
	{1, "1"},
	{9, "9"},
	{10, "A"},
	{35, "Z"},
	{36, "a"},
	{61, "z"},
	{62, "10"},
	{99, "1b"},
	{3844, "100"},
	{3860, "10G"},
	{4815162342, "5Frvgk"},
}

func TestEncodeBase62(t *testing.T) {
	for _, tc := range testcases {
		result := EncodeBase62(int64(tc.num))
		if result != tc.expected {
			t.Errorf("Got %v as %s expected %s", tc.num, result, tc.expected)
		}

	}
}
