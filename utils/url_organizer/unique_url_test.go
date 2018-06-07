package url_organizer

import "testing"

var testcases = []struct {
	original string
}{
	{"https://www.youtube.com/watch?v=vjhmJaeI5ZI"},
	{"http://www.youtube.com/watch?v=vjhmJaeI5ZI"},
	{"https://youtube.com/watch?v=vjhmJaeI5ZI"},
	{"http://youtube.com/watch?v=vjhmJaeI5ZI"},
	{"www.youtube.com/watch?v=vjhmJaeI5ZI"},
	{"youtube.com/watch?v=vjhmJaeI5ZI"},
}

const expected1 string = "https://www.youtube.com?v=vjhmJaeI5ZI"

func TestGetUniqueURL(t *testing.T) {
	for _, tc := range testcases {
		result := GetUniqueURL(tc.original)
		if result != expected1 {
			t.Errorf("Got %s as %s  expected %s", result, tc.original, expected1)
		}

	}
}
