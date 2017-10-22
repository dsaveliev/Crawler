package models

import (
	"testing"
)

func TestNewLink(t *testing.T) {
	for _, tc := range testCasesNewLink {
		link := NewLink(tc.URL)
		if link.Hostname != tc.Hostname {
			t.Fatalf("link.Hostname = %#v. Want: %#v",
				link.Hostname, tc.Hostname)
		}
		if link.Error != tc.Error {
			t.Fatalf("link.Error = %v. Want: %v",
				link.Error, tc.Error)
		}
	}
}

func TestParseURL(t *testing.T) {
	for _, tc := range testCasesParseURL {
		switch hostname, err := parseURL(tc.URL); {
		case err != nil:
			if tc.Ok {
				t.Fatalf("parseURL(%q) returned error %q. Error not expected.",
					tc.URL, err)
			}
		case !tc.Ok:
			t.Fatalf("parseURL(%q) = %q, %v. Expected error.",
				tc.URL, hostname, err)
		case hostname != tc.Hostname:
			t.Fatalf("parseURL(%q) = %q, %v. Want: %s",
				tc.URL, hostname, err, tc.Hostname)
		}
	}
}
