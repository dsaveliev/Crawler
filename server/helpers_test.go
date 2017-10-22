package server

import (
	"testing"
)

func TestValidateBody(t *testing.T) {
	for _, tc := range testCasesValidateBody {
		switch body, err := validateBody(tc.Body); {
		case err != nil:
			if tc.Ok {
				t.Fatalf("validateBody(%q) returned error %q. Error not expected.",
					tc.Body, err)
			}
		case !tc.Ok:
			t.Fatalf("validateBody(%q) = %q, %v.  Expected error.",
				tc.Body, body, err)
		case body != tc.Result:
			t.Fatalf("validateBody(%q) = %#v, %v.  Want %#v.",
				tc.Body, body, err, tc.Result)
		}
	}
}
