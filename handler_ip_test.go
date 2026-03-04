package main

import (
	"testing"
)

func TestParseClientFromXFF(t *testing.T) {
	testCases := [][]string{
		[]string{"", ""},
		[]string{"1.1.1.1", "1.1.1.1"},
		[]string{"1.1.1.1,2.2.2.2", "1.1.1.1"},
		[]string{" 1.1.1.1 , 2.2.2.2", "1.1.1.1"},
		[]string{"arlkjnlakrujaslduh", ""},
		[]string{"sfhsadfhafhahfr,1.1.1.1", ""},
		[]string{"1.1.1.1:1234,2.2.2.2:1234","1.1.1.1"},
	}

	for _, testCase := range testCases {
		xForwardedFor := testCase[0]
		want := testCase[1]

		clientIp := parseClientFromXFF(xForwardedFor)

		if clientIp != want {
			t.Errorf(`parseClientFromXFF(%s) = %s, want match for %#q`, xForwardedFor, clientIp, want)
		}
	}
}
