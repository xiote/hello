package main

import "testing"

func TestHello(t *testing.T) {
	cases := []struct {
		inName, wantMessage string
	}{
		{"Brian", "Hello, Brian, nice to meet you!"},
	}
	for _, c := range cases {
		gotMessage := Hello(c.inName)
		if gotMessage != c.wantMessage {
			t.Errorf("Hello(%q) == %q, want %q", c.inName, gotMessage, c.wantMessage)
		}
	}

}
