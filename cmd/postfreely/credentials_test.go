package main

import (
	"testing"
)

func TestParseCredentials(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedUsername string
		ExpectedPassword string
	}{
		{
			Value:            ":",
			ExpectedUsername: "",
			ExpectedPassword:  "",
		},



		{
			Value:            "username:password",
			ExpectedUsername: "username",
			ExpectedPassword:          "password",
		},
		{
			Value:            "joeblow:supeR.s3crET",
			ExpectedUsername: "joeblow",
			ExpectedPassword:         "supeR.s3crET",
		},
		{
			Value:            "dariush:KIng123",
			ExpectedUsername: "dariush",
			ExpectedPassword:         "KIng123",
		},
		{
			Value:            "malekeh:qu33n_POW3r",
			ExpectedUsername: "malekeh",
			ExpectedPassword:         "qu33n_POW3r",
		},



		{
			Value:            "username:",
			ExpectedUsername: "username",
			ExpectedPassword:          "",
		},
		{
			Value:            "joeblow:",
			ExpectedUsername: "joeblow",
			ExpectedPassword:         "",
		},
		{
			Value:            "dariush:",
			ExpectedUsername: "dariush",
			ExpectedPassword:         "",
		},
		{
			Value:            "malekeh:",
			ExpectedUsername: "malekeh",
			ExpectedPassword:         "",
		},



		{
			Value:            ":password",
			ExpectedUsername: "",
			ExpectedPassword:  "password",
		},
		{
			Value:            ":supeR.s3crET",
			ExpectedUsername: "",
			ExpectedPassword:  "supeR.s3crET",
		},
		{
			Value:            ":KIng123",
			ExpectedUsername: "",
			ExpectedPassword:  "KIng123",
		},
		{
			Value:            ":qu33n_POW3r",
			ExpectedUsername: "",
			ExpectedPassword:  "qu33n_POW3r",
		},
	}

	for testNumber, test := range tests {

		actualUsername, actualPassword, err := parseCredentials(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		{
			expected := test.ExpectedUsername
			actual   := actualUsername

			if expected != actual {
				t.Errorf("For test #%d, the actual username was not what was expected.", testNumber)
				t.Logf("EXPECTED USERNAME: %q", expected)
				t.Logf("ACTUAL USERNAME:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedPassword
			actual   := actualPassword

			if expected != actual {
				t.Errorf("For test #%d, the actual password was not what was expected.", testNumber)
				t.Logf("EXPECTED PASSWORD: %q", expected)
				t.Logf("ACTUAL PASSWORD:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}

func TestParseCredentials_fail(t *testing.T) {

	tests := []struct{
		Value string
	}{
		{
			Value: "",
		},



		{
			Value: "username",
		},
		{
			Value: "password",
		},
	}

	for testNumber, test := range tests {

		actualUsername, actualPassword, err := parseCredentials(test.Value)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		if expected, actual := "", actualUsername; expected != actual {
			t.Errorf("For test #%d, expected the username to be empty but actually wasn't.", testNumber)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("ACTUAL USERNAME: %q", actualUsername)
			t.Logf("ACTUAL PASSWORD: %q", actualPassword)
			continue
		}

		if expected, actual := "", actualPassword; expected != actual {
			t.Errorf("For test #%d, expected the password to be empty but actually wasn't.", testNumber)
			t.Logf("VALUE: %q", test.Value)
			t.Logf("ACTUAL USERNAME: %q", actualUsername)
			t.Logf("ACTUAL PASSWORD: %q", actualPassword)
			continue
		}
	}
}
