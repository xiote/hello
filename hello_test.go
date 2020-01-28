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

func TestValidate(t *testing.T) {
	cases := []struct {
		inName    string
		wantValid bool
	}{
		{"Brian", true},
		{"", false},
	}
	for _, c := range cases {
		gotValid := Validate(c.inName)
		if gotValid != c.wantValid {
			t.Errorf("Validate(%q) == %t, want %t", c.inName, gotValid, c.wantValid)
		}
	}

}
